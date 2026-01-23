/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2021 Albert Moky
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
package digest

/**
 *  Message Digest
 *  ~~~~~~~~~~~~~~
 *  MD5, SHA1, SHA256, Keccak256, RipeMD160, ...
 */
type MessageDigester interface {

	/**
	 *  Get digest of binary data
	 *
	 * @param data - binary data
	 * @return binary data
	 */
	Digest(data []byte) []byte
}

//
//  SHA-256
//

var sha256Digester MessageDigester = nil

func SetSHA256Digester(digester MessageDigester) {
	sha256Digester = digester
}

func SHA256(bytes []byte) []byte {
	return sha256Digester.Digest(bytes)
}

//
//  Keccak-256
//

var keccak256Digester MessageDigester = nil

func SetKECCAK256Digester(digester MessageDigester) {
	keccak256Digester = digester
}

func KECCAK256(bytes []byte) []byte {
	return keccak256Digester.Digest(bytes)
}

//
//  RipeMD-160
//
var ripemd160Digester MessageDigester = nil

func SetRIPEMD160Digester(digester MessageDigester) {
	ripemd160Digester = digester
}

func RIPEMD160(bytes []byte) []byte {
	return ripemd160Digester.Digest(bytes)
}
