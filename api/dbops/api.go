package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

// 添加用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name,pwd) VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	stmtIns.Exec(loginName, pwd)
	defer stmtIns.Close()
	return nil

}

// 获取用户信息
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM  users where login_name = ? ")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}

// 获取用户
func GetUser(loginName string)(*defs.User,error){
	stmtOut, err := dbConn.Prepare("select id,pwd from users where login_name=?")
	if err != nil{
		log.Printf("%s",err)
		return nil,err
	}
	var id int
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&id,&pwd)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}
	res := &defs.User{Id:id,LoginName:loginName,Pwd:pwd}
	defer stmtOut.Close()
	return res,nil
}

// 删除用户
func DeleteUser(loginName string, pwd string) error {
	stmDel, err := dbConn.Prepare("DELETE  FROM users where login_name=? and pwd=?")
	if err != nil {
		log.Printf("DeleteUser error:%s", err)
		return err
	}
	_, err = stmDel.Exec(loginName, pwd)
	if err != err {
		return err
	}
	defer stmDel.Close()
	return nil
}

// 添加视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()

	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(
		`INSERT INTO video_info (id, author_id, name,display_ctime) VALUES (?,?,?,?)`)

	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

//// 获取视频
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where  id =?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

// 删除视频
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_info where id=? ")
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

// 添加评论
func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO  comments(id,video_id,author_id,content) values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)

	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 查看全部评论
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT  comments.id,users.Login_name,comments.content FROM comments
									inner join users on comments.author_id = users.id
									where comments.video_id = ? and comments.time > FROM_UNIXTIME(?) and comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	fmt.Println(res)
	return res, nil
}

// 所有视频
func ListVideoInfo(uname string,from,to int )([]*defs.VideoInfo,error){
	stmtOut, err := dbConn.Prepare(`select  video_info.id,video_info.author_id,video_info.name,video_info.display_ctime from video_info
											INNER join users ON video_info.author_id=users.id where users.login_name=? and video_info,create_time>from_unixtime(?) and 
											video_info.create_time<= from_unixtime(?) order by video_info.create_time DESC `)
	var res []*defs.VideoInfo
	if err != nil{
		return res,err
	}
	rows,err := stmtOut.Query(uname,from,to)
	if err != nil{
		return res,err
	}
	for rows.Next(){
		var id,name,ctime string
		var aid int
		if err := rows.Scan(&id,&aid,&name,&ctime);err != nil{
			return res,err
		}
		vi := &defs.VideoInfo{Id:id,AuthorId:aid,Name:name,DisplayCtime:ctime}
		res = append(res,vi)

	}
	defer stmtOut.Close()
	return res,nil
}

//func DeleteVideoInfo(vid string) error{
//	stmtDel, err := dbConn.Prepare("delete from video_info where id = ?")
//	if err != nil{
//		return err
//	}
//	_,err = stmtDel.Exec(vid)
//	if err != nil{
//		return err
//	}
//	defer stmtDel.Close()
//	return nil
//}
//
//func AddNewComments(vid string,aid int,content string)error{
//
//}



