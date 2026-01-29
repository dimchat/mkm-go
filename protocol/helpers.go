/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2026 by Moky <albert.moky@gmail.com>
 *
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
package protocol

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Address Helper
 */
type AddressHelper interface {
	SetAddressFactory(factory AddressFactory)
	GetAddressFactory() AddressFactory

	ParseAddress(address interface{}) Address

	GenerateAddress(meta Meta, network EntityType) Address
}

var sharedAddressHelper AddressHelper = nil

func SetAddressHelper(helper AddressHelper) {
	sharedAddressHelper = helper
}

func GetAddressHelper() AddressHelper {
	return sharedAddressHelper
}

/**
 *  ID Helper
 */
type IDHelper interface {
	SetIDFactory(factory IDFactory)
	GetIDFactory() IDFactory

	ParseID(did interface{}) ID

	CreateID(name string, address Address, terminal string) ID

	GenerateID(meta Meta, network EntityType, terminal string) ID
}

var sharedIDHelper IDHelper = nil

func SetIDHelper(helper IDHelper) {
	sharedIDHelper = helper
}

func GetIDHelper() IDHelper {
	return sharedIDHelper
}

/**
 *  Meta Helper
 */
type MetaHelper interface {
	SetMetaFactory(version MetaType, factory MetaFactory)
	GetMetaFactory(version MetaType) MetaFactory

	CreateMeta(version MetaType, pKey VerifyKey, seed string, fingerprint TransportableData) Meta

	GenerateMeta(version MetaType, sKey SignKey, seed string) Meta

	ParseMeta(meta interface{}) Meta
}

var sharedMetaHelper MetaHelper = nil

func SetMetaHelper(helper MetaHelper) {
	sharedMetaHelper = helper
}

func GetMetaHelper() MetaHelper {
	return sharedMetaHelper
}

/**
 *  Document Helper
 */
type DocumentHelper interface {
	SetDocumentFactory(docType DocumentType, factory DocumentFactory)
	GetDocumentFactory(docType DocumentType) DocumentFactory

	CreateDocument(docType DocumentType, data string, signature TransportableData) Document

	ParseDocument(doc interface{}) Document
}

var sharedDocumentHelper DocumentHelper = nil

func SetDocumentHelper(helper DocumentHelper) {
	sharedDocumentHelper = helper
}

func GetDocumentHelper() DocumentHelper {
	return sharedDocumentHelper
}

/**
 *  PNF Helper
 */
type TransportableFileHelper interface {
	SetTransportableFileFactory(factory TransportableFileFactory)
	GetTransportableFileFactory() TransportableFileFactory

	ParseTransportableFile(pnf interface{}) TransportableFile

	CreateTransportableFile(data TransportableData, filename string,
		url URL, password DecryptKey) TransportableFile
}

var sharedTransportableFileHelper TransportableFileHelper = nil

func SetTransportableFileHelper(helper TransportableFileHelper) {
	sharedTransportableFileHelper = helper
}

func GetTransportableFileHelper() TransportableFileHelper {
	return sharedTransportableFileHelper
}
