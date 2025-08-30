package user

type FieldUpdatePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type FieldPassword struct {
	Password string `json:"password" validate:"required"`
}
