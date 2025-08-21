package dto

type FieldEmail struct {
	Email string `json:"email" query:"email"`
}

type FilterUser struct {
	PhoneNumber string `query:"phoneNumber"`
	Name        string `query:"name"`
	FieldEmail
	FieldPagination
}

type ResponseUserDetails struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FieldEmail
}

type UpdateUserDetails struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FieldEmail
}
