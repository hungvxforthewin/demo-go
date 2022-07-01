package main

import "fmt"

func main() {

   a := 5
   
   thayDoiTrenGiaTri(a)
   fmt.Println(a)
   
   thayDoiTrenDiaChi(&a)
   fmt.Println(a)

}

func thayDoiTrenGiaTri(a int) {
   a = 100
}
func thayDoiTrenDiaChi(a *int){
   *a = 200
}

// Giải thích :
// Một hàm thayDoiTrenGiaTri() truyền vào giá trị biến a và thay đổi biến a = 100
// một hàm thayDoiTrenDiaChi() truyền vào địa chỉ biến a và thay đổi giá trị lưu trên địa chỉ đó *a = 200

// Ta có thể thấy khi dùng hàm thayDoiTrenGiaTri() thì giá trị biến a chỉ thay đổi ở trong hàm đó và giá trị biến a ngoài hàm main không thay đổi, 
// thực chất là do ta chỉ chuyền "bản sao" giá trị của a vào nên ta không thể tác động để thay đổi a
// khi dùng hàm thayDoiTrenDiaChi() thì giá trị biến a có thể thay đổi vì ta truyền địa chỉ biến a vào và có thể thay đổi trực tiếp giá trị trên địa chỉ đó.
// Do đó tùy vào mục đích bài toán ta có thể cân nhắc khi nào sử dụng con trỏ