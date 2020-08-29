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

// SSH2Repository ...
type SSH2Repository interface {
	GetSessionByID(sessionID int64) *SessionBean
	QuerySessionList() []SessionBean
	SaveSession(session *SessionBean) int64
	DeleteSession(sessionID int64) int
}

// SessionBean User terminal session info bean
type SessionBean struct {
	ID            int64  `db:"id"`
	Name          string `db:"name"`
	Address       string `db:"address"`
	Username      string `db:"username"`
	Password      string `db:"password"`
	SSHPrivateKey string `db:"ssh_key"`
}

//
// --- Delegate repository. ---
//

// DelegateSSH2Repository ...
type DelegateSSH2Repository struct {
	// SSH2Repository
	mysql MysqlRepository
	csv   CsvRepository
}

// GetDelegateSSH2Repository ...
func GetDelegateSSH2Repository() *DelegateSSH2Repository {
	return &DelegateSSH2Repository{}
}

// GetSessionByID ...
func (repo *DelegateSSH2Repository) GetSessionByID(sessionID int64) *SessionBean {
	return repo.getOrginRepository().GetSessionByID(sessionID)
}

// QuerySessionList ...
func (repo *DelegateSSH2Repository) QuerySessionList() []SessionBean {
	return repo.getOrginRepository().QuerySessionList()
}

// SaveSession ...
func (repo *DelegateSSH2Repository) SaveSession(session *SessionBean) int64 {
	return repo.getOrginRepository().SaveSession(session)
}

// DeleteSession ...
func (repo *DelegateSSH2Repository) DeleteSession(sessionID int64) int {
	return repo.getOrginRepository().DeleteSession(sessionID)
}

// Gets orgin real repository instance.
func (repo *DelegateSSH2Repository) getOrginRepository() SSH2Repository {
	switch 1 {
	case 1:
		// return repo.mysql
		return nil
	default:
		// return repo.csv
		return nil
	}
}
