package validate

import "github.com/go-playground/validator"

var validate = validator.New()

// Validate 校验数据是否合法
func Validate(model interface{}) (error, string) {
	err := validate.Struct(model)
	if err != nil {
		params := ""
		for _, e := range err.(validator.ValidationErrors) {
			params += e.Field() + ","
		}
		if len(params) > 1 {
			params = params[0 : len(params)-1]
		}
		params += " 参数错误"
		return err, params
	}
	return nil, ""
}
