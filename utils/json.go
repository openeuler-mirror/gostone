package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/execption"
)

func Byte2Struct(b []byte, obj interface{}) {
	if err := json.Unmarshal(b, obj); err != nil {
		panic(execption.NewGoStoneError(execption.StatusInternalServerError, fmt.Sprintf("%v", err)))
	}
}

func Struct2Json(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(execption.NewGoStoneError(execption.StatusInternalServerError, fmt.Sprintf("%v", err)))
	}
	return string(b)
}

// RFC3339Milli describes a common time format used by some API responses.
const RFC3339Milli = "2006-01-02T15:04:05.000000Z"

type JSONRFC3339Milli time.Time

func (jt *JSONRFC3339Milli) UnmarshalJSON(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	t, err := time.Parse(RFC3339Milli, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339Milli(t)
	return nil
}

func (jt JSONRFC3339Milli) MarshalJSON() ([]byte, error) {
	msTime := time.Unix(time.Time(jt).Unix(), 0).UTC()
	var stamp = fmt.Sprintf("\"%s\"", msTime.Format(RFC3339Milli))
	return []byte(stamp), nil
}

const RFC3339MilliNoZ = "2006-01-02T15:04:05.999999"

type JSONRFC3339MilliNoZ time.Time

func (jt *JSONRFC3339MilliNoZ) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339MilliNoZ, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339MilliNoZ(t)
	return nil
}

func (jt JSONRFC3339MilliNoZ) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jt).Format(RFC3339MilliNoZ))
	return []byte(stamp), nil
}
