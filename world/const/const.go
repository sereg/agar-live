package _const

//EatRatio ratio when one object can eat another
const (
	FoodSize       = 3
	AliveStartSize = 15

	EatIncreaseRation     = 0.1
	EatSelfIncreaseRation = 0.9
	EatRatio              = 1.3

	ResurrectTime = 80

	SpeedRatio = 10
	StartSpeed = 12

	VisionRatio = 7
	StartVision = 10

	MinSizeAlive = 6

	GlueTime          = 600
	MinSizeSplit      = 20
	SplitRation       = 0.9
	SplitSpeed        = 5
	SplitDeceleration = 0.11
	SplitTime         = 60
	SplitMaxCount     = 10

	SplitRatio = 0.5
	BurstCount = 8

	PoisonColor = "#8eb021"
	PoisonSize  = 15
)

var (
	SplitDist = 0.0
	GridSize  = 70.0
)

func init() {
	v := float64(SplitSpeed)
	for {
		v -= SplitDeceleration
		if v <= 0 {
			break
		}
		SplitDist += v
	}
	println(int(SplitDist))
}
