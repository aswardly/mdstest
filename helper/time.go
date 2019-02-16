package helper

import "time"

//DefaultLocation is default location of application timezone
var DefaultLocation = time.FixedZone("UTC+7", 7 * 60 * 60)

//DatetimeFormat is the default date time format to use
var DatetimeFormat = "2006-01-02 15:04:05" 	  //mysql datetime format (RFC3339)