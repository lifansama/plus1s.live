package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"strings"
	"strconv"
)

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func plus1s(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("User-Agent"), "curl"){
		http.Redirect(w, r, "https://github.com/HFO4/plus1s.live", http.StatusFound)
		return 
	}
	i :=1
	for{
		if i==361{
			i=1
		}
		b, err := ioutil.ReadFile("pic/"+leftPad2Len(strconv.Itoa(i)  , "0", 3) +".txt")
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		write(w,str)
		i=i+1
	}
}


func write(w http.ResponseWriter,s string){
	flusher, ok := w.(http.Flusher)
	if !ok {
        	panic("Expected http.ResponseWriter to be an http.CloseNotifier")
	}
	fmt.Fprintf(w, s)
	fmt.Fprintf(w, "\033[2J\033[H")
	time.Sleep(100000000)
		flusher.Flush()
}

func main() {
	http.HandleFunc("/", plus1s)
	err := http.ListenAndServe(":1926", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
