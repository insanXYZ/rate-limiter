package utils

import "time"

func SetInterval(interval time.Duration, fnc func()) {
	tick := time.Tick(interval)
	for _ = range tick {
		fnc()
	}
}

func Filter[T comparable](items []T, condition func(item T) bool) []T {
	var res []T

	for _, v := range items {
		if condition(v) == true {
			res = append(res, v)
		}
	}

	return res
}
