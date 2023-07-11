/*Viết một chương trình thỏa mãn các yêu cầu sau:
- Tạo một mảng lữu trữ danh sách tên học sinh của một trường học
- Nhập vào số lượng học sinh cần lưu
- Nhập tên của từng học sinh (Tên có thể tùy ý viết hoa hoặc viết thường)
- in ra danh sách sinh viên trong mảng
- thêm một sinh viên vào mảng
- Tìm kiếm học sinh (yêu cầu nhập vào tên của học sinh), in ra danh sach học sinh nếu tìm thấy và ngược lại không tìm thấy
- Xóa phần tử trong mảng theo tên hoặc theo họ
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func printListStudent(students []string) {
	fmt.Println("List sutdents is: ")
	for _, student := range students {
		fmt.Println(student)
	}
}

func addStudent(students []string) []string {
	var name string
	fmt.Printf("Enter name of student to add: ")
	name = input()
	students = append(students, name)
	return students
}

func searchStudent(students []string) []string {
	var name string
	fmt.Printf("Enter name of student to search: ")
	fmt.Scanf("%s", &name)

	var search_student []string

	for _, student := range students {
		check := strings.Contains(strings.ToUpper(student), strings.ToUpper(name))
		if(check == true) {
			search_student = append(search_student, student)
		}
	}
	return search_student
}

func removeStudent(students []string) []string {
	var name string
	fmt.Printf("Enter name of student to remove: ")
	fmt.Scanf("%s", &name)

	for i, student := range students {
		check := strings.Contains(strings.ToUpper(student), strings.ToUpper(name))
		if(check == true) {
			students = append(students[:i], students[i+1:]...)
		}
	}
	return students
}

func inputStudent(listStudent []string) []string {
	var number_of_students int
	loop: fmt.Printf("Enter number of students: ")
	fmt.Scanf("%d", &number_of_students)

	if(number_of_students <= 0) {
		fmt.Println("Please enter a number of students > 0")
		goto loop
	}

	fmt.Println("Enter name of students: ")
	var name string

	for i:=0; i<number_of_students; i++ {
		name = input()
		listStudent = append(listStudent, name)
	}
	return listStudent
}

func main() {
	var listStudent []string
	/*
	Menu options:
	1. - tạo mảng sinh viên
	2. - in ra danh sách sinh viên trong mảng
	3. - thêm một sinh viên vào mảng
	4. - Tìm kiếm học sinh
	5. - Xóa phần tử 
	6. - Exit
	*/

	flag := true

	for flag {
	    var opt int
	    loop: 
		fmt.Println("Menu options:")
		fmt.Println("1. Tao mang sinh vien")
		fmt.Println("2. In ra danh sach sinh vien trong mang")
		fmt.Println("3. Them mot sinh vien vao mang")
		fmt.Println("4. Tim kiem sinh vien")
		fmt.Println("5. Xoa sinh vien")
		fmt.Println("6. Exit")
		fmt.Println("Enter option: ")
	    fmt.Scanf("%d", &opt)
		if(opt <= 0) {
			fmt.Println("Please enter a option > 0")
			goto loop
		}

		switch(opt) {
		case 1: listStudent = inputStudent(listStudent)
	    case 2: {
			if(listStudent == nil) {
				fmt.Println("Please create a array of students")
				break
			}
			printListStudent(listStudent)
		}
		case 3: {
			if(listStudent == nil) {
				fmt.Println("Please create a array of students")
				break
			}
			listStudent = addStudent(listStudent)
		}

	    case 4: {
			if(listStudent == nil) {
				fmt.Println("Please create a array of students")
				break
			}

			search_student := searchStudent(listStudent)
			if(search_student == nil) {
				fmt.Println("Not found student!")
			} else {
				fmt.Println("Found student!")
				printListStudent(search_student)
			}

		}
		case 5: {
			if(listStudent == nil) {
				fmt.Println("Please create a array of students")
				break
			}
			listStudent = removeStudent(listStudent)
		}
		case 6:
			flag = false
		}
	}
}