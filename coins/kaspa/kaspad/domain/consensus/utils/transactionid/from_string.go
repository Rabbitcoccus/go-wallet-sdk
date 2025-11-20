package transactionid

import (
	"github.com/Rabbitcoccus/go-wallet-sdk/coins/kaspa/kaspad/domain/consensus/model/externalapi"
)

// FromString creates a new DomainTransactionID from the given string
func FromString(str string) (*externalapi.DomainTransactionID, error) {
	hash, err := externalapi.NewDomainHashFromString(str)
	return (*externalapi.DomainTransactionID)(hash), err
}
