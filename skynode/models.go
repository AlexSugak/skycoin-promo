package skynode

// WalletMeta is a meta information about wallet
// It's a part of Wallet entity
type WalletMeta struct {
	Coin       string `json:"coin"`
	FileName   string `json:"filename"`
	Label      string `json:"label"`
	Type       string `json:"type"`
	Version    string `json:"version"`
	CryptoType string `json:"crypto_type"`
	Timestamp  string `json:"timestamp"`
	Encrypted  bool   `json:"encrypted"`
}

// WalletEntry is an entity about an address and public key of a wallet
// It's a part of Wallet entity
type WalletEntry struct {
	Address   string `json:"address"`
	PublicKey string `json:"public_key"`
}

// Wallet is an entity that represents a skycoin wallet
type Wallet struct {
	Meta    WalletMeta    `json:"meta"`
	Entries []WalletEntry `json:"entries"`
}
