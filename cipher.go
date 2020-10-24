package main
import "fmt"

	func main() {
	fmt.Println("What cipher are you looking for? Ceasarcipher, ROT13cipher or Vigenerecipher")
	   var input string
		fmt.Scanln(&input)
	if (input == "Ceasarcipher") { fmt.Println("Input message: ")
                         message := " "
                        fmt.Scanln(&message)
                for i := 0; i < len(message); i++ {
                        c := message[i]
                        if c >= 'a' && c <= 'z' {
                         c = c + 3
                        if c > 'z' {
                         c = c - 26
                                }
                        }
                        fmt.Printf("%c", c )
                }
              } else if (input == "ROT13cipher") {
			fmt.Println("Input text: ")
                         text := " "
                        fmt.Scanln(&text)
               	 	for i := 0; i < len(text); i++ {
                        c := text[i]
                        if c >= 'a' && c <= 'z' {
                         c = c + 13
                        if c > 'z' {
                         c = c - 26
                                }
                    }
                        fmt.Printf("%c" , c)
		}

        } else if (input == "Vigenerecipher") {
		fmt.Println("Input a message: ")
		message := " "
		fmt.Scanln(&message)
		fmt.Println("Input a Keyword: ")
		Keyword := " "
		fmt.Scanln(&Keyword)
		keyIndex := 0
	for i := 0; i<len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 'a'
			k := Keyword[keyIndex]
			c = (c + k)
		keyIndex++
		keyIndex %= len(Keyword)
		}
	 	fmt.Printf("%c", c)
	}
	}

}
