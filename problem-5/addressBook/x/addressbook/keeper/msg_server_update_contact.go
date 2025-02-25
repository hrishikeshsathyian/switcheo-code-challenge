package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"addressbook/x/addressbook/types"
)

func (k msgServer) UpdateContact(goCtx context.Context, msg *types.MsgUpdateContact) (*types.MsgUpdateContactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var contact = types.Contact{
		Creator:   msg.Creator,
		Id:        msg.Id,
		FirstName: msg.FirstName,
		LastName:  msg.LastName,
		Phone:     msg.Phone,
		Email:     msg.Email,
		Address:   msg.Address,
	}
	val, found := k.GetContact(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetContact(ctx, contact)
	return &types.MsgUpdateContactResponse{}, nil
}
