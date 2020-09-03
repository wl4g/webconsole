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
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GinCorsHolder CORS enhanced processor based on gin framework,
// such as: support for https://*.console.example.com Wildcard configuration.
type GinCorsHolder struct {
	AllowOrigins     []string `yaml:"allow-origins"`
	AllowCredentials bool     `yaml:"allow-credentials"`
	AllowMethods     []string `yaml:"allow-methods"`
	AllowHeaders     []string `yaml:"allow-headers"`
	ExposeHeaders    []string `yaml:"exposes-headers"`
	MaxAge           int      `yaml:"max-age"` // Seconds
}

// RegisterCorsProcessor ...
func (holder *GinCorsHolder) RegisterCorsProcessor(engine *gin.Engine) {
	engine.Use(holder.createCorsHandlerFunc())
}

// createCorsHandlerFunc ...
func (holder *GinCorsHolder) createCorsHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerNames []string
		for k := range c.Request.Header {
			headerNames = append(headerNames, k)
		}

		// headerNamesString := strings.Join(headerNames, ", ")
		// if headerNamesString != "" {
		// 	headerNamesString = fmt.Sprintf("Access-Control-Allow-Origin, Access-Control-Allow-Headers, %s", headerNamesString)
		// } else {
		// 	headerNamesString = "Access-Control-Allow-Origin, Access-Control-Allow-Headers"
		// }

		// Unconditional pass
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Cors OPTIONS Request")
		}

		// Sets default access control policy for CORS requests
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", holder.matchCorsOrigin(origin))
			c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(holder.AllowCredentials))
			c.Header("Access-Control-Allow-Methods", holder.matchCorsMethod(method))
			c.Header("Access-Control-Allow-Headers", holder.matchCorsHeaders(headerNames))
			c.Header("Access-Control-Expose-Headers", JoinAll(holder.ExposeHeaders, ",")) // 跨域关键设置让浏览器可以解析
			c.Header("Access-Control-Max-Age", strconv.Itoa(holder.MaxAge))
		}

		// Execute the next handler.
		c.Next()
	}
}

// matchCorsOrigin ...
func (holder *GinCorsHolder) matchCorsOrigin(requestOrigin string) string {
	if requestOrigin == "" {
		return ""
	}
	if holder.AllowOrigins == nil || len(holder.AllowOrigins) <= 0 {
		return ""
	}

	if StringsContains(holder.AllowOrigins, AllAllow) {
		/**
		 * Note: Chrome will prompt: </br>
		 * The value of the 'Access-Control-Allow-Origin' header in the
		 * response must not be the wildcard '*' when the request's
		 * credentials mode is 'include'. The credentials mode of
		 * requests initiated by the XMLHttpRequest is controlled by the
		 * withCredentials attribute.
		 */
		if !holder.AllowCredentials {
			return AllAllow // Rejected
		} else {
			return requestOrigin
		}
	}

	for _, allowedOrigin := range holder.AllowOrigins {
		if strings.EqualFold(requestOrigin, allowedOrigin) {
			return requestOrigin
		}
		// e.g: allowedOrigin => "http://*.aa.mydomain.com"
		if IsSameWildcardOrigin(allowedOrigin, requestOrigin, true) {
			return requestOrigin
		}
	}
	return ""
}

// matchCorsHeaders ...
func (holder *GinCorsHolder) matchCorsHeaders(requestHeaders []string) string {
	// if (isNull(requestHeaders)) {
	// 	return null;
	// }
	// if (requestHeaders.isEmpty()) {
	// 	return Collections.emptyList();
	// }
	// if (ObjectUtils.isEmpty(allowedHeaders)) {
	// 	return null;
	// }

	// boolean allowAnyHeader = allowedHeaders.contains(ALL);
	// List<String> result = new ArrayList<String>(requestHeaders.size());
	// for (String requestHeader : requestHeaders) {
	// 	if (StringUtils.hasText(requestHeader)) {
	// 		requestHeader = requestHeader.trim();
	// 		if (allowAnyHeader) {
	// 			result.add(requestHeader);
	// 		} else {
	// 			for (String allowedHeader : allowedHeaders) {
	// 				// e.g: allowedHeader => "X-Iam-*"
	// 				if (allowedHeader.contains(ALL)) {
	// 					String allowedHeaderPrefix = allowedHeader.substring(allowedHeader.indexOf(ALL) + 1);
	// 					if (startsWithIgnoreCase(requestHeader, allowedHeaderPrefix)) {
	// 						result.add(requestHeader);
	// 						break;
	// 					}
	// 				} else if (requestHeader.equalsIgnoreCase(allowedHeader)) {
	// 					result.add(requestHeader);
	// 					break;
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// return (result.isEmpty() ? null : result);

	return JoinAll(requestHeaders, ",")
}

// matchCorsMethod ...
func (holder *GinCorsHolder) matchCorsMethod(requestMethod string) string {
	if requestMethod == "" {
		return ""
	}
	if holder.AllowMethods == nil || len(holder.AllowMethods) <= 0 {
		return requestMethod
	}
	if StringsContains(holder.AllowMethods, requestMethod) {
		return JoinAll(holder.AllowMethods, ",")
	}
	return ""
}
