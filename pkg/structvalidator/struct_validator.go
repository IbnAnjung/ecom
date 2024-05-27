package structvalidator

import (
	coreerror "edot/ecommerce/error"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type IStructValidator interface {
	Validate(obj interface{}) error
}

type structValidator struct {
	validator *validator.Validate
}

func NewStructValidator() IStructValidator {
	return &structValidator{
		validator: validator.New(),
	}
}

func (v *structValidator) Validate(obj interface{}) error {
	err := v.validator.Struct(obj)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "fail validate data")
			return e
		}

		e := coreerror.NewValidationError()
		e.Errors = map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("val error", err.Error())
			e.Errors[err.Field()] = err.Error()
		}

		err = e

		return err
	}

	return nil
}
