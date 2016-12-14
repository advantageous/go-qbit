package logging

import (
	"fmt"
	"log"
)

type BasicLogger struct {
	emergency        *log.Logger
	alert            *log.Logger
	critical         *log.Logger
	error            *log.Logger
	warning          *log.Logger
	notice           *log.Logger
	info             *log.Logger
	debug            *log.Logger
	logLevel         LogLevel
	panicOnEmergency bool
}

func (l *BasicLogger) Info(args ...interface{}) {
	if INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Println(args ...interface{}) {
	if INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Debug(args ...interface{}) {
	if DEBUG <= l.logLevel {
		l.debug.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Warn(args ...interface{}) {
	if WARNING <= l.logLevel {
		l.warning.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Error(args ...interface{}) {
	if ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Alert(args ...interface{}) {
	if ALERT <= l.logLevel {
		l.alert.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Emergency(args ...interface{}) {

	if l.panicOnEmergency {
		panic(fmt.Sprintln(args))
	} else if EMERGENCY <= l.logLevel {
		l.emergency.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Notice(args ...interface{}) {
	if NOTICE <= l.logLevel {
		l.notice.Output(2, fmt.Sprintln(args...))
	}
}

func (l *BasicLogger) Critical(args ...interface{}) {
	if CRITICAL <= l.logLevel {
		l.critical.Output(2, fmt.Sprintln(args...))
	}
}

//////
func (l *BasicLogger) Infof(format string, args ...interface{}) {
	if INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Printf(format string, args ...interface{}) {

	if INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Debugf(format string, args ...interface{}) {
	if DEBUG <= l.logLevel {
		l.debug.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Warnf(format string, args ...interface{}) {
	if WARNING <= l.logLevel {
		l.warning.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Errorf(format string, args ...interface{}) {
	if ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Alertf(format string, args ...interface{}) {
	if ALERT <= l.logLevel {
		l.alert.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Emergencyf(format string, args ...interface{}) {

	if l.panicOnEmergency {
		panic(fmt.Sprintf(format, args))
	} else if EMERGENCY <= l.logLevel {
		l.emergency.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Noticef(format string, args ...interface{}) {
	if NOTICE <= l.logLevel {
		l.notice.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) Criticalf(format string, args ...interface{}) {
	if CRITICAL <= l.logLevel {
		l.critical.Output(2, fmt.Sprintf(format, args...))
	}
}

func (l *BasicLogger) InfoError(message string, err error) {
	if INFO <= l.logLevel {
		l.info.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) PrintError(message string, err error) {
	if ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) DebugError(message string, err error) {
	if DEBUG <= l.logLevel {
		l.debug.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) WarnError(message string, err error) {
	if WARNING <= l.logLevel {
		l.warning.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) ErrorError(message string, err error) {
	if ERROR <= l.logLevel {
		l.error.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) AlertError(message string, err error) {
	if ALERT <= l.logLevel {
		l.alert.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) EmergencyError(message string, err error) {
	if l.panicOnEmergency {
		panic(fmt.Sprintf(" GOT ERROR %s %v", message, err))
	} else if EMERGENCY <= l.logLevel {
		l.emergency.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) NoticeError(message string, err error) {
	if NOTICE <= l.logLevel {
		l.notice.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}

func (l *BasicLogger) CriticalError(message string, err error) {
	if CRITICAL <= l.logLevel {
		l.critical.Output(2, fmt.Sprintf(" GOT ERROR %s %v", message, err))
	}
}
