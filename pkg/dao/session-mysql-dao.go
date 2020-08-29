/**
 * Copyright 2017 ~ 2035 the original author or authors<Wanglsir@gmail.com, 983708408@qq.com>.
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
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "gzsm"
	PASS_WORD = "gzsm@%#jh?"
	// HOST      = "10.0.0.160"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DATABASE = "devops_dev"
	CHARSET  = "utf8"
)

// 初始化链接
func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}

// Session User sessions bean
type Session struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	Username string `db:"username"`
	Password string `db:"password"`
	SshKey   string `db:"ssh_key"`
}

// GetSessionById find session info by id
func GetSessionById(id int64) *Session {
	session := new(Session)
	row := MysqlDb.QueryRow("select id,name,address,username,password,IFNULL(ssh_key, '') from webconsole_session where id=?", id)
	if err := row.Scan(&session.Id, &session.Name, &session.Address, &session.Username, &session.Password, &session.SshKey); err != nil {
		log.Fatal("GetSessionById", err)
	}
	fmt.Println(session.Id, session.Name, session.Username)
	return session
}

// 查询Session列表
func SessionList() []Session {

	// 通过切片存储
	sessions := make([]Session, 0)
	rows, _ := MysqlDb.Query("SELECT * FROM `webconsole_session` limit ?", 100)

	// 遍历
	var session Session
	for rows.Next() {
		rows.Scan(&session.Id, &session.Name, &session.Address, &session.Username, &session.Password, &session.SshKey)
		sessions = append(sessions, session)
	}

	fmt.Println(sessions)
	return sessions
}

// 插入数据
func InsertSession(session *Session) int64 {
	ret, e := MysqlDb.Exec("insert INTO webconsole_session(name,address,username,password,ssh_key) values(?,?,?,?,?)", session.Name, session.Address, session.Username, session.Password, session.SshKey)
	if nil != e {
		log.Print("add fail", e)
		return 0
	}
	//影响行数
	rowsaffected, _ := ret.RowsAffected()
	id, _ := ret.LastInsertId()
	log.Printf("RowsAffected: %d", rowsaffected)
	return id
}

func DelSession(id int64) {
	result, _ := MysqlDb.Exec("delete from webconsole_session where id=?", id)
	rowsaffected, _ := result.RowsAffected()
	log.Printf("RowsAffected: %d", rowsaffected)
}
