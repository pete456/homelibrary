package handles

import (
	"fmt"
	"homelibrary/internal/parsing"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
)

func SendDirectory(w http.ResponseWriter, r *http.Request) {
	url := parsing.MediaURL(r.URL.String())
	defer func() {
		if re := recover(); re != nil {
			w.Write([]byte(fmt.Sprintf("404 - %s not found",r.URL.String())))
		}
	}()

	dirfilelist,err := ioutil.ReadDir(url)
	if err != nil {
		log.Panicf("Could NOT read %s directory. ERROR: %s",url,err.Error())
	}
	var namestring string
	for _,n := range dirfilelist {
		appendname := n.Name()
		if n.IsDir() {
			appendname += "/"
		}
		namestring += appendname+"\n"
	}
	log.Printf("Serving directory %s\n",url)
	w.Write([]byte(namestring))
}

func SendFile(w http.ResponseWriter, r *http.Request) {
	url := parsing.MediaURL(r.URL.String())
	log.Printf("Serving file %s\n",url)
	http.ServeFile(w,r,url)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	spliturl := strings.Split(r.URL.String(),"/")
	if parsing.NameIsPDF(spliturl[len(spliturl)-1]) {
		SendFile(w,r)
	} else {
		SendDirectory(w,r)
	}
}
