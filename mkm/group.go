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
	. "github.com/dimchat/mkm-go/protocol"
)

type Group struct {
	Entity

	_founder ID
}

func (group *Group) Init(identifier ID) *Group {
	if group.Entity.Init(identifier) != nil {
		// lazy load
		group._founder = nil
	}
	return group
}

func (group Group) GetDataSource() GroupDataSource {
	return group._delegate.(GroupDataSource)
}

func (group *Group) GetFounder() ID {
	if group._founder == nil {
		delegate := group.GetDataSource()
		group._founder = delegate.GetFounder(group.ID())
	}
	return group._founder
}

func (group Group) GetOwner() ID {
	delegate := group.GetDataSource()
	return delegate.GetOwner(group.ID())
}

func (group Group) GetMembers() []ID {
	delegate := group.GetDataSource()
	return delegate.GetMembers(group.ID())
}
