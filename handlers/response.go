package handlers

// New Register user respone to app
type RegisterUserRequest struct {
	FirstName string `json:"first_name" example:"Nyoman Ayu"`
	LastName  string `json:"last_name" example:"Carmenita"`
	Username  string `json:"username" example:"Carmen"`
	Password  string `json:"password" example:"12345678"`
	DateOfBirth string `json:"date_of_birth" example:"2002-01-01"`
}

// Response success register user
type UserRegisterResponse struct {
	FullName string `json:"full_name" example:"Nyoman Ayu Carmenita"`
	Username string `json:"username" example:"Carmen"`
	Age	  int    `json:"age" example:"22"`
}
type ResponseSuccessRegister struct {
	Message string       `json:"message" example:"success register"`
	User    UserRegisterResponse `json:"user"`
}

// Bad Request response
type ResponseBadRequest struct {
	Message string `json:"message" example:"name is required"`
}

// Internal server error; all error bad request and error not found
type ResponseInternalServerError struct {
	Message string `json:"message" example:"internal server error"`
	Detail  string `json:"detail" example:"error generated from err.Error() object"`
}

// Respone Login user
type LoginUserRequest struct {
	Username string `json:"email" example:"carmen@gmail.com"`
	Password string `json:"password" example:"12345678"`
}
// Respone User Login
type ResponseSuccessLogin struct {
	Message string `json:"message" example:"success login"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODg4ODQ4MDB9.3n8sKZl7j8mLh0v5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5z5"`
}

// Respone Failed Login
type ResponseFailedLogin struct {
	Message string `json:"message" example:"Failed login"`
}

// Respone Get User Profile
type ResponseSuccessGetProfile struct {
	Message string `json:"message" example:"success get profile"`
	User    struct {
		ID          uint   `json:"id" example:"1"`
		FirstName   string `json:"first_name" example:"Nyoman Ayu"`
		LastName    string `json:"last_name" example:"Carmenita"`
		Username    string `json:"username" example:"Carmen"`
		DateOfBirth string `json:"date_of_birth" example:"2002-01-01"`
	} `json:"user"`
}

// Respone Failed Get User Profile
type ResponseFailedGetProfile struct {
	Message string `json:"message" example:"Failed get profile"`
}

// Respone Update User Profile
type ResponseSuccessUpdateProfile struct {
	Message string `json:"message" example:"success update profile"`
	User    struct {
		ID          uint   `json:"id" example:"1"`
		FirstName   string `json:"first_name" example:"Nyoman Ayu"`
		LastName    string `json:"last_name" example:"Carmenita"`
		Username    string `json:"username" example:"Carmen"`
		DateOfBirth string `json:"date_of_birth" example:"2002-01-01"`
	} `json:"user"`
}

// Respone Failed Update User Profile
type ResponseFailedUpdateProfile struct {
	Message string `json:"message" example:"Failed update profile"`
}