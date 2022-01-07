package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestTimeSliceSort(t *testing.T) {
	var reviews_data_map = make(map[string]ReviewsData)
	reviews_data_map["1"] = ReviewsData{Date: time.Now().Add(12 * time.Hour)}
	reviews_data_map["2"] = ReviewsData{Date: time.Now()}
	reviews_data_map["3"] = ReviewsData{Date: time.Now().Add(24 * time.Hour)}
	//Sort the map by date
	date_sorted_reviews := make(timeSlice, 0, len(reviews_data_map))
	for _, d := range reviews_data_map {
		date_sorted_reviews = append(date_sorted_reviews, d)
	}
	fmt.Println(date_sorted_reviews)
	sort.Sort(date_sorted_reviews)
	fmt.Println(date_sorted_reviews)
}
