package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/yosuke7040/commandservice/domain/models/categories"
	"github.com/yosuke7040/commandservice/errs"
	"github.com/yosuke7040/commandservice/infra/sqlboiler/handler"
	"github.com/yosuke7040/commandservice/infra/sqlboiler/models"
	"log"
)

type categoryRepositorySQLBoiler struct{}

func NewcategoryRepositorySQLBoiler() categories.CategoryRepository {
	models.AddCategoryHook(boil.AfterInsertHook, CategoryAfterInsertHook)
	models.AddCategoryHook(boil.AfterUpdateHook, CategoryAfterUpdateHook)
	models.AddCategoryHook(boil.AfterDeleteHook, CategoryAfterDeleteHook)
	return &categoryRepositorySQLBoiler{}
}

func (rep *categoryRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	condition := models.CategoryWhere.Name.EQ(category.Name().Value())
	if exists, err := models.Categories(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists {
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", category.Name().Value()))
	}
}

func (rep *categoryRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	new_category := models.Category{
		ID:    0,
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}
	if err := new_category.Insert(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *categoryRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	up_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("カテゴリ番号:%sは存在しないため、更新できませんでした。", category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	up_model.ObjID = category.Id().Value()
	up_model.Name = category.Name().Value()
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *categoryRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	del_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("カテゴリ番号:%sは存在しないため、削除できませんでした。",
			category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func CategoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを登録しました。\n", category.ObjID, category.Name)
	return nil
}

func CategoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを変更しました。\n", category.ObjID, category.Name)
	return nil
}

func CategoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを削除しました。\n", category.ObjID, category.Name)
	return nil
}
