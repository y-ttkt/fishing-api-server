package util

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"io"
	"time"
)

type Identifier struct {
	Identifier string
}

type IdentifierGenerator interface {
	Generate() Identifier
}

var defaultGenerator IdentifierGenerator

func init() {
	defaultGenerator = newULIDGenerator(rand.Reader)
}

func GenerateIdentifier() Identifier {
	return defaultGenerator.Generate()
}

type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

func newULIDGenerator(reader io.Reader) *ULIDGenerator {
	return &ULIDGenerator{
		entropy: ulid.Monotonic(reader, 0),
	}
}

func (g *ULIDGenerator) Generate() Identifier {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), g.entropy)

	return Identifier{
		Identifier: id.String(),
	}
}
