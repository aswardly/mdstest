package middleware

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
type ContextLogger struct {
	logger *logrus.Logger	//embedded logrus.Logger instance
}

//NewLogger returns a new logger instance using the given logrus.Logger from param
func NewContextLogger(logger *logrus.Logger) *ContextLogger {
	return &ContextLogger{
		logger: logger,
	}
}

//Output returns the output written to by the logger
func (cl *ContextLogger) Output() io.Writer {
	return cl.logger.Out
}

//SetOutput sets the output to write log entries
func (cl *ContextLogger) SetOutput(w io.Writer) {
	cl.logger.SetOutput(w)
}

//Prefix is unimplemented
func (cl *ContextLogger) Prefix() string {
	//unimplemented
	return ""
}

//SetPrefix is unimplemented
func (cl *ContextLogger) SetPrefix(p string) {
	//unimplemented
}

//Level returns the current log level of the logger
func (cl *ContextLogger) Level() log.Lvl {
	return lookupLogrusLevelToGommonLogLevel(cl.logger.GetLevel())
}

//SetLevel sets the log level of the logger
func (cl *ContextLogger) SetLevel(level log.Lvl) {
	cl.logger.SetLevel(lookupGommonLogLevelToLogrusLevel(level))
}

//SetHeader is unimplemented
func (cl *ContextLogger) SetHeader(h string) {
	//unimplemented
}

//Print prints a log entry
func (cl *ContextLogger) Print(i ...interface{}) {
	cl.logger.Print(i)
}

//Printf prints a formatted string as a log entry
func (cl *ContextLogger) Printf(format string, args ...interface{}) {
	cl.logger.Printf(format, args)
}

//Printj is unimplemented
func (cl *ContextLogger) Printj(j log.JSON) {
	//unimplemented
}

//Debug prints a new log entry with log level Debug
func (cl *ContextLogger) Debug(i ...interface{}) {
	cl.logger.Debug(i)
}

//Debugf prints a formatted string as a log entry with log level Debug
func (cl *ContextLogger) Debugf(format string, args ...interface{}) {
	cl.logger.Debugf(format, args)
}

//Debugj is unimplemented
func (cl *ContextLogger) Debugj(j log.JSON) {
	//unimplemented
}

//Info prints a new log entry with log level Info
func (cl *ContextLogger) Info(i ...interface{}) {
	cl.logger.Info(i)
}

//Infof prints a formatted string as a log entry with log level Info
func (cl *ContextLogger) Infof(format string, args ...interface{}) {
	cl.logger.Infof(format, args)
}

//Infoj is unimplemented
func (cl *ContextLogger) Infoj(j log.JSON) {
	//unimplemented
}

//Warn prints a new log entry with log level Warn
func (cl *ContextLogger) Warn(i ...interface{}) {
	cl.logger.Warn(i)
}

//Warnf prints a formatted string as a log entry with log level Warn
func (cl *ContextLogger) Warnf(format string, args ...interface{}) {
	cl.logger.Warnf(format, args)
}

//Warnj is unimplemented
func (cl *ContextLogger) Warnj(j log.JSON) {
	//unimplemented
}

//Error prints a new log entry with log level Error
func (cl *ContextLogger) Error(i ...interface{}) {
	cl.logger.Error(i)
}

//Errorf prints a formatted string as a log entry with log level Error
func (cl *ContextLogger) Errorf(format string, args ...interface{}) {
	cl.logger.Errorf(format, args)
}

//Errorj is unimplemented
func (cl *ContextLogger) Errorj(j log.JSON) {
	//unimplemented
}

//Fatal prints a new log entry with log level Fatal
func (cl *ContextLogger) Fatal(i ...interface{}) {
	cl.logger.Fatal(i)
}

//Fatalf prints a formatted string as a log entry with log level Fatal
func (cl *ContextLogger) Fatalf(format string, args ...interface{}) {
	cl.logger.Fatalf(format, args)
}

//Fatalj is unimplemented
func (cl *ContextLogger) Fatalj(j log.JSON) {
	//unimplemented
}

//Panic prints a new log entry with log level Panic
func (cl *ContextLogger) Panic(i ...interface{}) {
	cl.logger.Panic(i)
}

//Panicf prints a formatted string as a log entry with log level Panic
func (cl *ContextLogger) Panicf(format string, args ...interface{}) {
	cl.logger.Panicf(format, args)
}

//Panicj is unimplemented
func (cl *ContextLogger) Panicj(j log.JSON) {
	//unimplemented
}