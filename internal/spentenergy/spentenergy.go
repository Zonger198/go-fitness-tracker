package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func Distance(steps int, height float64) float64 {
	return float64(steps) * (height * stepLengthCoefficient) / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}
