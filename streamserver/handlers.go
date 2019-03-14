package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//vid := p.ByName("vid-id")
	//vl := VIDEO_DIR + vid
	//
	//video, err := os.Open(vl)
	//if err != nil {
	//	log.Printf("Error when try to open file: %v", err)
	//	sendErrorResponse(w, http.StatusInternalServerError, "internal Error")
	//	return
	//}
	//w.Header().Set("Content-Type", "video/mp4")
	//http.ServeContent(w, r, "", time.Now(), video)
	//defer video.Close()
	log.Println("Entered the streamserver")
	targetUrl := "http://avenssi-videos1.oss-cn-beijing.aliyuncs.com/videos/" + p.ByName("vid-id")
	http.Redirect(w, r, targetUrl, 301)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 校验传入文件的大小是否正确
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is to big")
		return
	}
	// 获取文件
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	// 获取文件二进制
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read filr error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	// 获取文件名，并保存文件
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	ossfn := "videos/" + fn
	path := "./videos/" + fn
	bn := "avenssi-videos1"
	ret := UploadToOss(ossfn, path, bn)
	if !ret {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
