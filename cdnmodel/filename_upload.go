package model

import (
	"fmt"
	"time"
)

type UploadFilename struct {
	value     string
	timestamp time.Time
}

func ParseUploadFilename(value string) (*UploadFilename, error) {
	items := regexUploadFilename.FindStringSubmatch(value)
	fmt.Println("value :%s", value)
	if len(items) != 3 {
		return nil, fmt.Errorf("upload filename: invalid format")
	}
	sec := items[1]
	value = items[2]
	timestamp, err := parseTimestamp(sec)
	if err != nil {
		return nil, fmt.Errorf("upload filename: invalid timestamp %v", err)
	}
	return &UploadFilename{
		value:     value,
		timestamp: timestamp,
	}, nil
}

func NewUploadFilename(value string, timestamp time.Time) UploadFilename {
	return UploadFilename{
		value:     value,
		timestamp: timestamp,
	}
}

func (name UploadFilename) Name() string {
	return fmt.Sprintf("%d_%s", name.timestamp.Unix(), name.value)
}

func (name UploadFilename) Path() string {
	return pathByTimestamp(name.Name(), name.timestamp)
}
