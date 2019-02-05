package vips

// #include "bridge.h"
import "C"

import "errors"

var (
	// ErrUnsupportedImageFormat when image type is unsupported
	ErrUnsupportedImageFormat = errors.New("Unsupported image format")
)
