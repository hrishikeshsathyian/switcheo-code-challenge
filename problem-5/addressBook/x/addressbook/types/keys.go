package types

const (
	// ModuleName defines the module name
	ModuleName = "addressbook"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_addressbook"

	// ContactKey is used to uniquely identify contacts within the system.
	// It will be used as the beginning of the key for each contact, followed by their unique ID
	ContactKey = "Contact/value/"

	// This key will be used to keep track of the ID of the latest contact added to the store.
	ContactCountKey = "Contact/count/"
)

var (
	ParamsKey = []byte("p_addressbook")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
