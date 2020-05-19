// Package utils is use for all generic export types or func or etc
package utils

import (
	"errors"
	"unicode/utf8"
)

// Date struct type
type Date struct {
	year  int
	month int
	day   int
}

// SetYear setter methods
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth setter method
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay setter method
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

//getters method

// Year function to get year value
func (d *Date) Year() int {
	return d.year
}

// Month function to get month value
func (d *Date) Month() int {
	return d.month
}

// Day function to get day value
func (d *Date) Day() int {
	return d.day
}

// Event struct type used for events
type Event struct {
	title string
	Date
}

// SetTitle this is setter for the title
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 32 {
		return errors.New("too long string")
	}
	e.title = title
	return nil

}

// Title this is getter for title
func (e *Event) Title() string {
	return e.title

}
