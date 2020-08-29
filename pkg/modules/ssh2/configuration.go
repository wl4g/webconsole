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
package ssh2

import (
	"io/ioutil"
	"log"
	"time"
)

const (
	// DefaultSSH2APIBaseURI ssh2 console base URI.
	DefaultSSH2APIBaseURI = "/terminal/"
	// DefaultSSH2APIWebSocketURI ssh2 console websocket connection URI.
	DefaultSSH2APIWebSocketURI = DefaultSSH2APIBaseURI + "ws/:id"
	// DefaultSSH2APISessionQueryURI ssh2 console session query URI.
	DefaultSSH2APISessionQueryURI = DefaultSSH2APIBaseURI + "session/list"
	// DefaultSSH2APISessionAddURI ssh2 console session create URI.
	DefaultSSH2APISessionAddURI = DefaultSSH2APIBaseURI + "session/create"
	// DefaultSSH2APISessionDeleteURI ssh2 console session delete URI.
	DefaultSSH2APISessionDeleteURI = DefaultSSH2APIBaseURI + "session/delete"
)

const (
	TermLinux         = "linux"
	TermAnsi          = "ansi"
	TermScoAnsi       = "scoansi"
	TermXterm         = "xterm"
	TermXterm256Color = "xterm-256color"
	TermVt100         = "vt100"
	TermVt102         = "vt102"
	TermVt220         = "vt220"
	TermVt320         = "vt320"
	TermWyse50        = "wyse50"
	TermWyse60        = "wyse60"
	TermDumb          = "dumb"
)

var (
	//DefaultTerm ...
	DefaultTerm = TermXterm
	//DefaultConnTimeout ...
	DefaultConnTimeout = 15 * time.Second
	//DefaultLogger ...
	DefaultLogger = log.New(ioutil.Discard, "[webssh] ", log.Ltime|log.Ldate)
	//DefaultBuffSize ...
	DefaultBuffSize = uint32(8192)
)
