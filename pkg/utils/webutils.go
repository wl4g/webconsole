/**
 * Copyright 2017 ~ 2025 the original author or author<Wanglsir@gmail.com, 983708408@qq.com>.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package utils

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// IsSameWildcardOrigin ...
func IsSameWildcardOrigin(defWildcardURI string,
	requestURI string, checkScheme bool) bool {
	if defWildcardURI == "" || requestURI == "" {
		return false
	}
	if defWildcardURI == requestURI { // URL equaled?
		return true
	}

	// Scheme matched?
	schemeMatched := false
	uri1, err1 := url.Parse(defWildcardURI)
	uri2, err2 := url.Parse(requestURI)
	if err1 == nil && err2 == nil {
		schemeMatched = strings.EqualFold(uri1.Scheme, uri2.Scheme)
	}
	if checkScheme && !schemeMatched {
		return false
	}

	// Hostname equaled?
	hostname1, _ := ExtractWildcardEndpoint(defWildcardURI)
	hostname2, _ := ExtractWildcardEndpoint(requestURI)
	hostname1Len := len(hostname1)
	hostname2Len := len(hostname2)
	if strings.EqualFold(hostname1, hostname2) {
		return true
	}

	// Hostname wildcard matched?
	wildcardHostnameMatched := false
	parts1 := strings.Split(hostname1, ".")
	parts2 := strings.Split(hostname2, ".")
	parts1Len := len(parts1)
	parts2Len := len(parts2)
	for i := 0; i < parts1Len; i++ {
		if strings.EqualFold(parts1[i], AllAllow) {
			if i < (hostname1Len-1) && i < (hostname2Len-1) {
				compare1 := JoinString(parts1, ".", i+1, parts1Len)
				compare2 := JoinString(parts2, ".", i+1, parts2Len)
				if strings.EqualFold(compare1, compare2) {
					wildcardHostnameMatched = true
					break
				}
			}
		}
	}
	// Check scheme matched.
	if checkScheme && wildcardHostnameMatched {
		return schemeMatched
	}

	return wildcardHostnameMatched
}

/**
 * Extract domain text from {@link URI}.
 * Uri resolution cannot be used here because it may fail when there are
 * wildcards, e.g,
 * {@link URI#create}("http://*.aa.domain.com/api/v2/).gethost() is
 * null.
 *
 *  ExtractWildcardEndpoint("http://*.domain.com/v2/xx") => *.domain.com
 *  ExtractWildcardEndpoint("http://*.aa.domain.com:*") => *.aa.domain.com
 *  ExtractWildcardEndpoint("http://*.bb.domain.com:8080/v2/xx") => *.bb.domain.com
 *
 * @param wildcardURI
 * @return
 */
func ExtractWildcardEndpoint(wildcardURI string) (string, error) {
	if wildcardURI == "" {
		return "", nil
	}

	wildcardURI = strings.ToLower(strings.TrimSpace(SafeEncodeURL(wildcardURI)))
	noPrefix := wildcardURI[strings.Index(wildcardURI, URLSeparProto)+len(URLSeparProto) : len(wildcardURI)-1]
	slashIndex := strings.Index(noPrefix, URLSeparSlash)
	serverName := noPrefix
	if slashIndex > 0 {
		serverName = noPrefix[0:slashIndex]
	}

	// Check domain illegal?
	// e.g, http://*.domain.com:8080[allow]
	// http://*.domain.com:*[allow]
	// http://*.aa.*.domain.com[noallow]
	hostname := serverName
	if strings.Contains(serverName, URLSeparColon) {
		hostname = serverName[0:strings.Index(serverName, URLSeparColon)]
	}

	if strings.Index(hostname, AllAllow) != strings.LastIndex(hostname, AllAllow) {
		errmsg := fmt.Sprintf("Illegal serverName: %s, contains multiple wildcards!", serverName)
		return "", errors.New(errmsg)
	}

	return SafeDecodeURL(hostname)
}

// SafeEncodeURL URL encode by UTF-8
func SafeEncodeURL(requestURL string) string {
	lowerURL := strings.ToLower(strings.TrimSpace(requestURL))
	if !strings.Contains(lowerURL, URLSeparSlash) &&
		!strings.Contains(lowerURL, URLSeparQuest) &&
		!strings.Contains(lowerURL, URLSeparColon) {
		return url.QueryEscape(requestURL)
	}
	return requestURL
}

// SafeDecodeURL URL decode by UTF-8
func SafeDecodeURL(requestURL string) (string, error) {
	lowerURL := strings.ToLower(strings.TrimSpace(requestURL))
	if strings.Contains(lowerURL, URLSeparSlash) ||
		strings.Contains(lowerURL, URLSeparQuest) ||
		strings.Contains(lowerURL, URLSeparColon) {
		return url.QueryUnescape(requestURL)
	}
	return requestURL, nil
}

const (
	// AllAllow ...
	AllAllow = "*"

	// URLSchemeHTTPS URL scheme(HTTPS)
	URLSchemeHTTPS = "https"

	// URLSchemeHTTP URL scheme(HTTP)
	URLSchemeHTTP = "http"

	// URLSeparSlash URL separator(/)
	URLSeparSlash = "%2f"

	//URLSeparSlash2 URL double separator(//)
	URLSeparSlash2 = URLSeparSlash + URLSeparSlash

	//URLSeparQuest URL separator(?)
	URLSeparQuest = "%3f"

	//URLSeparColon URL colon separator(:)
	URLSeparColon = "%3a"

	//URLSeparProto Protocol separators, such as
	// https://my.domain.com=>https%3A%2F%2Fmy.domain.com
	URLSeparProto = URLSeparColon + URLSeparSlash + URLSeparSlash
)
