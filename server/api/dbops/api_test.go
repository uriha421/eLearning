package dbops

import "testing"

//the test of api

//init(dblogin, truncate tables) --> run tests -->
//clear data(truncate tables) -->

var tempvid string

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
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("del", testDeleteUser)
	t.Run("reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("test","1234")
	if err != nil{
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "1234" || err != nil{             //
		t.Errorf("Error of GetUser")    // 为什么不用输出err
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("test","1234")
	if err != nil{
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if err != nil{
		t.Errorf("Error of RegetUser:%v", err)
	}
	if pwd != ""{
		t.Errorf("Deleting user test failed:%v", err)
	}

}

func TestVideoWorkFlow(t *testing.T){
	t.Run("AddUser", testAddUser)
	t.Run("AddVideoInfo", testAddNewVideoInfo)
	t.Run("GetVideoInfo", testGetVideoInfo)
	t.Run("DelVideoInfo", testDelVideoInfo)
	t.Run("RegetVideoInfo", testRegetVideoInfo)
}

func testAddNewVideoInfo(t *testing.T) {
	vi, err := AddNewVideoInfo(1, "test_video")
	if err != nil {
		t.Errorf("Error of AddNewVideoInfo:%v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err = GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of AddGetVideoInfo:%v", err)
	}
}

func testDelVideoInfo(t *testing.T) {
	err = DelVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DelVideoInfo:%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}
