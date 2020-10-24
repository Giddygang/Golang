package main
import "fmt"

	func main() {
        fmt.Println("What decipher are you looking for? Ceasardecipher, ROT13decipher or Vigeneredecipher")
           var input string
                fmt.Scanln(&input)
        if (input == "Ceasardecipher") { fmt.Println("Input message: ")
                         message := " "
                        fmt.Scanln(&message)
                for i := 0; i < len(message); i++ {
                        c := message[i]
                        if c >= 'a' && c <= 'z' {
                         c = c - 3
                        if c < 'a' {
                         c = c + 26
                                }
                        }
                        fmt.Printf("%c", c )
                }
              } else if (input == "ROT13decipher") {
                        fmt.Println("Input text: ")
                         text := " "
                        fmt.Scanln(&text)
                        for i := 0; i < len(text); i++ {
                        c := text[i]
                        if c >= 'a' && c <= 'z' {
                         c = c - 13
                        if c < 'a' {
                         c = c + 26
                                }
                    }
                        fmt.Printf("%c" , c)
                }

        } else if (input == "Vigeneredecipher") {
                fmt.Println("Input a message: ")
                message := " "
                fmt.Scanln(&message)
                fmt.Println("Input a Keyword: ")
                Keyword := " "
                fmt.Scanln(&Keyword)
                keyIndex := 0
		var decipher string
        for i := 0; i<len(message); i++ {
                c := message[i]
                if c >= 'a' && c <= 'z' {
                      //  c += 'a'
                        k := Keyword[keyIndex]
                        c = (c - k) + 'a'
                 decipher += string (c)
		 keyIndex++
                keyIndex %= len(Keyword)
                }
          //      fmt.Println(decipher)
        }
		  fmt.Println(decipher)
    }

}

