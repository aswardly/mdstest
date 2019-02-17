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
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
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
		c.Logger().Errorf("%+v", errors.Wrap(err,"Validation error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "validation error")
	}

	err = userModel.SetPassword(userPassword)
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Set password error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Set password error")
	}

	//save user
	err = db.Create(userModel).Error
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Save user error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Save user error")
	}

	c.Logger().Infof("user with id: %v created", userModel.UserId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "User created",
	})
}

//UserUpdate is a handler for user update route
func UserUpdate(c echo.Context) error {

	//get PUT request variables
	userId := c.FormValue("user_id")
	userName := c.FormValue("user_name")
	userStatus := c.FormValue("user_status")
	userPassword := c.FormValue("user_password")
	repeatPassword := c.FormValue("repeat_password")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
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
		UserStatus:		userStatus,
		LastUpdated:	time.Now().In(helper.DefaultLocation),
	}

	err := userModel.ValidateEdit(db)
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Validation error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "validation error")
	}

	err = userModel.SetPassword(userPassword)
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Set password error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Set password error")
	}

	//update user
	err = db.Save(userModel).Error
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Update user error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Update user error")
	}

	c.Logger().Infof("user with id: %v updated", userModel.UserId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "User updated",
	})
}

//UserDelete is a handler for user delete route
func UserDelete(c echo.Context) error {
	//get param
	userId := c.Param("user_id")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting db instance from context")
	}

	userModel := &model.User{
		UserId: userId,
	}

	//get user data
	err := db.Table("users").Where("user_id = ?", userId).First(userModel).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Logger().Infof("User with id: %v not found", userId)
			return c.JSON(http.StatusOK, response.GenericResponse{
				ResponseCode: response.ResponseCodeError,
				ResponseMessage: "User not found",
			})
		}
		c.Logger().Errorf("%+v", errors.Wrap(db.Error, "Query error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user data")
	}

	//Soft delete, update status only
	userModel.UserStatus = model.UserStatusDeleted
	userModel.LastUpdated = time.Now().In(helper.DefaultLocation)

	//update user data
	err = db.Save(userModel).Error
	if err != nil {
		c.Logger().Errorf("%+v", errors.Wrap(err,"Delete user error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Delete user error")
	}

	c.Logger().Infof("user with id: %v deleted", userModel.UserId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "User deleted",
	})
}

//UserQuery is a handler for fetching user data
func UserQuery(c echo.Context) error {
	//get param
	userId := c.Param("user_id")

	//get db instance from context
	db, ok := c.Get("GORM").(*gorm.DB)
	if false == ok {
		c.Logger().Errorf("%+v", errors.New("Failed getting db instance from context"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed getting db instance from context")
	}

	userModel := new(model.User)

	//get user data
	err := db.Preload("UserSettings").Where("user_id = ?", userId).First(userModel).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Logger().Infof("User with id: %v not found", userId)
			return c.JSON(http.StatusOK, response.GenericResponse{
				ResponseCode: response.ResponseCodeError,
				ResponseMessage: "User not found",
			})
		}
		c.Logger().Errorf("%+v", errors.Wrap(db.Error, "Query error"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user data")
	}

	c.Logger().Infof("user with id: %v loaded", userModel.UserId)
	return c.JSON(http.StatusOK, response.GenericResponse{
		ResponseCode: response.ResponseCodeSuccess,
		ResponseMessage: "",
		Data: userModel,
	})
}
