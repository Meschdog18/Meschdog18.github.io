package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"html/template"
	"strings"
	"log"
)
var filename string
func main(){
 fmt.Println("heloo")
 RoutesSetup()

}

func upload(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, "file upload endpoint found")

	//max file upload size 10mb(can be changed)
	r.ParseMultipartForm(32 << 20)

	//the file being uploaded
	file, handler, err := r.FormFile("myFile")
	//err handleing
	if err != nil{
		fmt.Println("ERROR", err)
		return
	}
	defer file.Close()

	 fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    tempFile, err := ioutil.TempFile("static/MainVideoPlayer/Videos","upload-*.mp4")
   
    if err != nil{
    	fmt.Println(err)
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil{
    	fmt.Println(err)
    }
    tempFile.Write(fileBytes)
 //gets file name
 filename = tempFile.Name()
  realname := strings.SplitAfter(filename,"static/MainVideoPlayer/")
  
   t, err := template.ParseFiles("videolink.html")
    if err != nil {
        fmt.Println(err)
    }
    items := struct {
        Video string
        Desp string
        VideoID string
    }{
        Video: "My vid",
        Desp: "Cool Vid",
         VideoID: realname[1],
    }
    t.Execute(w, items)
	
}
func videos(w http.ResponseWriter, r *http.Request){
	files, err := ioutil.ReadDir("static/MainVideoPlayer/Videos")
	if err != nil{
		log.Fatal(err)
	}

	for _, f := range files{
		fmt.Println(f.Name())
	}
}
func RoutesSetup(){
	http.HandleFunc("/upload",upload)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/vid", http.FileServer(http.Dir("./static/MainVideoPlayer")))
	http.HandleFunc("/videos", videos)
	http.ListenAndServe(":8080",nil)
}