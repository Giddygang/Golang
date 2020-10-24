package main
import ("fmt" ; "time")
		var k int
		var c int
		type kelvin int
	func drawtable(c int , celsiustokelvin func() kelvin ) {
                for i := 40 ; i <= 100 ; i +=5 {
                        time.Sleep(time.Second * 2)
                        c = i
			k = c + 273
                fmt.Printf("|%vC %-12v | %15vK \n", c , " ",  k )
                }
        }
        func CelsiusToKelvin() kelvin {
                return kelvin(c + 273)
        }
        func main() {
                fmt.Println("============================================")
                fmt.Printf("| %-15v | %15v \n", "C" , "K"  )
                fmt.Println("============================================")
                drawtable(c , CelsiusToKelvin)
                fmt.Println("============================================")
}
