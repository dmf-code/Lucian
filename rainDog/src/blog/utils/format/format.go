package format

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

const (
	TimeFormat = "2006-01-02 15:04:05"
)

// 格式化时区时间
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//fmt.Println("MarshalJSON")
	//fmt.Println(t)
	//fmt.Println(t.Format(TimeFormat))
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(formatted), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	//fmt.Println("UnmarshalJSON")
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	//fmt.Println(now)
	*t = JSONTime{now}
	return
}

func (t JSONTime) Value() (driver.Value, error) {
	//fmt.Println("Value")
	//fmt.Println(t.Time)
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONTime) Scan(v interface{}) error {
	//fmt.Println("Scan")
	//fmt.Println(v)
	//fmt.Println(v.(time.Time))
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
