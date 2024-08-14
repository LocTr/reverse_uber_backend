package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 100),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("Password should not contain any whitespace"),
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
