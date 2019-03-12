package dbops

import "testing"

//the test of api

//init(dblogin, truncate tables) --> run tests -->
//clear data(truncate tables) -->

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain (m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", TestAddUser)
	t.Run("get", TestGetUser)
	t.Run("del", TestDeleteUser)
	t.Run("reget", TestRegetUser)
}

func TestAddUser(t *testing.T) {
	err := AddUserCredential("test","1234")
	if err != nil{
		t.Errorf("Error of AddUser:%v", err)
	}
}

func TestGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "1234" || err != nil{             //
		t.Errorf("Error of GetUser")    // 为什么不用输出err
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("test","1234")
	if err != nil{
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func TestRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if err != nil{
		t.Errorf("Error of RegetUser:%v", err)
	}
	if pwd != ""{
		t.Errorf("Deleting user test failed:%v", err)
	}

}


