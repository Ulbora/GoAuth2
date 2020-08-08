//Package rolecontrol ...
package rolecontrol

import "fmt"

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

//ControlledAsset ControlledAsset
type ControlledAsset struct {
	ControlledAsset string `json:"controlledAsset"`
	AllowedRole     string `json:"allowedRole"`
}

//ControlledURL ControlledURL
type ControlledURL struct {
	URL   string            `json:"url"`
	Asset []ControlledAsset `json:"assets"`
}

//AssetControl AssetControl
type AssetControl interface {
	AddControledURLs(urls *[]ControlledURL) bool
	GetControlledAsset(url string, ca string) (bool, string)
}

//OauthAssets OauthAssets
type OauthAssets struct {
	m map[string]*[]ControlledAsset
}

//AddControledURLs AddControledURLs
func (c *OauthAssets) AddControledURLs(urls *[]ControlledURL) bool {
	var rtn bool
	//creates a map by url of prohibited sections to be used in a particular url
	c.m = make(map[string]*[]ControlledAsset)
	//for example
	//url https://addTest/admin/user
	//could require role superUser
	for i := range *urls {
		var u = (*urls)[i]
		c.m[u.URL] = &u.Asset
	}
	if len(c.m) > 0 {
		rtn = true
	}
	fmt.Println("assess: ", c.m)
	fmt.Println("assess list: ", c.m["/ulbora/rs/clientAllowedUri/add"])

	return rtn
}

//GetControlledAsset GetControlledAsset
func (c *OauthAssets) GetControlledAsset(url string, ca string) (bool, string) {
	var suc bool
	var rtn string
	cas := c.m[url]
	if cas != nil {
		fmt.Println("cas: ", cas)
		for _, a := range *cas {
			//example
			//url https://addTest/admin/user
			//with controlled asset "admin"
			//could require role superUser
			if a.ControlledAsset == ca {
				suc = true
				rtn = a.AllowedRole
				break
			}
		}
	}
	return suc, rtn
}

//GetNewAssetControl GetNewAssetControl
func (c *OauthAssets) GetNewAssetControl() AssetControl {
	var ac AssetControl
	ac = c
	return ac
}
