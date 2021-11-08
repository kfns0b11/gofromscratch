package main

func main() {
	// defer func() {
	// 	if p := recover(); p != nil {
	// 		fmt.Println(p)
	// 	}
	// }()
	defer recover()
	panic("ooo")
}
