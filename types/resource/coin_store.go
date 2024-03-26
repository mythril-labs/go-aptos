package resource

import (
	"encoding/json"
	"regexp"

	"github.com/mythril-labs/go-aptos/types"
)

var (
	TypeCoinStorePattern = regexp.MustCompile(`^0x1::coin::CoinStore<(.+<.+,.+>|[^,]+)>$`)
)

type CoinStore struct {
	Type string
	ID   string
	Data struct {
		Coin struct {
			Value string `json:"value"`
		} `json:"coin"`
		Frozen bool `json:"frozen"`
	}
}

func AsCoinStore(resource types.Resource) (CoinStore, bool) {
	matches := TypeCoinInfoPattern.FindStringSubmatch(resource.Type)

	if len(matches) != 2 {
		return CoinStore{}, false
	}

	coinStore := CoinStore{Type: resource.Type, ID: matches[1]}

	if err := json.Unmarshal(resource.Data, &coinStore.Data); err != nil {
		return CoinStore{}, false
	}

	return coinStore, true
}
