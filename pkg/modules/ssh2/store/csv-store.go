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
package repository

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

const defaultCsvStoreFile = "~/sshconsole.csv"

// CsvStore ...
type CsvStore struct{}

// GetSessionByID ...
func (that *CsvStore) GetSessionByID(id int64) *SessionBean {
	rfile, _ := os.Open(defaultCsvStoreFile)
	r := csv.NewReader(rfile)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if len(row) < 5 {
			log.Println("can not format this row ,%V", row)
			continue
		}
		rowId, err := strconv.ParseInt(row[0], 10, 64)
		if rowId == id {
			session := new(SessionBean)
			session.ID = id
			session.Name = row[1]
			session.Address = row[2]
			session.Username = row[3]
			session.Password = row[4]
			session.SSHPrivateKey = row[5]
			return session
		}
	}
	return nil
}

// QuerySessionList ...
func (that *CsvStore) QuerySessionList() []SessionBean {
	// 通过切片存储
	sessions := make([]SessionBean, 0)
	rfile, _ := os.Open(defaultCsvStoreFile)
	r := csv.NewReader(rfile)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if len(row) < 5 {
			log.Println("can not format this row ,%V", row)
			continue
		}
		rowId, err := strconv.ParseInt(row[0], 10, 64)
		var session SessionBean
		session.ID = rowId
		session.Name = row[1]
		session.Address = row[2]
		session.Username = row[3]
		session.Password = row[4]
		session.SSHPrivateKey = row[5]
		sessions = append(sessions, session)
	}
	return sessions
}

// SaveSession ...
func (that *CsvStore) SaveSession(session *SessionBean) {
	file, _ := os.OpenFile(defaultCsvStoreFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	w := csv.NewWriter(file)
	id := strconv.FormatInt(session.ID, 10)
	_ = w.Write([]string{id, session.Name, session.Address, session.Username, session.Password, session.Password})
	w.Flush()
	_ = file.Close()
}

// DeleteSession ...
func (that *CsvStore) DeleteSession(sessionID int64) int {
	return 0
}
