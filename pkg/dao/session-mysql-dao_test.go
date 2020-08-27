/**
 * Copyright 2017 ~ 2025 the original author or authors.
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
package dao

import "testing"

func TestSessionSelect(t *testing.T) {
	SessionList();
}

func TestGetSessionById(t *testing.T) {
	GetSessionById(2);
}

func TestInsertSession(t *testing.T) {
	session := new(Session)
	session.Name = "test1";
	session.Address = "10.0.0.160:30022";
	session.Username = "sshconsole";
	session.Password = "123456";
	session.SshKey = "";
	InsertSession(session)
}
