package main

import "fmt"

type Employee struct {
	name    string
	address string
	job     string
	remark  string
}

func main() {

	// fmt.Println("Hello World")

	// for i := 0; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "=", "genap")
	// 	} else {
	// 		fmt.Println(i, "=", "ganjil")
	// 	}
	// }

	// nama := []string{"Andi", "Budi", "Cacing"}
	// for i := 0; i < len(nama); i++ {
	// 	fmt.Println(nama[i])
	// }

	var employee = []Employee{
		{name: "Delon", address: "DelonAddress", job: "DelonJob", remark: "DelonRemark"},
		{name: "KX", address: "KXAddress", job: "KXJob", remark: "KXRemark"},
		{name: "Michael", address: "MichaelAddress", job: "MichaelJob", remark: "MichaelRemark"},
		{name: "Lady", address: "LadyAddress", job: "LadyJob", remark: "LadyRemark"},
		{name: "Dzul", address: "DzulAddress", job: "DzulJob", remark: "DzulRemark"},
		{name: "Kevin", address: "KevinAddress", job: "KevinJob", remark: "KevinRemark"},
	}

	for _, x := range employee {
		fmt.Println(x)
	}

}
