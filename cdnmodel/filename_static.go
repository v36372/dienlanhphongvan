package model

import (
	"fmt"
	"strconv"
)

type StaticFilename struct {
	shapeWidth
	value string
}

func ParseStaticFilename(value, shape, width string) (*StaticFilename, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf("static filename: filename is empty")
	}

	w, _ := strconv.Atoi(width)
	w = normalizeWidth(w)

	s := normalizeShape(shape)
	return &StaticFilename{
		value: value,
		shapeWidth: shapeWidth{
			shape: s,
			width: w,
		},
	}, nil
}

func NewStaticFilename(value string) StaticFilename {
	return StaticFilename{value: value}
}

func (name StaticFilename) Name() string {
	return name.value
}

func (name StaticFilename) Path() string {
	return name.value
}

func (name StaticFilename) OldPath() string {
	return name.Path()
}

func (name StaticFilename) CachedPath() string {
	return fmt.Sprintf("%s_%s_%d", name.value, name.shape, name.width)
}
