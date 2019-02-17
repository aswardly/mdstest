package handler

import (
	"net/http"
	"time"

	"mdstest/helper"
	"mdstest/model"
	"mdstest/response"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

//UserSettingCreate is a handler for user setting creation route
func UserSettingCreate(c echo.Context) error {

	//get POST request variables
	userId := c.FormValue("user_id")
	settingKey := c.FormValue("setting_key")
	settingValue := c.FormValue("setting_value")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting db instance from context")
	}

	userSettingModel := &model.UserSetting{
		UserId:			userId,
		SettingKey:		settingKey,
		SettingValue:	settingValue,
		LastUpdated:	time.Now().In(helper.DefaultLocation),
	}

	err := userSettingModel.ValidateAdd(db)
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Validation error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "validation error")
	}

	//save user setting
	err = db.Create(userSettingModel).Error
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Save user setting error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Save user setting error")
	}

	c.Logger().Infof("user setting key: %v created for user id: %v", userSettingModel.SettingKey, userSettingModel.UserId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "User setting created",
	})
}

//UserSettingUpdate is a handler for user setting update route
func UserSettingUpdate(c echo.Context) error {

	//get POST request variables
	userId := c.FormValue("user_id")
	settingKey := c.FormValue("setting_key")
	settingValue := c.FormValue("setting_value")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting db instance from context")
	}

	userSettingModel := &model.UserSetting{
		UserId:			userId,
		SettingKey:		settingKey,
		SettingValue:	settingValue,
		LastUpdated:	time.Now().In(helper.DefaultLocation),
	}

	err := userSettingModel.ValidateEdit(db)
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Validation error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "validation error")
	}

	//save user setting
	err = db.Table("user_settings").Where("setting_key = ? AND user_id = ?", userSettingModel.SettingKey, userSettingModel.UserId).Update(userSettingModel).Error
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Update user setting error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Update user setting error")
	}

	c.Logger().Infof("user setting key: %v updated for user id: %v", userSettingModel.SettingKey, userSettingModel.UserId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "User setting updated",
	})
}

//UserSettingDelete is a handler for user setting deletion route
func UserSettingDelete(c echo.Context) error {

	//get param
	userId := c.Param("user_id")
	settingKey := c.Param("setting_key")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting db instance from context")
	}

	userSettingModel := &model.UserSetting{
		UserId:			userId,
		SettingKey:		settingKey,
	}

	//delete user setting
	rowsAffected , err := db.Table("user_settings").Where("setting_key = ? AND user_id = ?", settingKey, userId).Delete(userSettingModel).RowsAffected, db.Error
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Delete user setting error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Delete user setting error")
	}
	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, response.GenericResponse{
			ResponseCode: response.ResponseCodeSuccess,
			ResponseMessage: "User setting doesn't exist",
		})
	}

	c.Logger().Infof("user setting key: %v deleted for user id: %v", settingKey, userId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "User setting deleted",
	})
}