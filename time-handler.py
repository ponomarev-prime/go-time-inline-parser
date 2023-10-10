import re
import sys
import datetime

# Регулярное выражение для поиска дат в формате Unix Nanoseconds
date_regex = r'\d{19}'

# Функция для замены дат в строке на новый формат
def convert_dates(line):
    def replace_date(match):
        unix_time_ns = int(match.group())
        unix_time_s = unix_time_ns / 1e9
        dt = datetime.datetime.utcfromtimestamp(unix_time_s)
        formatted_date = dt.strftime('%Y.%m.%d_%H.%M.%S.%f')
        return formatted_date

    return re.sub(date_regex, replace_date, line)

# Обрабатываем ввод
for line in sys.stdin:
    converted_line = convert_dates(line.strip())
    print(converted_line)
