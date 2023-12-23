package main
import (
	"fmt"
	"math"
	"strconv"
	"os"
	"errors"
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
	v1, v2, err := calculatePower(10, 20)
	if err == nil {
		fmt.Println(v1, v2)
	}
	
	os.WriteFile("balance.txt", []byte(formattedFV), 0644) //0644 is the file permisson

	data, _ := os.ReadFile("balance.txt")
	fmt.Println(string(data))

	floatVal := "65.655"
	val, _ := strconv.ParseFloat(floatVal, 64)
	fmt.Println("floatval ", val)

	data, err = os.ReadFile("balance1.txt")
	if(err!=nil) {
		fmt.Println("error is", err)
		panic("in panic code below it not executed")
		fmt.Println("this is not executed")
	}
	fmt.Println("this is not executed")
}

func noReturn(text string) {
	fmt.Println(text)
}

func calculatePower(val1 int, val2 int) (int, int, error) {
	if(val1!=10) {
		return 1,1, errors.New("Custom error")
	}
	f1 := 10
	f2 := val1 / val2
	return f1, f2, nil
}