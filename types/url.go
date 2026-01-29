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
package types

import "net/url"

type URL interface {

	// String reassembles the [URL] into a valid URL string.
	// The general form of the result is one of:
	//
	//	scheme:opaque?query#fragment
	//	scheme://userinfo@host/path?query#fragment
	//
	// If u.Opaque is non-empty, String uses the first form;
	// otherwise it uses the second form.
	// Any non-ASCII characters in host are escaped.
	// To obtain the path, String uses u.EscapedPath().
	//
	// In the second form, the following rules apply:
	//   - if u.Scheme is empty, scheme: is omitted.
	//   - if u.User is nil, userinfo@ is omitted.
	//   - if u.Host is empty, host/ is omitted.
	//   - if u.Scheme and u.Host are empty and u.User is nil,
	//     the entire scheme://userinfo@host/ is omitted.
	//   - if u.Host is non-empty and u.Path begins with a /,
	//     the form host/path does not add its own /.
	//   - if u.RawQuery is empty, ?query is omitted.
	//   - if u.Fragment is empty, #fragment is omitted.
	String() string

	// Hostname returns u.Host, stripping any valid port number if present.
	//
	// If the result is enclosed in square brackets, as literal IPv6 addresses are,
	// the square brackets are removed from the result.
	Hostname() string

	// Port returns the port part of u.Host, without the leading colon.
	//
	// If u.Host doesn't contain a valid numeric port, Port returns an empty string.
	Port() string
}

func ParseURL(rawURL string) URL {
	uri, err := url.Parse(rawURL)
	if err != nil {
		//panic(err)
		return nil
	}
	return uri
}
