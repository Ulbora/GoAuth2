package oauth2mysql

import (
	"testing"
	"github.com/Ulbora/oauth2server/domain"
)

func TestInitialize(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	t.Log("Initialize MySql database")
	expected := 0
	actual := Initialize()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestAddClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	t.Log("Closing MySql database")
	c:= domain.Client{}// new(domain.Client)	
	c.Secret = "admin2"
	c.RedirectUrl = "http://gogle.com"
	c.Name = "admin";
	c.WebSite = "www.google.com"
	c.Email = "go@google.com"
	c.Enabled = true
	
	expected := true
	actual := AddClient(&c)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}



func TestCloseDb(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	t.Log("Closing MySql database")
	expected := 0
	actual := CloseDb()
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
