package log

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.TextFormatter{
	//	DisableColors:    true,
	//	DisableTimestamp: true,
	//	TimestampFormat:  "2006-01-02 15:03:04",
	//	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
	//		fileName := path.Base(frame.File)
	//		return frame.Function, fileName
	//	},
	//})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&MyFormatter{})
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// MyFormatter 自定义 formatter
type MyFormatter struct {
}

// Format implement the Formatter interface
func (mf *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = gray
	case log.WarnLevel:
		levelColor = yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = red
	case log.InfoLevel:
		levelColor = blue
	default:
		levelColor = blue
	}
	lino := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	b.WriteString(fmt.Sprintf("[%s] [%-15.15s] [\u001B[%dm%.4s\u001B[0m] - %s\n",
		entry.Time.Format("15:03:04.000"),
		lino,
		levelColor,
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	))
	return b.Bytes(), nil
}
