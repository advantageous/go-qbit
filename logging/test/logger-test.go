package test

import (
	"fmt"
	lg "github.com/advantageous/go-qbit/logging"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type TestLogger struct {
	emergency *log.Logger
	alert     *log.Logger
	critical  *log.Logger
	error     *log.Logger
	warning   *log.Logger
	notice    *log.Logger
	info      *log.Logger
	debug     *log.Logger
	logLevel  lg.LogLevel
	t         *testing.T
}

func (l *TestLogger) Info(args ...interface{}) {
	if lg.INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintln(args...))
	}
}

func (l *TestLogger) Println(args ...interface{}) {
	if lg.INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintln(args...))
	}
}

func (l *TestLogger) Debug(args ...interface{}) {
	if lg.DEBUG <= l.logLevel {
		l.debug.Output(2, fmt.Sprintln(args...))
	}
}

func (l *TestLogger) Warn(args ...interface{}) {
	if lg.WARNING <= l.logLevel {
		l.warning.Output(2, fmt.Sprintln(args...))
	}
}

func (l *TestLogger) Error(args ...interface{}) {
	if lg.ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintln(args...))
	}
	l.t.Error(args)
}

func (l *TestLogger) Alert(args ...interface{}) {
	if lg.ALERT <= l.logLevel {
		l.alert.Output(2, fmt.Sprintln(args...))
	}
	l.t.Fatal(args)
}

func (l *TestLogger) Emergency(args ...interface{}) {
	if lg.EMERGENCY <= l.logLevel {
		l.emergency.Output(2, fmt.Sprintln(args...))
	}
	l.t.Fatal(args)
}

func (l *TestLogger) Notice(args ...interface{}) {
	if lg.NOTICE <= l.logLevel {
		l.notice.Output(2, fmt.Sprintln(args...))
	}
}

func (l *TestLogger) Critical(args ...interface{}) {
	if lg.CRITICAL <= l.logLevel {
		l.critical.Output(2, fmt.Sprintln(args...))
	}
	l.t.Fatal(args)
}

//////
func (l *TestLogger) Infof(format string, args ...interface{}) {
	if lg.INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *TestLogger) Printf(format string, args ...interface{}) {

	if lg.INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *TestLogger) Debugf(format string, args ...interface{}) {
	if lg.DEBUG <= l.logLevel {
		l.debug.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *TestLogger) Warnf(format string, args ...interface{}) {
	if lg.WARNING <= l.logLevel {
		l.warning.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *TestLogger) Errorf(format string, args ...interface{}) {
	if lg.ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintf(format, args...))
	}
	l.t.Errorf(format, args)
}

func (l *TestLogger) Alertf(format string, args ...interface{}) {
	if lg.ALERT <= l.logLevel {
		l.alert.Output(2, fmt.Sprintf(format, args...))
	}
	l.t.Fatalf(format, args)
}

func (l *TestLogger) Emergencyf(format string, args ...interface{}) {
	if lg.EMERGENCY <= l.logLevel {
		l.emergency.Output(2, fmt.Sprintf(format, args...))
	}
	l.t.Fatalf(format, args)
}

func (l *TestLogger) Noticef(format string, args ...interface{}) {
	if lg.NOTICE <= l.logLevel {
		l.notice.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *TestLogger) Criticalf(format string, args ...interface{}) {
	if lg.CRITICAL <= l.logLevel {
		l.critical.Output(2, fmt.Sprintf(format, args...))
	}
	l.t.Fatalf(format, args)
}

func (l *TestLogger) InfoError(message string, err error) {
	if lg.INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
}

func (l *TestLogger) PrintError(message string, err error) {
	if lg.ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
	l.t.Errorf(" GOT lg.ERROR %s %v", message, err)
}

func (l *TestLogger) DebugError(message string, err error) {
	if lg.DEBUG <= l.logLevel {
		l.debug.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
}

func (l *TestLogger) WarnError(message string, err error) {
	if lg.WARNING <= l.logLevel {
		l.warning.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
}

func (l *TestLogger) ErrorError(message string, err error) {
	if lg.ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
	l.t.Errorf(" GOT lg.ERROR %s %v", message, err)
}

func (l *TestLogger) AlertError(message string, err error) {
	if lg.ALERT <= l.logLevel {
		l.alert.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
	l.t.Fatalf(" GOT lg.ERROR %s %v", message, err)
}

func (l *TestLogger) EmergencyError(message string, err error) {
	if lg.EMERGENCY <= l.logLevel {
		l.emergency.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
	l.t.Fatalf(" GOT lg.ERROR %s %v", message, err)
}

func (l *TestLogger) NoticeError(message string, err error) {
	if lg.NOTICE <= l.logLevel {
		l.notice.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
}

func (l *TestLogger) CriticalError(message string, err error) {
	if lg.CRITICAL <= l.logLevel {
		l.critical.Output(2, fmt.Sprintf(" GOT lg.ERROR %s %v", message, err))
	}
	l.t.Fatalf(" GOT lg.ERROR %s %v", message, err)
}

func NewTestSimpleLogger(name string, t *testing.T) lg.Logger {
	return NewTestLogger(name, t, lg.INFO, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stdout, os.Stdout, ioutil.Discard)
}

func NewTestDebugLogger(name string, t *testing.T) lg.Logger {
	return NewTestLogger(name, t, lg.DEBUG, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stdout, os.Stdout, os.Stdout)
}

func NewTestLogger(name string, t *testing.T, logLevel lg.LogLevel, emergency io.Writer, alert io.Writer, critical io.Writer,
	error io.Writer, warning io.Writer, notice io.Writer,
	info io.Writer, debug io.Writer) *TestLogger {

	logger := TestLogger{}
	logger.logLevel = logLevel
	logger.t = t
	logger.alert = log.New(alert,
		"lg.ALERT    : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.emergency = log.New(emergency,
		"lg.EMERGENCY: ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.critical = log.New(critical,
		"lg.CRITICAL : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.notice = log.New(notice,
		"lg.NOTICE   : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.debug = log.New(debug,
		"lg.DEBUG    : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.info = log.New(info,
		"lg.INFO     : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.warning = log.New(warning,
		"WARN     : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.error = log.New(error,
		"lg.ERROR    : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return &logger

}
