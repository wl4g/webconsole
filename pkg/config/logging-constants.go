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

const (
	// -------------------------------
	// Log constants.
	// -------------------------------

	// DefaultLogMain Default log service 'main'.
	DefaultLogMain = "main"

	// DefaultLogReceive Default log service 'receive'.
	DefaultLogReceive = "receive"

	// DefaultLogDir Default log directory.
	DefaultLogDir = "./log/"

	// DefaultLogLevel Default log level
	DefaultLogLevel = "INFO"

	// DefaultLogRetentionDays Default log retention days.
	DefaultLogRetentionDays = 7

	// DefaultLogMaxBackups Default log max backup numbers.
	DefaultLogMaxBackups = 30

	// DefaultLogMaxSize Default log max size(MB).
	DefaultLogMaxSize = 128
)
