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
// Repository(datasource) DAO properties
// ---------------------

// DataSourceProperties ...
type DataSourceProperties struct {
	Mysql MysqlProperties `yaml:"mysql"`
	Csv   CsvProperties   `yaml:"csv"`
}

// MysqlProperties ...
type MysqlProperties struct {
	// e.g: user:password@tcp(host:port)/database?charset=utf-8
	DbConnectStr       string `yaml:"dbconnectstr"`
	MaxOpenConns       int    `yaml:"max-open-conns"`
	MaxIdleConns       int    `yaml:"max-idle-conns"`
	ConnMaxLifetimeSec uint32 `yaml:"conn-max-lifetime-sec"` // Seconds
}

// CsvProperties ...
type CsvProperties struct {
	DataDir string `yaml:"data-dir"`
}

const (
	// DefaultCsvDataFile ...
	DefaultCsvDataFile = "webconsole.db.csv"
)
