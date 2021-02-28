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
package format

import (
	"encoding/hex"
	. "github.com/dimchat/mkm-go/types"
)

type HexCoder struct {
	DataCoder
}

func (coder HexCoder) Encode(data []byte) string {
	return hex.EncodeToString(data)
}

func (coder HexCoder) Decode(string string) []byte {
	bytes, err := hex.DecodeString(string)
	if err == nil {
		return bytes
	} else {
		//panic("failed to decode string for Hex")
		return nil
	}
}

var hexCoder DataCoder = new(HexCoder)

func SetHexCoder(coder DataCoder) {
	if coder != hexCoder {
		ObjectRetain(coder)
		ObjectRelease(hexCoder)
		hexCoder = coder
	}
}

func HexEncode(bytes []byte) string {
	return hexCoder.Encode(bytes)
}

func HexDecode(h string) []byte {
	return hexCoder.Decode(h)
}
