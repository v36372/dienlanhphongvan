package view

import (
	"fmt"
)

func NewImage(imageName string) string {
	return fmt.Sprintf("/images/original/%s", imageName)
}
