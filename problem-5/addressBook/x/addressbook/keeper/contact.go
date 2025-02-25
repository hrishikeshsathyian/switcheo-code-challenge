package keeper

import (
	"encoding/binary" // Helps convert between byte arrays and numeric types

	"addressbook/x/addressbook/types"

	"cosmossdk.io/store/prefix" // Creates a key-value store with prefixes for namespaces
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
*
1. Get the current contact count
2. Assign an ID to the new contact
3. Serialize and store the contact using the contact ID as the key
4. Update the contact count and return the new contact ID
*
*/
func (k Keeper) AppendContact(ctx sdk.Context, contact types.Contact) uint64 {
	// get current total contact count
	count := k.GetContactCount(ctx)
	// set the ID to the current count
	contact.Id = count
	// open the key-value store in the context of the blockchain
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	// serialise the contact into bytes for storage
	appendedValue := k.cdc.MustMarshal(&contact)
	// store the post in the key-value store, using the contactID as the key
	store.Set(GetContactIDBytes(contact.Id), appendedValue)
	// increment byte value of the contact count
	k.SetContactCount(ctx, count+1)
	// return ID of the new contact
	return count
}

// Retrieve total number of contacts in the store
// Access the key-value store
// Retrieve the contact count from a predefined key. If the key does not exist, return 0
// Convert the byte array to an unsigned integer and return the value
func (k Keeper) GetContactCount(ctx sdk.Context) uint64 {
	// open the key-value store in the context of the blockchain
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	// define the key for the contact count
	byteKey := types.KeyPrefix(types.ContactCountKey)
	// retrieve the byte calue stored under the post count key
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// convert the contact ID to a byte array to use as a key
func GetContactIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetContactCount(ctx sdk.Context, count uint64) {
	// open the key-value store in the context of the blockchain
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	// define the key for the contact count
	byteKey := types.KeyPrefix(types.ContactCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// Retrieve a contact from the store
func (k Keeper) GetContact(ctx sdk.Context, id uint64) (val types.Contact, found bool) {
	// open the key-value store in the context of the blockchain
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	b := store.Get(GetContactIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetContact(ctx sdk.Context, contact types.Contact) {
	// open the key-value store in the context of the blockchain
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// create a new store with the contact prefix
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	b := k.cdc.MustMarshal(&contact)
	store.Set(GetContactIDBytes(contact.Id), b)
}

func (k Keeper) RemoveContact(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))
	store.Delete(GetContactIDBytes(id))
}
