# Timestamp inline parser

Задача: в выводе команды `$ curl -s localhost:5050/mgsv/outer/heaven | jq .` присутствуют дататаймы с наносекундами в формате UNIX timestamp, нужно заменить дататймы на человеко-читаемые сохраняя вывод в json.

Для примера вывод команды сохранён в `input.json`.

Например, нужно из
```
      "start": 1696868796296203000,
      "end": 1696868802596088000,
```
получить
```
      "start": 2023.10.09_19.26.36.296203,
      "end": 2023.10.09_19.26.42.596088,
```

## Go way

Компилируем go приложение и используем команду
```
$ curl -s localhost:5050/mgsv/outer/heaven | jq . | ./time-handler 
```
или
```
$ cat input.json | jq . | go run time-handler.go
{
  "id": 1,
  "Snake": 1,
  "Shalashaka": 4,
  "Snake_r": {
    "begin": {
      "start": 2023.10.09_19.26.42.696503,
      "end": 2023.10.09_19.26.45.096412,
...
```

## Python way
[ломает json]

Компилировать ничего не нужно, просто используем скрипт
```
$ curl -s localhost:5050/mgsv/outer/heaven | jq . | python3 time-handler.py 
```
или
```
$ cat input.json | jq . | python3 time-handler.py | head -8
{
"id": 1,
"Snake": 1,
"Shalashaka": 4,
"Snake_r": {
"begin": {
"start": 2023.10.09_16.26.42.696503,
"end": 2023.10.09_16.26.45.096412,
```