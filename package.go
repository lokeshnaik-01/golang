package main
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World");
	fmt.Println(`back ticks can be used to print`);

	const investmentAmount = 1000
	// constant can't be changed
	var expectedReturnRate float64= 5.5
	var years float64= 10

	fmt.Print("Input the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println("Future value :",  futureValue, " years ", years)
	fmt.Printf("futureValue: %.2f number of years are %v \n", futureValue, years)

	formattedFV := fmt.Sprintf("Future value: %.2f", futureValue)
	fmt.Print(formattedFV, "\nNumber of years ", years)

}