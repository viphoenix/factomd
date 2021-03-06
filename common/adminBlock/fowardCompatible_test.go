package adminBlock_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/FactomProject/factomd/common/adminBlock"
	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/testHelper"
)

func TestNewForwardCompatibleEntry(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// Coinbase Descriptor
	for i := 0; i < 100; i++ {
		var outputs []interfaces.ITransAddress
		for c := 0; c < rand.Intn(64); c++ {
			outputs = append(outputs, factoid.RandomTransAddress())
		}
		a := NewCoinbaseDescriptor(outputs)

		// Coinbase Descriptor is forward compatible
		b := NewForwardCompatibleEntry(0)
		testHelper.TestMarshaling(a, b, rand.Intn(100), t)

	}

	// Coinbase Address
	for i := 0; i < 100; i++ {
		a := NewAddFactoidAddress(primitives.RandomHash(), factoid.RandomAddress())

		b := NewForwardCompatibleEntry(0)
		testHelper.TestMarshaling(a, b, rand.Intn(100), t)
	}

	// Efficiency
	for i := 0; i < 100; i++ {
		a := NewAddEfficiency(primitives.RandomHash(), uint16(rand.Intn(10000)))

		b := NewForwardCompatibleEntry(0)
		testHelper.TestMarshaling(a, b, rand.Intn(100), t)
	}
}
