package main

import "fmt"

// "time"

// func main() {
// 	fmt.Println("When's Saturday?")
// 	today := time.Now().Weekday()
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }

func main(){
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

// Đối với Go thì không có từ khóa "while"

// Chỉ sử dụng ngoặc tròn () khi khai báo điều kiện cần ưu tiên trong if 

// Nguyên tắc của "defer" là sẽ thực thi câu lệnh ngay sau khi hàm đó kết thúc

// Defer có tính chất stacking (last-in-first-out) nghĩa là câu lệnh cuối cùng được thêm vào stack sẽ được thực thi đầu tiên
