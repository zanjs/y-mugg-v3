package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
	"golang.org/x/crypto/bcrypt"
)

// UserController is
type UserController struct {
	Controller
}

// GetAll is get all users
func (ctl UserController) GetAll(c echo.Context) error {
	var (
		users []models.User
		err   error
	)
	users, err = models.GetUsers()
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}
	return ctl.ResponseSuccess(c, users)
}

// Get is get one user
func (ctl UserController) Get(c echo.Context) error {
	var (
		user       models.User
		pathparams models.PathParams
		err        error
	)
	pathparams = ctl.GetPathParam(c)
	user, err = models.GetUserById(pathparams.ID)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}
	return ctl.ResponseSuccess(c, user)
}

//Create is user
func (ctl UserController) Create(c echo.Context) error {
	user := new(models.User)

	user.Username = c.FormValue("username")
	user.Email = c.FormValue("email")
	user.Password = hashPassword(c.FormValue("password"))

	err := models.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, user)
}

//Update is update user
func (ctl UserController) Update(c echo.Context) error {
	// Parse the content
	user := new(models.User)

	user.Username = c.FormValue("username")
	user.Email = c.FormValue("email")
	user.Password = hashPassword(c.FormValue("password"))

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// update user data
	err = m.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

//Delete is user
func (ctl UserController) Delete(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteUser()
	return c.JSON(http.StatusNoContent, err)
}

func hashPassword(input string) string {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
