package main

import (
	"github.com/gorilla/mux"
	"github.com/jlaffaye/ftp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//TODO
// - move all product images into products directory in ftp
// - change project structure
// - refactor project/file
func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	log.Info("FTP SERVICE STARTED")
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/assets/products/{filename}", product).Methods("GET")

	log.Error(http.ListenAndServe(":"+port, router))
}

func product(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filename := params["filename"]

	c, err := ftp.Dial("ftp:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal("error FTP connection: ", err)
	}

	err = c.Login("username", "password")
	if err != nil {
		log.Fatal("error user/pass ", err)
	}

	// Do something with the FTP conn
	size, sizeError := c.FileSize("/" + filename)
	if sizeError != nil {
		log.Error(sizeError)
	}
	log.Error("size: ", size)

	ras, erro := c.Retr("/" + filename)
	if erro != nil {
		log.Error("error reader: ", erro)
	}

	defer ras.Close()

	buf, err := ioutil.ReadAll(ras)
	if err != nil {
		log.Error("lqlql: ", err)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	if _, err := w.Write(buf); err != nil {
		log.Println("unable to write image.")
	}

	if err := c.Quit(); err != nil {
		log.Fatal("error quit: ", err)
	}

}
