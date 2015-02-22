package validation

import (
	"fmt"
	"reflect"
)

func Valid(target interface{}) error {
	if target == nil {
		return fmt.Errorf("target is nil.")
	}

	t := reflect.TypeOf(target)
	v := reflect.ValueOf(target)
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Type.Kind() {
			case reflect.String:
				value := v.FieldByName(t.Field(i).Name)
				valid := t.Field(i).Tag.Get("valid")
				if valid == "required" {
					if value.String() == "" {
						return fmt.Errorf("%s is required.", t.Field(i).Name)
					}
				}
			default:
		}
	}

	return nil
}