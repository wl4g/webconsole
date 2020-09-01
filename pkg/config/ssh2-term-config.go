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

// SSH2TermProperties ...
type SSH2TermProperties struct {
	// 可设置 ansi/linux/vt100/xterm/dumb, 除dumb外其他都有颜色显示, 默认xterm
	PtyTermType             string `yaml:"pty-term-type"`
	PtyTermConnTimeoutSec   uint32 `yaml:"pty-term-conn-timeout-sec"` // Seconds
	PtyWSTransferBufferSize uint32 `yaml:"pty-ws-transfer-buffer-size"`
}

const (
	// DefaultPtyTermType ...
	DefaultPtyTermType = "xterm"

	// DefaultPtyTermConnTimeoutSec ...
	DefaultPtyTermConnTimeoutSec = uint32(5) // Seconds

	// DefaultPtyWSTransferBufferSize ...
	DefaultPtyWSTransferBufferSize = uint32(8192)
)
