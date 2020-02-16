//Package managers ...
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

//GenerateAccessToken GenerateAccessToken
func (m *OauthManager) GenerateAccessToken(pl *Payload) string {
	var token string
	m.Log.Debug("m.TokenParams: ", m.TokenParams)
	var theKey string
	var theIssuer string
	var theAudience string
	if m.TokenParams != nil && m.TokenParams.AccessTokenKey != "" {
		theKey = m.TokenParams.AccessTokenKey
	} else {
		theKey = m.Db.GetAccessTokenKey()
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
	//key := m.Db.GetAccessTokenKey()
	//if key != "" {
	pl.SecretKey = theKey
	pl.Subject = pl.Grant
	pl.Issuer = theIssuer
	pl.Audience = theAudience
	token = m.GenerateJwtToken(pl)
	//}
	return token
}
