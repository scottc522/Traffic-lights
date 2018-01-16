// Traffic light simulator
// b4005596 Scott Chapman

package main

// Imported packages

import (
	"fmt" // for console I/O

	"time" // for creating delays when lights change
)

func TrafficLight(id int, allowedToChange <-chan bool, setAllowedToChange chan<- bool, PeopleWantToCross <-chan bool) {
	for {
		//IGoGreen is set to the value in the allowedToChangeChanel
		IGoGreen := <-allowedToChange
		//Checks to see if this traffic light is allowed to change to green
		if IGoGreen == true {
			//Prints out is ID along with red as that will always be its starting state.
			fmt.Println(id, "is red")
			//Change to red amber and wait
			fmt.Println(id, " is red amber")
			time.Sleep(time.Second * 2)
			//change to green and wait
			fmt.Println(id, "is green")

			time.Sleep(time.Second * 6)
			//change to amber and wait
			fmt.Println(id, "is amber")

			time.Sleep(time.Second * 2)
			//change to red and wait
			fmt.Println(id, "is red")

			time.Sleep(time.Second * 2)
			//Select statement to check if people need to cross
			select {
			case <-PeopleWantToCross:
				//If channel contains data then it will tell them to cross and wait for 3 seconds
				fmt.Println("Please Cross the Road")
				time.Sleep(time.Second * 3)
				//Warn them that it will soon be unsafe to cross and wait
				fmt.Println("Please cross quicker")
				time.Sleep(time.Second * 3)
				//Tell them not to cross
				fmt.Println("Please do not cross")

			}
			//Pass true through the setAllowedToChange chanel which the other light will be listening on
			setAllowedToChange <- true

		}
		//If not allowed to go green
		if IGoGreen == false {
			//print out red
			fmt.Println(id, "is n Red")

		}
	}
}

func crossing(PeopleWantToCross chan<- bool) {
	//creates variable to hold button press
	var button int

	for {
		//reads a line of text from console into button variable
		fmt.Scanln(&button)
		//if that button is equal to the int 1 then pass true into the channel
		if button == 1 {
			PeopleWantToCross <- true
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////
//  Main program, create required channels, then start goroutines in parallel.
//////////////////////////////////////////////////////////////////////////////////

func main() {

	// Set up required channels
	//Allowed to be red will be used by one traffic light to tell the other one it is safe to change
	allowedToChange := make([]chan bool, 2)
	//PeopleWantToCross will be used to send an interupt to both traffic lights if someone whishes to cross the
	//road
	PeopleWantToCross := make(chan bool)
	//Set up the channels in the allowedToChange array
	for i := range allowedToChange { // Now set them up.
		allowedToChange[i] = make(chan bool)

	}
	//Let the console know that the traffic lights have started
	fmt.Println("\n Start Traffic light processors ...")
	for i := 0; i < 2; i++ {
		x := (i + 1) % 2
		go TrafficLight(i, allowedToChange[i], allowedToChange[x], PeopleWantToCross)
	}
	//Run the crossing function and pass in the PeopleWantToCross function
	go crossing(PeopleWantToCross)
	//Pass initial values into both channels to start the simulation
	allowedToChange[0] <- true
	allowedToChange[1] <- false

	for {
		//Run for ever
	}

} // end of main
