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
package store

import (
	"io"
	"xcloud-webconsole/pkg/logging"

	"go.uber.org/zap"
)

// SSH2Store ...
type SSH2Store interface {
	GetSessionByID(sessionID int64) *SessionBean
	QuerySessionList() []SessionBean
	SaveSession(session *SessionBean) int64
	DeleteSession(sessionID int64) int64
	io.Closer
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
	orgin *SSH2Store
}

var (
	singletonDelegate *DelegateSSH2Store
)

// GetDelegate Gets or create real store instance with orgin.
func GetDelegate() *DelegateSSH2Store {
	var err error
	if singletonDelegate == nil {
		singletonDelegate, err = newDelegateSSH2Store()
	}
	if err != nil {
		logging.Main.Panic("Unable get or create DelegateSSH2Store, cause by: %s",
			zap.String(err.Error(), ""))
		return nil
	}
	return singletonDelegate
}

func newDelegateSSH2Store() (*DelegateSSH2Store, error) {
	var orginStore SSH2Store
	switch 1 {
	case 1:
		if mysql, err1 := NewMysqlStore(); err1 == nil {
			orginStore = *mysql
		} else {
			return nil, err1
		}
	default:
		if csv, err2 := NewCsvStore(); err2 == nil {
			orginStore = *csv
		} else {
			return nil, err2
		}
	}
	return &DelegateSSH2Store{
		orgin: &orginStore,
	}, nil
}

// GetSessionByID ...
func (store *DelegateSSH2Store) GetSessionByID(sessionID int64) *SessionBean {
	return (*store.orgin).GetSessionByID(sessionID)
}

// QuerySessionList ...
func (store *DelegateSSH2Store) QuerySessionList() []SessionBean {
	return (*store.orgin).QuerySessionList()
}

// SaveSession ...
func (store *DelegateSSH2Store) SaveSession(session *SessionBean) int64 {
	return (*store.orgin).SaveSession(session)
}

// DeleteSession ...
func (store *DelegateSSH2Store) DeleteSession(sessionID int64) int64 {
	return (*store.orgin).DeleteSession(sessionID)
}

// Close ...
func (store *DelegateSSH2Store) Close() error {
	return (*store.orgin).Close()
}
