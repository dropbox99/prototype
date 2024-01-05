package log

import (
	"context"
	"runtime"

	"github.com/sirupsen/logrus"
)

type (
	logs struct{}

	ILogs interface {
		Trace(ctx context.Context, actName string, data interface{})
		Debug(ctx context.Context, actName string, data interface{})
		Info(ctx context.Context, actName string, data interface{})
		Warning(ctx context.Context, actName string, data interface{})
		Error(ctx context.Context, actName string, data interface{})
		Fatal(ctx context.Context, actName string, data interface{})

		Http(ctx context.Context, actName, url, method string, header, req, res interface{})
	}
)

func NewLog() ILogs {
	return &logs{}
}

func fieldLogs(ctx context.Context, data interface{}) logrus.Fields {
	pc, file, line, _ := runtime.Caller(1)
	return logrus.Fields{
		"service_name":    serviceName,
		"service_type":    serviceType,
		"service_code":    serviceCode,
		"version_release": versionRelease,
		"version_type":    versionType,
		"data":            data,
		"trace_id":        ctx.Value("trace-id"),
		"package":         runtime.FuncForPC(pc).Name(),
		"file":            file,
		"line":            line,
	}
}

func (lib *logs) Trace(ctx context.Context, actName string, data interface{}) {
	item := fieldLogs(ctx, data)
	if ActiveLogFile {
		log.WithFields(item).Trace(actName)
	}
	logrus.WithFields(item).Trace(actName)
}

func (lib *logs) Debug(ctx context.Context, actName string, data interface{}) {
	item := fieldLogs(ctx, data)
	if ActiveLogFile {
		log.WithFields(item).Debug(actName)
	}
	logrus.WithFields(item).Debug(actName)
}

func (lib *logs) Info(ctx context.Context, actName string, data interface{}) {
	item := fieldLogs(ctx, data)
	if ActiveLogFile {
		log.WithFields(item).Info(actName)
	}
	logrus.WithFields(item).Info(actName)
}

func (lib *logs) Warning(ctx context.Context, actName string, data interface{}) {
	item := fieldLogs(ctx, data)
	if ActiveLogFile {
		log.WithFields(item).Warning(actName)
	}
	logrus.WithFields(item).Warning(actName)
}

func (lib *logs) Error(ctx context.Context, actName string, data interface{}) {
	item := fieldLogs(ctx, data)
	if ActiveLogFile {
		log.WithFields(item).Error(actName)
	}
	logrus.WithFields(item).Error(actName)
}

func (lib *logs) Fatal(ctx context.Context, actName string, data interface{}) {
	item := fieldLogs(ctx, data)
	if ActiveLogFile {
		log.WithFields(item).Fatal(actName)
	}
	logrus.WithFields(item).Fatal(actName)
}

func (lib *logs) Http(ctx context.Context, actName, url, method string, header, req, res interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	item := logrus.Fields{
		"host_url":        url,
		"host_method":     method,
		"host_header":     header,
		"host_request":    req,
		"host_response":   res,
		"service_name":    serviceName,
		"service_type":    serviceType,
		"service_code":    serviceCode,
		"version_release": versionRelease,
		"version_type":    versionType,
		"trace_id":        ctx.Value("trace-id"),
		"package":         runtime.FuncForPC(pc).Name(),
		"file":            file,
		"line":            line,
	}

	if ActiveLogFile {
		log.WithFields(item).Info(actName)
	}
	logrus.WithFields(item).Info(actName)
}
