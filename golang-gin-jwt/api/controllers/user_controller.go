package controllers

import (
	"fmt"
	"net/http"

	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/helpers"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/response"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/dtos"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/infrastructure"
	"github.com/gin-gonic/gin"
)

type UserControllers struct {
	env infrastructure.Env
}
type DummyUser struct {
	Id       int
	Name     string
	Email    string
	Address  string
	Verified bool
	Role     string
	Password string
}

func NewUserControllers(env infrastructure.Env) UserControllers {
	return UserControllers{
		env: env,
	}
}
func (c UserControllers) SignUp(ctx *gin.Context) {
	var payload *dtos.SignUpInput
	payloadError := ctx.ShouldBind(payload)
	if payloadError != nil {
		response := response.BuildErrorResponse("Failed to process request", payloadError.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if payload.Password != payload.PasswordConfirm {
		response := response.BuildErrorResponse("Passwords do not match", "password mismatch", response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	response := response.BuildErrorResponse("fail to hash password", "failt", response.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
	// 	return
	// }
	// newUser := models.User{
	// 	Name:     payload.Name,
	// 	Email:    strings.ToLower(payload.Email),
	// 	Password: string(hashedPassword),
	// 	Photo:    &payload.Photo,
	// }
	//todo
	// check for email
	//create new user
	// result := Create(&newUser)
	// if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
	// 	response := response.BuildErrorResponse("User with that email already exists", result.Error{}, response.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusConflict, response)
	// 	return
	// } else if result.Error != nil {
	// 	response := response.BuildErrorResponse("Something bad happened", result.Error{}, response.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusBadGateway, response)
	// 	return
	// }
	response := response.BuildResponse(true, "OK", response.EmptyObj{})
	ctx.JSON(http.StatusCreated, response)
	return
}

func (c UserControllers) SignIn(ctx *gin.Context) {
	// var payload *dtos.SignInInput
	// payloadError := ctx.ShouldBind(payload)
	// if payloadError != nil {
	// 	response := response.BuildErrorResponse("Failed to process request", payloadError.Error(), response.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	// 	return
	// }
	// todo
	// 1. import user modal
	// var user models.User
	user := DummyUser{Id: 1, Name: "user", Email: "email@gmail.com", Address: "address", Password: "12345"}
	//2. get user by email from database(services)
	//3. compare password
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	// if err != nil {
	// 	response := response.BuildErrorResponse("Invalid email or Password", "fail", response.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	// 	return
	// }
	tokenString, err := helpers.GenerateToken(user.Id, user.Name, user.Email, user.Role, c.env.JWT_SECRET, c.env.JWT_EXPIRED_IN)
	if err != nil {
		response := response.BuildErrorResponse("enerating JWT Token failed", err.Error(), response.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return

	}
	cookie, err := ctx.Cookie("auth")

	if err != nil {
		cookie = "NotSet"
		ctx.SetCookie("auth", tokenString, 3600, "", "", false, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
	response := response.BuildResponse(true, "OK", tokenString)
	ctx.JSON(http.StatusCreated, response)
}

func (c UserControllers) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "users",
	})

}

func (c UserControllers) GetUser(ctx *gin.Context) {
	user := DummyUser{Name: "user", Email: "email@gmail.com", Address: "address"}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func (c UserControllers) LogoutUser(ctx *gin.Context) {
	_, err := ctx.Cookie("auth")
	if err != nil {
		ctx.SetCookie("auth", "", 0, "", "", false, true)
	}

}
