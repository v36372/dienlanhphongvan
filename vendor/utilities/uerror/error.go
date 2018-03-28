//https://github.com/dropbox/godropbox/blob/master/errors/errors.go
package uerror

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v3"
)

const (
	ServiceErrorNotFoundStr   = "service_error_not_found"
	ServiceErrorBadRequestStr = "service_error_bad_request"
)

var (
	ServiceErrorBadRequest = fmt.Errorf(ServiceErrorBadRequestStr)
	ServiceErrorNotFound   = fmt.Errorf(ServiceErrorNotFoundStr)
)

func IsQueryOneError(err error, c *gin.Context) bool {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(404)
			return true
		}
		c.AbortWithError(500, err)
		return true
	}
	return false
}

func IsStatusFail(status int, err error, c *gin.Context) bool {
	if status != 200 && status != 201 {
		c.AbortWithError(status, err)
		return true
	}
	return false
}

func IsLoggedIn(isLoggedIn bool, c *gin.Context) bool {
	if !isLoggedIn {
		c.AbortWithStatus(401)
		return false
	}
	return true
}

func IsNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound || err == redis.Nil || err == ServiceErrorNotFound {
		return true
	}
	return false
}

func GenStatusError(err error) (int, error) {
	if err == nil {
		return 200, nil
	}
	if IsNotFound(err) {
		return 404, StackTrace(err)
	}
	if IsErrorBadRequest(err) {
		return 400, StackTrace(err)
	}
	return 500, StackTrace(err)
}

func IsErrorBadRequest(err error) bool {
	if err == ServiceErrorBadRequest {
		return true
	}
	return false
}

func NewErrorField(key, messesage string) error {
	json := fmt.Sprintf("[{\"key\":\"%s\",\"message\":\"%s\"}]", key, messesage)
	return errors.New(json)
}

func StackTraceStr(err error) string {
	return fmt.Sprintf("%v\n%v", err, stackTrace(2))
}

func StackTraceWithData(data interface{}, err error) error {
	if err != nil {
		errMess := fmt.Sprint(err)
		data := fmt.Sprint(data)
		return errors.New("[ERROR]" + errMess + "\n" + "[DATA]" + data + "\n" + stackTrace(2))
	}
	return err
}

func StackTrace(err error) error {
	if err != nil {
		errMess := fmt.Sprint(err)
		return errors.New(errMess + "\n" + stackTrace(2))
	}
	return err
}

// Returns a copy of the error with the stack trace field populated and any
// other shared initialization; skips 'skip' levels of the stack trace.
//
// NOTE: This panics on any error.
// func StackTrace() (current, context string) {
func stackTrace(skip int) (current string) {
	// grow buf until it's large enough to store entire stack trace
	buf := make([]byte, 128)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, len(buf)*2)
	}

	// Returns the index of the first occurrence of '\n' in the buffer 'b'
	// starting with index 'start'.
	//
	// In case no occurrence of '\n' is found, it returns len(b). This
	// simplifies the logic on the calling sites.
	indexNewline := func(b []byte, start int) int {
		if start >= len(b) {
			return len(b)
		}
		searchBuf := b[start:]
		index := bytes.IndexByte(searchBuf, '\n')
		if index == -1 {
			return len(b)
		}
		return (start + index)
	}

	// Strip initial levels of stack trace, but keep header line that
	// identifies the current goroutine.
	var strippedBuf bytes.Buffer
	index := indexNewline(buf, 0)
	if index != -1 {
		strippedBuf.Write(buf[:index])
	}

	// Skip lines.
	for i := 0; i < skip; i++ {
		index = indexNewline(buf, index+1)
		index = indexNewline(buf, index+1)
	}

	isDone := false
	startIndex := index
	lastIndex := index
	for !isDone {
		index = indexNewline(buf, index+1)
		if (index - lastIndex) <= 1 {
			isDone = true
		} else {
			lastIndex = index
		}
	}
	strippedBuf.Write(buf[startIndex:index])
	// return strippedBuf.String(), string(buf[index:])
	return strippedBuf.String()
}

func IsErrorNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound || err == redis.Nil {
		return true
	}
	return false
}
