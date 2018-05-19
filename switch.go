package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write", i, "as")

	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("its the weekend")
		default:
			fmt.Println("its a weekdat")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("its before noon")
		default:
			fmt.Println("its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool:
				fmt.Println("is a bool")
			case int:
				fmt.Println("is a int")
			default:
				fmt.Println("")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}