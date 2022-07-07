package main

import (
	"fmt"
	"sync"
)

// 닫힘을 의미하는 Closure

// 일반적으로 함수 외부에서 작성된 데이터를 함수에서 사용하기 위해 매개변수를 사용하여 받은 값을 수정한 후
// 이를 리턴하는 방식을 사용한다.
// 그러나 Closure 기법을 사용하면 외부에서 작성된 지역변수를 함수 내에서 마음대로 접근하는 것이 가능하다.
// Golang에서 Closure는 익명 함수(Anonymous functions)와 비슷한 개념으로 이름 없이 한줄로 함수를 정의하고 싶을때 유용하다.
// 이를 통해 함수 안에서 함수를 선언 및 정의할 수 있고, 바깥쪽 함수에 선언된 변수에도 접근할 수도 있다.

// 1. 함수를 지역에서 선언하고 구현하여 외부 변수를 수정가능하게 함
func basic() {

	val := 0
	f := func() {
		val++
	}

	f()
	fmt.Println(val)
	f()
	fmt.Println(val)
}

// 생성기(generator) 란 하나의 분리된 캡슐화된 객체를 생성하는 패턴이라고 생각할 수 있다

func generator() func() int {
	var val int
	return func() int {
		val++
		return val
	}
}

func generatorII() func(a, b int) int {
	var sum int
	return func(a, b int) int {
		sum += a * b
		return sum
	}
}
func intSeq() func() int {

	i := 0
	return func() int {
		i += 1
		return i
	}

}

func main() {

	basic()

	gen1 := generator()
	gen2 := generator()

	fmt.Println(gen1(), gen1(), gen1(), gen2(), gen2(), gen1(), gen2())

	// Explain
	// generator는 int type의 val 변수를 가지고 함수를 반환하는 Closure 함수이다.
	// gen1과 gen2는 변수에 generator 함수를 담았고, 객체는 서로 분리된 val을 가지는 것을 확인할 수 있다.

	sum := func(a int, b int) int {
		return a + b
	}
	result := sum(1, 2)
	fmt.Println(result)

	cal := generatorII()
	fmt.Println(cal(1, 3))
	fmt.Println(cal(5, 2))

	// 또한 해당 변수 val은 함수 호출이 지난 후에 소멸하지 않고, 호출할 때마다 지속적으로 사용할 수 있기 때문에
	// 함수를 선언했을 때의 환경을 유지할 수 있다.
	// 이를 통해 프로그래밍의 흐름을 제어하고, 함수형 프로그래밍 장점을 가질 수 있다.

	nextInt := intSeq()
	fmt.Println(nextInt())
	// from val after nextInt() exec
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// init
	newInts := intSeq()
	fmt.Println(newInts())

	// Closure with Go routine
	// 외부 변수의 값을 복사하지 않고 주소를 참조함
	var wg sync.WaitGroup
	val := "memory 1"
	wg.Add(1)
	go func() {
		defer wg.Done()
		val = "memory 2"
	}()
	wg.Wait()
	fmt.Println(val) // memory 2

	// 주소를 참조하기 때문에 합류지점 (go 루틴에 값 복사)을 정해주지 않는다면
	// 일반적으로 for문의 고루틴이 시작도 하기전에
	// 메인 고루틴이 종료된다.
	// loop variable num captured by func literal
	for _, num := range []int{1, 2, 3} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Heap 메모리 영역에 슬라이스의 마지막 값 (3)이
			// 저장돼 있기 때문에 마지막 값(3)만 반복 출력
			fmt.Printf("%d ", num)
		}()
	}
	wg.Wait()

	// ->
	for _, num := range []int{1, 2, 3} {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("%d ", num) // Go루틴 실행 순서에 따라 배열의 순서가 출력에 보장되지는 않음
		}(num) // 참조에 의한 값 전달이 아니라 직접 복사해 주면 된다
	}
	wg.Wait()
}

// Reference
// https://judo0179.tistory.com/86
// https://roy-jang.tistory.com/31?category=963134
// https://roy-jang.tistory.com/41?category=963134
