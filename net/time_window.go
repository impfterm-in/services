// Kiebitz - Privacy-Friendly Appointment Scheduling
// Copyright (C) 2021-2021 The Kiebitz Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package net

import (
	"time"
)

type TimeWindow struct {
	From int64
	To   int64
	Type string
}

func (t *TimeWindow) EqualTo(tw *TimeWindow) bool {
	if t.Type != tw.Type || t.From != tw.From || t.To != tw.To {
		return false
	}
	return true
}

type TimeWindowFunc func(int64) TimeWindow

func Second(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	from := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	to := from.Add(time.Second * 1)
	return TimeWindow{
		To:   to.UnixNano(),
		Type: "second",
	}
}

func Minute(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	from := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	to := from.Add(time.Minute * 1)
	return TimeWindow{
		From: from.UnixNano(),
		To:   to.UnixNano(),
		Type: "minute",
	}
}

func QuarterHour(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	q := t.Minute() / 15
	from := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), q*15, 0, 0, t.Location())
	to := from.Add(time.Minute * 15)
	return TimeWindow{
		From: from.UnixNano(),
		To:   to.UnixNano(),
		Type: "quarterHour",
	}
}

func Hour(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	from := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	to := from.Add(time.Hour * 1)
	return TimeWindow{
		From: from.UnixNano(),
		To:   to.UnixNano(),
		Type: "hour",
	}
}

func Day(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	from := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	to := from.AddDate(0, 0, 1)
	return TimeWindow{
		From: from.UnixNano(),
		To:   to.UnixNano(),
		Type: "day",
	}
}

func Week(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	from := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	wd := (int(t.Weekday()) - 1) % 7 // weekday starting from Monday
	if wd < 0 {
		wd += 7
	}
	from = from.AddDate(0, 0, -wd)
	to := from.AddDate(0, 0, 7)
	return TimeWindow{
		From: from.UnixNano(),
		To:   to.UnixNano(),
		Type: "week",
	}
}

func Month(value int64) TimeWindow {
	t := time.Unix(value/1e9, value%1e9).UTC()
	from := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	to := from.AddDate(0, 1, 0)
	return TimeWindow{
		From: from.UnixNano(),
		To:   to.UnixNano(),
		Type: "month",
	}
}

func MakeTimeWindow(t int64, twType string) TimeWindow {
	switch twType {
	case "second":
		return Second(t)
	case "minute":
		return Minute(t)
	case "hour":
		return Hour(t)
	case "quarterHour":
		return QuarterHour(t)
	case "day":
		return Day(t)
	case "week":
		return Week(t)
	case "month":
		return Month(t)
	}
	return TimeWindow{}
}

func (t *TimeWindow) Copy() TimeWindow {
	return TimeWindow{
		Type: t.Type,
		From: t.From,
		To:   t.To,
	}
}
