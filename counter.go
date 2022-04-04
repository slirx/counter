package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/user"
	"path"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

func parse(t time.Time, d Diff) string {
	now := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	t2 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	diff := t1.Sub(t2)

	switch d {
	case DiffDays:
		days := int(math.Abs(diff.Hours()) / 24)
		return strconv.Itoa(days) + " day" + format(days)
	case DiffWeeks:
		weeks := int(math.Abs(diff.Hours()) / 24 / 7)
		return strconv.Itoa(weeks) + " week" + format(weeks)
	}

	return ""
}

func format(in int) string {
	if in != 1 {
		return "s"
	}

	return ""
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.Open(path.Join(usr.HomeDir, ".config", "counter", "config.yml"))
	if err != nil {
		log.Fatal(err.Error())
	}

	var conf Config
	decoder := yaml.NewDecoder(f)

	if err = decoder.Decode(&conf); err != nil {
		log.Fatal(err.Error())
	}

	if err = f.Close(); err != nil {
		log.Fatal(fmt.Errorf("closing config file: %w", err))
	}

	for _, item := range conf.Items {
		fmt.Printf("%s: %s\n", item.Label, parse(item.Date.Time, item.Diff))
	}
}
