package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var newGrades []int32
	for _, v := range grades {
		if v >= 38 {
			s := (v % 10)
			if s <= 5 {
				res := ((v / 10) * 10) + 5
				fmt.Println("res:", res, "v", v)
				if res-v < 3 {
					newGrades = append(newGrades, res)
				} else {
					newGrades = append(newGrades, v)
				}
			} else {
				res := ((v / 10) * 10) + 10
				fmt.Println("res:", res, "v", v)
				if res-v < 3 {
					newGrades = append(newGrades, res)
				} else {
					newGrades = append(newGrades, v)
				}
			}
		} else {
			newGrades = append(newGrades, v)
		}
	}
	return newGrades
}

func main() {
	a := []int32{73, 67, 38, 33}
	t := []int32{75, 67, 40, 33}

	res := gradingStudents(a)
	fmt.Println(t)
	fmt.Println(res)
}

