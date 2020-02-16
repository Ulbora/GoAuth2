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

//GenerateRefreshToken GenerateRefreshToken
func (m *OauthManager) GenerateRefreshToken(clientID int64, userID string, grantType string) string {
	var token string
	m.Log.Debug("m.TokenParams: ", m.TokenParams)
	var theKey string
	var theIssuer string
	var theAudience string
	if m.TokenParams != nil && m.TokenParams.RefreshTokenKey != "" {
		theKey = m.TokenParams.RefreshTokenKey
	} else {
		theKey = m.Db.GetRefreshTokenKey()
	}
	if m.TokenParams != nil && m.TokenParams.Issuer != "" {
		theIssuer = m.TokenParams.Issuer
	} else {
		theIssuer = tokenIssuer
	}
	if m.TokenParams != nil && m.TokenParams.Audience != "" {
		theAudience = m.TokenParams.Audience
	} else {
		theAudience = tokenAudience
	}
	//rtk := m.Db.GetRefreshTokenKey()
	//if rtk != "" {
	// here write code to generate refresh token
	var pl Payload
	pl.TokenType = refreshTokenType
	pl.UserID = userID
	pl.ClientID = clientID
	pl.Subject = grantType
	pl.Issuer = theIssuer
	pl.Audience = theAudience
	pl.ExpiresInMinute = refreshTokenLifeInMinutes //(600 * time.Minute) => (600 * 60) => 36000 minutes => 10 hours
	pl.Grant = grantType
	pl.SecretKey = theKey
	token = m.GenerateJwtToken(&pl)
	//}
	return token
}
