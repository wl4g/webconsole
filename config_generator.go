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
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Gets config content template
	contentTpl := getConfigContent("WEBCONSOLE.DEFAULT.YML.GO.TPL")
	// Gets default config content
	defaultConfigContent := getConfigContent("resources/webconsole.default.yml")

	// Replace out to webconsole.default.yml.go
	goContent := strings.ReplaceAll(contentTpl, "{CONTENT}", defaultConfigContent)

	if err := ioutil.WriteFile("pkg/config/webconsole.default.yml.go", []byte(goContent), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Successfully for pkg/config/webconsole.default.yml.go")
}

func getConfigContent(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		panic(err2)
	}

	return string(data)
}
