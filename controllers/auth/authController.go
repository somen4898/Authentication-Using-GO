package auth

import (
	"github.com/gin-gonic/gin"
	"go-auth/core/results"
	"go-auth/helpers/jwtHelper"
	"go-auth/models/user"
	"net/http"
	"strings"
)

func VerifyLogin(c *gin.Context) {
	c.JSON(http.StatusOK, results.Success("Session is there!"))
	c.Abort()
}

func Login(c *gin.Context) {
	var _u user.UserForCredential
	if err := c.Bind(&_u); err != nil {
		c.JSON(http.StatusBadRequest, results.Error("Please enter credentials."))
		c.Abort()
		return
	}
	u, err := user.GetOneByEmail(_u.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, results.Error("Please check your credentials."))
		c.Abort()
		return
	}
	if strings.Compare(_u.Password, u.Password) != 0 {
		c.JSON(http.StatusBadRequest, results.Error("Please check your credentials."))
		c.Abort()
		return
	}
	loginUser(c, u)
}

func Register(c *gin.Context) {
	var newUser user.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, results.Error("Please enter a valid user."))
		c.Abort()
		return
	}

	_, err := newUser.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, results.Error("Please enter a valid user."))
		c.Abort()
		return
	}
	loginUser(c, newUser)
}

func loginUser(c *gin.Context, u user.User) {
	token, err := jwtHelper.Sign(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, results.Error("Login failed."))
		c.Abort()
		return
	}
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, results.Success("Login Successful."))
	c.Abort()
}
