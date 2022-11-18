/*
Dofusdude

Testing AllItemsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dodugo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dofusdude/dodugo"
)

func Test_dodugo_AllItemsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AllItemsApiService GetItemsAllSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.AllItemsApi.GetItemsAllSearch(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
