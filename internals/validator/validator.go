package validator

import "strings"

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}
func (v *Validator) AddFieldError(key, message string) {

	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}
func PermittedInt(value int, permittedValues ...int) bool { 
    for i := range permittedValues { 
        if value == permittedValues[i] { 
            return true 
        } 
    } 
    return false
}

func Marksvalid(obtained int,total int) bool{
	return obtained <= total
}
func Notless(value int,n int) bool{
	return value>n
}


