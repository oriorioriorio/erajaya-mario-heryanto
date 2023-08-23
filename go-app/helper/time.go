package helper

import "time"

func GetTimeNowWithLocation(locationName string) (time.Time, error) {
	var timeNow time.Time

	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return timeNow, err
	}

	return time.Now().In(loc), nil
}
