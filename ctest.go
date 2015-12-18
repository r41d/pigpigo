package main // code.bitsetter.de/tk/gpigo

import ( 
	"fmt"
	"log"
)

// int sum(int a, int b) {
//   return a+b;
// }
import "C"

func main() {
	s, err := C.sum(21, 88)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hi", s )
}
