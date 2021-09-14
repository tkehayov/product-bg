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
	ftpENV := os.Getenv("FTP")
	username := os.Getenv("FTP_USERNAME")
	password := os.Getenv("FTP_PASSWORD")

	ftpConnection, err := connect(ftpENV)
	login(ftpConnection, username, password)
	ras := getFile(ftpConnection, filename)

	defer ras.Close()

	buf, err := ioutil.ReadAll(ras)
	if err != nil {
		log.Error("error reading data: ", err)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	if _, err := w.Write(buf); err != nil {
		log.Println("unable to write image.", err)
	}

	if err := ftpConnection.Quit(); err != nil {
		log.Fatal("error quit: ", err)
	}

}

func getFile(c *ftp.ServerConn, filename string) *ftp.Response {
	productPath := os.Getenv("PRODUCT_PATH")

	ras, erro := c.Retr(productPath + "/" + filename)
	if erro != nil {
		log.Error("error reader: ", erro)
	}
	return ras
}

func login(c *ftp.ServerConn, username string, password string) {
	errLogin := c.Login(username, password)
	if errLogin != nil {
		log.Fatal("error user/pass: ", errLogin)
	}
}

func connect(ftpENV string) (*ftp.ServerConn, error) {
	c, err := ftp.Dial(ftpENV, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal("error FTP connection: ", err)
	}
	return c, err
}
