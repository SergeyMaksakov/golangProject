package actioninfo

import (
	"log"
)

// DataParser определяет интерфейс для парсинга данных и получения информации об активности.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными и объект активности, 
// парсит каждую строку, формирует отчет и выводит его на экран.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// 1-2. Распарсить значение с помощью метода Parse
		err := dp.Parse(data)
		if err != nil {
			// 3. Обработать ошибку парсинга (залогировать и перейти к следующей итерации)
			log.Printf("ошибка парсинга данных: %v", err)
			continue
		}

		// 4-5. Сформировать строку с информацией об активности и вывести на экран
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации: %v", err)
			continue
		}

		println(info)
	}
}
