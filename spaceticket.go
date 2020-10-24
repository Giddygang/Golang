package main
import ("fmt"
        "math/rand"
	"time"
	)
 func main() {
	const adayinsecs = 86400
	const distance = 62100000
		spaceline := " "
		trip := " "
		fmt.Println("Spaceline        Days    Trip    Price")
		fmt.Println("======================================")
 for count := 0 ; count < 10 ;  count++ {
	switch rand.Intn(3) {
		case 0:
		spaceline = "Space Adventures"
		case 1:
		spaceline = "SpaceX"
		case 2:
		spaceline = "Virgin Galactic"
		}
		speed := rand.Intn(15) + 16
		days := distance/speed/adayinsecs
		price :=  20.0 + speed
	 num := rand.Intn(2) ; if num == 1 {
			trip = "One-way"
		} else {
			trip = "Round-trip"
			price = price * 2
	}
		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline ,  days ,  trip ,  price)
  		time.Sleep(time.Second * 3) 
  }

}
