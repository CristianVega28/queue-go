package utils

import (
	"errors"
	"reflect"
)

func ReflectStruct(stt any) (reflect.Type, error) {
	t := reflect.TypeOf(stt)

	if t.Kind() == reflect.Ptr {
		if t.Elem().Kind() != reflect.Struct {
			return nil, errors.New("It isnt struct")
		}
	} else {
		if t.Kind() != reflect.Struct {
			return nil, errors.New("It isnt struct")
		}
	}

	if t.Kind() == reflect.Ptr {
		return t.Elem(), nil
	} else {
		return t, nil
	}
}
