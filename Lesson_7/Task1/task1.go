package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

type MyStruct struct {
	I int  `yaml:"i" json:"i"`
	Y bool `yaml:"y" json:"y"`
}

var myMap = map[string]interface{}{
	"I": 15,
	"Y": false,
	"G": 3.4,
}

func NewMyStruct() *MyStruct {
	return &MyStruct{}
}

func main() {
	x := NewMyStruct()
	err := SetStructValues(x, myMap)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(x)

}

func SetStructValues(in interface{}, m map[string]interface{}) (err error) {
	if in == nil {
		return ErrInvalidSpecification
	}

	val := reflect.ValueOf(in)

	if val.Kind() != reflect.Ptr {
		return ErrInvalidSpecification
	}

	val = val.Elem()

	if val.Kind() != reflect.Struct {
		return ErrInvalidSpecification
	}

	for name, value := range m {
		structFieldValue := val.FieldByName(name)

		if structFieldValue.IsValid() {
			fieldType := structFieldValue.Type()
			mapValue := reflect.ValueOf(value)
			structFieldValue.Set(mapValue.Convert(fieldType))
		}

	}

	return nil
}
