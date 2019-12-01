//Package rolecontrol ...
package rolecontrol

import (
	"fmt"
	"testing"
)

var ac AssetControl

func TestOauthAssets_AddControledURLs(t *testing.T) {
	var cass ControlledAsset
	cass.ControlledAsset = "test"
	cass.AllowedRole = "admin"

	var cass2 ControlledAsset
	cass2.ControlledAsset = "test2"
	cass2.AllowedRole = "superAdmin"

	var caList = []ControlledAsset{cass, cass2}
	var cu ControlledURL
	cu.URL = "/test"
	cu.Asset = caList

	var cus = []ControlledURL{cu}

	var oa OauthAssets
	ac = oa.GetNewAssetControl()
	suc := ac.AddControledURLs(&cus)
	if !suc {
		t.Fail()
	}
}

func TestOauthAssets_GetControlledAsset(t *testing.T) {
	suc, role := ac.GetControlledAsset("/test", "test")
	fmt.Println("role: ", role)
	if !suc || role != "admin" {
		t.Fail()
	}
}

func TestOauthAssets_GetControlledAsset2(t *testing.T) {
	suc, role := ac.GetControlledAsset("/test", "test2")
	fmt.Println("role: ", role)
	if !suc || role != "superAdmin" {
		t.Fail()
	}
}

func TestOauthAssets_GetControlledAsset3(t *testing.T) {
	suc, role := ac.GetControlledAsset("/test1", "test2")
	fmt.Println("role: ", role)
	if suc {
		t.Fail()
	}
}

func TestOauthAssets_GetControlledAsset4(t *testing.T) {
	suc, role := ac.GetControlledAsset("/test", "test22")
	fmt.Println("role: ", role)
	if suc {
		t.Fail()
	}
}
