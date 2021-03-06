package main

import (
	"fmt"
	"github.com/aligoren/gotime"
)

func main() {
	
	s := gotime.Now{}

	fmt.Println(s.TimeNow()) // Now Time 14:55:29 (My current time)
	fmt.Println(s.DayNow()) // Day Now (14th)
	fmt.Println(s.WeekNow()) // Weekday (Thursday)
	fmt.Println(s.PreviousDay()) // Wednesday May 13
	fmt.Println(s.NextDay()) // Friday May 15
	fmt.Println(s.DayAgo(3)) // 3 Days Ago: Monday
	fmt.Println(s.DayLater(2)) // 2 Days Later: Saturday
	fmt.Println(s.MonthNow()) // Now Month (May)
	fmt.Println(s.MonthAgo(3)) // 3 Months Ago: February
	fmt.Println(s.MonthLater(3)) // 3 Months Later: August
	fmt.Println(s.YearNow()) // Year Now: 2015
	fmt.Println(s.YearAgo(5)) // 5 Years Ago: 2010
	fmt.Println(s.YearLater(5)) // 5 Years Later: 2020
	fmt.Println(s.DayOfWeek(3)) // Day of Week 3th: Wednesday
	fmt.Println(s.DayOfYear()) // Day of Year 134th
	s.WaitSec(3) // Wait for 3 Seconds
	s.WaitMS(300) // Wait for 300 Milliseconds
	s.WaitMicro(5000) // Wait for 5000 Microsecond
	s.WaitNS(3000) // Wait for 3000 Nanoseconds
	s.WaitMin(3) // Wait for 3 Minutes
	s.WaitHour(5) // Wait for 5 Hours
}