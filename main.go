// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leachim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2018-15-09 17:25<mklimoff222@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

// This is a funny package that does my age countdown (much like Chromium extension)
// The purpose of this thing is to make it work on Conky.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const PER_YEAR = uint64(uint64(31556952000000430))

func calculate(year, month, day int) (string, string) {
	birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	item := strings.Split(fmt.Sprintf("%.09f", float64(time.Now().Sub(birthday).Nanoseconds())/float64(PER_YEAR)), ".")
	return item[0], item[1]
}

func main() {
	param := ""
	if len(os.Args) > 1 {
		param = os.Args[1]
	}
	years, amount := calculate(1992, 12, 12)

	switch param {
	case "years":
		fmt.Print(years)
		return
	case "nanoseconds":
		fmt.Print(amount)
		return
	}
	fmt.Print(years, ".", amount)
	return
}
