package handlers

import (
	"errors"
	"reflect"

	"github.com/google/uuid"
)

type appHandler struct {
}

type AppHandler interface {
	ValidateFields(fields []string, data interface{}) error
	IsValidUUID(u string) bool
}

func NewAppHandler() AppHandler {
	return &appHandler{}
}

func (ah *appHandler) IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func (ah *appHandler) isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func (ah *appHandler) ValidateFields(fields []string, data interface{}) error {
	for _, field := range fields {
		if ah.isZero(reflect.ValueOf(data).FieldByName(field)) {
			return errors.New(field + " is required")
		}
	}
	return nil
}
