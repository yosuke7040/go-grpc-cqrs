package adapter

import (
	"github.com/yosuke7040/commandservice/domain/models/products"

	"github.com/fullness-MFurukawa/samplepb/pb"
)

type ProductAdapter interface {
	ToEntity(param *pb.ProductUpParam) (*products.Product, error)
	ToResult(result any) *pb.ProductUpResult
}
