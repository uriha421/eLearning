package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"github.com/eLearning/server/api/defs"
	"sync"
)

//插入一个session
func InserSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmsIns, err := dbConn.Prepare("insert into sessions (session_id, TTL, login_name) Value (?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmsIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return nil
	}

	defer stmsIns.Close()
	return nil
}

//通过session_id查找一个session 为什么要返回一个指针变量?直接操作返回结果
func RetrieveSession(sid string) (*defs.SimpleSession, error){
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL, login_name from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string

	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	//将ttl从string转成int
	res, err := strconv.ParseInt(ttl, 10, 64)
	if err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

//查找全部sessions
func RetreiverAllSessions() (*sync.Map, error){
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next(){
		var id       string
		var ttlstr   string
		var login_name string
		if er := rows.Scan(&id, &ttlstr, &login_name); er != nil{
			log.Printf("retrive sessions error: %s", er)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err != nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf(" session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

//删除一个session
func DeleteSession(sid string)(error){
	stmtDel, err := dbConn.Prepare("delete * from sessions where id = ?")
	if err != nil {
		return nil
	}

	_, err = stmtDel.Query(sid)
	if err != nil {
		return nil
	}

	return nil
}
