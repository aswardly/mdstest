package model

import (
	"mdstest/helper"

	"encoding/json"
	"time"
)

type User struct {
	UserId			string			`json:"user_id" gorm:"primary_key"`
	UserName		string			`json:"user_name"`
	UserPassword	string			`json:"user_password"`
	UserStatus		string			`json:"user_status"`
	LastUpdated		time.Time		`json:"last_updated"`
	UserSetting		[]UserSetting	`json:"user_setting"`
}

//MarshalJSON is a custom JSON marshaller for the user struct (satisfies interface json.Marshaler)
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User

	var lastUpdated string
	if u.LastUpdated.IsZero()== false {
		lastUpdated = u.LastUpdated.Format(helper.DatetimeFormat)
	}

	return json.Marshal(&struct{
		LastUpdated string `json:"lastUpdated"`
		*Alias
	}{
		//special treatment for fields containing datetime value (outputs to mysql datetime format/RFC3339). Do note that this limits time precision to seconds
		LastUpdated: lastUpdated,
		Alias	: (*Alias)(u),
	})
}

//UnmarshalJSON is a custom JSON unmarshaller for the user struct (satisfies interface json.Unmarshaler)
func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	user := &struct{
		LastUpdated string `json:"lastUpdated"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	var err error
	if err = json.Unmarshal(data, &user); err != nil {
		return err
	}
	//special treatment for fields containing datetime (parse from a mysql datetime format/RFC3339 to time.Time)
	var lastUpdated time.Time
	if user.LastUpdated != "" {
		lastUpdated, err = time.Parse(helper.DatetimeFormat, user.LastUpdated)
		if err != nil {
			return err
		}
	}
	u.LastUpdated = lastUpdated.In(helper.DefaultLocation)

	return nil
}
