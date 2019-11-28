package managers

import (
	"fmt"
	"testing"
)

func TestMockManager_AddClient(t *testing.T) {
	var man MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 1
	var m Manager
	m = &man
	var c Client

	suc, id := m.AddClient(&c)
	if !suc || id != 1 {
		t.Fail()
	}
}

func TestMockManager_UpdateClient(t *testing.T) {
	var man MockManager
	man.MockUpdateSuccess1 = true

	var m Manager
	m = &man
	var c Client

	suc := m.UpdateClient(&c)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_GetClient(t *testing.T) {
	var c Client
	c.ClientID = 1
	var man MockManager
	man.MockClient = c

	var m Manager
	m = &man

	clt := m.GetClient(2)
	if clt.ClientID != 1 {
		t.Fail()
	}
}

func TestMockManager_GetClientList(t *testing.T) {
	var c Client
	var cl = []Client{c}
	var man MockManager
	man.MockClientList = cl

	var m Manager
	m = &man

	clt := m.GetClientList()
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_GetClientSearchList(t *testing.T) {
	var c Client
	var cl = []Client{c}
	var man MockManager
	man.MockClientList = cl

	var m Manager
	m = &man

	clt := m.GetClientSearchList("test")
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_DeleteClient(t *testing.T) {
	var man MockManager
	man.MockDeleteSuccess1 = true

	var m Manager
	m = &man

	suc := m.DeleteClient(1)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_AddClientRedirectURI(t *testing.T) {
	var man MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 1
	var m Manager
	m = &man
	var c ClientRedirectURI

	suc, id := m.AddClientRedirectURI(&c)
	if !suc || id != 1 {
		t.Fail()
	}
}

func TestMockManager_GetClientRedirectURIList(t *testing.T) {
	var c ClientRedirectURI
	var cl = []ClientRedirectURI{c}
	var man MockManager
	man.MockClientRedirectURIList = cl

	var m Manager
	m = &man

	clt := m.GetClientRedirectURIList(1)
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_DeleteClientRedirectURI(t *testing.T) {
	var man MockManager
	man.MockDeleteSuccess1 = true

	var m Manager
	m = &man

	suc := m.DeleteClientRedirectURI(1)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_AddClientRole(t *testing.T) {
	var man MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 1
	var m Manager
	m = &man
	var c ClientRole

	suc, id := m.AddClientRole(&c)
	if !suc || id != 1 {
		t.Fail()
	}
}

func TestMockManager_GetClientRoleList(t *testing.T) {
	var c ClientRole
	var cl = []ClientRole{c}
	var man MockManager
	man.MockClientRoleList = cl

	var m Manager
	m = &man

	clt := m.GetClientRoleList(1)
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_DeleteClientRole(t *testing.T) {
	var man MockManager
	man.MockDeleteSuccess1 = true

	var m Manager
	m = &man

	suc := m.DeleteClientRole(1)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_AddClientAllowedURI(t *testing.T) {
	var man MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 1
	var m Manager
	m = &man
	var c ClientAllowedURI

	suc, id := m.AddClientAllowedURI(&c)
	if !suc || id != 1 {
		t.Fail()
	}
}

func TestMockManager_UpdateClientAllowedURI(t *testing.T) {
	var man MockManager
	man.MockUpdateSuccess1 = true

	var m Manager
	m = &man
	var c ClientAllowedURI

	suc := m.UpdateClientAllowedURI(&c)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_GetClientAllowedURI(t *testing.T) {
	var c ClientAllowedURI
	c.ID = 2
	var man MockManager
	man.MockClientAllowedURI = c

	var m Manager
	m = &man

	clt := m.GetClientAllowedURI(2)
	fmt.Println("clt: ", clt)
	if clt.ID != 2 {
		t.Fail()
	}
}

func TestMockManager_GetClientAllowedURIList(t *testing.T) {
	var c ClientAllowedURI
	var cl = []ClientAllowedURI{c}
	var man MockManager
	man.MockClientAllowedURIList = cl

	var m Manager
	m = &man

	clt := m.GetClientAllowedURIList(1)
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_DeleteClientAllowedURI(t *testing.T) {
	var man MockManager
	man.MockDeleteSuccess1 = true

	var m Manager
	m = &man

	suc := m.DeleteClientAllowedURI(1)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_AddClientRoleURI(t *testing.T) {
	var man MockManager
	man.MockInsertSuccess1 = true
	var m Manager
	m = &man
	var c ClientRoleURI

	suc := m.AddClientRoleURI(&c)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_GetClientRoleAllowedURIList(t *testing.T) {
	var c ClientRoleURI
	var cl = []ClientRoleURI{c}
	var man MockManager
	man.MockClientRoleURIList = cl

	var m Manager
	m = &man

	clt := m.GetClientRoleAllowedURIList(1)
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_DeleteClientRoleURI(t *testing.T) {
	var man MockManager
	man.MockDeleteSuccess1 = true

	var m Manager
	m = &man
	var cr ClientRoleURI

	suc := m.DeleteClientRoleURI(&cr)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_AddClientGrantType(t *testing.T) {
	var man MockManager
	man.MockInsertSuccess1 = true
	man.MockInsertID1 = 2
	var m Manager
	m = &man
	var c ClientGrantType

	suc, id := m.AddClientGrantType(&c)
	if !suc || id != 2 {
		t.Fail()
	}
}

func TestMockManager_GetClientGrantTypeList(t *testing.T) {
	var c ClientGrantType
	var cl = []ClientGrantType{c}
	var man MockManager
	man.MockClientGrantTypeList = cl

	var m Manager
	m = &man

	clt := m.GetClientGrantTypeList(1)
	if clt == nil || len(*clt) != 1 {
		t.Fail()
	}
}

func TestMockManager_DeleteClientGrantType(t *testing.T) {
	var man MockManager
	man.MockDeleteSuccess1 = true

	var m Manager
	m = &man

	suc := m.DeleteClientGrantType(1)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_AuthorizeAuthCode(t *testing.T) {
	var man MockManager
	man.MockAuthCodeAuthorizeSuccess = true
	man.MockAuthCode = 2
	man.MockAuthCodeString = "1234"

	var m Manager
	m = &man
	var ac AuthCode
	suc, cd, cds := m.AuthorizeAuthCode(&ac)
	if !suc || cd != 2 || cds != "1234" {
		t.Fail()
	}
}

func TestMockManager_CheckAuthCodeApplicationAuthorization(t *testing.T) {
	var man MockManager
	man.MockAuthCodeAuthorized = true

	var m Manager
	m = &man
	var ac AuthCode
	suc := m.CheckAuthCodeApplicationAuthorization(&ac)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_ValidateAuthCodeClientAndCallback(t *testing.T) {
	var man MockManager
	var ac AuthCodeClient
	ac.Valid = true
	man.MockAuthCodeClient = ac

	var m Manager
	m = &man
	var acode AuthCode
	acc := m.ValidateAuthCodeClientAndCallback(&acode)
	if !acc.Valid {
		t.Fail()
	}
}

func TestMockManager_AuthorizeImplicit(t *testing.T) {
	var man MockManager
	var im1 ImplicitReturn
	im1.ID = 2
	man.MockImplicitReturn = im1
	man.MockImplicitAuthorizeSuccess = true

	var m Manager
	m = &man
	var imp Implicit
	suc, im := m.AuthorizeImplicit(&imp)
	if !suc || im.ID != 2 {
		t.Fail()
	}
}

func TestMockManager_CheckImplicitApplicationAuthorization(t *testing.T) {
	var man MockManager

	man.MockImplicitAuthorized = true

	var m Manager
	m = &man
	var imp Implicit
	suc := m.CheckImplicitApplicationAuthorization(&imp)
	if !suc {
		t.Fail()
	}
}

func TestMockManager_ValidateImplicitClientAndCallback(t *testing.T) {
	var man MockManager
	var im1 ImplicitClient
	im1.Valid = true
	man.MockImplicitClient = im1

	var m Manager
	m = &man
	var imp Implicit
	im := m.ValidateImplicitClientAndCallback(&imp)
	if !im.Valid {
		t.Fail()
	}
}

func TestMockManager_GetAuthCodeToken(t *testing.T) {
	var man MockManager
	var tk1 Token
	tk1.AccessToken = "1234"
	man.MockToken = tk1
	man.MockAuthCodeTokenSuccess = true

	var m Manager
	m = &man
	var imp AuthCodeTokenReq
	suc, tk := m.GetAuthCodeToken(&imp)
	if !suc || tk.AccessToken != "1234" {
		t.Fail()
	}
}

func TestMockManager_GetCredentialsToken(t *testing.T) {
	var man MockManager
	var tk1 Token
	tk1.AccessToken = "1234"
	man.MockToken = tk1
	man.MockCredentialsTokenSuccess = true

	var m Manager
	m = &man
	var imp CredentialsTokenReq
	suc, tk := m.GetCredentialsToken(&imp)
	if !suc || tk.AccessToken != "1234" {
		t.Fail()
	}
}

func TestMockManager_GetPasswordToken(t *testing.T) {
	var man MockManager
	var tk1 Token
	tk1.AccessToken = "1234"
	man.MockToken = tk1
	man.MockPasswordTokenSuccess = true

	var m Manager
	m = &man
	var imp PasswordTokenReq
	suc, tk := m.GetPasswordToken(&imp)
	if !suc || tk.AccessToken != "1234" {
		t.Fail()
	}
}

func TestMockManager_GetAuthCodeAccesssTokenWithRefreshToken(t *testing.T) {
	var man MockManager
	var tk1 Token
	tk1.AccessToken = "1234"
	man.MockToken = tk1
	man.MockAuthCodeRefreshTokenSuccess = true

	var m Manager
	m = &man
	var imp RefreshTokenReq
	suc, tk := m.GetAuthCodeAccesssTokenWithRefreshToken(&imp)
	if !suc || tk.AccessToken != "1234" {
		t.Fail()
	}
}

func TestMockManager_GetPasswordAccesssTokenWithRefreshToken(t *testing.T) {
	var man MockManager
	var tk1 Token
	tk1.AccessToken = "1234"
	man.MockToken = tk1
	man.MockPasswordRefreshTokenSuccess = true

	var m Manager
	m = &man
	var imp RefreshTokenReq
	suc, tk := m.GetPasswordAccesssTokenWithRefreshToken(&imp)
	if !suc || tk.AccessToken != "1234" {
		t.Fail()
	}
}

func TestMockManager_ValidateAccessToken(t *testing.T) {
	var man MockManager

	man.MockValidateAccessTokenSuccess = true

	var m Manager
	m = &man
	var imp ValidateAccessTokenReq
	suc := m.ValidateAccessToken(&imp)
	if !suc {
		t.Fail()
	}
}


func TestMockManager_UserLogin(t *testing.T) {
	var man MockManager

	man.MockUserLoginSuccess = true

	var m Manager
	m = &man
	var l Login
	suc := m.UserLogin(&l)
	if !suc {
		t.Fail()
	}
}