package main

import "fmt"

type A struct {
	data []Data
}

type Data struct {
	Status bool
	Name   string
}

func main() {

	// 배열에서 range의 값은 복사, 인덱스는 참조이다
	// 값을 변경하기 위해서는 인덱스로 접근, 출력하는 용도는 값을 이용하면 된다
	arr := []int{1, 2, 3, 4}
	for i, v := range arr {

		if v == 3 {
			fmt.Println(v, arr[i])
			arr[i] = 5
			fmt.Println(v, arr[i])
		}
		if v == 4 {
			fmt.Println(v, arr[i])
			v = 5
			fmt.Println(v, arr[i])
		}

	}

	mm := make(map[int]int)

	mm[0] = 0
	mm[1] = 1

	// Golang의 데이터 구조 map은 값도 참조이다 (map 자체가 참조 형태이라서 그런것 같다)
	for _, v := range mm {
		fmt.Println(v)
		v = 3
		fmt.Println(v)
	}

	m := make(map[string]A)

	m["1"] = A{data: []Data{{false, "1"}, {false, "11"}}}
	m["2"] = A{data: []Data{{false, "2"}, {false, "22"}}}

	for k, v := range m {
		// v = 5
		for i := 0; i < len(v.data); i++ {

			if v.data[i].Name == "1" {
				fmt.Println(v.data[i].Status, m[k].data[i].Status)
				v.data[i].Status = true
				fmt.Println(v.data[i].Status, m[k].data[i].Status)
			}
		}

	}

	fmt.Println(m)

}
