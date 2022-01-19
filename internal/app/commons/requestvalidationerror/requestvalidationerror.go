package requestvalidationerror

import (
	"github.com/go-playground/validator/v10"
	"github.com/stoewer/go-strcase"
)

type ValidationField struct {
	Field   string
	Message string
}

var (
	RequiredMsg = "is required"
	MaxMsg      = "maximum : "
	MinMsg      = "minimum : "
	NoChange    = "no change"
)

func GetvalidationError(err error) []ValidationField {
	var validationFields []ValidationField
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range ve {

			switch validationError.Tag() {
			case "required":
				myField := validationError.Field()
				validationFields = append(validationFields, ValidationField{
					Field:   strcase.LowerCamelCase(myField),
					Message: "this " + RequiredMsg,
				})
			case "max":
				myField := validationError.Field()
				validationFields = append(validationFields, ValidationField{
					Field:   strcase.LowerCamelCase(myField),
					Message: MaxMsg + validationError.Param(),
				})
			case "min":
				myField := validationError.Field()
				validationFields = append(validationFields, ValidationField{
					Field:   strcase.LowerCamelCase(myField),
					Message: MinMsg + validationError.Param(),
				})
			}
		}
	}

	return validationFields
}
