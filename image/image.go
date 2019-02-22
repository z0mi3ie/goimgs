package image

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

var (
	validFileTypes = []string{
		".jpg",
		".jpeg",
		".png",
	}
)

// Data holds all associated data the DB needs to store for an image
// before saving that image to the file system
type Data struct {
	file     *multipart.FileHeader
	id       string
	fileType string
	ogName   string
}

// NewImageData return a Data struct containing necessary information
// to further process the file and request
func NewImageData(f *multipart.FileHeader) (*Data, error) {
	id, err := generateUUIDString()
	if err != nil {
		return nil, err
	}

	ft := filepath.Ext(f.Filename)
	validateFileType(ft)

	d := &Data{
		file:     f,
		id:       id,
		fileType: ft,
		ogName:   f.Filename,
	}
	return d, nil
}

// Filename is the name of the file when being served, returns a string
func (d *Data) Filename() string {
	return fmt.Sprintf("%s.%s", d.id, d.fileType)
}

// URL is the path to the file where it is being served, returns a string
func (d *Data) URL(path string) string {
	return fmt.Sprintf("%s/%s", path, d.Filename())
}

// generateUUIDString is a helper that generates a UUID V4 and returns its
// string format, or an error if one is encountered
func generateUUIDString() (string, error) {
	gid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return gid.String(), nil
}

func validateFileType(ft string) error {
	// If the given filetype is found in the valid file types, return
	// without error. If it is not found we return an error.
	for _, vf := range validFileTypes {
		if vf == ft {
			return nil
		}
	}
	return fmt.Errorf("%s is not a valid file type", ft)
}
