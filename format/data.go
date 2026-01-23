/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 Albert Moky
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * ==============================================================================
 */
package format

/**
 *  Data Coder
 *  ~~~~~~~~~~
 *  Hex, Base58, Base64, ...
 *
 *  1. encode binary data to string;
 *  2. decode string to binary data.
 */
type DataCoder interface {

	/**
	 *  Encode binary data to local string
	 *
	 * @param data - binary data
	 * @return local string
	 */
	Encode(data []byte) string

	/**
	 *  Decode local string to binary data
	 *
	 * @param string - local string
	 * @return binary data
	 */
	Decode(string string) []byte
}

//
//  Base-64
//

var base64Coder DataCoder = nil

func SetBase64Coder(coder DataCoder) {
	base64Coder = coder
}

func Base64Encode(bytes []byte) string {
	return base64Coder.Encode(bytes)
}

func Base64Decode(b64 string) []byte {
	return base64Coder.Decode(b64)
}

//
//  Base-58
//

var base58Coder DataCoder = nil

func SetBase58Coder(coder DataCoder) {
	base58Coder = coder
}

func Base58Encode(bytes []byte) string {
	return base58Coder.Encode(bytes)
}

func Base58Decode(b58 string) []byte {
	return base58Coder.Decode(b58)
}

//
//  Hex
//

var hexCoder DataCoder = nil

func SetHexCoder(coder DataCoder) {
	hexCoder = coder
}

func HexEncode(bytes []byte) string {
	return hexCoder.Encode(bytes)
}

func HexDecode(h string) []byte {
	return hexCoder.Decode(h)
}
