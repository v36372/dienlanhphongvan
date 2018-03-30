package view

import (
	"fmt"
)

func NewImage(imageName string) string {
	return fmt.Sprintf("/images/cached/%s?type=s&w=480", imageName)
}
