package session

import (
	"sync"
	"time"
	"github.com/eLearning/server/api/dbops"
	"github.com/eLearning/server/api/defs"
	"github.com/eLearning/server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMill() int64 {
	return time.Now().Unix()/1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

//将session从DB读取到缓存中
func LoadSessionsFromDB() {
	r, err := dbops.RetreiverAllSessions()
	if err != nil {
		return
	}

	//Range(func(k, v interface{})是什么
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})

}

//生成一个session
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMill()
	ttl := ct + 30 * 60 * 1000

	ss := &defs.SimpleSession{Username:un, TTL:ttl}
	sessionMap.Store(id, ss)
	dbops.InserSession(id, ttl, un)
	return id
}

//判断这个session是否过期
func IsSessionExpired(sid string)(string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := nowInMill()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}

		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		sessionMap.Store(sid, ss)
		return ss.Username, false
	}
	return "", true
}


