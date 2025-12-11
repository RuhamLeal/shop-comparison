package constants

import "project/internal/domain/types"

const (
	SpecificationTypeString types.SpecificationType = "string"
	SpecificationTypeInt    types.SpecificationType = "int"
	SpecificationTypeBool   types.SpecificationType = "bool"
)

const (
	PowerInWatts   types.SpecificationID = 1
	ConsumptionKwh types.SpecificationID = 2
	CapacityLiters types.SpecificationID = 3
	FrequencyMHz   types.SpecificationID = 4
	FrequencyGHz   types.SpecificationID = 5
	Threads        types.SpecificationID = 6
	TDPWatts       types.SpecificationID = 7
	USBC           types.SpecificationID = 8
	Waterproof     types.SpecificationID = 9
	NoiseDb        types.SpecificationID = 10
	CaloriesKcal   types.SpecificationID = 11
	WidthCm        types.SpecificationID = 12
	HeightCm       types.SpecificationID = 13
	DepthCm        types.SpecificationID = 14
	WeightKg       types.SpecificationID = 15
	VolumeLiters   types.SpecificationID = 16
)
