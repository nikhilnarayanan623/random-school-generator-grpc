package usecase

import (
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/utils"
)

func (s *schoolUseCase) createClass(name string, classChan chan<- domain.Class) {
	// select a random student count
	randStudentCount := utils.GetIntBetween(minStudents, maxStudent)

	// make slice of students with random count
	students := make([]domain.Student, randStudentCount)

	// fill the slice with students
	for i := range students {
		students[i] = s.schoolRepo.GetStudent()
	}

	updateRoleNumber(students)

	classChan <- domain.Class{
		Name:          name,
		Students:      students,
		TotalStudents: len(students),
	}
}

// to update the students roll number according to student name
func updateRoleNumber(students []domain.Student) {

	// sort and update the role number according to the sorted name
	sortAndUpdateRoleNumber(students, 0, len(students)-1)

	//shuffle the sorted students
	for i := range students {
		// select a random index from the previous the slice and swap with it
		j := utils.GetIntBetween(0, i)
		students[i], students[j] = students[j], students[i]
	}
}

// using quick sort for sorting; reason each time finding pivot it's actually same for roll number
func sortAndUpdateRoleNumber(arr []domain.Student, start, end int) {

	if start < end {
		pivotIdx := partition(arr, start, end)
		//  founded pivot index + 1 is the name's roll number
		arr[pivotIdx].RollNumber = uint(pivotIdx) + 1
		sortAndUpdateRoleNumber(arr, start, pivotIdx-1)
		sortAndUpdateRoleNumber(arr, pivotIdx+1, end)
	}

	if start == end {
		// if start and end is same; the roll number start or end + 1
		arr[end].RollNumber = uint(end) + 1
	}
}

func partition(arr []domain.Student, start, end int) int {
	pivot := arr[end].Name
	pIndex := start

	for i := start; i < end; i++ {
		if arr[i].Name < pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}

	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}
