package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"addressbook/x/addressbook/types"
)

func (k Keeper) ListContact(ctx context.Context, req *types.QueryListContactRequest) (*types.QueryListContactResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ContactKey))

	var contacts []types.Contact
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var contact types.Contact
		if err := k.cdc.Unmarshal(value, &contact); err != nil {
			return err
		}

		contacts = append(contacts, contact)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListContactResponse{Contact: contacts, Pagination: pageRes}, nil
}
