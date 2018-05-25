package skynode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SkyNode is main type for skynode access
type SkyNode struct {
	baseURL string
}

// NewSkyNode creates a new SkyNode
func NewSkyNode(baseURL string) SkyNode {
	return SkyNode{baseURL}
}

// GetSeed performs GET request to get a seed for wallet
func (s SkyNode) GetSeed() (string, error) {
	url := fmt.Sprintf("%s/wallet/newSeed", s.baseURL)
	return get(url, "seed")
}

// GetCsrfToken performs GET request to an auth token for skynode
// Csrf token lives only for 30sec
func (s SkyNode) GetCsrfToken() (string, error) {
	url := fmt.Sprintf("%s/csrf", s.baseURL)
	return get(url, "csrf_token")
}

// CreateWallet performs POST request to create a wallet in the skynode
func (s SkyNode) CreateWallet(name string, seed string, csrf string) (*Wallet, error) {
	form := url.Values{
		"label": {name},
		"seed":  {seed},
	}

	req, err := http.NewRequest("POST", (fmt.Sprintf("%s/wallet/create", s.baseURL)), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-CSRF-Token", csrf)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	w := &Wallet{}
	err = json.NewDecoder(resp.Body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

// TransferMoney transfers money from src wallet to dst wallet
func (s SkyNode) TransferMoney(walletID string, dst string, coins string, csrf string) error {
	form := url.Values{
		"id":    {walletID},
		"dst":   {dst},
		"coins": {coins},
	}

	req, err := http.NewRequest("POST", (fmt.Sprintf("%s/wallet/spend", s.baseURL)), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-CSRF-Token", csrf)

	_, err = http.DefaultClient.Do(req)
	return err
}

func get(url string, keyExtractor string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var dat map[string]interface{}
	if err = json.Unmarshal(body, &dat); err != nil {
		return "", err
	}

	return dat[keyExtractor].(string), nil
}
