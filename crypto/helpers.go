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
package crypto

/**
 * SymmetricKey Helper
 */
type SymmetricKeyHelper interface {
	SetSymmetricKeyFactory(algorithm string, factory SymmetricKeyFactory)
	GetSymmetricKeyFactory(algorithm string) SymmetricKeyFactory

	GenerateSymmetricKey(algorithm string) SymmetricKey

	ParseSymmetricKey(key interface{}) SymmetricKey
}

var sharedSymmetricKeyHelper SymmetricKeyHelper = nil

func SetSymmetricKeyHelper(helper SymmetricKeyHelper) {
	sharedSymmetricKeyHelper = helper
}

func GetSymmetricKeyHelper() SymmetricKeyHelper {
	return sharedSymmetricKeyHelper
}

/**
 *  PublicKey Helper
 */
type PublicKeyHelper interface {
	SetPublicKeyFactory(algorithm string, factory PublicKeyFactory)
	GetPublicKeyFactory(algorithm string) PublicKeyFactory

	ParsePublicKey(key interface{}) PublicKey
}

var sharedPublicKeyHelper PublicKeyHelper = nil

func SetPublicKeyHelper(helper PublicKeyHelper) {
	sharedPublicKeyHelper = helper
}

func GetPublicKeyHelper() PublicKeyHelper {
	return sharedPublicKeyHelper
}

/**
 *  PrivateKey Helper
 */
type PrivateKeyHelper interface {
	SetPrivateKeyFactory(algorithm string, factory PrivateKeyFactory)
	GetPrivateKeyFactory(algorithm string) PrivateKeyFactory

	GeneratePrivateKey(algorithm string) PrivateKey
	ParsePrivateKey(key interface{}) PrivateKey
}

var sharedPrivateKeyHelper PrivateKeyHelper = nil

func SetPrivateKeyHelper(helper PrivateKeyHelper) {
	sharedPrivateKeyHelper = helper
}

func GetPrivateKeyHelper() PrivateKeyHelper {
	return sharedPrivateKeyHelper
}
