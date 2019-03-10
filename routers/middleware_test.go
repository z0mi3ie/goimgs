package routers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	testImageID1 = "hellomeid1"
	testImageID2 = "hellomeid2"
)

func TestExtractDeleteImageQueryParams(t *testing.T) {
	testCases := []struct {
		tc             string
		params         string
		expectedParams DeleteImageQueryParams
	}{
		{
			tc:     "successfully parse two id query parameters",
			params: fmt.Sprintf("id=%s&id=%s", testImageID1, testImageID2),
			expectedParams: DeleteImageQueryParams{
				ID: []string{testImageID1, testImageID2},
			},
		},
	}

	for _, test := range testCases {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		path := fmt.Sprintf("localhost:8080?%s", test.params)
		ctx.Request, _ = http.NewRequest("GET", path, nil)
		DeleteImageQueryParamsMiddleware(ctx)
		qp, _ := ctx.Get(QueryParamsKey)
		assert.Equal(t, test.expectedParams, qp)
	}
}
