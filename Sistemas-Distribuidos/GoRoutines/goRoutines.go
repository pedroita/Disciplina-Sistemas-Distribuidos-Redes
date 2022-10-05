package main

//import "time"
import "sync"

// func main() {
// 	go helloWold()
// 	time.Sleep(time.Second)
// }

// func helloWold() {
// 	println("hello,Wold")
// }
//função anonima
func main() {
	go func(msg string) {
		println(msg)
		wait.Done()
	}("hello, Wold")
	//time.Sleep(time.Second)
	wait.Wait()
}
