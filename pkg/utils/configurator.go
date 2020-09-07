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
package utils

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// FieldProcessor ...
type FieldProcessor func(fieldInfo reflect.StructField)

// Configurator ...
type Configurator struct {
	vp *viper.Viper
}

// NewConfigurator ...
func NewConfigurator(defineFullConfigFile string, configFile string) *Configurator {
	if defineFullConfigFile == "" || configFile == "" {
		panic(errors.New("DefineFullCofnigFile and configFile must not is empty"))
	}

	vp := viper.New()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.AllowEmptyEnv(false)
	vp.SetConfigFile(defineFullConfigFile)

	// Enable auto population from environment variables (Priority: Set()/Flags/Config/Env/Default)
	// @see https://github.com/spf13/viper#working-with-environment-variables
	vp.AutomaticEnv()

	// Note: some strange logic problems are found in the test. After automaticenv() is enabled,
	// the parsing is based on the key defined in the configuration file??? For example: environment
	// variable: DATASOURCE.URI=uri111 (the default correct rule) must be configured at the same time
	// config.yml There are also: datasource.uri=uri222 In order to get the expected value of uri111,
	// otherwise, if the configuration file config.yml There is no definition datasource.uri You will not
	// be able to get the correct value of uri111 https://github.com/spf13/viper/blob/master/viper.go#L1904

	// Load parse the real runtime configuration file
	cf, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	vp.MergeConfig(cf)
	cf.Close()

	return &Configurator{vp: vp}
}

// SetEnvPrefix ...
func (c *Configurator) SetEnvPrefix(envPrefix string) {
	c.vp.SetEnvPrefix(envPrefix)
}

// Parse ...
func (c *Configurator) Parse(config interface{}) error {
	if err := c.vp.ReadInConfig(); err != nil {
		return err
	}
	if err := c.vp.Unmarshal(config); err != nil {
		return err
	}
	return nil
}

// GetViper ...
func (c *Configurator) GetViper() *viper.Viper {
	return c.vp
}
