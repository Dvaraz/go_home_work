#!/bin/bash
# Берем имя файла из аргумента, выбираем случайную строку из файла, формата: слово-определение.
filename=$1
# Разделяем слово и определение и присваеваем в переменные, затем приводим слово к нижнему регистру.

clean_line() {
    echo "$1" | sed -e 's/^[ \t]*//' -e 's/[ \t]*$//' -e 's/^"//' -e 's/"$//'
}
process_file() {
	local file_path="$1"
	if [[ -f "$file_path" ]]; 
    	then
		word_and_definition=$(shuf -n 1 "$file_path")
		IFS='-' read -r word def <<< "$word_and_definition"
		word=$(clean_line "$word")
		def=$(clean_line "$def")
		word="${word,,}"
	else
		echo "Ошибка файл не найден будут использованны стандартные слова и определения"
		sleep 1
		return 1
	fi
}

get_word_and_def() {
	words=("оператор" "конструкция" "объект")
	defs=("Это слово обозначает наименьшую автономную часть языка программирования" "def2" "def3" )
	random_index=$(( RANDOM % ${#words[@]} ))
	word="${words[$random_index]}"
	def="${defs[$random_index]}"
}

engine() {
	tries=10
	masked_word=$word
	while [ $tries -gt 0 ]
	do
	if [[ -z "$masked_word" ]]
	then
		echo "Вы выиграли!!!"
		echo "Это слово - $word"
		exit 0
	fi

	echo "$def"
	echo "У вас осталось $tries попыток! "
	echo ${word//[$masked_word]/\# }
	echo -n "Введите букву: " 
	read letter
	letter="${letter,}" 
	if [[ ${#letter} -gt 1 ]]
	then
		echo "Введите только одну букву"
	elif [[ "$word" == *"$letter"* ]]
	then
		masked_word=${masked_word//$letter/}
	else
		tries=$((tries - 1))
		echo "Нет такой буквы."
	fi
	done

	echo "У Вас закончились попытки, к сожелению Вы проиграли это было слово $word"
}
main () {
	process_file "$filename"
	if [ $? -eq 1 ]
	then
		get_word_and_def
	fi
	engine
}

main "$@"
