package core

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"kada/server/log"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Time time.Time

func Now() Time {
	return Time(time.Now())
}

func (o Time) IsZero() bool {
	t := time.Time(o)
	return t.IsZero()
}

func (o *Time) UnmarshalBSONValue(p bsontype.Type, data []byte) (err error) {
	decoder, err := bson.NewDecoder(bsonrw.NewBSONValueReader(p, data))
	if err != nil {
		return err
	}

	var t time.Time
	if err := decoder.Decode(&t); err != nil {
		return err
	}
	log.Debug("[UnmarshalBSONValue] %v, %v", t.UTC(), t.Local())
	*o = Time(t.Local())
	return
}

func (o Time) MarshalBSONValue() (bsontype.Type, []byte, error) {
	t := time.Time(o)
	log.Debug("[MarshalBSONValue] %v, %v", t.UTC(), t.Local())
	return bson.MarshalValue(time.Time(o).Local())
}

func (o *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*o = Time(now)
	return
}

func (o Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(o).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}
