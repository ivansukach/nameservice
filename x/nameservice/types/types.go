package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Whois struct{
	Value	string	`json:"value"`
	Owner	sdk.AccAddress	`json:"owner"`
	Price	sdk.Coins	`json:"price"`
}

// NewWhois returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}
