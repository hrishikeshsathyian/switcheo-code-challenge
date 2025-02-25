package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateContact{}

func NewMsgUpdateContact(creator string, firstName string, lastName string, phone string, email string, address string, id uint64) *MsgUpdateContact {
	return &MsgUpdateContact{
		Creator:   creator,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		Address:   address,
		Id:        id,
	}
}

func (msg *MsgUpdateContact) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
