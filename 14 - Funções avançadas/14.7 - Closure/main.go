package main

<<<<<<< HEAD
import "fmt"

func closure() func() {
	text := "AAAAA"

	funcao := func() {
		fmt.Println(text)
	}

	return  funcao
}

func main() {
	text := "BBBB"
	fmt.Println(text)

	newFunc := closure()
	newFunc()

}
=======
func main() {

}
>>>>>>> 686038d95814ea60fe1645f6ced9004428a9a07e
