package repository

import (
	"context"
	"database/sql"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
	"project/internal/domain/types"
	"project/internal/infra/sqlite"
)

type CategorySqlite struct {
	Conn *sql.DB
	DB   *sqlite.Queries
}

func NewCategorySqlite(dbConn *sql.DB) repository.Category {
	return &CategorySqlite{
		Conn: dbConn,
		DB:   sqlite.New(dbConn),
	}
}

func (c *CategorySqlite) GetAll() ([]*entity.Category, exceptions.RepositoryException) {
	ctx := context.Background()

	categories := []*entity.Category{}

	categoriesOutput, err := c.DB.GetAllCategories(ctx)

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	for _, category := range categoriesOutput {
		categoryEntity, entityErr := entity.NewCategory(entity.CategoryProps{
			ID:          types.CategoryID(category.ID),
			PublicID:    types.CategoryPublicID(category.PublicID),
			Name:        category.Name,
			Description: category.Description.String,
		})

		if entityErr != nil {
			return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
				Reason: sqlite.Reason(entityErr),
			})
		}

		categories = append(categories, categoryEntity)
	}

	return categories, nil
}

func (c *CategorySqlite) GetOneByPublicID(publicId types.CategoryPublicID) (*entity.Category, exceptions.RepositoryException) {
	ctx := context.Background()

	categoryOutput, err := c.DB.GetOneCategoryByPublicID(ctx, string(publicId))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	categoryEntity, entityErr := entity.NewCategory(entity.CategoryProps{
		ID:          types.CategoryID(categoryOutput.ID),
		PublicID:    types.CategoryPublicID(categoryOutput.PublicID),
		Name:        categoryOutput.Name,
		Description: categoryOutput.Description.String,
	})

	if entityErr != nil {
		return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(entityErr),
		})
	}

	return categoryEntity, nil
}
