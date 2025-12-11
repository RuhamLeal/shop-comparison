package repository

import (
	"context"
	"database/sql"
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
	. "project/internal/domain/types"
	"project/internal/infra/sqlite"
)

type Specificationqlite struct {
	Conn *sql.DB
	DB   *sqlite.Queries
}

func NewSpecificationqlite(dbConn *sql.DB) repository.Specification {
	return &Specificationqlite{
		Conn: dbConn,
		DB:   sqlite.New(dbConn),
	}
}

func (s *Specificationqlite) GetAllByGroupID(specGroupID SpecificationGroupID) ([]*entity.Specification, RepositoryException) {
	ctx := context.Background()

	specifications := []*entity.Specification{}

	specificationsOutput, err := s.DB.GetAllSpecifications(ctx, int64(specGroupID))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	for _, specificationOutput := range specificationsOutput {
		specificationEntity := &entity.Specification{
			ID:                    SpecificationID(specificationOutput.ID),
			PublicID:              SpecificationPublicID(specificationOutput.PublicID),
			Title:                 specificationOutput.Title,
			EspecificationGroupID: specGroupID,
			Type:                  SpecificationType(specificationOutput.Type),
		}

		specifications = append(specifications, specificationEntity)
	}

	return specifications, nil
}

func (s *Specificationqlite) GetOneByPublicID(publicId SpecificationPublicID) (*entity.Specification, RepositoryException) {
	ctx := context.Background()

	specOutput, err := s.DB.GetOneSpecificationByPublicID(ctx, string(publicId))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	return &entity.Specification{
		ID:                    SpecificationID(specOutput.ID),
		PublicID:              SpecificationPublicID(specOutput.PublicID),
		Title:                 specOutput.Title,
		EspecificationGroupID: SpecificationGroupID(specOutput.ID_2),
		Type:                  SpecificationType(specOutput.Type),
	}, nil
}
