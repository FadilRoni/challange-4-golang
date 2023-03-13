package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	no int
	name, address, occupation, reason string
}

type classroom struct {
	students []student
}

type biodata interface{
	getName(no int) string
	getAddress(no int) string
	getOccupation(no int) string
	getReason(no int) string
}

func (c classroom) getName(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.name
		}
	}
	return ""
}

func (c classroom) getAddress(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.address
		}
	}
	return ""
}

func (c classroom) getOccupation(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.occupation
		}
	}
	return ""
}

func (c classroom) getReason(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.reason
		}
	}
	return ""
}

func print(i int, b biodata) {
	fmt.Println("Nama\t: ", b.getName(i))
	fmt.Println("Alamat\t: ", b.getAddress(i))
  	fmt.Println("Pekerjaan\t: ", b.getOccupation(i))
  	fmt.Println("Alasan\t: ", b.getReason(i))
}

func main() {
	student_id := os.Args[1]

	i, _ := strconv.Atoi(student_id)

	var cr classroom

	var murid = []student {
		{
			no: 1,
			name: "adi",
			address: "Pekanbaru",
			occupation: "gojek",
			reason: "Pengen aja",
		},
		{
			no: 2,
			name: "david",
			address: "Pekanbaru",
			occupation: "operator",
			reason: "butuh ilmu",
		},
		{
			no: 3,
			name: "riski",
			address: "Pekanbaru",
			occupation: "admin",
			reason: "coba-coba",
		},
	}

	cr.students = murid

	print(i, cr)
}