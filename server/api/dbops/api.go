package dbops

import (
	"github.com/eLearning/server/api/defs"
	"github.com/eLearning/server/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//the operations of database

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?) ")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string, error){
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ? ")
	if err != nil {
		log.Printf("%s", err)
		return "", err    //why to return "" instead of nil?
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)      //先执行stmtOut.QueryRow(loginName)检索loginName，返回一个*Row。再执行Scan。
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser (loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ? ")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return err
}

// creat uuid --> creat time --> DB

func AddNewVideoInfo(aid int, video_name string)(*defs.VideoInfo, error){
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") //M D Y, HH:MM:SS

	stmtIns, err := dbConn.Prepare("insert into video_info (id, author_id, name, display_ctime) values (?,?,?,?)")
	if err != nil{
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, video_name, ctime)
	defer stmtIns.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: video_name, DisplayCtime: ctime}
	return res, nil
}

func GetVideoInfo(vid string)(*defs.VideoInfo, error){
	stmtGet, err := dbConn.Prepare("select author_id, name, display_ctime from video_info where id = ?")
	if err != nil {
		return nil, err
	}

	var (
		aid int
		name string
		dct string
	)

	err = stmtGet.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil, nil
	}

	defer stmtGet.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

//func ListVideoInfo(aid string)(*defs.VideoInfo error){
//	stmtList error := dbConn.Prepare("select  from")
//
//}

func DelVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_info where id = ?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid string, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtAdd, err := dbConn.Prepare("insert into comments value (id, video_id, author_id, content) values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtAdd.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmtAdd.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`select comments.id, users.Login_name, comments.content from comments 
		inner join users on comments.author_id = users.id 
		where video_id = ? and comments.time > FROM_UNIXTIME(?) and comments.time <= FROM_UNIXTIME(?)
		order by comments.time desc`) //为什么要加时间？？？
	if err != nil {
		return nil, err
	}

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)   //执行Query，然后返回一个*rows。rows是row的一个数组。
	if err != nil {
		return res, err
	}

	for rows.Next(){
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil{
			return res, err
		}

	c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
	res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}