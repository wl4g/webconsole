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

import (
	"fmt"
	"io/ioutil"

	tools "xcloud-webconsole/pkg/utils"

	"gopkg.in/yaml.v2"
)

// GlobalProperties ...
type GlobalProperties struct {
	Admin      AdminProperties      `yaml:"admin"`
	Server     ServerProperties     `yaml:"server"`
	DataSource DataSourceProperties `yaml:"datasource"`
	Logging    LoggingProperties    `yaml:"logging"`
}

var (
	// GlobalConfig ...
	GlobalConfig GlobalProperties
)

// InitGlobalConfig global config properties.
func InitGlobalConfig(path string) {
	// Create default config.
	GlobalConfig = *createDefaultProperties()

	conf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Read config '%s' error! %s", path, err)
		panic(err)
	}

	err = yaml.Unmarshal(conf, &GlobalConfig)
	if err != nil {
		fmt.Printf("Unmarshal config '%s' error! %s", path, err)
		panic(err)
	}

	// Post properties.
	afterPropertiesSet(&GlobalConfig)
}

// Create default configuration properties.
func createDefaultProperties() *GlobalProperties {
	globalConfig := &GlobalProperties{
		Admin: AdminProperties{
			Listen: DefaultAdminServeListen,
		},
		Server: ServerProperties{
			Listen: DefaultServeListen,
			Cors: CorsProperties{
				AllowOrigins:     DefaultCorsAllowOrigins,
				AllowCredentials: DefaultCorsAllowCredentials,
				AllowMethods:     DefaultCorsAllowMethods,
				AllowHeaders:     DefaultCorsAllowHeaders,
				ExposeHeaders:    DefaultCorsExposeHeaders,
				MaxAge:           DefaultCorsMaxAge,
			},
		},
		DataSource: DataSourceProperties{
			Mysql: MysqlProperties{
				DbConnectStr:    DefaultMysqlConnectStr,
				MaxOpenConns:    DefaultMysqlMaxOpenConns,
				MaxIdleConns:    DefaultMysqlMaxIdleConns,
				ConnMaxLifetime: DefaultMysqlConnMaxLifetime,
			},
			Csv: CsvProperties{},
		},
		Logging: LoggingProperties{
			DateFormatPattern: DefaultLogDateFormatPattern,
			LogItems: map[string]LogItemProperties{
				DefaultLogMain: {
					FileName: DefaultLogDir + DefaultLogMain + ".log",
					Level:    DefaultLogLevel,
					Policy: PolicyProperties{
						RetentionDays: DefaultLogRetentionDays,
						MaxBackups:    DefaultLogMaxBackups,
						MaxSize:       DefaultLogMaxSize,
					},
				},
				DefaultLogReceive: {
					FileName: DefaultLogDir + DefaultLogReceive + ".log",
					Level:    DefaultLogLevel,
					Policy: PolicyProperties{
						RetentionDays: DefaultLogRetentionDays,
						MaxBackups:    DefaultLogMaxBackups,
						MaxSize:       DefaultLogMaxSize,
					},
				},
			},
		},
	}
	return globalConfig
}

// MetricExclude settings after initialization
func afterPropertiesSet(globalConfig *GlobalProperties) {
}

// RefreshConfig Refresh global config.
func RefreshConfig(config *GlobalProperties) {
	tools.CopyProperties(&config, &GlobalConfig)
}
