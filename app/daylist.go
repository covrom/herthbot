package app

import (
	"time"

	"github.com/covrom/herthbot/store"
)

func StartStopTomorrow() (start, stop time.Time) {
	tnow := time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour)
	t6 := tnow.Add(6 * time.Hour)
	t18 := tnow.Add(18 * time.Hour)
	return t6, t18
}

// list opened tomorrow at 6:00 and closed at 18:00
func (a *App) NewDayList() {
	start, stop := StartStopTomorrow()
	a.currentList = store.NewDayList(start, stop)
}
