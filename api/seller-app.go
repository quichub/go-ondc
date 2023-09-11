package api

import (
	"ondc/model/operations"
)

type SellerApp interface {
	Search(request *operations.SearchRequest) (*operations.SearchResponse, error)
}
