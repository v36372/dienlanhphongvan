package view

import (
	"fmt"
)

func NewImage(imageName string) string {
	if len(imageName) == 0 {
		return ""
	}
	return fmt.Sprintf("/images/cached/%s?type=s&w=480", imageName)
}
