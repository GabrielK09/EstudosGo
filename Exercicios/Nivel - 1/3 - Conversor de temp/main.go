package main

import "fmt"

func main() {
	cInFahrenheit := func(c float64) (float64, error) {
		return (c * 9 / 5) + 32, nil
	}

	cInKelvin := func(c float64) (float64, error) {
		return c + 273.15, nil
	}

	fInKelvin := func(f float64) (float64, error) {
		return (f-32)*5/9 + 273.15, nil
	}

	fInKelvinResult, err := fInKelvin(1.00)

	if err != nil {
		panic(err)
	}

	fmt.Println(cInFahrenheit(1.00))
	fmt.Println(cInKelvin(1.00))
	fmt.Printf("%f.2\n", fInKelvinResult)

}
