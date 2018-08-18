package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initial(r *gin.Engine) {
	r.GET("/login", LoginPage)
	r.POST("/login", Login)
	r.GET("/register", RegisterPage)
	r.POST("/register", Register)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": "COBA"})
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	var err error
	c.ShouldBind(&loginRequest)

	err = loginRequest.ValidateFields()
	if err != nil {
		fmt.Println(err)
		return
	}

	err, loginResponse := loginRequest.CheckLogin()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "COBA"})
}

func Register(c *gin.Context) {

	var registerRequest RegisterRequest
	var err error
	c.ShouldBind(&registerRequest)

	err = registerRequest.ValidateFields()
	if err != nil {
		fmt.Println(err)
		return
	}

	err, RegisterResponse := registerRequest.InsertRegister()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, RegisterResponse)
}
