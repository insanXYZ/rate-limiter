package utils

import "time"

func SetInterval(interval time.Duration, fnc func()) {
	tick := time.Tick(interval)
	for _ = range tick {
		fnc()
	}
}
