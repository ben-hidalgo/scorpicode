package envconfig

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSetString(t *testing.T) {

	name := fmt.Sprintf("%s%d", "TestSetString", time.Now().UnixNano())
	s := "a"
	os.Setenv(name, "b")
	SetString(name, &s)
	if s != "b" {
		t.Error()
	}
	os.Unsetenv(name)
}

func TestSetInt(t *testing.T) {

	name := fmt.Sprintf("%s%d", "TestSetInt", time.Now().UnixNano())
	i := 1
	os.Setenv(name, "2")
	SetInt(name, &i)
	if i != 2 {
		t.Error()
	}
	os.Unsetenv(name)
}

func TestSetBool(t *testing.T) {

	name := fmt.Sprintf("%s%d", "TestSetInt", time.Now().UnixNano())
	b := false
	os.Setenv(name, "true")
	SetBool(name, &b)
	if b != true {
		t.Error()
	}
	os.Unsetenv(name)
}
