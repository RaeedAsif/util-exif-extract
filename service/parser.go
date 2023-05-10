package service

import (
	"os"

	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type GpsExif struct {
	exif *exif.GpsInfo
}

// NewGpsExif returns a new GpsExif instance
func NewGpsExif(path string) (*GpsExif, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	rawExif, err := exif.SearchAndExtractExifWithReader(f)
	if err != nil {
		return nil, err
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}

	ti := exif.NewTagIndex()

	err = exif.LoadStandardTags(ti)
	if err != nil {
		return nil, err
	}

	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return nil, err
	}

	ifd, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return nil, err
	}

	exif, err := ifd.GpsInfo()
	if err != nil {
		return nil, err
	}

	return &GpsExif{exif: exif}, nil
}

// GetLongitude returns the longitude of the image
func (gps *GpsExif) GetLongitude() float64 {
	return gps.exif.Longitude.Decimal()
}

// GetLatitude returns the latitude of the image
func (gps *GpsExif) GetLatitude() float64 {
	return gps.exif.Latitude.Decimal()
}
