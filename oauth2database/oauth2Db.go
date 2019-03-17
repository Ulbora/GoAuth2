package oauth2database

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

//Oauth2DB Oauth2DB
type Oauth2DB interface {
	AddClient(client *Client, uris *[]ClientRedirectURI) (bool, int64)
	UpdateClient(client *Client) bool
	GetClient(clientID int64)
	GetClients() *[]Client
	SearchClients(name string) *[]Client
	DeleteClient(clientID int64) bool
}

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
