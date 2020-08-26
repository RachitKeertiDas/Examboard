package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"path/filepath"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/lib/pq"
)

type Course struct {
	Instructor string   `json:Instructor`
	Code       string   `json:Code`
	Name       string   `json:Name`
	Conflicts  []string `json:Courses`
	Students   []string `json:Students`
}

type student struct {
	RollNo  string   `json:RollNo`
	Name    string   `json:Name`
	Courses []string `json:Courses`
}

type StudentArray struct {
	Students []student
}

var StudentData StudentArray

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "your_postgres_user"
	password = "your_postgres_password"
	dbname   = "your_postgres_db"
)

func StudentCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Search for Roll Number
	Roll := r.URL.Path[len("/api/student/getCourses/"):]
	fmt.Println(Roll)
	log.Println(Roll)
	var flag = 0
	for i := range StudentData.Students {
		if StudentData.Students[i].RollNo == Roll {
			log.Println("Found!")
			json_data, err := json.Marshal(StudentData.Students[i])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Server Error while fetching course details")
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(json_data)

			}
			flag = 1
			break
		}

	}
	//if not found return 404
	if flag == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Roll No Not Found"))
	}
	//search up the student Roll No, Jsonify and return JSON

}

func InstructorCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	course_file, err := os.Open("courses/" + r.URL.Path[len("/api/getCourseDetails/"):] + ".json")

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		byteStream, err := ioutil.ReadAll(course_file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(byteStream)
	}
}

func InstructorListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	instructor_file, err := os.Open("instructors/Abhinav.json")

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		byteStream, err := ioutil.ReadAll(instructor_file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(byteStream)
	}
}

func init_data() error {
	os.MkdirAll("courses", 0775)
	os.MkdirAll("students", 0775)
	os.MkdirAll("instructors", 0775)

	json_file, err := os.Open("students/students.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error Opening JSON File")
		return err
	}

	byteStream, _ := ioutil.ReadAll(json_file)
	err = json.Unmarshal(byteStream, &StudentData)

	//fmt.Println(StudentData)
	if err != nil {
		fmt.Println("JSON Syntax Error in students.json")
		fmt.Println(err)
		return err
	}
	return nil
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	course_name := r.URL.Path[len("/api/getCourseDetails/"):]
	var CourseQuery Course

	err := db.QueryRow("SELECT name,course_code,instructor_name,students,conflict_courses FROM courses WHERE course_code = $1", course_name).Scan(
		&CourseQuery.Name, &CourseQuery.Code, &CourseQuery.Instructor, pq.Array(&CourseQuery.Students), pq.Array(&CourseQuery.Conflicts))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error in SQL Query")
	}

	course_json, err := json.Marshal(CourseQuery)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in encoding into JSON")

	} else {

		w.WriteHeader(http.StatusOK)
		w.Write(course_json)
	}
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var StudentQuery student
	Roll := r.URL.Path[len("/api/student/getCourses/"):]
	fmt.Println(Roll)
	log.Println(Roll)
	err := db.QueryRow("SELECT * FROM students WHERE roll_no = $1", Roll).Scan(&StudentQuery.RollNo, &StudentQuery.Name, pq.Array(&StudentQuery.Courses))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error in SQL Query")
	}

	student_json, err := json.Marshal(StudentQuery)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in encoding into JSON")

	} else {

		w.WriteHeader(http.StatusOK)
		w.Write(student_json)
	}
}

func main() {

	///Initialize DataBase
	init_data()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/student/getCourses/{RollNo}", StudentHandler)
	r.Get("/api/getCourseDetails/{course}", CourseHandler)
	r.Get("/api/instructor/getCourseList/", InstructorListHandler)
	//r.Get("/api/student/getCourses/{course}",StudentHandler)
	//http.HandleFunc("/api/instructor/getCourses/", InstructorCourseHandler)
	//http.HandleFunc("/api/instructor/conflicts/",FetchConflictHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Printf("\n")
}

func InitDB(db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
