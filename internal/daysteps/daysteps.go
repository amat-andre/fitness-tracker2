package daysteps

import (
	"fmt"
	"time"
	"errors"
	"strings"
	"strconv"

	pd "github.com/Yandex-Practicum/tracker/internal/personaldata"
	sp "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var errInput = errors.New("input value error")

type DaySteps struct {
	pd.Personal
	Steps int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	lines := strings.Split(datastring, ",")
	if len(lines) != 2 {
		return fmt.Errorf("[Parse(%v)] %w: required input format %q", datastring, errInput, "678,0h50m")
	}

	steps, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("[Parse(%v)] invalid steps format: %w", lines[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("[Parse(%v)] %w: 'steps' must be greater '0'", steps, errInput)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(lines[1])
	if err != nil {
		return fmt.Errorf("[Parse(%v)] invalid duration format: %w", lines[1], err)
	}
	if duration <= 0 {
		return fmt.Errorf("[Parse(%v)] %w: 'duration' must be greater '0'", duration, errInput)
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := sp.Distance(ds.Steps, ds.Height)

	calories, err := sp.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("[ActionInfo] %w", err)
	}

	line := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return line, nil
}
