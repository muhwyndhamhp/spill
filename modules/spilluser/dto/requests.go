package dto

type PostRegisterReq struct {
	Alias string `json:"alias" form:"alias" validate:"required"`
	Bio   string `json:"bio" form:"bio" validate:"required"`
}
