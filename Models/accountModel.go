package models

type account struct {
	Id            uint    `gorm:"primaryKey"`
	First_Name    *string `json:"first_name" validate:"required, min=5, max=100"`
	Last_Name     *string `json:"last_name" validate:"required, min=5, max=100"`
	Password      *string `json:"Password" validate:"required, min=6, max=15"`
	Email         *string `json:"email" validate:"email, required"`
	Token         *string `json:"token"`
	User_type     *string `json:"user_type" validate:"required, eq=ADMIN|ew=USER"`
	Refresh_token *string `json:"refresh_token"`
	Account_id    *string `json:"accountid"`
}
