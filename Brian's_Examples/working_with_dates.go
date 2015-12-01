// Dates and times are managed in Go with the time package
// Its time type encapsulates like everything you need

package main

import(
	"fmt"
	"time"
)

func main(){
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC) // day hour, minute, second, nanosecond, constant for location
	fmt.Printf("Go lanuched at %s\n", t)

	// current date and time
	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("The month is", t.Month())
	fmt.Println("the day is", t.Day())
	fmt.Println("the weekday is", t.Weekday())

	tomorrow := t.AddDate(0,0,1) // year, month, day
	fmt.Printf("Tomorrow is %v, %v, %v\n",
	tomorrow.Weekday(), tomorrow.Month(), tomorrow.Year())

}