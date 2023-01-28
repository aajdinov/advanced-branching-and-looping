package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const totalDays = 192

type student struct {
	name       string
	grade      rune
	percentage float32
	attDays    int
	courses    course
}

type course struct {
	dataLogic    float32
	algorithms   float32
	calculus     float32
	finalProject float32
}

type studentList [12]student

func (s *studentList) generateMarks() {
	for i := 0; i < 12; i++ {
		name := fmt.Sprintf("Student %v", i+1)
		s[i] = student{name, ' ', 0, 82 + rand.Intn(111),
			course{50 + float32(rand.Intn(51)), 50 + float32(rand.Intn(51)), 50 +
				float32(rand.Intn(51)), 8 + float32(rand.Intn(18))}}
	}
}

func printStudentDetails(s []student) {
	fmt.Println("\nStudent Name\tPercentage\tAttendance\tProject\t  Final Grade")
	fmt.Println("---------------------------------------------------------------------")
	for i := 0; i < 12; i++ {
		fmt.Printf("%-10v\t%8.3v\t%6v\t%12v\t%8c\n", s[i].name, s[i].percentage, s[i].attDays, s[i].courses.finalProject, s[i].grade)
	}
}

func runBusinessLogic(s []student) {
	var grade rune
	for i := 0; i < 12; i++ {
		s[i].percentage, grade = calGrade(s[i])
		if grade != 'F' {
			calAttendance(s[i].attDays, &grade)
		}
		grade = calProject(s[i].courses.finalProject, grade)

		s[i].grade = grade
	}
}

func calGrade(s student) (float32, rune) {
	percentage := (s.courses.dataLogic + s.courses.algorithms + s.courses.calculus) / 3
	var grade rune

	switch {
	case percentage >= 90:
		grade = 'A'
	case percentage >= 80:
		grade = 'B'
	case percentage >= 70:
		grade = 'C'
	case percentage >= 60:
		grade = 'D'
	default:
		grade = 'F'
	}

	return percentage, grade
}

func calAttendance(attDays int, grade *rune) {
	if float32(attDays)/float32(totalDays) < 0.75 {
		*grade = 'F'
	}
}

func calProject(marks float32, grade rune) rune {
	percent := float32(marks) / 25 * 100
	if percent >= 80 && grade != 'A' {
		grade--
	} else if percent < 40 && grade != 'F' {
		grade++
	}
	return grade
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var Students studentList
	Students.generateMarks()
	students := Students[:]

	runBusinessLogic(students)

	sort.Slice(students, func(i, j int) bool {
		return students[i].grade < students[j].grade
	})

	printStudentDetails(students)
}
