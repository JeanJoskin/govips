// +build static_build
package vips

// #cgo pkg-config: --static vips
// #cgo LDFLAGS: -lexpat
import "C"
