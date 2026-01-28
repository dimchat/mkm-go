/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2026 Albert Moky
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
package ext

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	"net/url"
)

/**
 *  TED Helper
 */
type TransportableDataHelper interface {

	SetTransportableDataFactory(factory TransportableDataFactory)
	GetTransportableDataFactory() TransportableDataFactory

	ParseTransportableData(ted interface{}) TransportableData
}

var sharedTransportableDataHelper TransportableDataHelper = nil

func SetTransportableDataHelper(helper TransportableDataHelper) {
	sharedTransportableDataHelper = helper
}

func GetTransportableDataHelper() TransportableDataHelper {
	return sharedTransportableDataHelper
}

/**
 *  PNF Helper
 */
type TransportableFileHelper interface {

	SetTransportableFileFactory(factory TransportableFileFactory)
	GetTransportableFileFactory() TransportableFileFactory

	ParseTransportableFile(pnf interface{}) TransportableFile

	CreateTransportableFile(data TransportableData, filename string,
		                    url url.URL, password DecryptKey) TransportableFile
}

var sharedTransportableFileHelper TransportableFileHelper = nil

func SetTransportableFileHelper(helper TransportableFileHelper) {
	sharedTransportableFileHelper = helper
}

func GetTransportableFileHelper() TransportableFileHelper {
	return sharedTransportableFileHelper
}
