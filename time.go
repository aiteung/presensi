package presensi

import (
	"fmt"
	"time"
)

func GetFirstLastDateCurrentMonth() (firstOfMonth, lastOfMonth time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth = firstOfMonth.AddDate(0, 1, -1)

	fmt.Println(firstOfMonth)
	fmt.Println(lastOfMonth)
	return

}
