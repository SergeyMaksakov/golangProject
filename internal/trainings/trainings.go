package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training содержит детальную информацию о тренировке и персональные данные пользователя.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку с данными о тренировке формата "3456,Ходьба,3h00m" и сохраняет их в структуру.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных для тренировки")
	}

	// 1-3. Парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка конвертации шагов: %w", err)
	}
	if steps < 0 {
		return errors.New("количество шагов не может быть отрицательным")
	}
	t.Steps = steps

	// 4. Сохраняем тип тренировки
	t.TrainingType = parts[1]

	// 5. Парсим продолжительность (time.Duration)
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность тренировки должна быть больше нуля")
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает строку с подробной информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// 1. Вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Height)

	// 2. Вычисляем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	// 3. Рассчитываем калории в зависимости от типа тренировки
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	// 4. Формируем и возвращаем итоговую строку
	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	)

	return info, nil
}
