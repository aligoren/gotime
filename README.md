#GoTime

GoTime is a time toolkit for golang

#Usage:

Check main file in test folder or follow this code:

**Import**

~~~~{.go}
import "github.com/aligoren/gotime"
~~~~

**Current Type Functions:**

    TimeNow() => Now Time
    DayNow() => Day Now
    WeekNow() => Weekday Now
    PreviousDay() => Previous Day (DayNow -1)
    NextDay() => Next Day (DayNow + 1)
    DayAgo() => X Days Ago (3 Day Ago)
    DayLater() => X Days Later (5 Day Later)
    MonthNow() => Month Now
    MonthAgo() => X Months Ago
    MonthLater() => X Months Later
    YearNow() => Current Year
    YearAgo() => X Years Ago (3 Years Ago: 2012)
    YearLater() => X Years Later (3 Years Later: 2018)
    DayOfWeek() => X Day of Week (Day of Week 3th: Wednesday)
    DayOfYear() => Current Now: 134th Day of Year

**Wait Functions:**

    WaitMicro() => Wait for X Microseconds
    WaitNS() => Wait for X Nanoseconds
    WaitMS() => Wait for X Milliseconds
    WaitSec() => Wait for X Seconds
    WaitMin() => Wait for X Minutes
    WaitHour() => Wait for X Hours

#Sample Usage:

~~~~{.go}
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
~~~~