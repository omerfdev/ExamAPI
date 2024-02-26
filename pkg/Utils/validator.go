package utils

import (
	"exam-api/Error/errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateRequest(ctx echo.Context, request interface{}) (result interface{}, err error) {

	err = ctx.Bind(request)
	if err != nil {
		if _, ok := err.(*echo.HTTPError); ok {
			return nil, errors.ValidatorError.WrapDesc(err.(*echo.HTTPError).Message.(string))
		}
		return nil, errors.UnknownError.Wrap(err)
	}

	v := validator.New()

	v.RegisterStructValidation(producerValidations, types.SendProductsRequest{})

	if err = v.Struct(request); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			desc := ""
			for _, err := range err.(validator.ValidationErrors) {
				desc = fmt.Sprintf("Problem is: '%s' , Hint: '%s' ", err.Field(), err.ActualTag())
			}
			return nil, errors.ValidatorError.WrapDesc(desc)
		}
		return errors.UnknownError.Wrap(err), nil
	}

	return request, nil
}

func producerValidations(sl validator.StructLevel) {
	model := sl.Current().Interface().(types.SendProductsRequest)

	for i := 0; i < len(model); i++ {

		if model[i].Name == "" {
			panic(errors.ValidatorError.WrapDesc("Product Name is required!"))
		}

		if model[i].ImageUrl == "" {
			panic(errors.ValidatorError.WrapDesc("Product ImageUrl is required!"))
		}

	}
}
