package app

import (
	"time"

	"github.com/covrom/herthbot/store"
)

func StartStopTomorrow() (start, stop time.Time) {
	tnow := time.Now().Local().Truncate(24 * time.Hour).Add(24 * time.Hour)
	t6 := time.Date(tnow.Year(), tnow.Month(), tnow.Day(), 6, 0, 0, 0, time.Local)
	t18 := time.Date(tnow.Year(), tnow.Month(), tnow.Day(), 18, 0, 0, 0, time.Local)
	return t6, t18
}

// list opened tomorrow at 6:00 and closed at 18:00
func (a *App) NewDayList() {
	start, stop := StartStopTomorrow()
	a.currentList = store.NewDayList(start, stop)
}
