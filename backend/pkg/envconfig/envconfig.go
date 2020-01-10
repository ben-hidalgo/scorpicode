package envconfig

import (
	"fmt"
	"os"
	"strconv"
)

func SetString(name string, value *string) {
	v, ok := os.LookupEnv(name)
	if ok {
		*value = v
	}
}

func SetInt(name string, value *int) {
	v, ok := os.LookupEnv(name)
	if ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s from %s to int", name, v))
		}
		*value = i
	}
}

func SetBool(name string, value *bool) {
	v, ok := os.LookupEnv(name)
	if ok {
		b, err := strconv.ParseBool(v)
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s from %s to bool", name, v))
		}
		*value = b
	}
}
