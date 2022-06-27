package main

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Name    string
	Address string
	Job     string
	Remark  string
}

func main() {

	var employee = []Employee{
		{Name: "Delon", Address: "DelonAddress", Job: "DelonJob", Remark: "DelonRemark"},
		{Name: "KX", Address: "KXAddress", Job: "KXJob", Remark: "KXRemark"},
		{Name: "Michael", Address: "MichaelAddress", Job: "MichaelJob", Remark: "MichaelRemark"},
		{Name: "Lady", Address: "LadyAddress", Job: "LadyJob", Remark: "LadyRemark"},
		{Name: "Dzul", Address: "DzulAddress", Job: "DzulJob", Remark: "DzulRemark"},
		{Name: "Kevin", Address: "KevinAddress", Job: "KevinJob", Remark: "KevinRemark"},
	}

	if absen, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(err)
	} else {
		if absen > len(employee) {
			fmt.Println("Enter 0 - " + strconv.Itoa(len(employee)))
		} else {
			fmt.Println(employee[absen])
		}
	}
}
