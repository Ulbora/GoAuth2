//Package handlers ...
package handlers

import (
	"testing"
)

func TestUseMockWeb(t *testing.T) {
	h := UseMockWeb()
	if h == nil {
		t.Fail()
	}
}

func TestUseMockRest(t *testing.T) {
	h := UseMockRest()
	if h == nil {
		t.Fail()
	}
}
