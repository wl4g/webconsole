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

// SSH2Store ...
type SSH2Store interface {
	GetSessionByID(sessionID int64) *SessionBean
	QuerySessionList() []SessionBean
	SaveSession(session *SessionBean) int64
	DeleteSession(sessionID int64) int
	Destroy() error
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
// --- Delegate Store. ---
//

// DelegateSSH2Store ...
type DelegateSSH2Store struct {
	// SSH2Store
	mysql MysqlStore
	csv   CsvStore
}

// GetDelegateSSH2Store ...
func GetDelegateSSH2Store() *DelegateSSH2Store {
	return &DelegateSSH2Store{}
}

// GetSessionByID ...
func (store *DelegateSSH2Store) GetSessionByID(sessionID int64) *SessionBean {
	return store.getOrginStore().GetSessionByID(sessionID)
}

// QuerySessionList ...
func (store *DelegateSSH2Store) QuerySessionList() []SessionBean {
	return store.getOrginStore().QuerySessionList()
}

// SaveSession ...
func (store *DelegateSSH2Store) SaveSession(session *SessionBean) int64 {
	return store.getOrginStore().SaveSession(session)
}

// DeleteSession ...
func (store *DelegateSSH2Store) DeleteSession(sessionID int64) int {
	return store.getOrginStore().DeleteSession(sessionID)
}

// Destroy ...
func (store *DelegateSSH2Store) Destroy() {
	store.getOrginStore().Destroy()
}

// Gets orgin real store instance.
func (store *DelegateSSH2Store) getOrginStore() SSH2Store {
	// TODO
	switch 1 {
	case 1:
		// return store.mysql
		return nil
	default:
		// return store.csv
		return nil
	}
}
