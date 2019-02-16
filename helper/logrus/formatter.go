package logrus

import (
	"time"

	"github.com/sirupsen/logrus"
)

//TimezoneFormatter is a simple wrapper of logrus.Formatter for setting the timezone of the logged event
//example usage:
//  formatter := TimezoneFormatter{
//  	Formatter: &logrus.JSONFormatter{},
//  	Location: time.FixedZone("UTC+7", 7*60*60), //or use time.LoadLocation function
//  }
//  logrus.SetFormatter(formatter)
type TimezoneFormatter struct {
	Formatter logrus.Formatter 	//embedded logrus formatter
	Location *time.Location 	//location for setting the logged event's timezone
}

//Format changes the timezone of the event time then passes the event to the wrapped formatter
func (t TimezoneFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.In(t.Location)
	return t.Formatter.Format(e)
}