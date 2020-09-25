package mytarget_sdk

import (
	"encoding/json"
	"errors"
	"time"
)

type CustomFloat64 struct {
	Float64 float64
}

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}

func (cf *CustomFloat64) Value() float64 {
	return cf.Float64
}

type CustomTimestamp struct {
	Timestamp time.Time
}

func (ct *CustomTimestamp) Value() time.Time {
	return ct.Timestamp
}

func (ct *CustomTimestamp) UnmarshalJSON(data []byte) error {
	var ts string
	err := json.Unmarshal(data, &ts)
	if err != nil {
		return errors.New("CustomTimestamp: UnmarshalJSON: " + err.Error())
	}
	ct.Timestamp, err = time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		return errors.New("CustomTimestamp: UnmarshalJSON ParseTime: " + err.Error())
	}
	return nil
}

type CustomDate struct {
	Date time.Time
}

func (ct *CustomDate) Value() time.Time {
	return ct.Date
}

func (ct *CustomDate) UnmarshalJSON(data []byte) error {
	var ts string
	err := json.Unmarshal(data, &ts)
	if err != nil {
		return errors.New("CustomTimestamp: UnmarshalJSON: " + err.Error())
	}
	ct.Date, err = time.Parse("2006-01-02", ts)
	if err != nil {
		return errors.New("CustomTimestamp: UnmarshalJSON ParseTime: " + err.Error())
	}
	return nil
}
