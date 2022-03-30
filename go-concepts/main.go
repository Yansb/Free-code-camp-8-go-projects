package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var courses []Course

type Course struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func generateCourses() {
	course1 := Course{ID: "1", Name: "Go"}
	course2 := Course{ID: "2", Name: "Python"}
	course3 := Course{ID: "3", Name: "Java"}
	course4 := Course{ID: "4", Name: "C++"}
	course5 := Course{ID: "5", Name: "C#"}
	course6 := Course{ID: "6", Name: "Ruby"}
	course7 := Course{ID: "7", Name: "PHP"}
	course8 := Course{ID: "8", Name: "JavaScript"}
	course9 := Course{ID: "9", Name: "C"}

	courses = append(courses, course1, course2, course3, course4, course5, course6, course7, course8, course9)
}

func main() {
	generateCourses()
	e := echo.New()
	e.GET("/course", listCourses)
	e.POST("/course", createCourse)
	e.Logger.Fatal(e.Start(":9010"))
}

func listCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
}

func createCourse(c echo.Context) error {
	course := Course{}
	c.Bind(&course)
	err := persistCourse(course)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, course)
}

func persistCourse(course Course) error {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("insert into courses values ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(course.ID, course.Name)

	if err != nil {
		return err
	}

	return nil

}
