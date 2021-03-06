package jsonschema

import (
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/stashapp/stash/pkg/models"
)

type ImageFile struct {
	ModTime models.JSONTime `json:"mod_time,omitempty"`
	Size    int             `json:"size"`
	Width   int             `json:"width"`
	Height  int             `json:"height"`
}

type Image struct {
	Title      string          `json:"title,omitempty"`
	Checksum   string          `json:"checksum,omitempty"`
	Studio     string          `json:"studio,omitempty"`
	Rating     int             `json:"rating,omitempty"`
	OCounter   int             `json:"o_counter,omitempty"`
	Galleries  []string        `json:"galleries,omitempty"`
	Performers []string        `json:"performers,omitempty"`
	Tags       []string        `json:"tags,omitempty"`
	File       *ImageFile      `json:"file,omitempty"`
	CreatedAt  models.JSONTime `json:"created_at,omitempty"`
	UpdatedAt  models.JSONTime `json:"updated_at,omitempty"`
}

func LoadImageFile(filePath string) (*Image, error) {
	var image Image
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&image)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func SaveImageFile(filePath string, image *Image) error {
	if image == nil {
		return fmt.Errorf("image must not be nil")
	}
	return marshalToFile(filePath, image)
}
