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
	//"fmt"

	odb "github.com/Ulbora/GoAuth2/oauth2database"
)

//Client Client
type Client struct {
	ClientID     int64  `json:"clientId"`
	Secret       string `json:"secret"`
	Name         string `json:"name"`
	WebSite      string `json:"webSite"`
	Email        string `json:"email"`
	Enabled      bool   `json:"enabled"`
	Paid         bool   `json:"paid"`
	RedirectURIs *[]ClientRedirectURI
}

// //ClientRedirectURI ClientRedirectURI
// type ClientRedirectURI struct {
// 	ID       int64
// 	URI      string
// 	ClientID int64
// }

//AddClient AddClient
func (m *OauthManager) AddClient(client *Client) (bool, int64) {
	var suc bool
	var id int64
	//fmt.Println("client: ", client)
	//fmt.Println("uris: ", client.RedirectURIs)
	if len(*client.RedirectURIs) > 0 && (*client.RedirectURIs)[0].URI != "" {
		//fmt.Println("client: ", client)
		var c odb.Client
		c.Secret = generateClientSecret()
		c.Name = client.Name
		c.WebSite = client.WebSite
		c.Email = client.Email
		c.Enabled = client.Enabled
		c.Paid = client.Paid
		var uri []odb.ClientRedirectURI
		for _, u := range *client.RedirectURIs {
			var curi odb.ClientRedirectURI
			curi.URI = u.URI
			uri = append(uri, curi)
		}
		//fmt.Println("c: ", c)
		//fmt.Println("uri: ", uri)
		s, tid := m.Db.AddClient(&c, &uri)
		//fmt.Println("suc: ", s)
		suc = s
		id = tid
	}
	return suc, id
}

//UpdateClient UpdateClient
func (m *OauthManager) UpdateClient(client *Client) bool {
	// var suc bool
	var c odb.Client
	if client.Secret == "" {
		c.Secret = generateClientSecret()
	} else {
		c.Secret = client.Secret
	}
	c.ClientID = client.ClientID
	c.Name = client.Name
	c.WebSite = client.WebSite
	c.Email = client.Email
	c.Enabled = client.Enabled
	c.Paid = client.Paid
	//fmt.Println("c in update: ", c)

	suc := m.Db.UpdateClient(&c)

	return suc
}

//GetClient GetClient
func (m *OauthManager) GetClient(id int64) *Client {
	var rtn Client
	c := m.Db.GetClient(id)
	//fmt.Println("client in get: ", c)
	if c.ClientID != 0 {
		rtn.ClientID = c.ClientID
		rtn.Secret = c.Secret
		rtn.Name = c.Name
		rtn.WebSite = c.WebSite
		rtn.Email = c.Email
		rtn.Enabled = c.Enabled
		rtn.Paid = c.Paid
		uris := m.Db.GetClientRedirectURIList(id)
		var cruis []ClientRedirectURI
		for _, u := range *uris {
			var uri ClientRedirectURI
			uri.ID = u.ID
			uri.URI = u.URI
			uri.ClientID = u.ClientID
			cruis = append(cruis, uri)
		}
		rtn.RedirectURIs = &cruis
	}
	return &rtn
}

//GetClientList GetClientList
func (m *OauthManager) GetClientList() *[]Client {
	var rtn []Client
	cs := m.Db.GetClients()
	for _, c := range *cs {
		cc := populateClient(c)
		rtn = append(rtn, cc)
	}
	return &rtn
}

//GetClientSearchList GetClientSearchList
func (m *OauthManager) GetClientSearchList(name string) *[]Client {
	var rtn []Client
	cs := m.Db.SearchClients(name)
	for _, c := range *cs {
		cc := populateClient(c)
		rtn = append(rtn, cc)
	}
	return &rtn
}

//DeleteClient DeleteClient
func (m *OauthManager) DeleteClient(id int64) bool {
	suc := m.Db.DeleteClient(id)
	return suc
}

func populateClient(c odb.Client) Client {
	var cc Client
	cc.ClientID = c.ClientID
	cc.Name = c.Name
	cc.WebSite = c.WebSite
	cc.Email = c.Email
	cc.Enabled = c.Enabled
	return cc
}
