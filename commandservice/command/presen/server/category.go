package server

import (
	"context"
	"github.com/fullness-MFurukawa/samplepb/pb"
	"github.com/yosuke7040/commandservice/application/service"
	"github.com/yosuke7040/commandservice/presen/adapter"
)

// カテゴリ更新サーバの実装
type categoryServer struct {
	adapter adapter.CategoryAdapter // カテゴリ変換
	service service.CategoryService // カテゴリ更新サービス
	// 生成されたUnimplementedCategoryCommandServerをエンベデッド
	pb.UnimplementedCategoryCommandServer
}

// コンストラクタ
func NewcategoryServer(adapter adapter.CategoryAdapter, service service.CategoryService) pb.CategoryCommandServer {
	return &categoryServer{adapter: adapter, service: service}
}

func (ins *categoryServer) Create(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Add(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}

func (ins *categoryServer) Update(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Update(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(category), nil
	}
}

func (ins *categoryServer) Delete(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		if err := ins.service.Delete(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}

		return ins.adapter.ToResult(category), nil
	}
}
