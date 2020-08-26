package main

import (
	"fmt"

	"database/sql"
	"encoding/json"
	"io/ioutil"

	"github.com/lib/pq"
)

type student struct {
	RollNo  string   `json:RollNo`
	Name    string   `json:Name`
	Courses []string `json:Courses`
}

type StudentArray struct {
	Students []student
}

type Course struct {
	Instructor string   `json:Instructor`
	Code       string   `json:Code`
	Name       string   `json:Name`
	Conflicts  []string `json:Courses`
	Students   []string `json:Students`
}

var StudentData StudentArray

const (
	host     = "localhost"
	port     = 5432
	user     = "your_postgres_user"
	password = "your_postgres_password"
	dbname   = "your_postgres_db"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	CreateTablesinDB(db)

	AddStudentstoDBFromFile(db, "students/students.json")
	AddCoursestoDBFromFile(db, "./courses")

}

func CreateTablesinDB(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS EXAMS (
		ID serial PRIMARY KEY,
		course_code varchar(6),
		start_time timestamp NOT NULL,
		end_time timestamp NOT NULL
	);
	CREATE TABLE IF NOT EXISTS STUDENTS (
		roll_no varchar(14) PRIMARY KEY,
		name text,
		courses varchar(6)[]
	);
	CREATE TABLE IF NOT EXISTS INSTRUCTORS (
		instructor_id serial PRIMARY KEY,
		name text,
		email text
	);
	CREATE TABLE IF NOT EXISTS COURSES (
		course_code varchar(6) PRIMARY KEY,
		name text ,
		instructor_name text ,
		exam_ids integer[],
		students varchar(14)[],
		conflict_courses varchar(6)[]
	); `

	//In the schema above, instructor_name in courses is of type text. To be replaced by instructor id(integer).
	//This is currently for testing purposes only
	_, err := db.Exec(schema)

	if err != nil {
		fmt.Println(err)
	}

}

func AddStudentstoDBFromFile(db *sql.DB, path string) error {

	byteStream, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error opening students.json file")
		return err
	}

	err = json.Unmarshal(byteStream, &StudentData)

	//fmt.Println(StudentData)
	if err != nil {
		fmt.Println("JSON Syntax Error in students.json")
		fmt.Println(err)
		return err
	}

	stmt, err := db.Prepare("INSERT INTO students(name,roll_no,courses) VALUES($1,$2,$3)")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for i, student := range StudentData.Students {
		fmt.Println(student)

		_, err := stmt.Exec(student.Name, student.RollNo, pq.Array(student.Courses))

		if err != nil {
			fmt.Println(err)
			panic(err)
		} else {
			fmt.Printf("%s,%d,%s", "Row ", i, " inserted in DB.")
		}
	}

	return nil
}

func AddCoursestoDBFromFile(db *sql.DB, path string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return err
	}

	stmt, err := db.Prepare("INSERT INTO courses(course_code,name,instructor_name,students,conflict_courses) VALUES ($1,$2,$3,$4,$5)")

	for _, each_file := range files {
		course_file, err := ioutil.ReadFile("courses/" + each_file.Name())

		if err != nil {
			fmt.Println(err)
			return err
		}

		var course_json Course
		json.Unmarshal(course_file, &course_json)

		_, err = stmt.Exec(course_json.Code, course_json.Name, course_json.Instructor, pq.Array(course_json.Students), pq.Array(course_json.Conflicts))

		if err != nil {
			fmt.Println(err)
			return err
		} else {
			fmt.Printf("%s %s\n", "New row inserted for ", each_file.Name())
		}
	}
	return nil
}
