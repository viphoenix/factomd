package state

import (
	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type HardGrant struct {
	dbh     uint32
	amount  uint64
	address interfaces.IAddress
}

// Return the Hard Coded Grants. Buried in an func so other code cannot easily address the array and change it
func getHardCodedGrants() []HardGrant {
	var hardcodegrants = [...]HardGrant{
		// waiting for "real-ish" data from brian
		HardGrant{10, 1, factoid.NewAddress(primitives.ConvertUserStrToAddress("FA3oajkmHMfqkNMMShmqpwDThzMCuVrSsBwiXM2kYFVRz3MzxNAJ"))}, // Pay Clay 1 at dbheight 10
	}
	// Passing an array to a function creates a copy and I return a slide anchored to that copy
	// so the caller does not have the address of the array itself.
	// Closest thing to a constant array I can get in Go
	return func(x [len(hardcodegrants)]HardGrant) []HardGrant { return x[:] }(hardcodegrants)
}

//return a (possibly empty) of coinbase payouts to be scheduled at this height
func (s *State) GetGrantPayoutsFor(currentDBHeight uint32) []interfaces.ITransAddress {

	outputs := make([]interfaces.ITransAddress, 0)
	// this is only but temporary, once the hard coded grants are payed this code will go away
	// I can't modify the grant list because in simulation it is shared across nodes so for now I just
	// scan the whole list once every 25 blocks
	// I opted for one list knowing it will have to be different for testnet vs mainnet because making it
	// network sensitive just add complexity to the code.
	// there is no need for activation height because the grants have inherent activation heights per grant
	for _, g := range getHardCodedGrants() { // check every hardcoded grant
		if g.dbh == currentDBHeight { // if it's ready {...
			o := factoid.NewOutAddress(g.address, g.amount) // Create a payout
			outputs = append(outputs, o)                    // and add it to the list
		}
	}
	return outputs
}