package skynode

// NodeAPI is a type of a service that provides access to skynode
type NodeAPI interface {
	GetSeed() (string, error)
	GetCsrfToken() (string, error)
	CreateWallet(string, string, string) (*Wallet, error)
}
