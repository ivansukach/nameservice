package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"log"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "nameservice/DeleteName", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc = codec.New()

func init() {
	log.Println("PANIC")
	RegisterCodec(ModuleCdc)
}
