// Pipeline Simulator OOO simple version 1.go
// The Super-Scalar Processor Simulator - simple out-of-order version, runs forever.
// A.Oram 2017
//fidninjjnfjfjfj
package main

//iigijb/dinagn
// Imported packages

import (
	"fmt" // for console I/O
	// for randomly creating opcodes
	"time" // for the random number generator and 'executing' opcodes
)

func TrafficLight(id int, allowedToBeRed <-chan bool, setAllowedToBeRed chan<- bool, PeopleWantToCross <-chan bool) {
	for {
		IGoGreen := <-allowedToBeRed

		if IGoGreen == true {

			fmt.Println(id, "is red")
			fmt.Println(id, " is red amber")
			for i := 0; i < 1; i++ {

				time.Sleep(time.Second * 1)
			}
			fmt.Println(id, "is green")
			for i := 0; i < 3; i++ {

				time.Sleep(time.Second * 1)
			}
			fmt.Println(id, "is amber")
			for i := 0; i < 1; i++ {

				time.Sleep(time.Second * 1)
			}
			fmt.Println(id, "is red")
			for i := 0; i < 1; i++ {

				time.Sleep(time.Second)
			}
			select {
			case <-PeopleWantToCross:
				fmt.Println("Please Cross the Road")
				time.Sleep(time.Second * 3)
				fmt.Println("Please cross quicker")
				time.Sleep(time.Second * 3)
				fmt.Println("Please do not cross")
			default:
				setAllowedToBeRed <- true
			}
			setAllowedToBeRed <- true

		}
		if IGoGreen == false {
			fmt.Println(id, "RECIVED", IGoGreen)
			fmt.Println(id, "is n Red")
			for i := 0; i < 6; i++ {

				time.Sleep(time.Second * 1)
			}
			fmt.Println(id, "Is on Red")
		}
	}
}

func crossing(PeopleWantToCross chan<- bool) {
	var button int

	for {

		fmt.Scanln(&button)
		if button == 1 {
			PeopleWantToCross <- true
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////
//  Main program, create required channels, then start goroutines in parallel.
//////////////////////////////////////////////////////////////////////////////////

func main() {
	// launch two go routines. Both want to listening on a channel. When one goes Red it tells the other to start going green
	//When revices red signal
	//Go amber for 3 seconds
	//Go green for 6
	//Go amber red for 3 second
	//Go red
	//Send signal

	// Set up required channels
	allowedToBeRed := make([]chan bool, 2)
	PeopleWantToCross := make(chan bool)

	for i := range allowedToBeRed { // Now set them up.
		allowedToBeRed[i] = make(chan bool)

	}

	fmt.Println("\n Start Traffic light processors ...")
	for i := 0; i < 2; i++ {
		x := (i + 1) % 2
		go TrafficLight(i, allowedToBeRed[i], allowedToBeRed[x], PeopleWantToCross)
	}
	go crossing(PeopleWantToCross)

	allowedToBeRed[0] <- true
	allowedToBeRed[1] <- false

	for {

	}

} // end of main
