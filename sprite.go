package dast8x8

import (
	"github.com/reiver/go-rgba32"
	"github.com/reiver/go-rgba32sprite8x8"

	"image"
	"reflect"
	"unsafe"
)

func Sprite(id uint8) image.Image {
	var sprite [][rgba32.ByteSize]uint8

	switch id {
	case 0:
		sprite = sprite00[:]
	case 1:
		sprite = sprite01[:]
	default:
		return nil
	}


	var header reflect.SliceHeader
	{
		h := (*reflect.SliceHeader)(unsafe.Pointer(&sprite))

		header.Data = h.Data
		header.Len  = h.Len*4
		header.Cap  = h.Cap*4
	}

	var p []uint8 = *(*[]uint8)(unsafe.Pointer(&header))

	return rgba32sprite8x8.Slice(p)
}
