package sonyflake

import (
	"strconv"
	"time"
)

var ErrInvalidSnoyflake = "snoyflake init failed"

var defaultSonyflake *Sonyflake

type ID uint64

// Claims .
type Claims struct {
	ID        ID
	Msb       uint64
	Time      uint64
	Sequence  uint64
	MachineID uint64
}

func init() {
	var st Settings
	st.StartTime = time.Now()

	defaultSonyflake = NewSonyflake(st)

	if defaultSonyflake == nil {
		panic(ErrInvalidSnoyflake)
	}
}

// NextID .
func NextID() ID {
	id, _ := defaultSonyflake.NextID()
	return ID(id)
}

// Uint64 .
func (f ID) Uint64() uint64 {
	return uint64(f)
}

// Int64 .
func (f ID) Int64() int64 {
	return int64(f)
}

// String .
func (f ID) String() string {
	return strconv.FormatUint(uint64(f), 10)
}

// Decode .
func (f ID) Decode() Claims {
	r := Decompose(uint64(f))
	return Claims{
		ID:        f,
		Msb:       r["msb"],
		Time:      r["time"],
		Sequence:  r["sequence"],
		MachineID: r["machine_id"],
	}
}

// ParseInt64 .
func ParseInt64(i int64) ID {
	return ID(i)
}

// ParseUint64 .
func ParseUint64(i uint64) ID {
	return ID(i)
}

// ParseString .
func ParseString(s string) (ID, error) {
	f, err := strconv.ParseUint(s, 10, 64)
	return ID(f), err
}
