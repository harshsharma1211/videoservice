package handler

import (
	"fmt"
	"net/http"
	"uservideoservice/model"

	"github.com/gin-gonic/gin"
)

type IUser interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	VerifyUser() gin.HandlerFunc
}

func NewUserHandler() IUser {
	return &User{
		userData: make(map[string]model.User),
	}
}

type User struct {
	userData map[string]model.User
}

func (u *User) AddUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.userData[user.Name] = user
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (u *User) GetUser(ctx *gin.Context) {
	// Extract query parameters from the URL
	username := ctx.GetHeader("username")
	password := ctx.GetHeader("password")
	// Check if required parameters are provided
	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	token, err := createToken(username, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to login"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *User) VerifyUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// Extract query parameters from the URL
		token := ctx.GetHeader("token")
		// Check if required parameters are provided
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
			ctx.Abort()
			return
		}

		username, password, err := getClaimsFromToken(token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to verify user"})
			ctx.Abort()
			return
		}
		fmt.Println("username ", username)
		fmt.Println("password ", password)
		if user, ok := u.userData[username]; !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to verify user"})
			ctx.Abort()
			return

		} else {
			if user.Password == password {
				ctx.Next()
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to verify user"})
		ctx.Abort()
		return
	}
}
