package common

import (
	"encoding/json"
	"errors"
	"time"
)

const CustomTimestampFormat = "2006-01-02 15:04:05"
const CustomDateFormat = "2006-01-02"

type CustomFloat64 struct {
	Float64 float64
}

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64@UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64@UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}

func (cf *CustomFloat64) MarshalJSON() ([]byte, error) {
	jsonData, err := json.Marshal(cf.Float64)
	if err != nil {
		return nil, errors.New("CustomFloat64@MarshalJSON: " + err.Error())
	}
	return jsonData, err
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
		return errors.New("CustomTimestamp@UnmarshalJSON: " + err.Error())
	}
	ct.Timestamp, err = time.Parse(CustomTimestampFormat, ts)
	if err != nil {
		return errors.New("CustomTimestamp@UnmarshalJSON ParseTime: " + err.Error())
	}
	return nil
}

func (ct *CustomTimestamp) MarshalJSON() ([]byte, error) {
	if ct.Timestamp.IsZero() {
		return []byte(`""`), nil
	}
	formatted := ct.Timestamp.Format(CustomTimestampFormat)
	jsonData, err := json.Marshal(formatted)
	if err != nil {
		return nil, errors.New("CustomTimestamp@MarshalJSON: " + err.Error())
	}
	return jsonData, err
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
		return errors.New("CustomDate@UnmarshalJSON: " + err.Error())
	}
	ct.Date, err = time.Parse(CustomDateFormat, ts)
	if err != nil {
		return errors.New("CustomDate@UnmarshalJSON ParseTime: " + err.Error())
	}
	return nil
}

func (ct *CustomDate) MarshalJSON() ([]byte, error) {
	if ct.Date.IsZero() {
		return []byte(`""`), nil
	}
	jsonData, err := json.Marshal(ct.Date.Format(CustomDateFormat))
	if err != nil {
		return nil, errors.New("CustomDate@MarshalJSON: " + err.Error())
	}
	return jsonData, err
}
