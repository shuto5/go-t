package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id     int
	Title  string
	Detail string
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/detail", Detail)
	http.ListenAndServe("localhost:8080", nil)
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

var todo Todo

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	//データベースにから最新のTODOを取得
	todos := GetTodos()

	// HTMLテンプレートにデータを渡してレスポンスを生成
	err := tmpl.Execute(w, todos)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetTodos() []Todo {
	db := dbConn()

	// "todos"テーブルからすべてのレコードを選択
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	// レコードの情報を格納するスライス
	var todos []Todo

	// レコードを反復処理して取得
	for rows.Next() {
		var todo Todo

		// レコードの各列の値をスキャンしてtodo構造体に割り当てる
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Detail)
		if err != nil {
			log.Fatalln(err)
		}

		// 取得したToDoをスライスに追加
		todos = append(todos, todo)
	}

	return todos
}

func Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/create.html"))
	// POSTメソッドでリクエストが送信された場合のみ処理を実行
	if r.Method == "POST" {
		// フォームから入力値を取得
		title := r.FormValue("title")
		detail := r.FormValue("detail")

		// データベースに接続
		db := dbConn()

		// データベースにToDoを追加
		_, err := db.Exec("INSERT INTO todos (title, detail) VALUES (?, ?)", title, detail)
		if err != nil {
			log.Fatalln(err)
		}

		// トップページにリダイレクト
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/edit.html"))
	// URLクエリパラメータからidを取得
	id := r.URL.Query().Get("id")

	// データベースに接続
	db := dbConn()

	// 指定されたIDのTodoを取得
	//var todo Todo
	err := db.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Detail)
	if err != nil {
		log.Fatalln(err)
	}

	// HTMLテンプレートにデータを渡してレスポンスを生成
	err = tmpl.ExecuteTemplate(w, "edit.html", todo)
	if err != nil {
		log.Fatalln(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/delete.html"))
	// GETリクエストの場合、削除確認画面を表示する
	if r.Method == "GET" {
		// URLクエリパラメータからidを取得
		id := r.URL.Query().Get("id")

		// データベースに接続
		db := dbConn()

		// IDに基づいてTodoを取得
		todo, err := getTodoID(db, id)
		if err != nil {
			log.Fatalln(err)
		}

		err = tmpl.ExecuteTemplate(w, "delete.html", todo)
		if err != nil {
			log.Fatalln()
		}

	} else if r.Method == "POST" {
		// POSTリクエストの場合、削除を実行する
		id := r.FormValue("id")
		//データベースに接続
		db := dbConn()
		//IDに基づいてTodoを削除
		_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
		if err != nil {
			log.Fatalln(err)
		}

		// トップページにリダイレクト
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func getTodoID(db *sql.DB, id string) (Todo, error) {
	// データベースからIDに一致するTodoを取得
	err := db.QueryRow("SELECT id, title, detail FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Detail)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func getTodoByID(id string) (Todo, error) {

	db := dbConn()

	err := db.QueryRow("SELECT id, title, detail FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Detail)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func Update(w http.ResponseWriter, r *http.Request) {
	// POSTメソッドでリクエストが送信された場合のみ処理を実行
	if r.Method == "POST" {
		// フォームから入力値を取得
		id := r.FormValue("id")
		title := r.FormValue("title")
		detail := r.FormValue("detail")

		// データベースに接続
		db := dbConn()
		//Todoを更新
		_, err := db.Exec("UPDATE todos SET title = ?, detail = ? WHERE id = ?", title, detail, id)
		if err != nil {
			log.Fatalln(err)
		}
		// トップページにリダイレクト
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

}

func Detail(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/detail.html"))
	// リクエストからIDを取得
	id := r.URL.Query().Get("id")

	// IDをもとにデータベースから該当のタスクを取得
	todo, err := getTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTMLテンプレートにデータを渡してレスポンスを生成
	err = tmpl.ExecuteTemplate(w, "detail.html", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
