/**
 * Copyright 2017 ~ 2025 the original author or authors<Wanglsir@gmail.com, 983708408@qq.com>.
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
package config

// ---------------------
// Web Console server properties
// ---------------------

// ServerProperties ...
type ServerProperties struct {
	Listen   string             `yaml:"listen"`
	Cors     CorsProperties     `yaml:"cors"`
	SSH2Term SSH2TermProperties `yaml:"ssh2-term"`
}

// CorsProperties ...
type CorsProperties struct {
	AllowOrigins     string `"yaml:"allow-origins"`
	AllowCredentials bool   `"yaml:"allow-credentials"`
	AllowMethods     string `"yaml:"allow-methods"`
	AllowHeaders     string `"yaml:"allow-headers"`
	ExposeHeaders    string `"yaml:"expose-headers"`
	MaxAge           int    `"yaml:"max-age"` // Seconds
}

const (
	// -------------------------------
	// WebConsole server constants.
	// -------------------------------

	// DefaultServeListen ...
	DefaultServeListen = ":16088"

	// DefaultCorsAllowOrigins ...
	DefaultCorsAllowOrigins = "http://localhost:16088,https://webconsole.wl4g.debug,https://webconsole.wl4g.com,"

	// DefaultCorsAllowCredentials ...
	DefaultCorsAllowCredentials = false

	// DefaultCorsAllowMethods ...
	DefaultCorsAllowMethods = "GET,POST,OPTIONS,PUT,DELETE,UPDATE"

	// DefaultCorsAllowHeaders ...
	DefaultCorsAllowHeaders = "Authorization,Content-Length,X-CSRF-Token,Token,session,X_Requested_With,Accept,Origin,Host,Connection,Accept-Encoding,Accept-Language,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Pragma"

	// DefaultCorsExposeHeaders ...
	DefaultCorsExposeHeaders = "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma"

	// DefaultCorsMaxAge ...
	DefaultCorsMaxAge = 172800
)
