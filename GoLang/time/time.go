package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Golang")

	// lấy thời gian hiện tại
	presentTime := time.Now()
	// fmt.Println(presentTime)

	//Đổi định dạng thời gian
	fmt.Println(presentTime.Format(time.RFC3339Nano))

	//Tạo thời gian
	// createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	// fmt.Println(createdDate)
	// fmt.Println(presentTime.Format("01-02-2006 Monday"))
}
