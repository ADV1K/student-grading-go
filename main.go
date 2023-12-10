package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parseCSV(filePath string) []student {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var result []student
	for i, row := range data {
		if i == 0 {
			continue // skip header row
		}
		result = append(result, student{
			firstName:  row[0],
			lastName:   row[1],
			university: row[2],
			test1Score: parseInt(row[3]),
			test2Score: parseInt(row[4]),
			test3Score: parseInt(row[5]),
			test4Score: parseInt(row[6]),
		})
	}
	return result
}

func calculateGrade(students []student) (result []studentStat) {
	for _, s := range students {
		score := float32(s.test1Score+s.test2Score+s.test3Score+s.test4Score) / 4
		var grade Grade
		switch {
		case score >= 70:
			grade = A
		case score >= 50:
			grade = B
		case score >= 35:
			grade = C
		default:
			grade = F
		}

		result = append(result, studentStat{
			student:    s,
			finalScore: score,
			grade:      grade,
		})
	}
	return
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topper := gradedStudents[0]
	for _, s := range gradedStudents {
		if s.finalScore > topper.finalScore {
			topper = s
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	universityToppers := make(map[string]studentStat)
	for _, s := range gs {
		topper, present := universityToppers[s.student.university]
		if !present || s.finalScore > topper.finalScore {
			universityToppers[s.student.university] = s
		}
	}
	return universityToppers
}
