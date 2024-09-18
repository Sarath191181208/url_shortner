package validator

import "regexp"

var CharsDigitsRX = regexp.MustCompile(`^[\d\w]{3,32}$`)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
  return len(v.Errors) == 0
}

func (v *Validator) AddError(key, error string) {
	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = error
	}
}

func (v *Validator) Check(isTrue bool, key string, error string) {
	if isTrue {
		return
	}
	v.AddError(key, error)
}

func (v *Validator) Matches(s string, rx *regexp.Regexp) bool {
  return rx.MatchString(s)
}
