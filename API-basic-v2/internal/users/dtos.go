package users

type CreateUserDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserDTO struct {
	ID    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type DeleteUserDTO struct {
	ID int `json:"id" binding:"required"`
}

type ObtainUserDTO struct {
	ID int `json:"id" binding:"required"`
}
