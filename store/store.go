package store

import (
	"encoding/json"
	"fmt"
	"os"
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
		if len(fld) > 50 {
			return fmt.Errorf("each part of name must be shorter than 50 characters")
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

type DayListStore struct {
	Name      string
	StartedAt string
	StopAt    string
	Peoples   []string
}

func (dl *DayList) Save(fn string) error {
	ds := DayListStore{
		Name:      dl.Name,
		StartedAt: dl.StartedAt.Format(time.RFC3339),
		StopAt:    dl.StopAt.Format(time.RFC3339),
	}
	dl.EnumPeople(func(idx int, name string) {
		ds.Peoples = append(ds.Peoples, name)
	})
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(ds)
}

func (dl *DayList) Load(fn string) error {
	ds := &DayListStore{}
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(ds)
	if err != nil {
		return err
	}
	dl.Name = ds.Name
	dl.StartedAt, err = time.Parse(time.RFC3339, ds.StartedAt)
	if err != nil {
		return err
	}
	dl.StopAt, err = time.Parse(time.RFC3339, ds.StopAt)
	if err != nil {
		return err
	}
	dl.Head = nil
	dl.Tail = nil
	for _, name := range ds.Peoples {
		if err := dl.TailAddPeople(name); err != nil {
			return err
		}
	}
	return nil
}
