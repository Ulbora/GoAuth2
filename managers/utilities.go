package managers

import (
	"math/rand"
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

const (
	shifter = 4
)

func generateRandonAuthCode() string {
	var text = ""
	var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		t := possible[rand.Intn(len(possible))]
		text += string(t)
	}
	return text
}

func generateClientSecret() string {
	var text = ""
	var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < 50; i++ {
		rand.Seed(time.Now().UnixNano())
		t := possible[rand.Intn(len(possible))]
		text += string(t)
	}
	return text
}

func hashUser(username string) string {
	var rtn string
	for i := 0; i < len(username); i++ {
		//fmt.Println("username[i]: ", username[i])
		c := username[i]
		//fmt.Println("c before: ", c)
		c += shifter
		//fmt.Println("c: ", c)
		char := string(c)
		//fmt.Println("char: ", char)
		rtn += char
	}
	return rtn
}

func unHashUser(username string) string {
	var rtn string
	for i := 0; i < len(username); i++ {
		//fmt.Println("username[i]: ", username[i])
		c := username[i]
		//fmt.Println("c before: ", c)
		c -= shifter
		//fmt.Println("c: ", c)
		char := string(c)
		//fmt.Println("char: ", char)
		rtn += char
	}
	return rtn
}
