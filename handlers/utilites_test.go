//Package handlers ...
package handlers

import (
	"testing"

	db "github.com/Ulbora/dbinterface"
)

func TestUseHanders_UseWebHandler(t *testing.T) {
	var dbi db.Database
	h := UseWebHandler(dbi, false, "url")
	if h == nil {
		t.Fail()
	}
}

func TestUseHanders_UseRestHandler(t *testing.T) {
	var dbi db.Database
	var assets = "WwogICB7CiAgICAgICJ1cmwiOiIvdWxib3JhL3JzL2NsaWVudEFsbG93ZWRVcmkvYWRkIiwKICAgICAgImFzc2V0cyI6WwogICAgICAgICB7CiAgICAgICAgICAgICJjb250cm9sbGVkQXNzZXQiOiJ1bGJvcmEiLAogICAgICAgICAgICAiYWxsb3dlZFJvbGUiOiJzdXBlckFkbWluIgogICAgICAgICB9CiAgICAgIF0KICAgfSwKICAgewogICAgICAidXJsIjoiL3VsYm9yYS9ycy9jbGllbnRBbGxvd2VkVXJpL3VwZGF0ZSIsCiAgICAgICJhc3NldHMiOlsKICAgICAgICAgewogICAgICAgICAgICAiY29udHJvbGxlZEFzc2V0IjoidWxib3JhIiwKICAgICAgICAgICAgImFsbG93ZWRSb2xlIjoic3VwZXJBZG1pbiIKICAgICAgICAgfQogICAgICBdCiAgIH0sCiAgIHsKICAgICAgInVybCI6Ii91bGJvcmEvcnMvY2xpZW50Um9sZS9hZGQiLAogICAgICAiYXNzZXRzIjpbCiAgICAgICAgIHsKICAgICAgICAgICAgImNvbnRyb2xsZWRBc3NldCI6InN1cGVyQWRtaW4iLAogICAgICAgICAgICAiYWxsb3dlZFJvbGUiOiJzdXBlckFkbWluIgogICAgICAgICB9CiAgICAgIF0KICAgfQpd"
	h := UseRestHandler(dbi, assets, false, "url")
	if h == nil {
		t.Fail()
	}
}
