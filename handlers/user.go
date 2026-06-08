package handlers

import (
	"net/http"
	"time"

	"book-rental-api/configs"
	"book-rental-api/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// RegisterUser handles user registration
// @Summary Register a new user
// @Description Register user baru
// @Tags Users
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "Register user baru"
// @Success 201 {object} handlers.ResponseSuccessRegister "Response berhasil register user baru"
// @Failure 400 {object} handlers.ResponseBadRequest "Bad request error"
// @Failure 500 {object} handlers.ResponseInternalServerError "Internal server error; semua error lain seperti error bad request dan error not found masuk ke sini"
// @Router /users/register [post]
func RegisterUser(c echo.Context) error {

	var user models.User

	// Bind request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	// Save user to database
	if err := configs.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to register",
			"error":   err.Error(),
		})
	}

	// Return success response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register",
		"user": map[string]interface{}{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"address":    user.Address,
		},
	})
}

// LoginUser handles user login
// @Summary User login
// @Description Login user yang sudah terdaftar
// @Tags Users
// @Accept json
// @Produce json
// @Param request body handlers.LoginUserRequest true "Login user"
// @Success 200 {object} handlers.ResponseSuccessLogin "Response berhasil login user"
// @Failure 401 {object} handlers.ResponseFailedLogin "Failed login error"
// @Router /users/login [post]
func LoginUser(c echo.Context) error {

	// Create JWT claims
	claims := jwt.MapClaims{
		"email": "carmen@gmail.com",
		"exp":   jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to generate token",
		})
	}

	// Return token in response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "success login",
		"token":   signedToken,
	})
}

// GetUserProfile returns logged-in user's profile
// @Summary Get user profile
// @Description Menampilkan informasi user yang sedang login
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} handlers.ResponseSuccessGetProfile "Response berhasil mendapatkan profile user"
// @Failure 401 {object} handlers.ResponseFailedGetProfile "Failed get profile error"
// @Success 503 {object} handlers.ResponseSuccessUpdateProfile "Response berhasil update profile user"
// @Failure 504 {object} handlers.ResponseFailedUpdateProfile "Failed update profile error"
// @Router /users/profile [get]
func GetUserProfile(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"full_name": "Carmen Nyoman",
		"address":   "Bandung",
		"date_of_birth": "2000-08-09",
	})
}