package dbops

import "testing"

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of addUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Error("Error of GetUser")
	}

}
func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of deleteUser: %v", err)
	}
}
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of regetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test faile")
	}
}


//func TestVideoWorkFlow(t *testing.T){
//	clearTables()
//	t.Run("PrepareUser",testAddUser)
//	t.Run("AddVideo",testAddVideoInfo)
//	t.Run("GetVideo",testGetVideoInfo)
//	t.Run("DelVideo",testDeleteVideoInfo)
//	t.Run("RegetVidoe",testRegetVideoInfo)
//}
//
//
//func testAddVideoInfo(t *testing.T){
//	vi, err := AddNewVideo(1,"my_video")
//	if err != nil{
//		t.Errorf("Error od AddVideoInfo: %v",err)
//	}
//	tempvid = vi.Id
//}
//
//func testGetVideoInfo(t *testing.T) {
//	_, err := GetVideoInfo(tempvid)
//	if err != nil{
//		t.Errorf("Error of GetVideoInfo: %v",err)
//	}
//}
//
//func testDeleteVideoInfo(t *testing.T) {
//	err := DeleteVideoInfo(tempvid)
//	if err != nil{
//		t.Errorf("Error of DeleteVideoInfo:%v",err)
//	}
//}
//
//func testRegetVideoInfo(t *testing.T){
//	vi,err := GetVideoInfo(tempvid)
//	if err != nil || vi != nil{
//		t.Errorf("Error of RegetVideoInfo: %v",err)
//	}
//}



func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}