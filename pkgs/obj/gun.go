package obj

type Gun struct {
	BulletsNum int
	// Angle               float64
	// RotateSpeed         float64
	// RequiredCoolingTime float64
	// CoolingTimeLeft     float64
}

func (gun *Gun) Shoot() bool {

	gun.BulletsNum -= 1

}
