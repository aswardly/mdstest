package model

import (
	"mdstest/helper"

	"encoding/json"
	"time"
)

type UserSetting struct {
	SettingId		int			`json:"setting_id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId			string		`gorm:"foreignKey:User"`
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
