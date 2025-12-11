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

type SpecificationGroupSqlite struct {
	Conn *sql.DB
	DB   *sqlite.Queries
}

func NewSpecificationGroupSqlite(dbConn *sql.DB) repository.SpecificationGroup {
	return &SpecificationGroupSqlite{
		Conn: dbConn,
		DB:   sqlite.New(dbConn),
	}
}

func (s *SpecificationGroupSqlite) GetAll() ([]*entity.SpecificationGroup, exceptions.RepositoryException) {
	ctx := context.Background()

	specificationGroups := []*entity.SpecificationGroup{}

	specificationGroupsOutput, err := s.DB.GetAllSpecificationGroups(ctx)

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	for _, specificationGroupOutput := range specificationGroupsOutput {
		specificationGroups = append(specificationGroups, &entity.SpecificationGroup{
			ID:          types.SpecificationGroupID(specificationGroupOutput.ID),
			PublicID:    types.SpecificationGroupPublicID(specificationGroupOutput.PublicID),
			Name:        specificationGroupOutput.Name,
			Description: specificationGroupOutput.Description.String,
		})
	}

	return specificationGroups, nil
}

func (s *SpecificationGroupSqlite) GetOneByPublicID(publicId types.SpecificationGroupPublicID) (*entity.SpecificationGroup, exceptions.RepositoryException) {
	ctx := context.Background()

	specificationGroupOutput, err := s.DB.GetOneSpecificationGroupByPublicID(ctx, string(publicId))

	if err != nil {
		return nil, exceptions.Repo(err, exceptions.RepositoryOpts{
			Reason: sqlite.Reason(err),
		})
	}

	return &entity.SpecificationGroup{
		ID:          types.SpecificationGroupID(specificationGroupOutput.ID),
		PublicID:    types.SpecificationGroupPublicID(specificationGroupOutput.PublicID),
		Name:        specificationGroupOutput.Name,
		Description: specificationGroupOutput.Description.String,
	}, nil
}
