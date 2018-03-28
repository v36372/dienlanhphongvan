package model

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"time"
)

var (
	validImageWidths = []int{60, 120, 240, 320, 480, 640, 960, 1200, 1600, 1800}
)

var (
	// name- unixTimestamp
	regexNewImageFilename = regexp.MustCompile(`^(.+)-(\d+)$`)
	// name- unixTimestamp-shape-width
	regexOldImageFilename = regexp.MustCompile(`^.+-(\d+)-(o|s)-(\d+)$`)
	//  unixTimestamp_name
	regexUploadFilename = regexp.MustCompile(`^(\d+)_(.+)$`)
)

type Filename interface {
	Name() string
	Path() string
	OldPath() string
	CachedPath() string
	Width() int
	Shape() string
}

type shapeWidth struct {
	shape string
	width int
}

func (v shapeWidth) Shape() string {
	return v.shape
}

func (v shapeWidth) Width() int {
	return v.width
}

func pathByTimestamp(name string, timestamp time.Time) string {
	year, month, day := strconv.Itoa(timestamp.Year()),
		strconv.Itoa(int(timestamp.Month())),
		strconv.Itoa(timestamp.Day())
	return path.Join(year, month, day, name)
}

func oldPathByTimestamp(name string, timestamp time.Time) string {
	year, month :=
		strconv.Itoa(timestamp.Year()),
		strconv.Itoa(int(timestamp.Month()))

	return path.Join(year, month, name)
}

func parseTimestamp(value string) (ret time.Time, err error) {
	var sec int64
	sec, err = strconv.ParseInt(value, 10, 64)
	if err != nil {
		return
	}
	if sec < 1356973200 {
		err = fmt.Errorf("require since 2013")
		return
	}
	ret = time.Unix(sec, 0)
	return
}

func normalizeShape(shape string) string {
	switch shape {
	case "o", "original":
		return "o"
	case "s", "square":
		return "s"
	default:
		return "s"
	}
}

func normalizeWidth(value int) int {
	if value == 0 {
		value = 240
	}
	for _, width := range validImageWidths {
		if value <= width {
			return width
		}
	}
	return validImageWidths[len(validImageWidths)-1]
}
