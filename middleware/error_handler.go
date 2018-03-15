package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"dienlanhphongvan/utilities/ulog"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ErrorInternalServer   string = "INTERNAL_SERVER_ERROR"
	ErrorNotFound         string = "DATA_NOT_FOUND"
	ErrorUnauthenticated  string = "UNAUTHENTICATED"
	errorPermissionDenied string = "PERMISSION_DENIED"
	ErrorSpammer          string = "Dude, you are running so fast"
	ERR_INTERNAL_ERROR    string = "INTERNAL_SERVER_ERROR"
	ErrorReachLimit       string = "REACH_LIMIT"
)

type warningData struct {
	URL    string
	Method string
}

type Errs struct {
	Errors []Err `json:"errors"`
}

type Err struct {
	Key  string `json:"key"`
	Mess string `json:"message"`
}

type InternalServerError500 struct {
	Err error
}

type PermissionDeniedError403 struct {
	Err error
}

type HandleError400 struct {
	Err error
}

func (msg InternalServerError500) Error() string {
	return msg.Err.Error()
}

func (msg PermissionDeniedError403) Error() string {
	return msg.Err.Error()
}

func (msg HandleError400) Error() string {
	return msg.Err.Error()
}

func ManualLogError(err error, mess string) {
	if err != nil {
		ulog.Logger().LogErrorManual(err, mess)
	}
}

func errorInternalServer(errLog ulog.ErrLog) map[string]interface{} {
	ulog.Logger().LogErrorHttp(errLog)

	return gin.H{
		"ERROR_CODE": ErrorInternalServer,
	}
}

func errorInvalidRequest(errLog ulog.ErrLog) interface{} {
	ulog.Logger().LogErrorHttp(errLog)
	dat := []Err{}
	if er := json.Unmarshal([]byte(errLog.Error.Error()), &dat); er == nil {

		return Errs{Errors: dat}
	} else {

		return gin.H{"errors": errLog.Error.Error()}
	}
}

func errorSpammerDetected(c *gin.Context, err error) map[string]interface{} {
	ulog.Logger().LogWarning(err, warningData{URL: c.Request.URL.String(), Method: c.Request.Method}, "Spammer detected")
	return gin.H{
		"ERROR_CODE": ErrorSpammer,
	}
}

func errorReachLimit(c *gin.Context, err error) map[string]interface{} {
	ulog.Logger().LogWarning(err, warningData{URL: c.Request.URL.String(), Method: c.Request.Method}, "Spammer detected")
	return gin.H{
		"ERROR_CODE": ErrorReachLimit,
	}
}

func ErrorHandler(transactionName string) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)

		//Before logic
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		// After logic
		if err := c.Errors.Last(); err != nil {
			if c.Writer.Status() != 404 {
				/*clientId := utils.GetClient(c)*/
				/*userId := ""
				if  != "" {
					if currUser, _ := getCurrentUser(c); currUser != nil {
						userId = int(currUser.Id)
					}
				}*/

				switch c.Writer.Status() {
				case 500:
					fmt.Println("")
					fmt.Println("[ERROR]===========")
					fmt.Println("[ERROR] Link:", c.Request.RequestURI)
					fmt.Println("[ERROR] Method:", c.Request.Method)
					fmt.Println("[ERROR] Latency:", latency)
					fmt.Println("[ERROR] Error:", err)
					fmt.Println("[ERROR] Status:", c.Writer.Status())
					/*fmt.Println("[ERROR] Client:", clientId)*/
					fmt.Println("[ERROR] Header:", c.Request.Header)
					fmt.Println("[ERROR] Body:", bodyString)
					fmt.Println("[ERROR]===========")
				case 429:
					c.JSON(429, errorSpammerDetected(c, err))
				case 401:
				case 400:
					fmt.Println("")
					fmt.Println("[ERROR]===========")
					fmt.Println("[ERROR] Link:", c.Request.RequestURI)
					fmt.Println("[ERROR] Method:", c.Request.Method)
					fmt.Println("[ERROR] Latency:", latency)
					fmt.Println("[ERROR] Error:", err)
					fmt.Println("[ERROR] Status:", c.Writer.Status())
					/*fmt.Println("[ERROR] Client:", clientId)*/
					fmt.Println("[ERROR] Header:", c.Request.Header)
					fmt.Println("[ERROR] Body:", bodyString)
					fmt.Println("[ERROR]===========")

				case 403:
					mes := fmt.Sprintf("%s", err)
					if mes == ErrorReachLimit {
						c.JSON(403, errorReachLimit(c, err))
					}
				}
			}
		} else {
			switch c.Writer.Status() {
			case 403:
				c.JSON(403, gin.H{
					"ERROR_CODE": errorPermissionDenied,
				})
			}
		}

		/*if latency >= 1000*time.Millisecond {
			clientId := utils.GetClient(c)
			fmt.Println("")
			fmt.Println("[WARNING]===========")
			fmt.Println("[WARNING] Link:", c.Request.RequestURI)
			fmt.Println("[WARNING] Method:", c.Request.Method)
			fmt.Println("[WARNING] Latency:", latency)
			fmt.Println("[WARNING] Client:", clientId)
			fmt.Println("[WARNING] Body:", bodyString)
			fmt.Println("[WARNING] Status:", c.Writer.Status())
			fmt.Println("[WARNING]===========")
		}
		*/
	}
}
