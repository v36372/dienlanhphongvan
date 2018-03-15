package ulog

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
	"utilities/file"
	"utilities/uerror"

	"github.com/Sirupsen/logrus"
	eParser "github.com/go-errors/errors"
	"github.com/streadway/amqp"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logNameRequest = "requests.log"
	logNameWarning = "warning.log"
	logNameError   = "error.log"
	logNameInfo    = "info.log"
	logNameDebug   = "debug.log"
	defaultLogDir  = "./log"
)

type ErrLog struct {
	Link    string
	Method  string
	Latency time.Duration
	Error   error
	Status  int
	Client  int
	UserID  int
	Header  http.Header
	Body    string
}

type Ulogger struct {
	LogPath   string
	Request   *lumberjack.Logger
	Warning   *logrus.Logger
	Error     *logrus.Logger
	Info      *logrus.Logger
	Debug     *logrus.Logger
	init      bool
	debugMode bool
}

var (
	defaultLogger Ulogger
)

func Logger() Ulogger {
	if !defaultLogger.init {
		InitDefaultLogger(defaultLogDir, false)
		log.Println("[WARNING] log-path is not set, default path './log' will be used")
	}
	return defaultLogger
}

type Fields logrus.Fields

type Http struct {
	Link    string
	Method  string
	Latency time.Duration
	Status  int
	Client  int
	UserID  int
	Header  http.Header
	Body    string
	Err     error
}

func (v Http) Fields() Fields {
	ret := Fields{
		"link":    v.Link,
		"method":  v.Method,
		"latency": v.Latency,
		"status":  v.Status,
		"client":  v.Client,
		"userId":  v.UserID,
		"header":  v.Header,
		"body":    v.Body,
	}
	if v.Err != nil {
		ret["error"] = v.Err.Error()
	}
	return ret
}

func InitDefaultLogger(logPath string, debugMode bool) {
	err := file.CreateDir(logPath)
	if err != nil {
		panic(err)
	}
	defaultLogger = Ulogger{
		LogPath:   logPath,
		init:      true,
		Request:   newLogWriter(path.Join(logPath, logNameRequest)),
		Warning:   newLogger(path.Join(logPath, logNameWarning), logrus.WarnLevel),
		Error:     newLogger(path.Join(logPath, logNameError), logrus.ErrorLevel),
		Info:      newLogger(path.Join(logPath, logNameInfo), logrus.InfoLevel),
		Debug:     newLogger(path.Join(logPath, logNameDebug), logrus.DebugLevel),
		debugMode: debugMode,
	}
}

// Create logger with path <prefix_name>_<error/info/warning>.log
func NewLogger(logPath, prefixName string, debugMode bool) Ulogger {
	if logPath == "" {
		logPath = defaultLogger.LogPath
	}
	requests := []string{logNameRequest}
	warnings := []string{logNameWarning}
	errors := []string{logNameError}
	infos := []string{logNameInfo}
	debugs := []string{logNameDebug}
	if len(prefixName) > 0 {
		requests = append([]string{prefixName}, requests...)
		warnings = append([]string{prefixName}, warnings...)
		errors = append([]string{prefixName}, errors...)
		infos = append([]string{prefixName}, infos...)
		debugs = append([]string{prefixName}, debugs...)
	}

	return Ulogger{
		LogPath:   logPath,
		init:      true,
		Request:   newLogWriter(path.Join(logPath, strings.Join(requests, "_"))),
		Warning:   newLogger(path.Join(logPath, strings.Join(warnings, "_")), logrus.WarnLevel),
		Error:     newLogger(path.Join(logPath, strings.Join(errors, "_")), logrus.ErrorLevel),
		Info:      newLogger(path.Join(logPath, strings.Join(infos, "_")), logrus.InfoLevel),
		Debug:     newLogger(path.Join(logPath, strings.Join(debugs, "_")), logrus.DebugLevel),
		debugMode: debugMode,
	}
}

func HookDefaultLogerWithFluent(fluent Fluent) {
	hook, err := newHookFluent(fluent)
	if err != nil {
		panic(err)
	}
	defaultLogger.AddHook(hook)
}

func HookLoggerWithFluent(log *Ulogger, fluent Fluent) {
	if log == nil {
		panic(fmt.Errorf("logger: nil"))
	}
	hook, err := newHookFluent(fluent)
	if err != nil {
		panic(err)
	}
	log.AddHook(hook)
}

func newLogger(fileName string, level logrus.Level) *logrus.Logger {
	return &logrus.Logger{
		Out:       newLogWriter(fileName),
		Formatter: new(logrus.JSONFormatter),
		Level:     level,
		Hooks:     make(logrus.LevelHooks),
	}
}

func newLogWriter(log string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   log,
		MaxSize:    30, // megabytes
		MaxBackups: 100,
		MaxAge:     30, // days
	}
}

func (l *Ulogger) AddHook(hook logrus.Hook) {
	if l.Error != nil {
		l.Error.Hooks.Add(hook)
	}
	if l.Warning != nil {
		l.Warning.Hooks.Add(hook)
	}
	if l.Info != nil {
		l.Info.Hooks.Add(hook)
	}
	if l.Debug != nil {
		l.Debug.Hooks.Add(hook)
	}
}

func (l Ulogger) LogErrorManual(err error, mess string) {
	err = GetErrorStackTrace(err)
	errMess := fmt.Sprint(err)
	l.Error.WithFields(logrus.Fields{
		"time": time.Now(),
		"mess": mess,
	}).Error(errMess)
}

func (l Ulogger) LogError(msg string, fields Fields) {
	l.Error.WithFields(logrus.Fields(fields)).
		Error(msg)
}

func (l Ulogger) LogInfo(msg string, fields Fields) {
	l.Info.WithFields(logrus.Fields(fields)).
		Info(msg)
}

func (l Ulogger) LogWarn(msg string, fields Fields) {
	l.Warning.WithFields(logrus.Fields(fields)).
		Warn(msg)
}

func (l Ulogger) LogDebug(msg string, fields Fields) {
	l.Debug.WithFields(logrus.Fields(fields)).
		Debug(msg)
}

func (l Ulogger) LogErrorObjectManual(err error, mess string, data interface{}) {
	dataStr := fmt.Sprintf("%+v", data)
	err = GetErrorStackTrace(err)
	errMess := fmt.Sprint(err)
	l.Error.WithFields(logrus.Fields{
		"time": time.Now(),
		"data": dataStr,
		"mess": mess,
	}).Error(errMess)
}

func (l Ulogger) LogDebugObject(mess string, data interface{}) {
	if !l.debugMode {
		return
	}
	dataStr := fmt.Sprintf("%+v", data)
	l.Debug.WithFields(logrus.Fields{
		"time": time.Now(),
		"data": dataStr,
	}).Debug(mess)

}

func (l Ulogger) LogInfoObject(mess string, data interface{}) {
	dataStr := fmt.Sprintf("%+v", data)
	l.Info.WithFields(logrus.Fields{
		"time": time.Now(),
		"data": dataStr,
	}).Info(mess)
}

func (l Ulogger) LogInfoStat(duration time.Duration, msg string) {
	l.Info.WithFields(logrus.Fields{
		"time":     time.Now(),
		"duration": duration,
	}).Info(msg)
}

func (l Ulogger) LogDebugStat(duration time.Duration, msg string) {
	if !l.debugMode {
		return
	}
	l.Debug.WithFields(logrus.Fields{
		"time":     time.Now(),
		"duration": duration,
	}).Debug(msg)
}

func (l Ulogger) LogErrorStat(duration time.Duration, msg string, err error) {

	l.Error.WithFields(logrus.Fields{
		"time":     time.Now(),
		"duration": duration,
		"error":    err,
	}).Error(msg)
}

func (l Ulogger) LogWarning(err error, data interface{}, mess string) {
	err = GetErrorStackTrace(err)
	dataStr := fmt.Sprintf("%+v", data)
	errMess := fmt.Sprint(uerror.StackTrace(err))
	l.Warning.WithFields(logrus.Fields{
		"data": dataStr,
		"err":  errMess,
	}).Warn(mess)
}

func (l Ulogger) LogRequestObject(req *http.Request, body string, input interface{}, data interface{}, msg string) {
	l.Info.WithFields(logrus.Fields{
		"body":   body,
		"input":  input,
		"res":    data,
		"link":   req.RequestURI,
		"method": req.Method,
		"client": req.Header.Get("client"),
		"header": req.Header,
	}).Info(msg)
}

func getLogDataMessageAmqp(msg *amqp.Delivery, funcName string, latency time.Duration, err error) logrus.Fields {
	logData := logrus.Fields{
		"funcName": funcName,
		"latency":  latency,
	}
	if err != nil {
		logData["error"] = uerror.StackTrace(err)
	}
	if msg != nil {
		logData["msgBody"] = string(msg.Body)
		logData["msgType"] = msg.Type
		logData["msgContentType"] = msg.ContentType
	}
	return logData
}

func (l Ulogger) LogErrorMessageAmqp(msg *amqp.Delivery, funcName string, note string, latency time.Duration, err error) {
	logData := getLogDataMessageAmqp(msg, funcName, latency, err)
	l.Error.WithFields(logData).Error(note)
}

func (l Ulogger) LogInfoMessageAmqp(msg *amqp.Delivery, funcName string, note string, latency time.Duration) {
	logData := getLogDataMessageAmqp(msg, funcName, latency, nil)
	l.Info.WithFields(logData).Info(note)
}

func (l Ulogger) LogWarningMessageAmqp(msg *amqp.Delivery, funcName string, note string, latency time.Duration) {
	logData := getLogDataMessageAmqp(msg, funcName, latency, nil)
	l.Warning.WithFields(logData).Warn(note)
}

func (l Ulogger) LogDebugMessageAmqp(msg *amqp.Delivery, funcName string, note string, latency time.Duration) {
	if !l.debugMode {
		return
	}
	logData := getLogDataMessageAmqp(msg, funcName, latency, nil)
	l.Debug.WithFields(logData).Debug(note)
}

func GetErrorStackTrace(e error) error {
	return errors.New(eParser.Wrap(e, 0).ErrorStack())
}

func BackupBody(req *http.Request) string {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	if len(bodyString) > 2000 {
		bodyString = bodyString[:2000]
	}
	return bodyString
}

func (l Ulogger) LogErrorHttp(errLog ErrLog) {
	errMess := fmt.Sprint(errLog.Error)
	l.Error.WithFields(logrus.Fields{
		"time":   time.Now(),
		"status": errLog.Status,
		"method": errLog.Method,
		"header": errLog.Header,
		"link":   errLog.Link,
		"userId": errLog.UserID,
		"client": errLog.Client,
		"body":   errLog.Body,
	}).Error(errMess)
}
