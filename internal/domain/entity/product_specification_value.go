package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type SpecValue struct {
	StringValue *string
	IntValue    *int64
	BoolValue   *bool
}

type ProductSpecificationValue struct {
	ID              int64
	ProductID       ProductID
	SpecificationID SpecificationID
	Value           *SpecValue
}

type ProductSpecificationValueProps struct {
	ID              int64
	ProductID       ProductID
	SpecificationID SpecificationID
	Value           *SpecValue
}

type ComparisonProductSpecificationValues struct {
	Left     *ProductSpecificationValue
	Right    *ProductSpecificationValue
	Insights []*Insight
}

func NewProductSpecificationValue(props ProductSpecificationValueProps) (*ProductSpecificationValue, exceptions.EntityException) {
	productSpecificationValue := &ProductSpecificationValue{
		ID:              props.ID,
		ProductID:       props.ProductID,
		SpecificationID: props.SpecificationID,
		Value:           props.Value,
	}

	err := productSpecificationValue.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return productSpecificationValue, nil
}

func (s *ProductSpecificationValue) validateBeforeCompare(other *ProductSpecificationValue) error {
	if s.ID <= 0 || other.ID <= 0 {
		return errors.New("Cannot compare products with ID <= 0")
	}
	if s.ID == other.ID {
		return errors.New("Cannot compare the same product")
	}

	if s.ProductID <= 0 || other.ProductID <= 0 {
		return errors.New("Cannot compare products with ID <= 0")
	}
	if s.ProductID == other.ProductID {
		return errors.New("Cannot compare the same product")
	}

	if s.SpecificationID <= 0 || other.SpecificationID <= 0 {
		return errors.New("Cannot compare products with ID <= 0")
	}
	if s.SpecificationID != other.SpecificationID {
		return errors.New("Cannot compare products with different specifications")
	}
	return nil
}

func (s *ProductSpecificationValue) validate() error {
	if s.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if s.ProductID <= 0 {
		return errors.New("ProductID field must be greater than 0")
	}

	if s.SpecificationID <= 0 {
		return errors.New("SpecificationID field must be greater than 0")
	}

	hasString := s.Value.StringValue != nil
	hasInt := s.Value.IntValue != nil
	hasBool := s.Value.BoolValue != nil

	if !hasString && !hasInt && !hasBool {
		return errors.New("at least one value (String, Int, or Bool) must be provided")
	}

	return nil
}

func (s *ProductSpecificationValue) Compare(other *ProductSpecificationValue) (*ComparisonProductSpecificationValues, exceptions.EntityException) {
	if err := s.validateBeforeCompare(other); err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	comparisonCallback, exists := comparisons[s.SpecificationID]

	if !exists {
		return nil, exceptions.Entity(errors.New("no comparison callback found"), exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	result, err := comparisonCallback(s, other)

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return result, nil
}

var comparisons = map[SpecificationID]func(
	left *ProductSpecificationValue,
	right *ProductSpecificationValue,
) (*ComparisonProductSpecificationValues, error){
	constants.PowerInWatts: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Power requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "has higher power output"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may deliver better performance on heavy workloads"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "potentially consumes more energy"}),
			)
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "has lower power output"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "may be less performant in maximum load contexts"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may consume less energy"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both products deliver the same wattage"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "no meaningful difference in maximum power"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.ConsumptionKwh: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Consumption requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "consumes less energy"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may reduce long-term electricity costs"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "tends to be more eco-friendly"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "better for 24/7 operation scenarios"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "consumes more energy"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may increase electricity costs over time"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "less suitable for energy-efficient installations"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both products have identical energy consumption"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "no meaningful difference in long-term cost or usage impact"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.CapacityLiters: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Capacity requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "has greater internal capacity"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "can store more items simultaneously"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "better for families or heavy-duty usage"}),
			)
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "has smaller internal capacity"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "less suitable for large storage needs"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both products offer identical storage capacity"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.FrequencyMHz: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Frequency MHz requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "has higher operating frequency (MHz)"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may execute tasks faster depending on architecture"}),
			)
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "has lower operating frequency (MHz)"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may perform slower under computational bursts"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both operate at the same MHz frequency"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.FrequencyGHz: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Frequency GHz requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "operates at higher GHz"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may offer superior single-core performance"}),
			)
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "operates at lower GHz"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may have inferior single-core responsiveness"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both operate at the same GHz frequency"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.Threads: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Threads requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "supports more concurrent execution threads"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may deliver better performance in parallel workloads"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "better suited for multitasking and background processing"}),
			)
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "supports fewer threads"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may struggle in multi-threaded applications"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both support the same thread count"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.TDPWatts: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("TDP requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "has lower thermal design power"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may operate cooler and more quietly"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "likely requires less robust cooling solutions"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "has higher thermal design power"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may generate more heat under load"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "might require better ventilation or cooling"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both products have identical TDP values"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.USBC: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.BoolValue == nil || right.Value.BoolValue == nil {
			return nil, errors.New("USB-C requires boolean values")
		}

		l, r := *left.Value.BoolValue, *right.Value.BoolValue
		insights := []*Insight{}

		switch {
		case l && !r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "includes USB-C support"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "is compatible with modern charging and connectivity standards"}),
			)
		case !l && r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "does not include USB-C support"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both share the same USB-C capability"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.Waterproof: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.BoolValue == nil || right.Value.BoolValue == nil {
			return nil, errors.New("Waterproof requires boolean values")
		}

		l, r := *left.Value.BoolValue, *right.Value.BoolValue
		insights := []*Insight{}

		switch {
		case l && !r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "is waterproof"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "offers better resistance against accidental liquid exposure"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "may be more durable in humid or outdoor conditions"}),
			)
		case !l && r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is not waterproof"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may require extra care in wet environments"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both products share the same waterproof capability"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.NoiseDb: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Noise (dB) requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "operates more quietly"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "more suitable for silent environments"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: false, Message: "operates louder"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may be less comfortable for long-duration usage"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both have identical noise levels"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.CaloriesKcal: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Calories requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "contains fewer calories"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "more adequate for low-calorie diets"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may support weight management goals"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "contains more calories"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both contain the same caloric value"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.WidthCm: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Width requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is slimmer"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may fit better in small spaces"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is slimmer"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is wider"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may require more installation space"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both share the same width"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.HeightCm: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Height requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is shorter"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "more adequate for compact environments"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is shorter"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is taller"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may offer greater internal capacity depending on design"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both share the same height"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.DepthCm: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Depth requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is less deep"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "easier to accommodate in shallow installations"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is less deep"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is deeper"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "requires more installation clearance"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both share the same depth"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.WeightKg: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Weight requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is lighter"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "easier to transport or install"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "offers greater portability"}),
			)
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is lighter"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "is heavier"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may feel more robust depending on build quality"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both weigh the same"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},

	constants.VolumeLiters: func(left, right *ProductSpecificationValue) (*ComparisonProductSpecificationValues, error) {

		if left.Value.IntValue == nil || right.Value.IntValue == nil {
			return nil, errors.New("Volume requires int64 values")
		}

		l, r := *left.Value.IntValue, *right.Value.IntValue
		insights := []*Insight{}

		switch {
		case l > r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "offers more internal volume"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Favorable: true, Message: "better for storage, transport or operational space"}),
			)
		case l < r:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "offers smaller internal volume"}),
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "may limit operational or storage capacity"}),
			)
		default:
			insights = append(insights,
				NewInsight(InsightProps{ProductID: left.ProductID, Neutral: true, Message: "both offer the same volume"}),
			)
		}

		return &ComparisonProductSpecificationValues{Left: left, Right: right, Insights: insights}, nil
	},
}
