package application

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

//lookupLogrusLevelToGommonLogLevel perform lookup of a logrus.Level to gommon/log.Lvl
func lookupLogrusLevelToGommonLogLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.TraceLevel:
		return log.OFF
	case logrus.PanicLevel:
		return 6 //log.panicLevel
	case logrus.FatalLevel:
		return 7 //log.fatalLevel
	}
	//none matched, defaults to DEBUG
	return log.DEBUG
}

//lookupGommonLogLevelToLogrusLevel perform lookup of a gommon/log.Lvl to logrus.Level
func lookupGommonLogLevelToLogrusLevel(level log.Lvl) logrus.Level {
	switch level {
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	case log.OFF:
		return logrus.TraceLevel
	case 6: //log.panicLevel
		return logrus.PanicLevel
	case 7: //log.fatalLevel
		return logrus.FatalLevel
	}
	//none matched, defaults to DebugLevel
	return logrus.DebugLevel
}

//Logger is meant to be a generic logger used by echo.Context (adheres to echo.Logger interface)
//uses logrus for logging purpose
type Logger struct {
	logger *logrus.Logger	//embedded logrus.Logger instance
}

//NewLogger returns a new logger instance using the given logrus.Logger from param
func NewLogger(logger *logrus.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

//Output returns the output written to by the logger
func (l *Logger) Output() io.Writer {
	return l.logger.Out
}

//SetOutput sets the output to write log entries
func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

//Prefix is unimplemented
func (l *Logger) Prefix() string {
	//unimplemented
	return ""
}

//SetPrefix is unimplemented
func (l *Logger) SetPrefix(p string) {
	//unimplemented
}

//Level returns the current log level of the logger
func (l *Logger) Level() log.Lvl {
	return lookupLogrusLevelToGommonLogLevel(l.logger.GetLevel())
}

//SetLevel sets the log level of the logger
func (l *Logger) SetLevel(level log.Lvl) {
	l.logger.SetLevel(lookupGommonLogLevelToLogrusLevel(level))
}

//SetHeader is unimplemented
func (l *Logger) SetHeader(h string) {
	//unimplemented
}

//Print prints a log entry
func (l *Logger) Print(i ...interface{}) {
	l.logger.Print(i)
}

//Printf prints a formatted string as a log entry
func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}

//Printj is unimplemented
func (l *Logger) Printj(j log.JSON) {
	//unimplemented
}

//Debug prints a new log entry with log level Debug
func (l *Logger) Debug(i ...interface{}) {
	l.logger.Debug(i)
}

//Debugf prints a formatted string as a log entry with log level Debug
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args)
}

//Debugj is unimplemented
func (l *Logger) Debugj(j log.JSON) {
	//unimplemented
}

//Info prints a new log entry with log level Info
func (l *Logger) Info(i ...interface{}) {
	l.logger.Info(i)
}

//Infof prints a formatted string as a log entry with log level Info
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args)
}

//Infoj is unimplemented
func (l *Logger) Infoj(j log.JSON) {
	//unimplemented
}

//Warn prints a new log entry with log level Warn
func (l *Logger) Warn(i ...interface{}) {
	l.logger.Warn(i)
}

//Warnf prints a formatted string as a log entry with log level Warn
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args)
}

//Warnj is unimplemented
func (l *Logger) Warnj(j log.JSON) {
	//unimplemented
}

//Error prints a new log entry with log level Error
func (l *Logger) Error(i ...interface{}) {
	l.logger.Error(i)
}

//Errorf prints a formatted string as a log entry with log level Error
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args)
}

//Errorj is unimplemented
func (l *Logger) Errorj(j log.JSON) {
	//unimplemented
}

//Fatal prints a new log entry with log level Fatal
func (l *Logger) Fatal(i ...interface{}) {
	l.logger.Fatal(i)
}

//Fatalf prints a formatted string as a log entry with log level Fatal
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args)
}

//Fatalj is unimplemented
func (l *Logger) Fatalj(j log.JSON) {
	//unimplemented
}

//Panic prints a new log entry with log level Panic
func (l *Logger) Panic(i ...interface{}) {
	l.logger.Panic(i)
}

//Panicf prints a formatted string as a log entry with log level Panic
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args)
}

//Panicj is unimplemented
func (l *Logger) Panicj(j log.JSON) {
	//unimplemented
}