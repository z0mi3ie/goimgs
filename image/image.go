package image

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
	"github.com/z0mi3ie/goimgs/utils"
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

// MetaData holds the metadata associated with an image
type MetaData struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	OgName      string `json:"ogName"`
	Description string `json:"description"`
	Title       string `json:"title"`
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

// NewImageMetaData returns a struct for an image metadata
func NewImageMetaData(id, url, ogName, description, title string) MetaData {
	return MetaData{
		ID:          id,
		URL:         url,
		OgName:      ogName,
		Description: description,
		Title:       title,
	}
}

// ID returns the image's globally unique ID
func (d *Data) ID() string {
	return d.id
}

// Filename is the name of the file when being served, returns a string
func (d *Data) Filename() string {
	return fmt.Sprintf("%s%s", d.id, d.fileType)
}

// URL is the path to the file where it is being served, returns a string
func (d *Data) URL(path string) string {
	return fmt.Sprintf("%s/%s", utils.NormalizePath(path), d.Filename())
}

// File returns the multipart.Fileheader for direct processing
func (d *Data) File() *multipart.FileHeader {
	return d.file
}

// OGName returns the original name the file was uploaded as
func (d *Data) OGName() string {
	return d.ogName
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
