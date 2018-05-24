package skynode

import "github.com/AlexSugak/skycoin-promo/skynode"

// SkyNode is a type of a service that provides access to skynode
type SkyNode interface {
	GetSeed() (string, error)
	GetCsrfToken() (string, error)
	CreateWallet(string, string, string) (*skynode.Wallet, error)
}
