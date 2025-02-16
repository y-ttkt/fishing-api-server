package profile

import (
	"fmt"
	"github.com/google/uuid"
)

type Image string

func NewImage(value string) Image {
	return Image(value)
}

func MakeHashName(fileExtension string) Image {
	hashName := fmt.Sprintf("%s.%s", uuid.New().String(), fileExtension)
	return Image(hashName)
}

func (i Image) Value() string {
	return string(i)
}
