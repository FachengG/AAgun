package obj

type Gun struct {
	BulletsNum int
	// TODO: Add more features:
	// Angle               float64
	// RotateSpeed         float64
	// RequiredCoolingLifetime float64
	// CoolingLifetimeLeft     float64
}

func (gun *Gun) Shoot() bool {
	if gun.BulletsNum > 0 {
		gun.BulletsNum -= 1
		return true
	}
	return false
}
