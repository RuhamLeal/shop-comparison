package entity_test

import (
	domain_entity "project/internal/domain/entity"
	"testing"
)

func TestNewInsight(t *testing.T) {
	tests := []struct {
		name  string
		props domain_entity.InsightProps
	}{
		{
			name: "Should create a valid Favorable insight",
			props: domain_entity.InsightProps{
				ProductID: 10,
				Favorable: true,
				Neutral:   false,
				Message:   "Sales are increasing",
			},
		},
		{
			name: "Should create a valid Neutral insight",
			props: domain_entity.InsightProps{
				ProductID: 12,
				Favorable: false,
				Neutral:   true,
				Message:   "Market is stable",
			},
		},
		{
			name: "Should create a valid Unfavorable insight",
			props: domain_entity.InsightProps{
				ProductID: 15,
				Favorable: false,
				Neutral:   false,
				Message:   "Sales dropped",
			},
		},
		{
			name: "Should map empty Message correctly",
			props: domain_entity.InsightProps{
				ProductID: 20,
				Favorable: true,
				Message:   "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insight := domain_entity.NewInsight(tt.props)

			if insight == nil {
				t.Error("Expected insight instance, but got nil")
				return
			}

			if insight.ProductID != tt.props.ProductID {
				t.Errorf("Expected ProductID %v, got %v", tt.props.ProductID, insight.ProductID)
			}
			if insight.Favorable != tt.props.Favorable {
				t.Errorf("Expected Favorable %v, got %v", tt.props.Favorable, insight.Favorable)
			}
			if insight.Neutral != tt.props.Neutral {
				t.Errorf("Expected Neutral %v, got %v", tt.props.Neutral, insight.Neutral)
			}
			if insight.Message != tt.props.Message {
				t.Errorf("Expected Message %q, got %q", tt.props.Message, insight.Message)
			}
		})
	}
}
