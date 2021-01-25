package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Song struct {
	Id     int
	Name   string
	Artist string
	Album  string
	Genre  string
	year   int
}

var tpl *template.Template

var db *sql.DB

func main() {

	tpl, _ = template.ParseGlob("templates/*.html")
	var err error

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/Playlist")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/updateresult/", updateResultHandler)
	http.HandleFunc("/delete/,", deleteHandler)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("****Browser handler running****")
	stmt := "SELECT * FROM Playlist.songs"
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var songs []Song
	for rows.Next() {
		var s Song
		err = rows.Scan(&s.Id, &s.Name, &s.Artist, &s.Album, &s.Genre, &s.year)
		if err != nil {
			panic(err)
		}
		songs = append(songs, s)

	}
	tpl.ExecuteTemplate(w, "select.html", songs)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("***Inserthandler running***")
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}
	r.ParseForm()
	name := r.FormValue("nameName")
	artist := r.FormValue("artistName")
	album := r.FormValue("albumName")
	genre := r.FormValue("genreName")
	year := r.FormValue("yearName")
	var err error
	if name == "" || artist == "" || album == "" || year == "" || genre == "" {
		fmt.Println("Error inserting the row: ", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data")
		return
	}

	var ins *sql.Stmt

	ins, err = db.Prepare("INSERT INTO Playlist.songs(Name,Artist,Album,Genre,year) VALUES(?,?,?,?,?);")

	if err != nil {
		panic(err)
	}
	defer ins.Close()

	res, err := ins.Exec(name, artist, album, genre, year)

	changes, _ := res.RowsAffected()
	if err != nil || changes != 1 {
		fmt.Println("Error inserting values: ", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting values")
		return

	}

	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("ID of the recent song inserted", lastInserted)
	fmt.Println("Number of records affected: ", rowsAffected)
	tpl.ExecuteTemplate(w, "insert.html", "Song added :)")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete request handler running")
	r.ParseForm()
	id := r.FormValue("Id")

	del, err := db.Prepare("DELETE FROM Playlist.songs WHERE (Id=?);")

	if err != nil {
		panic(err)
	}

	defer del.Close()
	var res sql.Result
	res, err = del.Exec(id)
	changes, _ := res.RowsAffected()
	fmt.Println("Rows affected: ", changes)

	if err != nil {
		fmt.Println(w, "Error deleting song")
		return
	}
	fmt.Println("Err: ", err)
	tpl.ExecuteTemplate(w, "result.html", "Song has been deleted :(")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("***Update handler running***")
	r.ParseForm()
	id := r.FormValue("Id")
	row := db.QueryRow("SELECT * FROM Playlist.songs WHERE Id=?;", id)
	var s Song
	err := row.Scan(&s.Id, &s.Name, &s.Artist, &s.Album, &s.Genre, &s.year)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/browse", 307)
		return
	}
	tpl.ExecuteTemplate(w, "update.html", s)
}

func updateResultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("***UpdateResultHandler running***")
	r.ParseForm()
	id := r.FormValue("Id")
	name := r.FormValue("nameName")
	artist := r.FormValue("artistName")
	album := r.FormValue("albumName")
	genre := r.FormValue("genreName")
	year := r.FormValue("yearName")
	upstmt := "UPDATE Playlist.songs SET Name=?, Artist=?, Album = ?, Genre=?,year=? WHERE (Id=?);"

	stmt, err := db.Prepare(upstmt)
	if err != nil {
		fmt.Println("Error preparing statment")
		panic(err)
	}
	fmt.Println("db.Prepare error: ", err)
	fmt.Println("db.Prepare statement: ", stmt)
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(name, artist, album, genre, year, id)
	changes, _ := res.RowsAffected()
	if err != nil || changes != 1 {
		fmt.Println(err)
		tpl.ExecuteTemplate(w, "result.html", "Error")
		return
	}

	tpl.ExecuteTemplate(w, "result.html", "Playlist updayed :)")

}

func songSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "songsearch.html", nil)
		return
	}
	r.ParseForm()
	min := r.FormValue("minyear")
	max := r.FormValue("maxyear")
	stmt, err := db.Prepare("SELECT * FROM Playlist.songs WHERE year >= ? && year <= ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(min, max)

	var songs []Song

	for rows.Next() {
		var s Song
		err := rows.Scan(&s.Id, &s.Name, &s.Artist, &s.Album, &s.Genre, &s.year)
		if err != nil {
			panic(err)
		}
		songs = append(songs, s)
	}
	tpl.ExecuteTemplate(w, "songsearch.html", songs)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")

}
