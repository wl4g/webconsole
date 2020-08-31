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
	"context"
	"flag"
	"fmt"
)

const (
	defaultConfigPath = "/etc/webconsole.yml"
)

var (
	webConsole = &WebConsole{}
)

func main() {
	conf := defaultConfigPath
	// Parsing configuration path
	flag.StringVar(&conf, "c", defaultConfigPath, "WebConsole config path.")
	flag.Usage()
	flag.Parse()
	fmt.Printf("Initializing config path for '%s'\n", conf)

	// Start server...
	webConsole.StartServe(context.Background(), conf)
}
