package main

import (
	"strings"
	"time"
)

type Diff string

const (
	DiffDays  Diff = "days"
	DiffWeeks Diff = "weeks"
)

type Item struct {
	Label string `yaml:"label"`
	Date  Date   `yaml:"date"`
	Diff  Diff   `yaml:"diff"`
}

type Date struct {
	Time time.Time
}

func (d *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string

	err := unmarshal(&buf)
	if err != nil {
		return nil
	}

	t, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}

func (d Date) MarshalYAML() (interface{}, error) {
	return d.Time.Format("2006-01-02"), nil
}

type Config struct {
	Items []Item `yaml:"items"`
}
