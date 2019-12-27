//Package compresstoken ...
package compresstoken

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"encoding/base64"
	"io"
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

//JwtCompress JwtCompress
type JwtCompress struct {
}

//CompressJwt CompressJwt
func (c *JwtCompress) CompressJwt(jwt string) string {
	//compress jwt with zlib and package with base64
	var rtn string
	var b bytes.Buffer
	w, err := zlib.NewWriterLevel(&b, flate.BestCompression)
	if err == nil {
		w.Write([]byte(jwt))
		w.Close()
		rtn = base64.StdEncoding.EncodeToString(b.Bytes())
	}
	return rtn
}

//UnCompressJwt UnCompressJwt
func (c *JwtCompress) UnCompressJwt(cjwt string) string {
	//uncompress jwt with zlib after converting from base64
	var rtn string
	var b bytes.Buffer
	decoded, derr := base64.StdEncoding.DecodeString(cjwt)
	if derr == nil {
		b.Write(decoded)
		r, err := zlib.NewReader(&b)
		if err == nil {
			var out bytes.Buffer
			io.Copy(&out, r)
			r.Close()
			rtn = out.String()
		}
	}
	return rtn
}
