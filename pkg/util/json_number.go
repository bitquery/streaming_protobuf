package util

import (
	"fmt"
	"strconv"
)

type JsonNumber uint64

func (number *JsonNumber) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		return
	}
	str := string(data)
	if len(data) >= 2 && data[0] == 34 && data[len(data)-1] == 34 {
		str = string(data[1 : len(data)-1])
	}
	if len(str) == 0 {
		return
	}
	anum, err := strconv.ParseUint(str, 10, 64)
	*number = JsonNumber(anum)
	return
}

func (number JsonNumber) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", number)), nil
}

func (number JsonNumber) Uint64() uint64 {
	return uint64(number)
}

func (number JsonNumber) Int64() int64 {
	return int64(number)
}

func (number JsonNumber) String() string {
	return strconv.FormatUint(uint64(number), 10)
}
