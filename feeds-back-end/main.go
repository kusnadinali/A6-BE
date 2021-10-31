package main

import (
	"database/sql"
	"encoding/json"
	"time"
	//"encoding/json"
	_ "expvar"
	"fmt"
	"net/http"
	"os"

	//"time"

	"github.com/go-kit/kit/log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "Kgdwfjrc123"
	dbname = "dummy"
)

//sementara di main dulu
type Postingan struct {
	postingan_id      string    //`json:"post_id"`
	username          string    //`json:"username"`
	profile_image     string    //`json:"profile_image"`
	postingan_caption string    //`json:"postingan_caption"`
	postingan_image   string    //`json:"postingan_image"`
	date_post         time.Time //`json:"date_post"`
	date_time         time.Time //`json:"date_time"`
	total_like        int
	total_comment     int
}

func main() {

	initDb()
	defer db.Close()

	logger := log.NewLogfmtLogger(os.Stdout)

	http.HandleFunc("/ ", GETHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}

// function to open connection with database
func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)

	//will use this later once we set the OS Environment

	/*host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}*/
	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "postgres"
	conf[dbpass] = "Kgdwfjrc123"
	conf[dbname] = "dummy"
	return conf
}

// //the handlers to each route
func GETHandler(w http.ResponseWriter, r *http.Request) {
	initDb()

	rows, err := db.Query("SELECT * FROM postingan")
	if err != nil {
		panic(err)
	}

	// var response []JsonResponse
	var feeds []Postingan

	for rows.Next() {
		var post Postingan
		//scan all values at the columns and store than at postingan variable
		rows.Scan(&post.postingan_id, &post.username, &post.profile_image, &post.postingan_caption, &post.postingan_image, &post.date_post, &post.date_time, &post.total_like, &post.total_comment)
		//add the person into the list.
		feeds = append(feeds, post)
	}

	postBytes, _ := json.MarshalIndent(feeds, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(postBytes)

	// var response = JsonResponse{Type: "succes", Data: postingan}

	// json.NewEncoder(w).Encode(response)

	defer rows.Close()
	defer db.Close()

}
