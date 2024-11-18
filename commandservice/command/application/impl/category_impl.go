package impl

import (
	"context"
	"github.com/yosuke7040/commandservice/application/service"
	"github.com/yosuke7040/commandservice/domain/models/categories"
)

type categoryServiceImpl struct {
	rep categories.CategoryRepository
	transaction
}

func NewcategoryServiceImpl(rep categories.CategoryRepository) service.CategoryService {
	return &categoryServiceImpl{rep: rep}
}

func (ins *categoryServiceImpl) Add(ctx context.Context, category *categories.Category) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		defer ins.complete(tran, err)
	}()

	if err = ins.rep.Exists(ctx, category); err != nil {
		return err
	}
	if err = ins.rep.Create(ctx, category); err != nil {
		return err
	}
	return nil
}

func (ins *categoryServiceImpl) Update(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.complete(tran, err)
	}()
	// カテゴリを更新する
	if err = ins.rep.UpdateById(ctx, tran, category); err != nil {
		return err
	}
	return err
}

// カテゴリを削除する
func (ins *categoryServiceImpl) Delete(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.complete(tran, err)
	}()
	// カテゴリを削除する
	if err = ins.rep.DeleteById(ctx, tran, category); err != nil {
		return err
	}
	return err
}
