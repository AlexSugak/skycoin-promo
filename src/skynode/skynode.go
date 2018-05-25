package skynode

import (
	"github.com/shopspring/decimal"
)

// NodeAPI is a type of a service that provides access to skynode
type NodeAPI interface {
	GetSeed() (string, error)
	GetCsrfToken() (string, error)
	CreateWallet(string, string, string) (*Wallet, error)
	TransferMoney(string, string, decimal.Decimal, string) error
}
