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
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"xcloud-webconsole/pkg/config"
	utils "xcloud-webconsole/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlRepository ...
type MysqlRepository struct {
	mysqlDB *sql.DB
}

var repo *MysqlRepository

// InitMysqlRepository ...
func (that *MysqlRepository) InitMysqlRepository() {
	mysqlConfig := config.GlobalConfig.Repository.Mysql
	log.Print("Connecting to MySQL of configuration: " + utils.ToJSONString(mysqlConfig))

	mydb, dberr := utils.OpenMysqlConnection(
		mysqlConfig.DbConnectStr,
		mysqlConfig.MaxOpenConns,
		mysqlConfig.MaxIdleConns,
		mysqlConfig.ConnMaxLifetime,
	)
	that.mysqlDB = *mydb
	//defer mysqlDB.Close();

	if dberr != nil {
		panic("Cannot connect to mysql. " + dberr.Error())
	}
}

// GetSessionByID find session info by id
func (that *MysqlRepository) GetSessionByID(id int64) *SessionBean {
	session := new(SessionBean)
	row := that.mysqlDB.QueryRow("select id,name,address,username,password,IFNULL(ssh_key, '') from webconsole_session where id=?", id)
	if err := row.Scan(&session.ID, &session.Name, &session.Address, &session.Username, &session.Password, &session.SSHPrivateKey); err != nil {
		log.Fatal("GetSessionById", err)
	}
	fmt.Println(session.ID, session.Name, session.Username)
	return session
}

// QuerySessionList ...
func (that *MysqlRepository) QuerySessionList() []SessionBean {
	// 通过切片存储
	sessions := make([]SessionBean, 0)
	rows, _ := that.mysqlDB.Query("SELECT * FROM `webconsole_session` limit ?", 100)

	// 遍历
	var session SessionBean
	for rows.Next() {
		rows.Scan(&session.ID, &session.Name, &session.Address, &session.Username, &session.Password, &session.SSHPrivateKey)
		sessions = append(sessions, session)
	}

	fmt.Println(sessions)
	return sessions
}

// SaveSession ...
func (that *MysqlRepository) SaveSession(session *SessionBean) int64 {
	ret, e := that.mysqlDB.Exec("insert INTO webconsole_session(name,address,username,password,ssh_key) values(?,?,?,?,?)", session.Name, session.Address, session.Username, session.Password, session.SSHPrivateKey)
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

// DeleteSession ...
func (that *MysqlRepository) DeleteSession(ID int64) {
	result, _ := that.mysqlDB.Exec("delete from webconsole_session where id=?", ID)
	rowsaffected, _ := result.RowsAffected()
	log.Printf("RowsAffected: %d", rowsaffected)
}
