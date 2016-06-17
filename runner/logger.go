package runner

import (
	"fmt"
	logPkg "log"
	"os"
)

type logFunc func(string, ...interface{})

var logger = logPkg.New(os.Stdout, "", 0)

func newLogFunc(prefix string) func(string, ...interface{}) {
	return func(format string, v ...interface{}) {
		if prefix != "app" {
			format = fmt.Sprintf("[eyeball] \033[35m%s\033[0m", format)
		} else {
			format = fmt.Sprintf("[app out] %s", format)
		}
		logger.Printf(format, v...)
	}
}

func fatal(err error) {
	logger.Fatal(err)
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	appLog(string(p))

	return len(p), nil
}
