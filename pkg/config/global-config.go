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

	"xcloud-webconsole/pkg/utils"
	"xcloud-webconsole/pkg/config"
	"gopkg.in/yaml.v2"
)

// GlobalProperties ...
type GlobalProperties struct {
	Server  ServerProperties  `yaml:"server"`
	Logging LoggingProperties `yaml:"logging"`
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
		Server: {
			Listen: constant.DefaultServeListen,
			Cors:{
				AllowOrigin      :constant.DefaultCorsOrigin,
				AllowCredentials :constant.DefaultCorsCredentials,
				AllowMethods     :constant.DefaultCorsMethods,
				AllowHeaders     :constant.DefaultCorsHeaders,
				ExposeHeaders    :constant.DefaultCorsExposeHeaders,
				MaxAge           :constant.DefaultCorsMaxAge,
			}
		},
		Logging: LoggingProperties{
			LogItems: map[string]LogItemProperties{
				constant.DefaultLogMain: {
					FileName: constant.DefaultLogDir + constant.DefaultLogMain + ".log",
					Level:    constant.DefaultLogLevel,
					Policy: PolicyProperties{
						RetentionDays: constant.DefaultLogRetentionDays,
						MaxBackups:    constant.DefaultLogMaxBackups,
						MaxSize:       constant.DefaultLogMaxSize,
					},
				},
				constant.DefaultLogReceive: {
					FileName: constant.DefaultLogDir + constant.DefaultLogReceive + ".log",
					Level:    constant.DefaultLogLevel,
					Policy: PolicyProperties{
						RetentionDays: constant.DefaultLogRetentionDays,
						MaxBackups:    constant.DefaultLogMaxBackups,
						MaxSize:       constant.DefaultLogMaxSize,
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
	common.CopyProperties(&config, &GlobalConfig)
}
