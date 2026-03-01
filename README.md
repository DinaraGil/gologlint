# gologlint
Линтер для проверки лог-записей в Go-проектах. Разработан как плагин для **golangci-lint**.
## Установка и подключение

### Создать `.custom-gcl.yml` в корне проекта

```yaml
version: main

plugins:
  - module: github.com/DinaraGil/gologlint/plugin
    path: /path/to/gologlint/plugin
```

### Создать `.golangci.yml` в корне проекта

```yaml
linters:
  disable-all: true
  enable:
    - gologlint
```

### Собрать кастомную версию golangci-lint командой

```bash
golangci-lint custom -v
```
Будет создан исполняемый файл:

```
custom-gcl.exe
```

### Запуск линтера

Проверка всего проекта:

```bash
./custom-gcl run
```

Проверка конкретного файла:

```bash
./custom-gcl run main.go
```
## Клонирование проекта
```
git clone https://github.com/DinaraGil/gologlint.git
cd gologlint
```
## Тестовый файл в example/
Build линтера и запуск для ./example
```
 go build -o bin/gologlint.exe ./cmd/gologlint
./bin/gologlint.exe ./example
 ```

## Тестирование
Unit-тесты для каждого правила в tests/