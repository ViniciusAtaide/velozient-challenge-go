package domain

type Validatable interface {
	PasswordCardCreate | PasswordCardUpdate | IdPathParam
}

// PasswordCardCreate model info
// @Description PasswordCard information
// @Description with url name username and password
type PasswordCardCreate struct {
	URL      string `json:"url" validate:"required,url"`
	Name     string `json:"name" validate:"required,min=3"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

// PasswordCardUpdate model info
// @Description PasswordCard information
// @Description with url name username and password
type PasswordCardUpdate struct {
	URL      string `json:"url" validate:"omitempty,url"`
	Name     string `json:"name" validate:"omitempty,min=3"`
	Username string `json:"username" validate:"omitempty,min=3,max=32"`
	Password string `json:"password" validate:"omitempty,min=8"`
}

type IdPathParam struct {
	Id string `json:"id" validate:"required"`
}
