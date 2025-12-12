package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/qulDev/golang-restful-api/helper"
	"github.com/qulDev/golang-restful-api/model/domain"
	"github.com/qulDev/golang-restful-api/model/web"
	"github.com/qulDev/golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)

	defer helper.CommitOrRollback(tx)

	category, errFind := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(errFind)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)

	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)

	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)

}
