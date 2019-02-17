package handler

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"mdstest/helper"
	"mdstest/model"
)

//UserCreate is a handler for user creation route
func UserCreate(c echo.Context) error {

	//get POST request variables
	userId := c.FormValue("user_id")
	userName := c.FormValue("user_name")
	userPassword := c.FormValue("user_password")
	repeatPassword := c.FormValue("repeat_password")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Error(errors.New("Failed getting db instance from context"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting db instance from context")
	}

	//specific route validation, check password
	if userPassword != repeatPassword {
		c.Logger().Error("Repeat password does not match")
		return echo.NewHTTPError(http.StatusBadRequest, "Repeat password does not match")
	}

	userModel := &model.User{
		UserId: 		userId,
		UserName:		userName,
		UserStatus:		model.UserStatusActive,
		LastUpdated:	time.Now().In(helper.DefaultLocation),
	}

	err := userModel.ValidateAdd(db)
	if err != nil {
		c.Logger().Error(errors.Wrap(err,"Validation error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "validation error")
	}

	err = userModel.SetPassword(userPassword)
	if err != nil {
		c.Logger().Error(errors.Wrap(err,"Set password error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Set password error")
	}

	//save user
	err = db.Create(userModel).Error
	if err != nil {
		c.Logger().Error(errors.Wrap(err,"Save user error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Save user error")
	}

	c.Logger().Infof("user with id: %v created", userModel.UserId)
	return c.JSON(http.StatusOK, "User created")
}