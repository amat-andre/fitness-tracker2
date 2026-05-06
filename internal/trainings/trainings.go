package trainings

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"

	pd "github.com/amat-andre/fitness-tracker2/internal/personaldata"
	sp "github.com/amat-andre/fitness-tracker2/internal/spentenergy"
)

var (
	errInput        = errors.New("input value error")
	errTrainingType = errors.New("неизвестный тип тренировки")
)

type Training struct {
	pd.Personal
	Steps int
	TrainingType string
	Duration time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	lines := strings.Split(datastring, ",")
	if len(lines) != 3 {
		return fmt.Errorf("[Parse(%v)] %w: required input format %q", datastring, errInput, "3456,Ходьба,3h00m")
	}

	steps, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("[Parse(%v)] invalid steps format: %w", lines[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("[Parse(%v)] %w: 'steps' must be greater '0'", steps, errInput)
	}
	t.Steps = steps

	t.TrainingType = lines[1]

	duration, err := time.ParseDuration(lines[2])
	if err != nil {
		return fmt.Errorf("[Parse(%v)] invalid duration format: %w", lines[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("[Parse(%v)] %w: 'duration' must be greater '0'", duration, errInput)
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := sp.Distance(t.Steps, t.Height)
	meanSpeed := sp.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories float64
		err error
	)
	switch t.TrainingType {
	case "Ходьба":
		calories, err = sp.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("[ActionInfo] %w", err)
		}
	case "Бег":
		calories, err = sp.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("[ActionInfo] %w", err)
		}
	default:
		return "", errTrainingType
	}

	line := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", 
	t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
	return line, nil
}
