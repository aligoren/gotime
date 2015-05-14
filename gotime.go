package gotime

import (
	"fmt"
	"time"
)

type Now struct {
	dnow int
	ynow int
	mnow time.Month
}

func (n Now) TimeNow() string {

	nowHours := time.Now().Hour()
	nowMinutes := time.Now().Minute()
	nowSecond := time.Now().Second()
	nowTime := [3]int{nowHours, nowMinutes, nowSecond}
	timeReturn := fmt.Sprintf("%d:%d:%d", nowTime[0],nowTime[1],nowTime[2])
	
	return timeReturn
}

func (n Now) DayNow() int {
	n.dnow = time.Now().Day()
	return n.dnow

}

func (n Now) WeekNow() time.Weekday {
	weekNow := time.Now().Weekday()

	return weekNow
}

func (n Now) PreviousDay() string {
	day := time.Now().Day() - 1
	month := time.Now().Month()
	ago := time.Now().Weekday() - 1
	preday := fmt.Sprintf("%s %s %d", ago, month, day)
	return preday
}

func (n Now) NextDay() string {
	day := time.Now().Day() + 1
	month := time.Now().Month()
	later := time.Now().Weekday() + 1
	preday := fmt.Sprintf("%s %s %d", later, month, day)
	return preday
}

func (n Now) DayAgo(ago time.Weekday) time.Weekday {
	dago := time.Now().Weekday() - ago

	return dago
}

func (n Now) DayLater(later time.Weekday) time.Weekday {
	dlater := time.Now().Weekday() + later

	return dlater
}

func (n Now) MonthNow() time.Month {
	/*
	DataType
	reflect.TypeOf(time.Now().Month())
	*/
	n.mnow = time.Now().Month()
	return n.mnow
}

func (n Now) MonthAgo(ago time.Duration) string {

	months := [12]string{"January", "February", "March",
	"April", "May", "June", "July", "August", "September",
	 "October", "November", "November"}
	ast := time.Duration(time.Now().Month()) - ago
	var newTime time.Duration = time.Duration(ast)
	switch newTime.String() {
		case "1ns":return months[0]
		case "2ns":return months[1]
		case "3ns":return months[2]
		case "4ns":return months[3]
		case "5ns":return months[4]
		case "6ns":return months[5]
		case "7ns":return months[6]
		case "8ns":return months[7]
		case "9ns":return months[8]
		case "10ns":return months[9]
		case "11ns":return months[10]
		case "12ns":return months[11]
		default:return ""
	}
}

func (n Now) MonthLater(later time.Duration) string {
	months := [12]string{"January", "February", "March",
	"April", "May", "June", "July", "August", "September",
	 "October", "November", "November"}
	ast := time.Duration(time.Now().Month()) + later
	var newTime time.Duration = time.Duration(ast)
	switch newTime.String() {
		case "1ns":return months[0]
		case "2ns":return months[1]
		case "3ns":return months[2]
		case "4ns":return months[3]
		case "5ns":return months[4]
		case "6ns":return months[5]
		case "7ns":return months[6]
		case "8ns":return months[7]
		case "9ns":return months[8]
		case "10ns":return months[9]
		case "11ns":return months[10]
		case "12ns":return months[11]
		default:return ""
	}
}

func (n Now) YearNow() int {
	n.ynow = time.Now().Year()
	return n.ynow
}

func (n Now) YearAgo(ago int) int {
	yago := time.Now().Year() - ago

	return yago
}

func (n Now) YearLater(later int) int {
	ylater := time.Now().Year() + later

	return ylater
}

func (n Now) DayOfWeek(wn time.Weekday) time.Weekday {
	doweek := time.Weekday(wn)

	return doweek
}

func (n Now) DayOfYear() int {
	yday := time.Now().YearDay()

	return yday
}

func (n Now) WaitSec(wait time.Duration) {

	timer := time.NewTimer(time.Second * wait)
	<- timer.C
	
}

func (n Now) WaitMS(wait time.Duration) {
	timer := time.NewTimer(time.Millisecond * wait)
	<- timer.C
}

func (n Now) WaitNS(wait time.Duration) {
	timer := time.NewTicker(time.Nanosecond * wait)
	<- timer.C
}

func (n Now) WaitMicro(wait time.Duration) {
	timer := time.NewTimer(time.Microsecond * wait)
	<- timer.C
}

func (n Now) WaitMin(wait time.Duration) {
	timer := time.NewTimer(time.Minute * wait)
	<- timer.C
}

func (n Now) WaitHour(wait time.Duration) {
	timer := time.NewTimer(time.Hour * wait)
	<- timer.C
}