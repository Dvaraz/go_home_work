package hw2

//  todo: Дан номер месяца (1 — январь, 2 — февраль, ...). Вывести название соответствующего
//  времени года ("зима", "весна" и т.д.)

func Task11(month int) string {
	if month <= 0 || month > 12 {
		return "Не правильный месяц, попробуйте от 1 до 12"
	}
	var season string
	if month >= 3 && month < 6 {
		season = "весна"
	} else if month >= 6 && month < 9 {
		season = "лето"
	} else if month >= 9 && month < 12 {
		season = "осень"
	} else {
		season = "зима"
	}
	return season
}
