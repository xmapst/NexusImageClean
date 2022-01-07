package main

import "time"

type ReviewsData struct {
	Date           time.Time
	ID             string
	Name           string
	repositoryName string
}

type timeSlice []ReviewsData

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].Date.Before(p[j].Date)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
