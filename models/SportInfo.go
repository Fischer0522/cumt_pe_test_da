package models

type SportInfo struct {
	Height float64
	Weight float64
	VitalCapacity int64
	Meter50 float64 `gorm:"column:meter_50"`
	Jump int64
	SitAndReach float64
	Meter800 string `gorm:"column:meter_800"`
	Meter1000 string `gorm:"column:meter_1000"`
	SitUp  int64
	PullUp int64
	LeftVision float64
	RightVision float64
	HeightWeightScore int64
	VitalCapacityScore int64
	Meter50Score int64
	JumpScore int64
	SitAndReachScore int64
	SitUpPullUpScore int64
	TotalScore float64
	PEScore float64
	PassPercent float64
	GroupType string
	Grade string
	Class string
	Department string
	province string
	Sex string

}

func (SportInfo) TableName() string {
	return "student_test_score_2020_da"
}
