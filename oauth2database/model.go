package oauth2database

import (
	"time"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//Client Client
type Client struct {
	ClientID int64  `json:"clientId"`
	Secret   string `json:"secret"`
	Name     string `json:"name"`
	WebSite  string `json:"webSite"`
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	Paid     bool   `json:"paid"`
}

//ClientRedirectURI ClientRedirectURI
type ClientRedirectURI struct {
	ID       int64
	URI      string
	ClientID int64
}

//ClientAllowedURI ClientAllowedURI
type ClientAllowedURI struct {
	ID       int64
	URI      string
	ClientID int64
}

//ClientRole ClientRole
type ClientRole struct {
	ID       int64
	Role     string
	ClientID int64
}

//ClientScope ClientScope
type ClientScope struct {
	ID       int64
	Scope    string
	ClientID int64
}

//ClientRoleURI ClientRoleURI
type ClientRoleURI struct {
	ClientRoleID       int64
	ClientAllowedURIID int64
}

//RoleURI RoleURI
type RoleURI struct {
	ClientRoleID       int64
	Role               string
	ClientAllowedURIID int64
	ClientAllowedURI   string
	ClientID           int64
}

//RefreshToken RefreshToken
type RefreshToken struct {
	ID    int64
	Token string
}

//AccessToken AccessToken
type AccessToken struct {
	ID             int64
	Token          string
	Expires        time.Time
	RefreshTokenID int64
}

//start on AuthCode here

//AuthorizationCode AuthorizationCode
type AuthorizationCode struct {
	AuthorizationCode int64
	ClientID          int64
	UserID            string
	Expires           time.Time
	AccessTokenID     int64
	RandonAuthCode    string
	AlreadyUsed       bool
}

//AuthCodeScope AuthCodeScope
type AuthCodeScope struct {
	ID                int64
	Scope             string
	AuthorizationCode int64
}

//AuthCodeRevolk AuthCodeRevolk
type AuthCodeRevolk struct {
	ID                int64
	AuthorizationCode int64
}
