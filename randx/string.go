package randx

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

type StringGenerator struct {
	Base       string
	baseLength int
	indexBits  int
	indexMask  int64
	indexMax   int
	source     rand.Source
}

const (
	Base58 = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"
)

func NewStringGenerator(base string) *StringGenerator {
	y := 1.0
	for math.Pow(2, y) < float64(len(base)) {
		y = y + 1
	}
	bits := int(y)

	return &StringGenerator{
		Base:       base,
		baseLength: len(base),
		indexBits:  bits,
		indexMask:  (1 << bits) - 1,
		indexMax:   63 / bits,
		source:     rand.NewSource(time.Now().UnixNano()),
	}
}

func (sg *StringGenerator) Generate(length int) string {
	builder := strings.Builder{}
	builder.Grow(length)

	for i, cache, remain := length-1, sg.source.Int63(), sg.indexMax; i >= 0; {
		if remain == 0 {
			cache = sg.source.Int63()
			remain = sg.indexMax
		}

		index := int(cache & sg.indexMask)
		if index < sg.baseLength {
			builder.WriteByte(sg.Base[index])
			i--
		}

		cache >>= sg.indexBits
		remain--
	}

	return builder.String()
}
