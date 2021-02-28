/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2020 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  User Document
 *  ~~~~~~~~~~~~~
 *  This interface is defined for authorizing other apps to login,
 *  which can generate a temporary asymmetric key pair for messaging.
 */
type BaseVisa struct {
	BaseDocument
	IVisa

	_key EncryptKey
}

func (doc *BaseVisa) Init(dict map[string]interface{}) *BaseVisa {
	if doc.BaseDocument.Init(dict) != nil {
		// lazy load
		doc.setKey(nil)
	}
	return doc
}

func (doc *BaseVisa) InitWithID(identifier ID, data []byte, signature []byte) *BaseVisa {
	if doc.BaseDocument.InitWithType(VISA, identifier, data, signature) != nil {
		// lazy load
		doc.setKey(nil)
	}
	return doc
}

func (doc *BaseVisa) Release() int {
	cnt := doc.BaseDocument.Release()
	if cnt == 0 {
		// this object is going to be destroyed,
		// release children
		doc.setKey(nil)
	}
	return cnt
}

func (doc *BaseVisa) setKey(key EncryptKey) {
	if key != doc._key {
		ObjectRetain(key)
		ObjectRelease(doc._key)
		doc._key = key
	}
}

//-------- IVisa

func (doc *BaseVisa) Key() EncryptKey {
	if doc._key == nil {
		info := doc.self().GetProperty("key")
		pKey := PublicKeyParse(info)
		if pKey != nil {
			key, ok := pKey.(EncryptKey)
			if ok {
				doc.setKey(key)
			}
		}
	}
	return doc._key
}

func (doc *BaseVisa) SetKey(key EncryptKey) {
	if key == nil {
		doc.self().SetProperty("key", nil)
	} else {
		info, ok := key.(Map)
		if ok {
			doc.self().SetProperty("key", info.GetMap(false))
		}
	}
	doc.setKey(key)
}

func (doc *BaseVisa) Avatar() string {
	url := doc.self().GetProperty("avatar")
	if url == nil {
		return ""
	}
	return url.(string)
}

func (doc *BaseVisa) SetAvatar(url string) {
	doc.self().SetProperty("avatar", url)
}

/**
 *  Group Document
 *  ~~~~~~~~~~~~~~
 */
type BaseBulletin struct {
	BaseDocument
	IBulletin

	_assistants []ID
}

func (doc *BaseBulletin) Init(dict map[string]interface{}) *BaseBulletin {
	if doc.BaseDocument.Init(dict) != nil {
		// lazy load
		doc.setAssistants(nil)
	}
	return doc
}

func (doc *BaseBulletin) InitWithID(identifier ID, data []byte, signature []byte) *BaseBulletin {
	if doc.BaseDocument.InitWithType(BULLETIN, identifier, data, signature) != nil {
		// lazy load
		doc.setAssistants(nil)
	}
	return doc
}

func (doc *BaseBulletin) Release() int {
	cnt := doc.BaseDocument.Release()
	if cnt == 0 {
		// this object is going to be destroyed,
		// release children
		doc.setAssistants(nil)
	}
	return cnt
}

func (doc *BaseBulletin) setAssistants(bots []ID) {
	if bots != nil {
		for _, item := range bots {
			ObjectRetain(item)
		}
	}
	if doc._assistants != nil {
		for _, item := range doc._assistants {
			ObjectRelease(item)
		}
	}
	doc._assistants = bots
}

//-------- IBulletin

func (doc *BaseBulletin) Assistants() []ID {
	if doc._assistants == nil {
		assistants := doc.self().GetProperty("assistants")
		if assistants != nil {
			doc.setAssistants(IDConvert(assistants))
		}
	}
	return doc._assistants
}

func (doc *BaseBulletin) SetAssistants(bots []ID) {
	if bots == nil {
		doc.self().SetProperty("assistants", nil)
	} else {
		doc.self().SetProperty("assistants", IDRevert(bots))
	}
	doc.setAssistants(bots)
}
