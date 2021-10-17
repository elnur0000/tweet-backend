package utils

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func NewCustomValidator() *CustomValidator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validator := validator.New()
	en_translations.RegisterDefaultTranslations(validator, trans)
	customValidator := &CustomValidator{
		validator: validator,
		trans:     trans,
	}
	return customValidator
}

func (cv *CustomValidator) Validate(s interface{}) error {
	values := reflect.ValueOf(s).Elem()
	typ := values.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if err := cv.validator.StructPartial(s, field.Name); err != nil {
			translatedErr := err.(validator.ValidationErrors).Translate(cv.trans)
			errMessage := getFirstErrorMessage(translatedErr)
			return echo.NewHTTPError(http.StatusBadRequest, strings.ToLower(string(errMessage[0]))+errMessage[1:])
		}
	}

	return nil
}

func getFirstErrorMessage(m validator.ValidationErrorsTranslations) string {
	for k := range m {
		return m[k]
	}
	return ""
}
