package config

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type AppTools struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Validate    *validator.Validate
}
