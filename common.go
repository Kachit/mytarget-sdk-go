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

type CustomTimestamp struct {
	Timestamp time.Time
}

func (ct *CustomTimestamp) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &ct.Timestamp)
	if err != nil {
		return errors.New("CustomTimestamp: UnmarshalJSON: " + err.Error())
	}
	return nil
}
