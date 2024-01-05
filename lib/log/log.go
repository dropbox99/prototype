package log

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"prototype/lib/env"
	"runtime"
	"strconv"
	"strings"
	"time"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	year  int
	month time.Month
	day   int

	log       = logrus.New()
	borderLog = logrus.New()

	formatDate    string
	ActiveLogFile bool
	ActiveLogHook bool
	file          *os.File

	serviceName    = env.String("MainSetup.ServiceName", "")
	serviceType    = env.String("MainSetup.ServiceType", "")
	serviceCode    = env.String("MainSetup.ServiceCode", "")
	versionRelease = env.String("Version.ReleaseVersion", "")
	versionType    = env.String("Version.VersionType", "")
)

func init() {
	var err error
	year, month, day = time.Now().Local().Date()
	logPrettyPrint, _ := strconv.ParseBool(env.String("Logging.logPrint.PrettyPrint", "false"))

	ActiveLogHook, _ = strconv.ParseBool(env.String("GrayLog.Active", "false"))
	if ActiveLogHook {
		address := env.String("GrayLog.Address", "") + ":" + env.String("GrayLog.Port", "")
		logHook := graylog.NewGraylogHook(address, map[string]interface{}{"service": serviceName})
		if ActiveLogFile {
			log.AddHook(logHook)
		} else {
			logrus.AddHook(logHook)
			borderLog.AddHook(logHook)
		}
	}
	formatDate = fmt.Sprintf("%s_%d_%s_%d.log", env.String("Logging.logFile.FileName", "log"), day, month, year)
	ActiveLogFile, _ = strconv.ParseBool(env.String("Logging.logFile.Active", "false"))

	if ActiveLogFile {
		log.SetOutput(os.Stderr)
		logPrettyPrintFile, _ := strconv.ParseBool(env.String("Logging.logFile.PrettyPrint", "false"))
		file, err = os.OpenFile(formatDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			logrus.Info("Failed create file log print file")
		}

		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: env.String("Logging.logFile.TimeFormat", "01-02-2006 15:04:05"),
			PrettyPrint:     logPrettyPrintFile,
		})

	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: env.String("Logging.logPrint.TimeFormat", "01-02-2006 15:04:05"),
		PrettyPrint:     logPrettyPrint,
	})

	borderLog.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
	})

}

type (
	logging struct {
		TraceID string

		SkipCaller int
	}

	ILogging interface {
		Start(c *gin.Context)
		End(c *gin.Context)
		SetTraceID(traceID string)
		Trace(actName string, data interface{})
		Debug(actName string, data interface{})
		Info(actName string, data interface{})
		Warning(actName string, data interface{})
		Error(actName string, data interface{})
		Fatal(actName string, data interface{})
		Http(actName string, url, method string, header, req, res interface{})
	}

	bodyLogWriter struct {
		gin.ResponseWriter
		body *bytes.Buffer
	}
)

func NewLoggingWithTraceID(traceID string, SkipCaller int) ILogging {
	return &logging{
		TraceID:    traceID,
		SkipCaller: SkipCaller,
	}
}

func NewLogging() ILogging {
	return &logging{}
}

func fieldLog(trace_id string, SkipCaller int, data interface{}) logrus.Fields {
	pc, file, line, _ := runtime.Caller(SkipCaller)
	return logrus.Fields{
		"service_name":    serviceName,
		"service_type":    serviceType,
		"service_code":    serviceCode,
		"version_release": versionRelease,
		"version_type":    versionType,
		"data":            data,
		"trace_id":        trace_id,
		"package":         runtime.FuncForPC(pc).Name(),
		"file":            file,
		"line":            line,
	}
}

func (lib *logging) SetTraceID(traceID string) {
	lib.TraceID = traceID
	lib.SkipCaller = 2
}

func (lib *logging) Trace(actName string, data interface{}) {
	item := fieldLog(lib.TraceID, lib.SkipCaller, data)
	if ActiveLogFile {
		log.WithFields(item).Trace(actName)
	}
	logrus.WithFields(item).Trace(actName)
}

func (lib *logging) Debug(actName string, data interface{}) {
	item := fieldLog(lib.TraceID, lib.SkipCaller, data)
	if ActiveLogFile {
		log.WithFields(item).Debug(actName)
	}
	logrus.WithFields(item).Debug(actName)
}

func (lib *logging) Info(actName string, data interface{}) {
	item := fieldLog(lib.TraceID, lib.SkipCaller, data)
	if ActiveLogFile {
		log.WithFields(item).Info(actName)
	}
	logrus.WithFields(item).Info(actName)
}
func (lib *logging) Warning(actName string, data interface{}) {
	item := fieldLog(lib.TraceID, lib.SkipCaller, data)
	if ActiveLogFile {
		log.WithFields(item).Warning(actName)
	}
	logrus.WithFields(item).Warning(actName)
}
func (lib *logging) Error(actName string, data interface{}) {
	item := fieldLog(lib.TraceID, lib.SkipCaller, data)
	if ActiveLogFile {
		log.WithFields(item).Error(actName)
	}
	logrus.WithFields(item).Error(actName)
}
func (lib *logging) Fatal(actName string, data interface{}) {
	item := fieldLog(lib.TraceID, lib.SkipCaller, data)
	if ActiveLogFile {
		log.WithFields(item).Fatal(actName)
	}
	logrus.WithFields(item).Fatal(actName)
}
func (lib *logging) Http(actName, url, method string, header, req, res interface{}) {
	pc, file, line, _ := runtime.Caller(lib.SkipCaller - 1)
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
		"trace_id":        lib.TraceID,
		"package":         runtime.FuncForPC(pc).Name(),
		"file":            file,
		"line":            line,
	}

	if ActiveLogFile {
		log.WithFields(item).Info(actName)
	}
	logrus.WithFields(item).Info(actName)
}

func (lib *logging) Start(c *gin.Context) {
	pc, file, line, _ := runtime.Caller(lib.SkipCaller - 1)

	//Request method
	reqMethod := c.Request.Method

	// set trace id
	traceID := c.GetHeader("Trace-ID")

	// Request body
	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	//Request routing
	reqUri := c.Request.RequestURI

	// request IP
	clientIP := c.ClientIP()

	item := logrus.Fields{
		"uri":             reqUri,
		"method":          reqMethod,
		"ip":              clientIP,
		"header":          c.Request.Header,
		"body":            fmt.Sprintf("%v", string(bodyBytes)),
		"service_name":    serviceName,
		"service_type":    serviceType,
		"service_code":    serviceCode,
		"version_release": versionRelease,
		"version_type":    versionType,
		"trace_id":        traceID,
		"package":         runtime.FuncForPC(pc).Name(),
		"file":            file,
		"line":            line,
	}

	lib.TraceID = traceID
	lib.SkipCaller = 2

	if ActiveLogFile {
		log.WithFields(item).Info("START")
	}
	logrus.WithFields(item).Info("START")
}

func (lib *logging) End(c *gin.Context) {
	pc, file, line, _ := runtime.Caller(lib.SkipCaller - 1)

	//Request method
	reqMethod := c.Request.Method

	// set trace id
	traceID := c.GetHeader("Trace-ID")

	// Request body
	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	//Request routing
	reqUri := c.Request.RequestURI

	// request IP
	clientIP := c.ClientIP()

	// status code
	statusCode := c.Writer.Status()

	// response body
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	item := logrus.Fields{
		"uri":             reqUri,
		"method":          reqMethod,
		"ip":              clientIP,
		"status_code":     statusCode,
		"response":        strings.Trim(blw.body.String(), "\n"),
		"service_name":    serviceName,
		"service_type":    serviceType,
		"service_code":    serviceCode,
		"version_release": versionRelease,
		"version_type":    versionType,
		"trace_id":        traceID,
		"package":         runtime.FuncForPC(pc).Name(),
		"file":            file,
		"line":            line,
	}

	if ActiveLogFile {
		log.WithFields(item).Info("END")
	}
	logrus.WithFields(item).Info("END")
}
