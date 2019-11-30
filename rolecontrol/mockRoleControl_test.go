package rolecontrol

import (
	"fmt"
	"testing"
)



func TestMockOauthAssets_AddControledURLs(t *testing.T) {
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

	var oa MockOauthAssets
	
	ac := oa.GetNewAssetControl()
	suc := ac.AddControledURLs(&cus)
	if !suc {
		t.Fail()
	}
}

func TestMockOauthAssets_GetControlledAsset(t *testing.T) {
	var oa MockOauthAssets
	oa.MockSuccess = true 
	oa.MockAllowedRole = "admin"
	ac := oa.GetNewAssetControl()
	suc, role := ac.GetControlledAsset("/test", "test")
	fmt.Println("role: ", role)
	if !suc || role != "admin" {
		t.Fail()
	}
}

func TestMockOauthAssets_GetControlledAsset2(t *testing.T) {
	var oa MockOauthAssets
	oa.MockSuccess = true 
	oa.MockAllowedRole = "superAdmin"
	ac := oa.GetNewAssetControl()
	suc, role := ac.GetControlledAsset("/test", "test2")
	fmt.Println("role: ", role)
	if !suc || role != "superAdmin" {
		t.Fail()
	}
}

func TestMockOauthAssets_GetControlledAsset3(t *testing.T) {
	var oa MockOauthAssets
	oa.MockSuccess = false 
	//oa.MockAllowedRole = "superAdmin"
	ac := oa.GetNewAssetControl()
	suc, role := ac.GetControlledAsset("/test1", "test2")
	fmt.Println("role: ", role)
	if suc  {
		t.Fail()
	}
}


func TestMockOauthAssets_GetControlledAsset4(t *testing.T) {
	var oa MockOauthAssets
	oa.MockSuccess = false 
	//oa.MockAllowedRole = "superAdmin"
	ac := oa.GetNewAssetControl()
	suc, role := ac.GetControlledAsset("/test", "test22")
	fmt.Println("role: ", role)
	if suc  {
		t.Fail()
	}
}