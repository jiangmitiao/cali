package rcali

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"io"
	"sync"
	"time"
)

var mux = sync.Mutex{}
var loginSessionMap = make(map[string]string)
var loginSessionTime = make(map[string]time.Time)
var timeOut int64 = 60 * 60 * 2

func init() {
	loginSessionMap = make(map[string]string)
	loginSessionTime = make(map[string]time.Time)
	timeOut = 60 * 60 * 2

	go func() {
		for {
			time.Sleep(time.Second * 30)
			freshMap()
		}
	}()
}

func SetLoginUser(loginSession string, userId string) {
	loginSessionMap[loginSession] = userId
	loginSessionTime[loginSession] = time.Now()
}

func GetUserIdByLoginSession(loginSession string) (string, time.Time) {
	return loginSessionMap[loginSession], loginSessionTime[loginSession]
}

func DeleteLoginSession(loginSession string) {
	mux.Lock()
	delete(loginSessionMap, loginSession)
	delete(loginSessionTime, loginSession)
	mux.Unlock()
}

func DeleteLoginUserId(userId string)  {
	mux.Lock()
	for loginSession, v := range loginSessionMap {
		if v==userId {
			delete(loginSessionMap, loginSession)
			delete(loginSessionTime, loginSession)
			break
		}
	}
	mux.Unlock()
}

func freshMap() {
	var timeNow = time.Now()
	mux.Lock()
	for k, v := range loginSessionTime {
		if timeNow.Unix()-v.Unix() > timeOut {
			DEBUG.Debug("delete session ", k)
			delete(loginSessionMap, k)
			delete(loginSessionTime, k)
		}
	}
	mux.Unlock()
}

func Sha3_256(in string) string {
	m := sha3.New256()
	io.WriteString(m, in)
	return hex.EncodeToString(m.Sum(nil))
}
