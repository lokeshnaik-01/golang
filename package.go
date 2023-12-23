package main
import (
	"fmt"
	"math"
	"strconv"
	"os"
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

	noReturn("\nThis function call doesn't return any value")
	fmt.Println(calculatePower(10, 20))

	
	os.WriteFile("balance.txt", []byte(formattedFV), 0644) //0644 is the file permisson

	data, _ := os.ReadFile("balance.txt")
	fmt.Println(string(data))

	floatVal := "65.655"
	val, _ := strconv.ParseFloat(floatVal, 64)
	fmt.Println("floatval ", val)
}

func noReturn(text string) {
	fmt.Println(text)
}

func calculatePower(val1 int, val2 int) (int, int) {
	f1 := 10
	f2 := val1 / val2
	return f1, f2
}