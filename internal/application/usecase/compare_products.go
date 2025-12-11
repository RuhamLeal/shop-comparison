package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type CompareProducts struct {
	ProductRepository repository.Product
	code              string
}

func NewCompareProducts(
	productRepository repository.Product,
) *CompareProducts {
	return &CompareProducts{
		code:              "CompareProducts",
		ProductRepository: productRepository,
	}
}

func (u *CompareProducts) Execute(input dto.CompareProductsInput) (*dto.CompareProductsOutput, exceptions.UsecaseException) {
	leftProduct, repoErr := u.ProductRepository.GetOneByPublicId(input.LeftPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product",
		})
	}

	rightProduct, repoErr := u.ProductRepository.GetOneByPublicId(input.RightPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product",
		})
	}

	result, entityErr := leftProduct.Compare(rightProduct)

	if entityErr != nil {
		return nil, exceptions.Usecase(entityErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 500,
			Message:    "Error comparing products",
		})
	}

	return u.toCompareProductsOutput(result)
}

func (u *CompareProducts) toCompareProductsOutput(result *entity.ComparisonProductsResult) (*dto.CompareProductsOutput, exceptions.UsecaseException) {
	output := &dto.CompareProductsOutput{
		Price: &dto.PriceComparisonOutput{
			Left:  result.PriceComparisonResult.Left,
			Right: result.PriceComparisonResult.Right,
		},
		Rating: &dto.RatingComparisonOutput{
			Left:  result.RatingComparisonResult.Left,
			Right: result.RatingComparisonResult.Right,
		},
		Specifications: []*dto.SpecificationsComparisonOutput{},
	}

	for _, insight := range result.PriceComparisonResult.Insights {
		output.Price.Insights = append(output.Price.Insights, &dto.InsightOutput{
			Favorable: insight.Favorable,
			Neutral:   insight.Neutral,
			Message:   insight.Message,
		})
	}

	for _, insight := range result.RatingComparisonResult.Insights {
		output.Rating.Insights = append(output.Rating.Insights, &dto.InsightOutput{
			Favorable: insight.Favorable,
			Neutral:   insight.Neutral,
			Message:   insight.Message,
		})
	}

	for _, specificationResult := range result.SpecificationsComparisonResults {
		outputSpecification := &dto.SpecificationsComparisonOutput{
			Type: specificationResult.Left.Type,
			Left: &dto.SpecificationComparisonOutput{
				StringValue: specificationResult.Left.Value.StringValue,
				IntValue:    specificationResult.Left.Value.IntValue,
				BoolValue:   specificationResult.Left.Value.BoolValue,
			},
			Right: &dto.SpecificationComparisonOutput{
				StringValue: specificationResult.Right.Value.StringValue,
				IntValue:    specificationResult.Right.Value.IntValue,
				BoolValue:   specificationResult.Right.Value.BoolValue,
			},
		}

		for _, insight := range specificationResult.Insights {
			outputSpecification.Insights = append(outputSpecification.Insights, &dto.InsightOutput{
				Favorable: insight.Favorable,
				Neutral:   insight.Neutral,
				Message:   insight.Message,
			})
		}

		output.Specifications = append(output.Specifications, outputSpecification)
	}

	return output, nil
}
