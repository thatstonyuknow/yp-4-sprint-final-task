package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

)

var (
	ErrConvToInt        = errors.New("converting into int error")
	ErrWrongLenthSlice  = errors.New("wrong lenth of slice error")
	ErrParseDate		= errors.New("parsing date error")
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, ErrWrongLenthSlice
	}

	stepNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, ErrConvToInt
	}

	t, err := time.ParseDuration(parts[2])
    if err != nil {
        return 0, "", 0, ErrParseDate
    }

	training := map[string]string{
		"Бег":	"Running",
		"Ходьба":	"Walking",
	}

	if _, ok := training[parts[1]]; !ok {
		return 0, "", 0, ErrParseDate
	}
	
	return stepNumber, parts[1], t, nil	
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	return (float64(steps) * lenStep) / float64(mInKm)
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 {
		return 0
	}
	distance := distance(steps)
	return distance / duration.Hours()
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, activity, t, err := parseTraining(data)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
		return ""
	}

	distanceVal := distance(steps)
    meanSpd := meanSpeed(steps, t)

	switch activity {
	case "Бег":
		calories := RunningSpentCalories(steps, weight, t)
		message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, t.Hours(), distanceVal, meanSpd, calories)
		return  message

	case "Ходьба":
		calories := WalkingSpentCalories(steps, weight, height, t)
		message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, t.Hours(), distanceVal, meanSpd, calories)
		return  message
		
	default:
		return "неизвестный тип тренировки"

	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)
	raw := ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight
	if raw < 0 {
		raw = 0
	}
	return raw
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)

	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * float64(minInH)

}