package consumer

import (
	"time"
	"utilities/ulog"

	"github.com/streadway/amqp"
)

type (
	subscribeEndpoint   func(message *amqp.Delivery) (resp interface{}, erro error)
	subscribeMiddleware func(next subscribeEndpoint) subscribeEndpoint
)

func GetConsumerEndpoint(logMessagePattern, logType string, logger *ulog.Ulogger, handle subscribeEndpoint) subscribeEndpoint {
	return use(
		withMiddlewares(
			infoLogger(logMessagePattern, logType, logger),
			errorLogger(logMessagePattern, logType, logger),
		),
		handle,
	)

}

func withMiddlewares(middlewares ...subscribeMiddleware) []subscribeMiddleware {
	if len(middlewares) == 0 {
		return nil
	}
	rets := make([]subscribeMiddleware, len(middlewares))
	for index, middleware := range middlewares {
		rets[index] = middleware
	}
	return rets
}

func use(middlewares []subscribeMiddleware, endpoint subscribeEndpoint) subscribeEndpoint {
	if len(middlewares) == 0 {
		return endpoint
	}
	ret := endpoint
	for i := len(middlewares) - 1; i >= 0; i -= 1 {
		ret = middlewares[i](ret)
	}
	return ret
}

func infoLogger(logMessagePattern, logType string, logger *ulog.Ulogger) subscribeMiddleware {
	return func(next subscribeEndpoint) subscribeEndpoint {
		return func(message *amqp.Delivery) (resp interface{}, err error) {
			start := time.Now()
			defer func() {
				latency := time.Now().Sub(start)
				logger.LogInfo(logMessagePattern, ulog.Fields{
					"@type":            logType,
					"latency":          latency.String(),
					"latency_microsec": latency.Nanoseconds() / 1000,
					"rabbit_msg":       resp,
				})
			}()
			resp, err = next(message)
			return
		}
	}
}

func errorLogger(logMessagePattern, logType string, logger *ulog.Ulogger) subscribeMiddleware {
	return func(next subscribeEndpoint) subscribeEndpoint {
		return func(message *amqp.Delivery) (resp interface{}, err error) {
			start := time.Now()
			defer func() {
				if err != nil {
					latency := time.Now().Sub(start)
					logger.LogError(logMessagePattern, ulog.Fields{
						"@type":            logType,
						"err":              err.Error(),
						"latency":          latency.String(),
						"latency_microsec": latency.Nanoseconds() / 1000,
						"rabbit_msg":       resp,
					})
				}
			}()
			resp, err = next(message)
			return
		}
	}
}
