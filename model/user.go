package model

import (
	"encoding/json"
	"fmt"
	"time"

	"mdstest/helper"

	"github.com/jinzhu/gorm"
)

const StatusActive = "A"
const StatusInactive = "I"
const StatusDeleted = "D"

var StatusMap =  map[string]string {
	"A" : "Active",
	"I" : "Inactive",
	"D" : "Deleted",
}

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


//validators

//validateStatus performs validation on field UserStatus
func (u *User) validateStatus() *ValidationError {
	if _, ok := StatusMap[u.UserStatus]; ok == false {
		return &ValidationError{
			ErrorField: "UserStatus",
			ErrorMsg: fmt.Sprintf("Invalid status value: %v", u.UserStatus),
		}
	}
	return nil
}

//validateAdd performs validation on the model for new user case
func (u *User) validateAdd(db gorm.DB) *ValidationError {
	var exist int
	//check if user with same id exists
	db.Where("user_id = ?", u.UserId).Count(&exist)

	if exist >= 1 {
		return &ValidationError{
			ErrorField: "UserId",
			ErrorMsg: "User already exists",
		}
	}

	return u.validateStatus()
}

//validateEdit performs validation on the model for edit user case
func (u *User) validateEdit(db gorm.DB) *ValidationError {
	var exist int
	//check if user with same id exists
	db.Where("user_id = ?", u.UserId).Count(&exist)

	if exist == 0 {
		return &ValidationError{
			ErrorField: "UserId",
			ErrorMsg: "User doesn't exist",
		}
	}

	return u.validateStatus()
}