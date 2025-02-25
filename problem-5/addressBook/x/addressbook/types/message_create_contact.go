package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateContact{}

func NewMsgCreateContact(creator string, firstName string, lastName string, phone string, email string, address string) *MsgCreateContact {
	return &MsgCreateContact{
		Creator:   creator,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Address:   address,
	}
}

func (msg *MsgCreateContact) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
