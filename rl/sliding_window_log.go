package rl

import (
	"fmt"
	"rate-limiter/utils"
	"strconv"
	"time"
)

type SlidingWindowLog struct {
	cache    []int64 // time.Now().Unix()
	capacity int
}

func NewSlidingWindowLog(capacity int) *SlidingWindowLog {
	return &SlidingWindowLog{
		cache:    make([]int64, 0),
		capacity: capacity,
	}
}

func (s *SlidingWindowLog) Add() int {
	s.cache = append(s.cache, time.Now().Unix())
	s.filterPreviousWindow()

	if len(s.cache) > s.capacity {
		return 429
	}

	return 200
}

func (s *SlidingWindowLog) filterPreviousWindow() {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location()).Unix()

	s.cache = utils.Filter(s.cache, func(i int64) bool {
		return i > start
	})

	fmt.Println("len cache = " + strconv.Itoa(int(len(s.cache))))
}
