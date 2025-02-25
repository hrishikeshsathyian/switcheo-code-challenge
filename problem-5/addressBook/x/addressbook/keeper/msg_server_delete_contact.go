package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"addressbook/x/addressbook/types"
)

func (k msgServer) DeleteContact(goCtx context.Context, msg *types.MsgDeleteContact) (*types.MsgDeleteContactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetContact(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.RemoveContact(ctx, msg.Id)
	return &types.MsgDeleteContactResponse{}, nil
}
