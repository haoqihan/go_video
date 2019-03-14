package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string

// 清空数据库
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

// 开始test先执行他
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()

}

// 增删改查用户
func TestUserWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)

}

// 添加用户
func testAddUser(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of addUser: %v", err)
	}
}

// 查看用户
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Error("Error of GetUser")
	}

}

// 删除用户
func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of deleteUser: %v", err)
	}
}

// 再次查看
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of regetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test faile")
	}
}

// 增删改查视频
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

// 添加视频
func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

// 查看视频
func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

// 删除视频
func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

// 再次查看
func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

// 添加和读取评论
func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

// 添加评论
func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

// 查看全部评论
func testListComments(t *testing.T) {
	vid := "12345"
	from := 1549955446

	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	fmt.Println(to)
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("111111")
		fmt.Printf("comment: %d,%v \n", i, ele)
	}

}
