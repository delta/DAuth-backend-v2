package model

type Resource struct {
	Name string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validation:"required"`
	Password string `json:"password" form:"password" validation:"required"`
}

type IsAuthRequest struct {
	Email    string `json:"email" form:"email" validation:"required"`
	Password string `json:"password" form:"password" validation:"required"`
}

type EditProfile struct {
	Name        string `json:"name" form:"name" validation:"required"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" validation:"required"`
	Gender      string `json:"gender" form:"gender" validation:"required"`
}

type Student struct {
	RollNumber string `json:"roll_number"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Webmail    bool   `json:"webmail"`
}
