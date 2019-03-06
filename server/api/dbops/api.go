package dbops

//the operations of database
import "log"

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?) ")
	if err != nil {
		return err
	}

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)(string, error){
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ? ")
	if err != nil {
		log.Printf("%s", err)
		return "", err    //why to return "" instead of nil?
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)      //先执行stmtOut.QueryRow(loginName)检索loginName，返回一个*Row。再执行Scan。
	stmtOut.Close()

	return pwd, nil
}

func DeleteUser (loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ? ")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()

	return err
}
