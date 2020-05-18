// Package utils is use for all generic export types or func or etc
package utils

import "errors"

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
