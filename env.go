package env

import (
	"fmt"
	"os"
	"strconv"
)

func Str(key string, choices ...string) (value string, err error) {
	value, err = get(key)
	if err != nil {
		return
	}
	if len(choices) > 0 {
		if !in(value, choices) {
			err = fmt.Errorf("%s is invalid, choices = %v", key, choices)
		}
	}
	return
}

func Int(key string, choices ...int) (value int, err error) {
	s, err := get(key)
	if err != nil {
		return
	}
	value, err = strconv.Atoi(s)
	if err != nil {
		err = fmt.Errorf("%s is not a valid number", key)
		return
	}
	if len(choices) > 0 {
		if !in(value, choices) {
			err = fmt.Errorf("%s is invalid, choices = %v", key, choices)
		}
	}
	return
}

func Bool(key string) (value bool, err error) {
	s, err := get(key)
	if err != nil {
		return
	}
	value, err = strconv.ParseBool(s)
	if err != nil {
		err = fmt.Errorf("%s is not a valid boolean value", key)
	}
	return
}

func get(key string) (s string, err error) {
	s = os.Getenv(key)
	if s == "" {
		err = fmt.Errorf("%s is not specified", key)
	}
	return
}

func in[T comparable](value T, array []T) bool {
	for _, item := range array {
		if value == item {
			return true
		}
	}
	return false
}
