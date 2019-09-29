package managers

import (
	"fmt"
	"testing"
)

func Test_generateClientSecret(t *testing.T) {
	sec := generateClientSecret()
	fmt.Println("secret :", sec)
	if len(sec) < 50 {
		t.Fail()
	}
}

func Test_hashUser(t *testing.T) {
	h := hashUser("kenz")
	fmt.Println("hash :", h)
	if h != "oir~" {
		t.Fail()
	}
}

func Test_unHashUser(t *testing.T) {
	h := unHashUser("oir~")
	fmt.Println("hash :", h)
	if h != "kenz" {
		t.Fail()
	}
}
