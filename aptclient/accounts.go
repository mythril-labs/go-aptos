package aptclient

import (
	"context"
	"strconv"

	"github.com/mythril-labs/go-aptos/types"
)

var (
	pathGetAccountResources = "/accounts/{address}/resources"
	pathGetAccountResource  = "/accounts/{address}/resource/{resource_type}"
)

func (c *Client) GetAccountResources(
	ctx context.Context,
	address string,
	queryParams GetAccountResourcesQueryParams,
) ([]types.Resource, *Metadata, error) {
	req := c.client.R().SetContext(ctx).SetPathParam("address", address)

	if queryParams.Limit > 0 {
		req.SetQueryParam("limit", strconv.Itoa(queryParams.Limit))
	}

	if len(queryParams.Start) > 0 {
		req.SetQueryParam("start", queryParams.Start)
	}

	var resources []types.Resource
	req.SetResult(&resources)

	resp, err := req.Get(pathGetAccountResources)
	if err != nil {
		return nil, nil, err
	}

	if resp.IsError() {
		return nil, nil, handleErrResp(resp.Body())
	}

	return resources, handleRspHdr(resp.Header()), nil
}
