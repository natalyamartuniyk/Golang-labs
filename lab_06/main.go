package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Server struct {
	Host       string   `yaml:"host"`
	Port       int      `yaml:"port"`
	Debug      bool     `yaml:"debug"`
	AllowedIPs []string `yaml:"allowed_ips"`
}

func ToYAML(v any) (string, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() != reflect.Struct {
		return "", errors.New("not a struct")
	}

	var result string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		tag := field.Tag.Get("yaml")

		if fieldValue.Kind() == reflect.Slice {
			result += fmt.Sprintf("%s:\n", tag)

			for j := 0; j < fieldValue.Len(); j++ {
				result += fmt.Sprintf("- %v\n", fieldValue.Index(j))
			}
		} else {
			result += fmt.Sprintf("%s: %v\n", tag, fieldValue)
		}
	}
	return result, nil
}

func main() {
	server := Server{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
		AllowedIPs: []string{
			"192.168.1.1",
			"10.0.0.1",
		},
	}

	yaml, err := ToYAML(server)
	if err != nil {
		panic(err)
	}
	fmt.Println(yaml)
}
