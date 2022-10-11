package main

var students = []*Student{}

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grade int32  `json:"grade"`
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "Student1", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "Student2", Grade: 3})
	students = append(students, &Student{Id: "s003", Name: "Student3", Grade: 4})
}
