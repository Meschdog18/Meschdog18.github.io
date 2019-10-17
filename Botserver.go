package main

import (
	"fmt"
	"log"
    "net/http"

    "encoding/json"

)
type dat struct {
	Red   string `json:"red"`
	Green string `json:"green"`
	Blue string `json:"blue"`
}
var red string
var green string
var blue string
func main(){
	handleRequests()
}
func handleRequests(){
	http.HandleFunc("/", incomingPOST)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func incomingPOST(w http.ResponseWriter, r *http.Request){
	 var data dat
	
	switch r.Method{
	case "POST":
		
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
			log.Println(err)
		}
		red = data.Red
		green = data.Green
		blue = data.Blue
		fmt.Println(data.Red)
		fmt.Println(data.Green)
		fmt.Println(data.Blue)
     	 
		
	default:
	 fmt.Fprintf(w, "red" + red)
	 fmt.Fprintf(w, "green" + green)
	 fmt.Fprintf(w, "blue" + blue)
	
		
	}
 		
	 		

	      
}