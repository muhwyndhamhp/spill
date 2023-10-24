package dto

type PostRegisterReq struct {
	Alias string `json:"alias" form:"alias" validate:"required"`
	Bio   string `json:"bio" form:"bio" validate:"required"`
}

type PostUsersCompaniesUpsertReq struct {
	CompanyName string `json:"company_name" form:"company_name" validate:"required"`
}
