package main

import (
	"fmt"
	// "os"
)

// func g() int {
// 	fmt.Println("g()")
// 	return 0
// }
// func f(int) {
// 	fmt.Println("f()")
// }

func f(nums [6]int) {
	fmt.Printf("%v\n", nums)
}

// again some test commits
func main() {

	nums := [...]int{0, 1, 2, 3, 4, 5}
	nums[0] += 1
	defer f(nums)

	// defer f(g())
	// fmt.Println("Привіт, світ!")

	// GetFileContent()

	// path := "/home/andrey/password"
	// getFile(path)
}

// func GetFileContent() string {
// 	file, _ := os.Open("/home/bee/Документи/test.txt")
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	file.Read(data)
// 	dataString := string(data)
// 	//fmt.Printf(dataString)
// 	if dataString == "Hello Go!" {
// 		return "Mayank"
// 	} else {
// 		return "Anshul"
// 	}
// }

// func getFile(path string) {
// 	f, _ := os.Open(path)
// 	defer f.Close()

// 	s := bufio.NewScanner(f)
// 	for s.Scan() {
// 		fmt.Println(s.Text())
// 	}
// }
