package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zanjs/y-mugg-v3/app/models"
)

// QueryStartEndTime is
func QueryStartEndTime(q models.QueryParams) models.QueryParamsTime {
	if q.EndTime == "" {
		q.EndTime = "2099-01-01 00:00:00"
		fmt.Println("endTime 为空")
	}

	if q.StartTime == "" {

		passingTime := "360"

		fmt.Println("StartTime 为空")
		now := time.Now()
		d, _ := time.ParseDuration("-" + passingTime + "h")
		d15 := now.Add(d)
		stime := d15.String()
		timeArr := strings.Split(stime, "-")

		year := timeArr[0]
		month := timeArr[1]
		day := strconv.Itoa(d15.Day())

		if len(day) == 1 {
			day = "0" + day
		}

		fmt.Println("15天:", year, month, day)
		fmt.Println(stime)
		fmt.Println(timeArr[0])
		fmt.Println(day)

		q.StartTime = year + "-" + month + "-" + day + " 00:00:00"
		fmt.Println(q.StartTime)
	}

	timeLayout := "2006-01-02 15:04:05"

	startTime, _ := time.Parse(timeLayout, q.StartTime)
	endTime, _ := time.Parse(timeLayout, q.EndTime)

	queryTime := models.QueryParamsTime{}
	queryTime.EndTime = endTime
	queryTime.StartTime = startTime

	return queryTime
}

// QueryStartDay is
func QueryStartDay(q models.QueryParams) models.QueryParamsTime {
	if q.EndTime == "" {
		q.EndTime = "2099-01-01 00:00:00"
		fmt.Println("endTime 为空")
	}

	if q.Day == 0 {
		q.Day = 30
	}

	passingTime := strconv.Itoa(q.Day * 24)

	fmt.Println("StartTime 为空")
	now := time.Now()
	d, _ := time.ParseDuration("-" + passingTime + "h")
	d15 := now.Add(d)
	stime := d15.String()
	timeArr := strings.Split(stime, "-")

	year := timeArr[0]
	month := timeArr[1]
	day := strconv.Itoa(d15.Day())

	if len(day) == 1 {
		day = "0" + day
	}

	fmt.Println(q.Day, "天:", year, month, day)
	fmt.Println(stime)
	fmt.Println(timeArr[0])
	fmt.Println(day)

	q.StartTime = year + "-" + month + "-" + day + " 00:00:00"
	fmt.Println(q.StartTime)

	timeLayout := "2006-01-02 15:04:05"

	startTime, _ := time.Parse(timeLayout, q.StartTime)
	endTime, _ := time.Parse(timeLayout, q.EndTime)

	queryTime := models.QueryParamsTime{}
	queryTime.EndTime = endTime
	queryTime.StartTime = startTime

	fmt.Println(startTime)
	fmt.Println("queryTime TIME")
	fmt.Println(queryTime)

	return queryTime
}
