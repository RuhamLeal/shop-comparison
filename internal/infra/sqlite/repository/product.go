package repository

import (
	"context"
	"database/sql"
	"errors"
	"project/internal/domain/aggregate"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
	"project/internal/domain/types"
	"project/internal/infra/sqlite"
)

type ProductSqlite struct {
	Conn *sql.DB
	DB   *sqlite.Queries
}

func NewProductSqlite(dbConn *sql.DB) repository.Product {
	return &ProductSqlite{
		Conn: dbConn,
		DB:   sqlite.New(dbConn),
	}
}

func (p *ProductSqlite) CreateOne(product *entity.Product) exceptions.RepositoryException {
	ctx := context.Background()

	result, err := p.DB.CreateOneProduct(ctx, sqlite.CreateOneProductParams{
		PublicID:    string(product.PublicID),
		Name:        string(product.Name),
		Description: sql.NullString{String: product.Description, Valid: product.Description != ""},
		Price:       product.Price,
		Rating:      int64(product.Rating),
		ImageUrl:    sql.NullString{String: product.ImageURL, Valid: product.ImageURL != ""},
		CategoryID:  int64(product.CategoryID),
	})

	if err != nil {
		return exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	id, err := result.LastInsertId()

	if err != nil {
		return exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	product.ID = types.ProductID(id)

	return nil
}

func (p *ProductSqlite) DeleteOne(product *entity.Product) exceptions.RepositoryException {
	ctx := context.Background()

	err := p.DB.DeleteOneProduct(ctx, int64(product.ID))

	if err != nil {
		return exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	return nil
}

func (p *ProductSqlite) ExistsByName(name types.ProductName, publicId types.ProductPublicID) (bool, exceptions.RepositoryException) {
	ctx := context.Background()

	id, err := p.DB.CheckIfProductExists(ctx, sqlite.CheckIfProductExistsParams{
		Name:     string(name),
		PublicID: publicId,
	})

	if err != nil {
		return false, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	if id == 0 {
		return false, nil
	}

	return true, nil
}

func (p *ProductSqlite) GetAll(paginationInput entity.PaginatorInput) ([]*entity.Product, entity.PaginatorOutput, exceptions.RepositoryException) {
	ctx := context.Background()

	paginatorOutput := &entity.PaginatorOutput{Total: 0}

	productsOutput, err := p.DB.GetAllProducts(ctx, sqlite.GetAllProductsParams{
		Limit:  paginationInput.Limit,
		Offset: paginationInput.Skip,
	})

	if err != nil {
		return nil, *paginatorOutput, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	productsList := make([]*entity.Product, 0, len(productsOutput))

	for _, productOutput := range productsOutput {
		product, entityErr := entity.NewProduct(entity.ProductProps{
			ID:                  types.ProductID(productOutput.ID),
			PublicID:            types.ProductPublicID(productOutput.PublicID),
			CategoryID:          types.CategoryID(productOutput.CategoryID),
			Name:                types.ProductName(productOutput.Name),
			Description:         productOutput.Description.String,
			Price:               productOutput.Price,
			Rating:              int8(productOutput.Rating),
			ImageURL:            productOutput.ImageUrl.String,
			SpecificationValues: []*entity.ProductSpecificationValue{},
		})

		if entityErr != nil {
			return nil, *paginatorOutput, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
				Reason: sqlite.Reason(entityErr),
			})
		}

		if paginatorOutput.Total == 0 {
			paginatorOutput.Total = productOutput.ProductsQuantity
		}

		productsList = append(productsList, product)
	}

	return productsList, *paginatorOutput, nil
}

func (p *ProductSqlite) GetAllByCategoryID(categoryId types.CategoryID, paginationInput entity.PaginatorInput) ([]*entity.Product, entity.PaginatorOutput, exceptions.RepositoryException) {
	ctx := context.Background()

	paginatorOutput := &entity.PaginatorOutput{Total: 0}

	productsOutput, err := p.DB.GetAllProductsByCategoryId(ctx, sqlite.GetAllProductsByCategoryIdParams{
		CategoryID: int64(categoryId),
		Limit:      paginationInput.Limit,
		Offset:     paginationInput.Skip,
	})

	if err != nil {
		return nil, *paginatorOutput, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	productsList := make([]*entity.Product, 0, len(productsOutput))

	for _, productOutput := range productsOutput {
		product, entityErr := entity.NewProduct(entity.ProductProps{
			ID:                  types.ProductID(productOutput.ID),
			PublicID:            types.ProductPublicID(productOutput.PublicID),
			CategoryID:          categoryId,
			Name:                types.ProductName(productOutput.Name),
			Description:         productOutput.Description.String,
			Price:               productOutput.Price,
			Rating:              int8(productOutput.Rating),
			ImageURL:            productOutput.ImageUrl.String,
			SpecificationValues: []*entity.ProductSpecificationValue{},
		})

		if entityErr != nil {
			return nil, *paginatorOutput, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
				Reason: sqlite.Reason(entityErr),
			})
		}

		if paginatorOutput.Total == 0 {
			paginatorOutput.Total = productOutput.ProductsQuantity
		}

		productsList = append(productsList, product)
	}

	return productsList, *paginatorOutput, nil
}

func (p *ProductSqlite) GetOneByPublicId(publicId types.ProductPublicID) (*entity.Product, exceptions.RepositoryException) {
	ctx := context.Background()

	productOutput, err := p.DB.GetOneProductByPublicId(ctx, string(publicId))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	product, entityErr := entity.NewProduct(entity.ProductProps{
		ID:                  types.ProductID(productOutput.ID),
		PublicID:            types.ProductPublicID(productOutput.PublicID),
		CategoryID:          types.CategoryID(productOutput.CategoryID),
		Name:                types.ProductName(productOutput.Name),
		Description:         productOutput.Description.String,
		Price:               productOutput.Price,
		Rating:              int8(productOutput.Rating),
		ImageURL:            productOutput.ImageUrl.String,
		SpecificationValues: []*entity.ProductSpecificationValue{},
	})

	if entityErr != nil {
		return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(entityErr),
		})
	}

	productSpecsOutput, err := p.DB.GetAllProductSpecificationValuesByProductID(ctx, int64(product.ID))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	for _, productSpecOutput := range productSpecsOutput {
		specValue := &entity.SpecValue{}
		if productSpecOutput.StringValue.Valid {
			specValue.StringValue = &productSpecOutput.StringValue.String
		}

		if productSpecOutput.IntValue.Valid {
			specValue.IntValue = &productSpecOutput.IntValue.Int64
		}

		if productSpecOutput.BoolValue.Valid {
			var boolVal bool
			if productSpecOutput.BoolValue.Int64 == 1 {
				boolVal = true
				specValue.BoolValue = &boolVal
			} else {
				specValue.BoolValue = &boolVal
			}
		}
		productSpecEntity, entityErr := entity.NewProductSpecificationValue(entity.ProductSpecificationValueProps{
			ID:              productSpecOutput.ID,
			ProductID:       types.ProductID(productSpecOutput.ProductID),
			SpecificationID: types.SpecificationID(productSpecOutput.SpecificationID),
			Type:            types.SpecificationType(productSpecOutput.Type),
			Value:           specValue,
		})

		if entityErr != nil {
			return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
				Reason: sqlite.Reason(entityErr),
			})
		}

		product.SpecificationValues = append(product.SpecificationValues, productSpecEntity)
	}

	return product, nil
}

func (p *ProductSqlite) GetOneByPublicIdWithSpecificationGroups(publicId types.ProductPublicID) (*aggregate.ProductWithSpecificationsGroups, exceptions.RepositoryException) {
	ctx := context.Background()

	outputs, err := p.DB.GetOneProductWithSpecificationsByPublicId(ctx, string(publicId))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	if len(outputs) == 0 {
		return nil, exceptions.Repo(errors.New("product not found"))
	}

	product, entityErr := entity.NewProduct(entity.ProductProps{
		ID:                  types.ProductID(outputs[0].ProductID),
		PublicID:            types.ProductPublicID(outputs[0].ProductPublicID),
		CategoryID:          types.CategoryID(outputs[0].ProductCategoryID),
		Name:                types.ProductName(outputs[0].ProductName),
		Description:         outputs[0].ProductDescription.String,
		Price:               outputs[0].ProductPrice,
		Rating:              int8(outputs[0].ProductRating),
		ImageURL:            outputs[0].ProductImageUrl.String,
		SpecificationValues: []*entity.ProductSpecificationValue{},
	})

	if entityErr != nil {
		return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(entityErr),
		})
	}

	aggregate := &aggregate.ProductWithSpecificationsGroups{
		Product:              product,
		SpecificationsGroups: []*entity.SpecificationGroup{},
	}

	addedGroups := make(map[types.SpecificationGroupID]*entity.SpecificationGroup)

	for _, output := range outputs {
		specValue := &entity.SpecValue{}
		if output.SpecificationStringValue.Valid {
			specValue.StringValue = &output.SpecificationStringValue.String
		}

		if output.SpecificationIntValue.Valid {
			specValue.IntValue = &output.SpecificationIntValue.Int64
		}

		if output.SpecificationBoolValue.Valid {
			var boolVal bool
			if output.SpecificationBoolValue.Int64 == 1 {
				boolVal = true
				specValue.BoolValue = &boolVal
			} else {
				specValue.BoolValue = &boolVal
			}
		}
		productSpecEntity, entityErr := entity.NewProductSpecificationValue(entity.ProductSpecificationValueProps{
			ProductID:       types.ProductID(output.ProductID),
			SpecificationID: types.SpecificationID(output.SpecificationID),
			Type:            types.SpecificationType(output.SpecificationType),
			Value:           specValue,
		})

		if entityErr != nil {
			return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
				Reason: sqlite.Reason(entityErr),
			})
		}

		product.SpecificationValues = append(product.SpecificationValues, productSpecEntity)

		addedSpecificationGroup, exists := addedGroups[types.SpecificationGroupID(output.SpecificationGroupID)]

		if !exists {
			specificationGroupEntity, entityErr := entity.NewSpecificationGroup(entity.SpecificationGroupProps{
				ID:             types.SpecificationGroupID(output.SpecificationGroupID),
				PublicID:       types.SpecificationGroupPublicID(output.SpecificationGroupPublicID),
				Name:           output.SpecificationGroupName,
				Description:    output.SpecificationGroupDescription.String,
				Specifications: []*entity.Specification{},
			})

			if entityErr != nil {
				return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
					Reason: sqlite.Reason(entityErr),
				})
			}

			specificationEntity, entityErr := entity.NewSpecification(entity.SpecificationProps{
				ID:                    types.SpecificationID(output.SpecificationID),
				PublicID:              types.SpecificationPublicID(output.SpecificationPublicID),
				Title:                 output.SpecificationTitle,
				EspecificationGroupID: types.SpecificationGroupID(output.SpecificationGroupID),
				Type:                  types.SpecificationType(output.SpecificationType),
			})

			if entityErr != nil {
				return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
					Reason: sqlite.Reason(entityErr),
				})
			}

			specificationGroupEntity.Specifications = append(specificationGroupEntity.Specifications, specificationEntity)

			addedGroups[types.SpecificationGroupID(output.SpecificationGroupID)] = specificationGroupEntity

			aggregate.SpecificationsGroups = append(aggregate.SpecificationsGroups, specificationGroupEntity)
		} else {
			specificationEntity, entityErr := entity.NewSpecification(entity.SpecificationProps{
				ID:                    types.SpecificationID(output.SpecificationID),
				PublicID:              types.SpecificationPublicID(output.SpecificationPublicID),
				Title:                 output.SpecificationTitle,
				EspecificationGroupID: types.SpecificationGroupID(output.SpecificationGroupID),
				Type:                  types.SpecificationType(output.SpecificationType),
			})

			if entityErr != nil {
				return nil, exceptions.Repo(entityErr, exceptions.RepositoryOpts{
					Reason: sqlite.Reason(entityErr),
				})
			}

			addedSpecificationGroup.Specifications = append(addedSpecificationGroup.Specifications, specificationEntity)
		}
	}

	return aggregate, nil
}

func (p *ProductSqlite) UpdateOne(product *entity.Product) exceptions.RepositoryException {
	ctx := context.Background()

	err := p.DB.UpdateOneProduct(ctx, sqlite.UpdateOneProductParams{
		ID:          int64(product.ID),
		Name:        string(product.Name),
		Description: sql.NullString{String: product.Description, Valid: product.Description != ""},
		Price:       product.Price,
		Rating:      int64(product.Rating),
		ImageUrl:    sql.NullString{String: product.ImageURL, Valid: product.ImageURL != ""},
	})

	if err != nil {
		return exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	return nil
}
