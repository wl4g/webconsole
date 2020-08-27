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

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

const db_file = "~/sshconsole.csv";

// 根据id找session
func GetSessionByIdCsv(id int64) *Session {
	rfile, _ := os.Open(db_file)
	r := csv.NewReader(rfile)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if len(row)<5{
			log.Println("can not format this row ,%V",row)
			continue
		}
		rowId, err := strconv.ParseInt(row[0], 10, 64)
		if rowId == id{
			session := new(Session)
			session.Id = id
			session.Name = row[1]
			session.Address = row[2]
			session.Username = row[3]
			session.Password = row[4]
			session.SshKey = row[5]
			return session
		}
	}
	return nil;
}

// 查询Session列表
func SessionListCsv() []Session {

	// 通过切片存储
	sessions := make([]Session, 0)

	rfile, _ := os.Open(db_file)
	r := csv.NewReader(rfile)

	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if len(row)<5{
			log.Println("can not format this row ,%V",row)
			continue
		}
		rowId, err := strconv.ParseInt(row[0], 10, 64)
		var session Session
		session.Id = rowId
		session.Name = row[1]
		session.Address = row[2]
		session.Username = row[3]
		session.Password = row[4]
		session.SshKey = row[5]
		sessions=append(sessions,session)
	}
	return sessions;
}

// 插入数据
func InsertSessionCsv(session *Session) {
	file, _ := os.OpenFile(db_file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	w := csv.NewWriter(file)
	id := strconv.FormatInt(session.Id,10)
	_ = w.Write([]string{id, session.Name, session.Address, session.Username, session.Password, session.Password})
	w.Flush()
	_ = file.Close()
}

