/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2021 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	. "github.com/dimchat/mkm-go/protocol"
	"strings"
)

type GeneralIDFactory struct {
	IDFactory

	_identifiers map[string]ID
}

func (factory *GeneralIDFactory) Init() *GeneralIDFactory {
	factory._identifiers = make(map[string]ID)
	return factory
}

func (factory *GeneralIDFactory) CreateID(name string, address Address, terminal string) ID {
	identifier := concat(name, address, terminal)
	id := factory._identifiers[identifier]
	if id == nil {
		id = NewIdentifier(identifier, name, address, terminal)
		factory._identifiers[identifier] = id
	}
	return id
}

func (factory *GeneralIDFactory) ParseID(identifier string) ID {
	id := factory._identifiers[identifier]
	if id == nil {
		id = parse(identifier)
		if id != nil {
			factory._identifiers[identifier] = id
		}
	}
	return id
}

func concat(name string, address Address, terminal string) string {
	str := address.String()
	if name != "" {
		str = name + "@" + str
	}
	if terminal != "" {
		str = str + "/" + terminal
	}
	return str
}

func parse(identifier string) ID {
	var name string
	var address Address
	var terminal string
	// split ID string
	pair := strings.Split(identifier, "/")
	if len(pair) == 1 {
		// no terminal
		terminal = ""
	} else {
		// got terminal
		terminal = pair[1]
	}
	// split name & address
	pair = strings.Split(pair[0], "@")
	if len(pair) == 1 {
		// got address without name
		name = ""
		address = AddressParse(pair[0])
	} else {
		// got name & address
		name = pair[0]
		address = AddressParse(pair[1])
	}
	if address == nil {
		return nil
	}
	return NewIdentifier(identifier, name, address, terminal)
}

func BuildGeneralIDFactory() IDFactory {
	factory := IDGetFactory()
	if factory == nil {
		factory = new(GeneralIDFactory).Init()
		IDSetFactory(factory)
	}
	return factory
}
