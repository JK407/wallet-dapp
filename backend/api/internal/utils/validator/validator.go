package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct 通用结构体验证
func ValidateStruct(obj interface{}) error {
	err := validate.Struct(obj)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err // 系统级的错误处理
		}

		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// 检查错误类型是否为'required'
			if err.Tag() == "required" {
				errors = append(errors, fmt.Sprintf("%s is required", err.Field()))
			}
		}

		if len(errors) > 0 {
			return fmt.Errorf(strings.Join(errors, ", "))
		}
	}
	return nil
}
