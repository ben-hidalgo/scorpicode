package envconfig

import (
	"fmt"
	"os"
	"strconv"
)

// SetString populates value if name exists
func SetString(name string, value *string) {
	v, ok := os.LookupEnv(name)
	if ok {
		*value = v
	}
}

// SetInt populates value if name exists
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

// SetInt32 populates value if name exists
func SetInt32(name string, value *int32) {
	v, ok := os.LookupEnv(name)
	if ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s from %s to int32", name, v))
		}
		*value = int32(i)
	}
}

// SetBool populates value if name exists
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
