package dast8x8

import (
	"github.com/reiver/go-rgba32"
	"github.com/reiver/go-rgba32sprite8x8"
)

var sprite00 [rgba32sprite8x8.ByteSize/rgba32.ByteSize][rgba32.ByteSize]uint8 = [rgba32sprite8x8.ByteSize/rgba32.ByteSize][rgba32.ByteSize]uint8{
//	0	1	2	3	4	5	6	7
	nada,	nada,	red,	red,	red,	nada,	nada,	nada,
	nada,	red,	red,	red,	red,	red,	nada,	nada,
	nada,	red,	black,	white,	black,	black,	nada,	nada,
	nada,	red,	red,	skin,	skin,	skin,	nada,	nada,
	red,	red,	black,	black,	black,	red,	red,	nada,
	red,	nada,	red,	red,	red,	nada,	yellow,	nada,
	nada,	black,	black,	nada,	black,	black,	nada,	nada,
	red,	red,	red,	nada,	red,	red,	red,	nada,
}
