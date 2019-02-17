package model

import (
	"encoding/json"
	"fmt"
	"time"

	"mdstest/helper"

	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const UserStatusActive = "A"
const UserStatusInactive = "I"
const UserStatusDeleted = "D"

var UserStatusMap =  map[string]string {
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

//validateStatus performs validation on field UserStatus
func (u *User) ValidateStatus() error {
	if _, ok := UserStatusMap[u.UserStatus]; ok == false {
		return ValidationError{
			ErrorField: "UserStatus",
			ErrorMsg: fmt.Sprintf("Invalid status value: %v", u.UserStatus),
		}
	}
	return nil
}

//validateAdd performs validation on the model for new user case
func (u *User) ValidateAdd(db *gorm.DB) error {
	//check if user with same id exists
	var count int
	db.Table("users").Where("user_id = ?", u.UserId).Count(&count)

	if count >= 1 {
		return ValidationError{
			ErrorField: "UserId",
			ErrorMsg: "User already exists",
		}
	}

	return u.ValidateStatus()
}

//validateEdit performs validation on the model for edit user case
func (u *User) ValidateEdit(db *gorm.DB) error {
	//check if user with same id exists
	var count int
	db.Table("users").Where("user_id = ?", u.UserId).Count(&count)

	if count == 0 {
		return ValidationError{
			ErrorField: "UserId",
			ErrorMsg: "User doesn't exist",
		}
	}

	return u.ValidateStatus()
}

//SetPassword sets bcrypt hash for field password (default cost = 10)
func (u *User) SetPassword(plaintext string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "Error generating password hash")
	}
	u.UserPassword = string(hashedPassword)
	return nil
}

func (*User) TableName() string {
	return "users"
}