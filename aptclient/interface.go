package aptclient

import (
	"context"

	"github.com/mythril-labs/go-aptos/types"
)

type IClient interface {
	GetAccountResources(
		ctx context.Context,
		address string,
		queryParams GetAccountResourcesQueryParams,
	) ([]types.Resource, *Metadata, error)

	GetAccountResource(
		ctx context.Context,
		address string,
		resourceType string,
		queryParams GetAccountResourceQueryParams,
	) (*types.Resource, *Metadata, error)
}
