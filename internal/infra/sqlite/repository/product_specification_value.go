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

type ProductSpecificationValueSqlite struct {
	Conn *sql.DB
	DB   *sqlite.Queries
}

func NewProductSpecificationValueSqlite(dbConn *sql.DB) repository.ProductSpecificationValue {
	return &ProductSpecificationValueSqlite{
		Conn: dbConn,
		DB:   sqlite.New(dbConn),
	}
}

func (p *ProductSpecificationValueSqlite) CreateOne(productSpec *entity.ProductSpecificationValue) exceptions.RepositoryException {
	ctx := context.Background()

	var stringVal string
	var intVal int64
	var boolVal int64

	if productSpec.Value.StringValue != nil {
		stringVal = *productSpec.Value.StringValue
	}

	if productSpec.Value.IntValue != nil {
		intVal = *productSpec.Value.IntValue
	}

	if productSpec.Value.BoolValue != nil {
		if *productSpec.Value.BoolValue {
			boolVal = 1
		} else {
			boolVal = 0
		}

	}

	productSpecOutput, err := p.DB.CreateOneProductSpecificationValue(ctx, sqlite.CreateOneProductSpecificationValueParams{
		ProductID:       int64(productSpec.ProductID),
		SpecificationID: int64(productSpec.SpecificationID),
		StringValue:     sql.NullString{String: stringVal, Valid: productSpec.Value.StringValue != nil},
		IntValue:        sql.NullInt64{Int64: intVal, Valid: productSpec.Value.IntValue != nil},
		BoolValue:       sql.NullInt64{Int64: boolVal, Valid: productSpec.Value.BoolValue != nil},
	})

	if err != nil {
		return exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	id, err := productSpecOutput.LastInsertId()

	if err != nil {
		return exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	productSpec.ID = id

	return nil
}

func (p *ProductSpecificationValueSqlite) FindManyByProductID(productID types.ProductID) ([]*entity.ProductSpecificationValue, exceptions.RepositoryException) {
	ctx := context.Background()

	productSpecs := []*entity.ProductSpecificationValue{}

	productSpecsOutput, err := p.DB.GetAllProductSpecificationValuesByProductID(ctx, int64(productID))

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

		productSpecs = append(productSpecs, productSpecEntity)
	}

	return productSpecs, nil
}
