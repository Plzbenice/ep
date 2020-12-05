package ep

import (
	"math"
	"time"
)

//get the list of date range start and stop date
func GetDateRange(start, stop string) ([]string, error) {
	list := make([]string, 0)

	startTime, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}

	stopTime, err := time.Parse("2006-01-02", stop)
	if err != nil {
		return nil, err
	}

	var flag int
	diffDay := int64(math.Abs(startTime.Sub(stopTime).Hours() / 24))
	if stopTime.Before(startTime) {
		flag = -1
	} else {
		flag = 1
	}

	for i := 0; i <= int(diffDay); i++ {
		list = append(list, startTime.AddDate(0, 0, i*flag).Format("2006-01-02"))
	}

	return list, nil
}
