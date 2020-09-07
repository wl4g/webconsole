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
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// TestUserBean ...
type TestUserBean struct {
	Name  string          `config:"name"`
	Age   int             `config:"age" default:"18"`
	Roles []*TestRoleBean `config:"roles"`
}

// TestUserBean ...
type TestRoleBean struct {
	Name  string `config:"name" default:"anon"`
	Alias string `config:"alias"`
}

// TestCustomerBean ...
type TestCustomerBean struct {
	Name  string          `config:"name" default:"MyCompany"`
	Type  int             `config:"type" default:"1"`
	Users []*TestUserBean `config:"users"`
}

// Reflect sets fields
func TestConfiguratorSample1(t *testing.T) {
	os.Setenv("MYPREFIX_NAME", "Company333")
	os.Setenv("MYPREFIX_TYPE", "3")

	// Create config all sample tmpfile.
	defineFullConfigFile := createSampleTmpFile("config.all.yaml",
		`Name: Company1
Type: 1
Users:
  - Name: jack
    Age: 50
    Roles:
      - Name:  administrator
        Alias: admin
`)

	// Create use config sample tmpfile.
	configFile := createSampleTmpFile("config.yaml",
		`Name: Company222
Type: 2
Users:
  - Name: jack
    Roles:
      - Name:  administrator
`)

	// Create configurator
	c := NewConfigurator(defineFullConfigFile, configFile)
	c.SetEnvPrefix("MYPREFIX") // Sets auto use env variables prefix

	// Load & parse
	customer := &TestCustomerBean{}
	if err := c.Parse(customer); err != nil {
		panic(err)
	}

	// Gets parsed configuration
	if json, err := ToJSONString(customer); err != nil {
		panic(err)
	} else {
		fmt.Println("Final load configuration: " + json)
	}

}

func createSampleTmpFile(filename string, content string) string {
	filename = "/tmp/" + filename
	f, _ := os.Create(filename)
	defer f.Close()
	err := ioutil.WriteFile(filename, []byte(content), os.ModeTemporary)
	if err != nil {
		panic(err)
	}
	return filename
}
