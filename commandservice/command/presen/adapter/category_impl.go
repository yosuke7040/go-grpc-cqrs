package adapter

import (
	"github.com/fullness-MFurukawa/samplepb/pb"
	"github.com/yosuke7040/commandservice/domain/models/categories"
	"github.com/yosuke7040/commandservice/errs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryAdapaterImpl struct{}

func NewcategoryAdapaterImpl() CategoryAdapter {
	return &categoryAdapaterImpl{}
}

func (ins *categoryAdapaterImpl) ToEntity(param *pb.CategoryUpParam) (*categories.Category, error) {
	switch param.GetCrud() {
	case pb.CRUD_INSERT: // データの追加
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		category, err := categories.NewCategory(name)
		if err != nil {
			return nil, err
		}
		return category, nil
	case pb.CRUD_UPDATE: // データの変更
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, name), nil
	case pb.CRUD_DELETE: // データの削除
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		category := categories.BuildCategory(id, nil)
		return category, nil
	default:
		return nil, errs.NewDomainError("不明な操作を受信しました。")
	}
}

func (ins *categoryAdapaterImpl) ToResult(result any) *pb.CategoryUpResult {
	var up_category *pb.Category
	var up_err *pb.Error
	switch v := result.(type) {
	case *categories.Category:
		if v.Name() == nil {
			up_category = &pb.Category{Id: v.Id().Value(), Name: ""}
		} else {
			up_category = &pb.Category{Id: v.Id().Value(), Name: v.Name().Value()}
		}
	case *errs.DomainError: // 実行結果がerrs.DomainErrorの場合
		up_err = &pb.Error{Type: "Domain Error", Message: v.Error()}
	case *errs.CRUDError: // 実行結果がerrs.CRUDErrorの場合
		up_err = &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError: // 実行結果がerrs.Internalerrorの場合
		up_err = &pb.Error{Type: "Internal Error", Message: "只今、サービスを提供できません。"}
	}
	// pb.CategoryUpResultのインスタンスを生成して返す
	return &pb.CategoryUpResult{
		Category:  up_category,
		Error:     up_err,
		Timestamp: timestamppb.Now(),
	}
}
