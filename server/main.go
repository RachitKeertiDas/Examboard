package main

import (

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
)

type course struct {
	Name      string   `json:Name`
	Code      string   `json:Code`
	StartSeg  int      `json:StartSeg`
	EndSeg    int      `json:EndSeg`
	Exams     []string `json:Exams`
	Students  []string `json:Students`
	Conflicts []string `json:Conflicts`
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

func CreateCourseFile(hello course) int {
	courseFile, err := os.OpenFile("CS1353"+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating course file")
	}
	fmt.Fprintf(courseFile, "Course")
	return 1
}


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

func main() {

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

	init_data()
	r.Get("/api/student/getCourses/{RollNo}", StudentCourseHandler)
	r.Get("/api/getCourseDetails/{course}", InstructorCourseHandler)
	r.Get("/api/instructor/getCourseList/", InstructorListHandler)
	//r.Get("/api/student/getCourses/{course}",StudentHandler)
	//r.Patch("/api")
	//http.HandleFunc("/api/instructor/getCourses/", InstructorCourseHandler)
	//http.HandleFunc("/api/instructor/conflicts/",FetchConflictHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Printf("\n")
}
