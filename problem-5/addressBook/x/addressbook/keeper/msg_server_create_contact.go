package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"addressbook/x/addressbook/types"
)

func (k msgServer) CreateContact(goCtx context.Context, msg *types.MsgCreateContact) (*types.MsgCreateContactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var contact = types.Contact{
		Creator:   msg.Creator,
		FirstName: msg.FirstName,
		LastName:  msg.LastName,
		Phone:     msg.Phone,
		Email:     msg.Email,
		Address:   msg.Address,
		Nickname:  msg.Nickname,
	}
	id := k.AppendContact(
		ctx,
		contact,
	)
	return &types.MsgCreateContactResponse{
		Id: id,
	}, nil
}
