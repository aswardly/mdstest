package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"mdstest/helper"

	"encoding/json"
	"time"
)

type UserSetting struct {
	SettingId		int			`json:"setting_id" gorm:"primary_key:AUTO_INCREMENT"`
	UserId			string
	User			*User		`gorm:"ForeignKey:UserId;AssociationForeignKey:UserId;-"`
	SettingKey		string		`json:"setting_key"`
	SettingValue	string		`json:"setting_value"`
	LastUpdated		time.Time	`json:"last_updated"`
}

//MarshalJSON is a custom JSON marshaller for the user setting struct (satisfies interface json.Marshaler)
func (us *UserSetting) MarshalJSON() ([]byte, error) {
	type Alias UserSetting

	var lastUpdated string
	if us.LastUpdated.IsZero()== false {
		lastUpdated = us.LastUpdated.Format(helper.DatetimeFormat)
	}

	return json.Marshal(&struct{
		LastUpdated string `json:"lastUpdated"`
		*Alias
	}{
		//special treatment for fields containing datetime value (outputs to mysql datetime format/RFC3339). Do note that this limits time precision to seconds
		LastUpdated: lastUpdated,
		Alias	: (*Alias)(us),
	})
}

//UnmarshalJSON is a custom JSON unmarshaller for the user setting struct (satisfies interface json.Unmarshaler)
func (us *UserSetting) UnmarshalJSON(data []byte) error {
	type Alias UserSetting
	userSetting := &struct{
		LastUpdated string `json:"lastUpdated"`
		*Alias
	}{
		Alias: (*Alias)(us),
	}
	var err error
	if err = json.Unmarshal(data, &userSetting); err != nil {
		return err
	}
	//special treatment for fields containing datetime (parse from a mysql datetime format/RFC3339 to time.Time)
	var lastUpdated time.Time
	if userSetting.LastUpdated != "" {
		lastUpdated, err = time.Parse(helper.DatetimeFormat, userSetting.LastUpdated)
		if err != nil {
			return err
		}
	}
	us.LastUpdated = lastUpdated.In(helper.DefaultLocation)

	return nil
}

//validateAdd performs validation on the model for new user setting case
func (us *UserSetting) ValidateAdd(db *gorm.DB) error {
	//check if setting with same key exists for the given user
	var count int
	db.Table("user_settings").Where("setting_key = ? AND user_id = ?", us.SettingKey ,us.UserId).Count(&count)

	if db.Error != nil {
		return errors.Wrap(db.Error, "Query error")
	}

	if count >= 1 {
		return ValidationError{
			ErrorField: "UserId",
			ErrorMsg: "User setting already exists",
		}
	}
	return nil
}

//validateEdit performs validation on the model for new user setting case
func (us *UserSetting) ValidateEdit(db *gorm.DB) error {
	//check if setting with same key exists for the given user
	var count int
	db.Table("user_settings").Where("setting_key = ? AND user_id = ?", us.SettingKey ,us.UserId).Count(&count)

	if db.Error != nil {
		return errors.Wrap(db.Error, "Query error")
	}

	if count == 0 {
		return ValidationError{
			ErrorField: "UserId",
			ErrorMsg: "User setting doesn't exist",
		}
	}
	return nil
}