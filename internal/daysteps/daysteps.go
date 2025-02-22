package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

var (
	ErrConvToInt       = errors.New("converting into int error")
	ErrWrongLenthSlice = errors.New("wrong lenth of slice error")
	ErrParseDate       = errors.New("parsing date error")
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, ErrWrongLenthSlice
	}
	// removing spaces
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])

	stepNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, ErrConvToInt
	}

	t, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, ErrParseDate
	}

	return stepNumber, t, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, t, err := parsePackage(data)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := (StepLength * float64(steps)) / 1000

	calories := spentcalories.WalkingSpentCalories(steps, weight, height, t)

	message := fmt.Sprintf("Количество шагов: %d\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

	return message
}
