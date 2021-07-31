package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

const (
	RequestInfoContextName = "REQUEST_INFO_CONTEXT_NAME"
)

func Init(isLocal bool) {
	if isLocal {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

type RequestInfo struct {
	RequestId     string
	BrowserUserId string
	RequestURL    string
	ClientIP      string

	RemoteAddr string
	UserAgent  string
	Referer    string
}

func Debugf(ctx context.Context, format string, a ...interface{}) {
	logrus.WithFields(GetRequestLogFields(ctx)).Debugf(format, a...)
}

func Infof(ctx context.Context, format string, a ...interface{}) {
	logrus.WithFields(GetRequestLogFields(ctx)).Infof(format, a...)
}

func Errorf(ctx context.Context, format string, a ...interface{}) {
	logrus.WithFields(GetRequestLogFields(ctx)).Errorf(format, a...)
}

func Fatalf(ctx context.Context, format string, a ...interface{}) {
	logrus.WithFields(GetRequestLogFields(ctx)).Fatalf(format, a...)
}

func GetRequestLogMap(ctx context.Context) map[string]string {
	value := ctx.Value(RequestInfoContextName)
	if value == nil {
		return map[string]string{}
	}
	if v, ok := value.(RequestInfo); ok {
		return map[string]string{
			"requestId":     v.RequestId,
			"browserUserId": v.BrowserUserId,
			"requestURL":    v.RequestURL,
			"remoteAddr":    v.RemoteAddr,
			"userAgent":     v.UserAgent,
			"referer":       v.Referer,
		}
	}
	panic("REQUEST_INFO_CONTEXT_NAME value must be RequestInfo")
}

func GetRequestLogFields(ctx context.Context) logrus.Fields {
	value := ctx.Value(RequestInfoContextName)
	if value == nil {
		return logrus.Fields{}
	}
	if v, ok := value.(RequestInfo); ok {
		return logrus.Fields{
			"requestId":     v.RequestId,
			"browserUserId": v.BrowserUserId,
			"requestURL":    v.RequestURL,
			"clientIP":      v.ClientIP,
			"remoteAddr":    v.RemoteAddr,
			"userAgent":     v.UserAgent,
			"referer":       v.Referer,
		}
	}
	panic("REQUEST_INFO_CONTEXT_NAME value must be RequestInfo")
}