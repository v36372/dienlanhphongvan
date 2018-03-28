package model

import (
	"fmt"
	"strconv"
	"time"
)

type ImageFilename struct {
	shapeWidth
	value     string
	timestamp time.Time
}

func ParseImageFilename(value, shape, width string) (*ImageFilename, error) {
	var sec string
	if regexNewImageFilename.MatchString(value) {
		items := regexNewImageFilename.FindStringSubmatch(value)
		value = items[1]
		sec = items[2]
	} else if regexOldImageFilename.MatchString(value) {
		items := regexNewImageFilename.FindStringSubmatch(value)
		value = items[1]
		sec = items[2]
		shape = items[3]
		width = items[4]
	} else {
		return nil, fmt.Errorf("image filename: invalid format")
	}
	timestamp, err := parseTimestamp(sec)
	if err != nil {
		return nil, fmt.Errorf("image filename: invalid timestamp %v", err)
	}
	w, _ := strconv.Atoi(width)
	w = normalizeWidth(w)
	s := normalizeShape(shape)
	return &ImageFilename{
		value:     value,
		timestamp: timestamp,
		shapeWidth: shapeWidth{
			shape: s,
			width: w,
		},
	}, nil
}

func NewImageFilename(value string, timestamp time.Time) ImageFilename {
	return ImageFilename{
		value:     value,
		timestamp: timestamp,
	}
}

func (name ImageFilename) Name() string {
	return fmt.Sprintf("%s-%d", name.value, name.timestamp.Unix())
}

func (name ImageFilename) Path() string {
	return pathByTimestamp(name.Name(), name.timestamp)
}

func (name ImageFilename) OldPath() string {
	return oldPathByTimestamp(name.Name(), name.timestamp)
}

func (name ImageFilename) CachedPath() string {
	return fmt.Sprintf("%s_%s_%d", pathByTimestamp(name.Name(), name.timestamp), name.shape, name.width)
}
