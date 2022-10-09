package api

import (
	"github.com/go-playground/validator/v10"
	"pro.qbitty/simplebank/util"
)

var currencyValid validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}