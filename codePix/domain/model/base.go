package model

import (
	"time" // pacote de validacao do GO

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequireByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid: "uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
