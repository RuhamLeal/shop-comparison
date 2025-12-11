package entity

import . "project/internal/domain/types"

type Insight struct {
	ProductID ProductID
	Favorable bool
	Neutral   bool
	Message   string
}

type InsightProps struct {
	ProductID ProductID
	Favorable bool
	Neutral   bool
	Message   string
}

func NewInsight(props InsightProps) *Insight {
	insight := &Insight{
		ProductID: props.ProductID,
		Favorable: props.Favorable,
		Message:   props.Message,
	}

	return insight
}
