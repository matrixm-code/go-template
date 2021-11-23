package http

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gtemplate/internal/template-api/conf"
	"gtemplate/tools"
	"reflect"
	"regexp"
	"strings"
)

type Validator struct {
	c *conf.AppConfig
}

func NewValidator(c *conf.AppConfig) *Validator {
	return &Validator{c: c}
}

func (v *Validator) SetValidate() {
	var err error
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = val.RegisterValidation("rfv", v.relateFiledValue())
		err = val.RegisterValidation("cpuwr", v.cpuWithRegex())
		err = val.RegisterValidation("memwr", v.memWithRegex())
		err = val.RegisterValidation("volwr", v.vomuleWithRegex())
		if err != nil {
			return
		}

	}
}

func (v Validator) relateFiledValue() validator.Func {
	return func(fl validator.FieldLevel) bool {
		param := strings.Split(fl.Param(), `:`)
		paramField := param[0]
		paramValue := param[1]
		if paramField == `` {
			return true
		}
		var paramFieldValue reflect.Value
		if fl.Parent().Kind() == reflect.Ptr {
			paramFieldValue = fl.Parent().Elem().FieldByName(paramField)
		} else {
			paramFieldValue = fl.Parent().FieldByName(paramField)
		}
		if paramFieldValue.Kind() == reflect.String {
			return paramFieldValue.String() == paramValue
		}
		return false
	}
}

func (v Validator) cpuWithRegex() validator.Func  {
	return func(fl validator.FieldLevel) bool {
		reg, err := regexp.Compile("(^[1-9][0-9]*)(g|G)$")
		if err != nil {
			return false
		}
		return reg.MatchString(fl.Field().String())
	}
}

func (v Validator) memWithRegex() validator.Func  {
	return func(fl validator.FieldLevel) bool {
		reg, err := regexp.Compile("^[1-9][0-9]*$")
		if err != nil {
			return false
		}
		return reg.MatchString(fl.Field().String())
	}
}

func (v Validator) vomuleWithRegex() validator.Func {
	return func(fl validator.FieldLevel) bool {
		volume := fl.Field().String()
		reg, err := regexp.Compile("^((/[^/,]+)+=(/[^/,]+)+,)*((/[^/,]+)+=(/[^/,]+)+)$")
		if err != nil {
			return false
		}
		if reg.MatchString(volume) {
			repeat := make([]string, 0)
			insPath := strings.Split(volume, ",")
			for _,path := range insPath {
				if path == "/home/shared" {
					return false
				}
				if tools.IsKeyInSlice(repeat, path) {
					return false
				}
				repeat = append(repeat, path)
			}
		}
		return true
	}
}