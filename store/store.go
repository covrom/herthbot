package store

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type People struct {
	Name string
	Next *People
}

type DayList struct {
	Name      string
	StartedAt time.Time
	StopAt    time.Time
	Head      *People
	Tail      *People
}

func NewDayList(start, stop time.Time) *DayList {
	dl := &DayList{
		Name:      start.Format("2 January 2006"),
		StartedAt: start,
		StopAt:    stop,
	}
	return dl
}

func IsCorrectName(name string) error {
	flds := strings.Fields(name)
	if len(flds) < 2 || len(flds) > 5 {
		return fmt.Errorf("name must contain at least 2 words and no more than 5")
	}
	for _, fld := range flds {
		for _, r := range fld {
			if !unicode.IsLetter(r) {
				return fmt.Errorf("name must contain only letters")
			}
		}
	}
	return nil
}

func (dl *DayList) TailAddPeople(name string) error {
	if err := IsCorrectName(name); err != nil {
		return err
	}
	p := &People{
		Name: name,
	}
	if dl.Tail != nil {
		dl.Tail.Next = p
	}
	if dl.Head == nil {
		dl.Head = p
	}
	dl.Tail = p
	return nil
}

func (dl *DayList) EnumPeople(f func(idx int, name string)) {
	i := 0
	for p := dl.Head; p != nil; p = p.Next {
		f(i, p.Name)
		i++
	}
}

func (dl *DayList) String() string {
	sb := &strings.Builder{}
	if dl.Head != nil {
		fmt.Fprintf(sb, "%s queue:\n", dl.Name)
		dl.EnumPeople(func(idx int, name string) {
			fmt.Fprintf(sb, "%d. %s\n", idx+1, name)
		})
	} else {
		fmt.Fprintf(sb, "%s queue is empty.", dl.Name)
	}
	return sb.String()
}
