package aptclient

import (
	"context"
)

var (
	addressPancake = "0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa"
)

func (ts *ClientTestSuite) TestGetAccountResources() {
	ctx := context.Background()

	resources, metadata, err := ts.client.GetAccountResources(
		ctx,
		addressPancake,
		GetAccountResourcesQueryParams{Limit: 3},
	)

	ts.Assert().NoError(err)
	ts.Assert().NotNil(metadata)
	ts.Assert().NotNil(resources)
	ts.Assert().Equal(3, len(resources))
}
