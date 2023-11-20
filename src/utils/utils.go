package utils

import "fmt"

func LoadVariables(N *int, T *int, E *int) {
	fmt.Print("Enter N (number of items to load): ")
	_, err := fmt.Scanln(N)
	if err != nil {
		fmt.Println("Error reading N:", err)
		return
	}

	fmt.Print("Enter T (number of threads): ")
	_, err = fmt.Scanln(T)
	if err != nil {
		fmt.Println("Error reading T:", err)
		return
	}

	fmt.Print("Enter E (% of WRITE): ")
	_, err = fmt.Scanln(E)
	if err != nil {
		fmt.Println("Error reading T:", err)
		return
	}
}
