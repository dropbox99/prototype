package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"prototype/lib/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging(log log.ILogs) gin.HandlerFunc {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {
		// set trace id
		traceID := c.GetHeader("Trace-ID")

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "trace-id", traceID)
		c.Request = c.Request.WithContext(ctx)

		// Request body
		var bodyBytes []byte

		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		//Request routing
		reqUri := c.Request.RequestURI

		// response body
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		//Process request
		c.Next()

		log.Http(
			ctx,
			"Result",
			reqUri,
			c.Request.Method,
			c.Request.Header,
			fmt.Sprintf("%v", string(bodyBytes)),
			fmt.Sprintf("%v", blw.body.String()),
		)
	}
}
