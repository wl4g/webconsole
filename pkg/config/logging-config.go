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
// Logging properties
// ---------------------

// LoggingProperties Multi log properties.
type LoggingProperties struct {
	DateFormatPattern string                       `yaml:"dateformat-pattern"`
	LogItems          map[string]LogItemProperties `yaml:"log-items"`
}

// LogItemProperties LogItem properties.
type LogItemProperties struct {
	FileName string           `yaml:"file"`
	Level    string           `yaml:"level"`
	Policy   PolicyProperties `yaml:"policy"`
}

// PolicyProperties Logging archive policy.
type PolicyProperties struct {
	RetentionDays int `yaml:"retention-days"`
	MaxBackups    int `yaml:"max-backups"`
	MaxSize       int `yaml:"max-size"`
}

const (
	// -------------------------------
	// Log constants.
	// -------------------------------

	// DefaultLogMain Default log service 'main'.
	DefaultLogMain = "main"

	// DefaultLogReceive Default log service 'receive'.
	DefaultLogReceive = "receive"

	// DefaultLogDir Default log directory.
	DefaultLogDir = "/mnt/disk1/log/webconsole/"

	// DefaultLogDateFormatPattern Default log date format pattern
	DefaultLogDateFormatPattern = "06-01-02 15:04:05"

	// DefaultLogLevel Default log level
	DefaultLogLevel = "INFO"

	// DefaultLogRetentionDays Default log retention days.
	DefaultLogRetentionDays = 30

	// DefaultLogMaxBackups Default log max backup numbers.
	DefaultLogMaxBackups = 30

	// DefaultLogMaxSize Default log max size(MB).
	DefaultLogMaxSize = 512
)
