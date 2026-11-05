package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                        = 1000 // количество метров в километре.
	minInH                       = 60   // количество минут в часе.
	stepLengthCoefficient        = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает сожженные калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры для расчета калорий ходьбы")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	// (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient
	calories := ((weight * meanSpeed * durationInMinutes) / float64(minInH)) * walkingCaloriesCoefficient
	return calories, nil
}

// RunningSpentCalories рассчитывает сожженные калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры для расчета калорий бега")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	// (weight * meanSpeed * durationInMinutes) / minInH
	calories := (weight * meanSpeed * durationInMinutes) / float64(minInH)
	return calories, nil
}

// MeanSpeed рассчитывает среднюю скорость перемещения (км/ч).
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// 1. Проверить, что продолжительность duration больше 0 (а также шаги и рост)
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}

	// 2. Вычислить дистанцию с помощью Distance()
	distance := Distance(steps, height)

	// 3. Вычислить и вернуть среднюю скорость (дистанция в часах)
	durationHours := duration.Hours()
	if durationHours <= 0 {
		return 0
	}

	return distance / durationHours
}

// Distance рассчитывает пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	// Длина шага = рост * stepLengthCoefficient
	stepLength := height * stepLengthCoefficient

	// Пройденная дистанция в метрах = шаги * длина шага
	distanceInMeters := float64(steps) * stepLength

	// Переводим в километры
	return distanceInMeters / float64(mInKm)
}
