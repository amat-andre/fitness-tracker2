package spentenergy

import (
	"fmt"
	"time"
	"errors"
)

var errInput = errors.New("input value error")

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("%w: 'steps' must be greater '0'", errInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: 'weight' must be greater '0'", errInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: 'height' must be greater '0'", errInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: 'duration' must be greater '0'", errInput)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	calories := (weight * meanSpeed * durationInMin) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("%w: 'steps' must be greater '0'", errInput)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("%w: 'weight' must be greater '0'", errInput)
	}
	if height <= 0 {
		return 0, fmt.Errorf("%w: 'height' must be greater '0'", errInput)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("%w: 'duration' must be greater '0'", errInput)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	calories := (weight * meanSpeed * durationInMin) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
} 

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}
