package managers

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

import (
	"time"

	jwt "github.com/gbrlsnchs/jwt/v3"
)

//Payload Payload
type Payload struct {
	TokenType       string
	UserID          string
	ClientID        int64
	Subject         string
	Issuer          string
	ExpiresInMinute time.Duration
	Grant           string
	SecretKey       string
	RoleURIs        []string
	ScopeList       []string
}

//JwtPayload JwtPayload
type jwtPayload struct {
	Payload   jwt.Payload
	TokenType string   `json:"tokenType,omitempty"`
	UserID    string   `json:"userId,omitempty"`
	ClientID  int64    `json:"clientId,omitempty"`
	Grant     string   `json:"grant,omitempty"`
	RoleURIs  []string `json:"roleURIs,omitempty"`
	ScopeList []string `json:"scopeList,omitempty"`
}

//GenerateJwtToken GenerateJwtToken
func (m *OauthManager) GenerateJwtToken(pl *Payload) string {
	var rtn string
	now := time.Now()
	var jwtPl jwt.Payload
	jwtPl.Subject = pl.Subject
	jwtPl.Issuer = pl.Issuer
	jwtPl.ExpirationTime = jwt.NumericDate(now.Add(pl.ExpiresInMinute * time.Minute))
	jwtPl.IssuedAt = jwt.NumericDate(now)
	

	var jpl jwtPayload
	jpl.Payload = jwtPl
	jpl.TokenType = pl.TokenType
	if pl.UserID != "" {
		jpl.UserID = pl.UserID
	}
	jpl.ClientID = pl.ClientID
	if pl.Grant != "" {
		jpl.Grant = pl.Grant
	}
	jpl.RoleURIs = pl.RoleURIs
	jpl.ScopeList = pl.ScopeList

	hs := jwt.NewHS256([]byte(pl.SecretKey))
	token, err := jwt.Sign(jpl, hs)
	if err == nil {
		rtn = string(token)
	}
	return rtn
}
