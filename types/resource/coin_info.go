package resource

import (
	"encoding/json"
	"regexp"

	"github.com/mythril-labs/go-aptos/types"
)

var (
	TypeCoinInfoPattern = regexp.MustCompile(`^0x1::coin::CoinInfo<(.+<.+,.+>|[^,]+)>$`)
)

type CoinInfo struct {
	Type string
	ID   string

	Data struct {
		Decimals uint8  `json:"decimals"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
	}
}

func AsCoinInfo(resource types.Resource) (CoinInfo, bool) {
	matches := TypeCoinInfoPattern.FindStringSubmatch(resource.Type)

	if len(matches) != 2 {
		return CoinInfo{}, false
	}

	coinInfo := CoinInfo{Type: resource.Type, ID: matches[1]}

	if err := json.Unmarshal(resource.Data, &coinInfo.Data); err != nil {
		return CoinInfo{}, false
	}

	return coinInfo, true
}
