package types

// MinterKey is the key to use for the keeper store.

var MinterStateKey = []byte{0x02}

const (
	// ModuleName defines the module name
	ModuleName = "cfeminter"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfeminter"

	InflationCollectorName = "inflation_collector"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
