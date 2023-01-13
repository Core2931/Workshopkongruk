package main

import (
	"fmt"
	"goworkshop/calculator"
)

// Structure list เก็บข้อมูลได้หลายชนิด
type Product struct {
	name        string
	price       float64
	description string
}

func main() {
	// import package
	fmt.Println(calculator.Add(20, 300))

	product1 := Product{name: "pen", price: 12, description: "write"}
	fmt.Println("Product Name", product1.name)
	var score float32
	// Array
	list_name := [...]string{}
	list_str := [3]string{"str1", "str2", "str3"}
	// Slice Dynamic Array
	slice_name := []string{}
	slice_name = append(slice_name, "append1")
	// Map = key:value
	maps_name := map[string]string{}
	maps_name["BTC"] = "bitcoin"
	maps_str_int := map[string]int{}
	maps_str_int["USD"] = 37
	maps_country := map[string]string{"TH": "ไทย", "JP": "ญี่ปุ่น"}
	// values = value, check = boolean
	values, check := maps_country["TH"]
	if check {
		fmt.Println(check, values)
	} else {
		fmt.Println("info Not found")
	}

	// Loop
	for key, value := range maps_country {
		fmt.Println("key:", key, "value:", value)
	}
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, values := range numbers {
		fmt.Println("Value", values)
	}

	//  Call Functions
	sum_nums := unlimitParam(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("Sum unlimit:", sum_nums)
	myTotal, status := getTotal(50, 100)
	fmt.Println("myTotal:", myTotal, "status:", status)

	fmt.Println(maps_country["TH"])
	fmt.Println(list_str)
	fmt.Print(list_name, slice_name)
	fmt.Print("ป้อนคะแนนสอบของคุณ : ")
	fmt.Scanf("%f", &score)

	if score >= 80 && score <= 100 {
		fmt.Printf("Get A you score : %.2f ", score)
	} else if score >= 70 && score < 80 {
		fmt.Printf("Get B you score : %.2f ", score)
	} else if score >= 60 && score < 70 {
		fmt.Printf("Get C you score : %.2f ", score)
	} else {
		fmt.Printf("Get D you score : %.2f ", score)
	}
}

func getTotal(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}

func unlimitParam(num ...int) int {
	total := 0
	for _, values := range num {
		total += values
	}
	return total
}
