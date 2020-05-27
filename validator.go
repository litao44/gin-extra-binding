package ebinding

import "github.com/gin-gonic/gin/binding"

var Validator = binding.Validator

func validate(obj interface{}) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
