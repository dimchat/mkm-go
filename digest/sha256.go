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

import (
	"crypto/sha256"
)

type SHA256Digester struct {}

func (digester SHA256Digester) Init() DataDigester {
	return digester
}

//-------- IDataDigester

func (digester SHA256Digester) Digest(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

//
//  Instance of DataDigester
//
var sha256Digester DataDigester = new(SHA256Digester)

func SHA256SetDigester(digester DataDigester) {
	sha256Digester = digester
}

func SHA256(bytes []byte) []byte {
	return sha256Digester.Digest(bytes)
}
