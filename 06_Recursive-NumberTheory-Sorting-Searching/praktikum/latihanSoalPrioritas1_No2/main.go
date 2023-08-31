package main

import (
	"fmt"
	"math"
)

// Definisikan struct Student
type Student struct {
	name  string
	score float64
}

// Definisikan struct School
type School struct {
	students []Student
}

// Method untuk menghitung rata-rata skor
func (s *School) CalculateAverageScore() float64 {
	totalScore := 0.0
	for _, student := range s.students {
		totalScore += student.score
	}
	return totalScore / float64(len(s.students))
}

// Method untuk mencari siswa dengan skor minimum
func (s *School) FindMinScoreStudent() Student {
	minScore := math.MaxFloat64
	var minStudent Student

	for _, student := range s.students {
		if student.score < minScore {
			minScore = student.score
			minStudent = student
		}
	}
	return minStudent
}

// Method untuk mencari siswa dengan skor maksimum
func (s *School) FindMaxScoreStudent() Student {
	maxScore := 0.0
	var maxStudent Student

	for _, student := range s.students {
		if student.score > maxScore {
			maxScore = student.score
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	school := School{}
	
	for i := 1; i <= 5; i++ {
		var name string
		var score float64
		
		fmt.Printf("Input %d Student's name: ", i)
		fmt.Scanln(&name)
		
		fmt.Printf("Input %d Student's score: ", i)
		fmt.Scanln(&score)
		
		student := Student{
			name:  name,
			score: score,
		}
		
		school.students = append(school.students, student)
	}

	averageScore := school.CalculateAverageScore()
	minScoreStudent := school.FindMinScoreStudent()
	maxScoreStudent := school.FindMaxScoreStudent()

	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Min Score of Students: %s (%.0f)\n", minScoreStudent.name, minScoreStudent.score)
	fmt.Printf("Max Score of Students: %s (%.0f)\n", maxScoreStudent.name, maxScoreStudent.score)
}