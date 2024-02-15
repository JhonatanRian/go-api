package handlers

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func ErrorMissingField(name, typ string) error {
	return fmt.Errorf("Missing field %s of type %s", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return ErrorMissingField("role", "string")
	}
	if r.Company == "" {
		return ErrorMissingField("company", "string")
	}
	if r.Location == "" {
		return ErrorMissingField("location", "string")
	}
	if r.Link == "" {
		return ErrorMissingField("link", "string")
	}
	if r.Remote == nil {
		return ErrorMissingField("remote", "bool")
	}
	if r.Salary <= 0 {
		return ErrorMissingField("salary", "int64")
	}
	return nil
}
