package adapter

import (
	"github.com/fullness-MFurukawa/samplepb/pb"
	"github.com/yosuke7040/commandservice/domain/models/categories"
)

type CategoryAdapter interface {
	// CategoryUpParamからCategoryに変換する
	ToEntity(param *pb.CategoryUpParam) (*categories.Category, error)
	// 実行結果からCategoryUpResultに変換する
	ToResult(result any) *pb.CategoryUpResult
}
