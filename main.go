package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", HandleFuncAdd)
	http.HandleFunc("/minus", HandleFuncMinus)
	http.HandleFunc("/divide", HandleFuncDivide)
	http.ListenAndServe(":8080", nil)

}

// , err :=
//
//	if err != nil {
//		fmt.Fprint(w,"параметр b не может быть равен нулю ")
//		return
//	}
func HandleFuncAdd(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	result := a + b
	fmt.Fprint(w, result)

}

func HandleFuncMinus(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	minusresult := a - b
	fmt.Fprint(w, minusresult)

}

func HandleFuncDivide(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	b, err := strconv.Atoi(p2)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}
	if b == 0 {
		err := errors.New("параметр b не может быть равен нулю ")
		fmt.Fprint(w, err.Error())
		return

	}

	divideresult := a / b
	fmt.Fprint(w, divideresult)

}

// multiple := a * b
// fmt.Fprint(w, ";"+" ", multiple)

// divide := a / b
// if err != nil {
// 	fmt.Fprint(w, "оппа ошибка")
// 	return
// }

// fmt.Fprint(w, ";"+" ", divide)

// if result := a - b; result != 0  {
// fmt.Fprint(w, resurelt)
// } else if result  := a * b; result != 0 {
// fmt.Fprint(w, result)
// }

// var num1 int
// var sign string
// var num2 int

// fmt.Print("введите первое число:  ")
// fmt.Scan(&num1)
// fmt.Print("Cимволы: ")
// fmt.Scan(&sign)
// fmt.Print("введите второе число:  ")
// fmt.Scan(&num2)

// if sign == "+" {
// 	fmt.Println("ответ: " + strconv.Itoa(num1+num2))
// } else if sign == "-" {
// 	fmt.Println("ответ: " + strconv.Itoa(num1-num2))
// } else if sign == "/" {
// 	fmt.Println("ответ: " + strconv.Itoa(num1/num2))
// } else if sign == "*" {
// 	fmt.Println("ответ: " + strconv.Itoa(num1*num2))
// }
