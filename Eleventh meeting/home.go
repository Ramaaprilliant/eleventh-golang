package main

import (
	//"html/template"
	"fmt"
	"log"
	"net/http"
	"strconv"
	//"io/ioutil"
	"regexp"
	//"os"
	//"encoding/json"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func hometable(w http.ResponseWriter, r *http.Request){
	db,err := sql.Open("sqlite3", "./userDB.db")
	checkErr(err)
	
	//
	stmt,err := db.Prepare("insert into 'USERS'(username,password,age,authority) values( ?, ?, ?, ?)")
	checkErr(err)
	res,err := stmt.Exec("admin","pass12",01,"Administrator" )
	checkErr(err)
	id,err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	rows, err := db.Query("SELECT * FROM USERS")
	checkErr(err)
    var uid int
    var username string
	var password string
	var age int
	var authority string
	 
	for rows.Next() {
		err = rows.Scan(&uid, &username, &password, &age, &authority)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(age)
		fmt.Println(authority)
	}

	rows.Close()
	db.Close()

}

func cekLogin(w http.ResponseWriter, r *http.Request){
	if len(r.Form["username"][0]) < 5 {
		fmt.Fprintf(w, "User have yet to fill Username")
	}
	if len(r.Form["password"][0]) < 8 {
		fmt.Fprintf(w, "User have yet to fill Password")
	}
	_ , err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		log.Fatal("error running service : ", err)
	}
	if m, _ := regexp.MatchString("^([\\w\\.\\_]{2,10})@(\\w{1,}).([a-z]{2,4})$", r.Form.Get("email")); !m {
		
	}
}

func main() {
	http.HandleFunc("/home", hometable)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("error running service : ", err)
	}

}

func Adduser() {

}

func checkErr(err error) {
	if err != nil{
		panic(err)
	}
}