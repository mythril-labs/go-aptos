package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mythril-labs/go-aptos/types"
)

func TestAsCoinInfo(t *testing.T) {
	t.Run("it should return correct coinInfo", func(t *testing.T) {
		resource := types.Resource{
			Type: "0x1::coin::CoinInfo<0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa::swap::LPToken<0x3c1d4a86594d681ff7e5d5a233965daeabdc6a15fe5672ceeda5260038857183::vcoins::V<0x1::aptos_coin::AptosCoin>, 0x8b2df69c9766e18486c37e3cfc53c6ce6e9aa58bbc606a8a0a219f24cf9eafc1::sui_launch_token::SuiLaunchToken>>",
			Data: []byte(`{"decimals":8,"name":"Pancake-vAPT-SLT-LP","supply":{"vec":[{"aggregator":{"vec":[]},"integer":{"vec":[{"limit":"340282366920938463463374607431768211455","value":"18946986"}]}}]},"symbol":"Cake-LP"}`),
		}

		coinInfo, ok := AsCoinInfo(resource)

		assert.True(t, ok)
		assert.Equal(t, "0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa::swap::LPToken<0x3c1d4a86594d681ff7e5d5a233965daeabdc6a15fe5672ceeda5260038857183::vcoins::V<0x1::aptos_coin::AptosCoin>, 0x8b2df69c9766e18486c37e3cfc53c6ce6e9aa58bbc606a8a0a219f24cf9eafc1::sui_launch_token::SuiLaunchToken>", coinInfo.ID)
		assert.Equal(t, 8, coinInfo.Data.Decimals)
		assert.Equal(t, "Pancake-vAPT-SLT-LP", coinInfo.Data.Name)
	})
}
