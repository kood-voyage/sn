package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	Email  = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	Int    = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Base64 = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
)

var (
	rxEmail        = regexp.MustCompile(Email)
	rxAlpha        = regexp.MustCompile("^[a-zA-Z]+$")
	rxAlphaNum     = regexp.MustCompile("^[a-zA-Z0-9]+$")
	rxInt          = regexp.MustCompile(Int)
	rxBase64       = regexp.MustCompile(Base64)
	rxHasLowerCase = regexp.MustCompile("^[a-z]+$")
	rxHasUpperCase = regexp.MustCompile("^[A-Z]+$")
)

func Validate(data interface{}) error {
	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("validate")
		if tag != "" {
			if err := validateField(val.Field(i).Interface(), tag); err != nil {
				return fmt.Errorf("%s %s", val.Type().Field(i).Name, err)
			}
		}
	}
	return nil
}

func validateField(val interface{}, tag string) error {
	instructions := strings.Split(tag, "|")
	for _, instruction := range instructions {
		//split to function and parameters
		p := strings.SplitN(instruction, ":", 2)
		function := p[0]
		//check if parameter exists
		var parameter string
		if len(p) > 1 {
			parameter = p[1]
		}

		if err := runValidationFunction(function, parameter, val); err != nil {
			return err
		}
	}
	return nil
}

func runValidationFunction(function, parameter string, val interface{}) error {
	switch function {
	case "required":
		return required(val)
	case "min_len":
		return lengthValidation(val, parameter, func(length, threshold int) bool {
			return length < threshold
		}, "minimum of")
	case "max_len":
		return lengthValidation(val, parameter, func(length, threshold int) bool {
			return length > threshold
		}, "maximum of")
	case "email":
		return email(val)
	case "alpha":
		return alpha(val)
	case "alphanumeric":
		return alphaNum(val)
	case "integer":
		return integer(val)
	case "base64":
		return base64(val)
	case "lowercase":
		return hasLowerCase(val)
	case "uppercase":
		return hasUpperCase(val)
	case "contains":
		return contains(val, parameter)
	default:
		return errors.New("unkown validation function: " + function)
	}
}

func ValidateString(val interface{}, regex *regexp.Regexp, errorMsg string) error {
	s, ok := val.(string)
	if !ok {
		return errors.New("only string allowed")
	}
	if !regex.MatchString(s) {
		return errors.New(errorMsg)
	}
	return nil
}

func email(val interface{}) error {
	return ValidateString(val, rxEmail, "invalid format")
}

func alpha(val interface{}) error {
	return ValidateString(val, rxAlpha, "string contains more than just letters")
}

func alphaNum(val interface{}) error {
	return ValidateString(val, rxAlphaNum, "string contains more than just letters and numbers")
}

func integer(val interface{}) error {
	return ValidateString(val, rxInt, "string containts more than just integers")
}

func base64(val interface{}) error {
	return ValidateString(val, rxBase64, "string is not valid base64")
}

func hasLowerCase(val interface{}) error {
	return ValidateString(val, rxHasLowerCase, "only lowercase letters allowed")
}

func hasUpperCase(val interface{}) error {
	return ValidateString(val, rxHasUpperCase, "only uppercase letters allowed")
}

func required(val interface{}) error {
	if isZero(val) {
		return errors.New("value is required")
	}
	return nil
}

func isZero(val interface{}) bool {
	switch reflect.ValueOf(val).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val == 0
	default:
		return reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
	}
}
func lengthValidation(val interface{}, parameter string, compare func(int, int) bool, validationType string) error {
	// Convert parameter to int
	i, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	switch reflect.TypeOf(val).Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		length := reflect.ValueOf(val).Len()
		if compare(length, i) {
			return errors.New("length must be " + validationType + " " + strconv.Itoa(i))
		}
	default:
		return fmt.Errorf("%s validation only applicable to arrays, slices, maps, and strings", validationType)
	}

	return nil
}

func contains(val interface{}, param string) error {
	switch reflect.TypeOf(val).Kind() {
	case reflect.String:
		value := strings.ToLower(val.(string))
		stringValues := strings.Split(strings.ToLower(param), ",")

		allowedValues := make(map[string]struct{})
		for _, word := range stringValues {
			allowedValues[word] = struct{}{}
		}

		if _, ok := allowedValues[value]; ok {
			return nil
		}
		return fmt.Errorf("%s must be one of %v", value, param)
	default:
		return fmt.Errorf("only string allowed")
	}
}
