#!/bin/bash

# Настройка переменных
FILE_NAME="cloud-swagger.json"
NEW_VERSION="v3.6.2"
GENERATOR="openapi-generator-cli-7.18.0.jar"
CONFIG="config/genConfig.yml"

if [ ! -f ".swagger/$FILE_NAME" ]; then
    echo "Ошибка: .swagger/$FILE_NAME не найден!"
    exit 1
fi

if [ ! -f $CONFIG ]; then
    echo "Ошибка: $CONFIG не найден!"
    exit 1
fi

echo "Начало генерации клиента..."

# Обновление версии в genConfig.yml
echo "Обновление версии в $CONFIG..."
sed -i "s/packageVersion: .*/packageVersion: $NEW_VERSION/" $CONFIG

# Очистка предыдущей генерации
echo "Очистка предыдущей генерации..."
rm -rf new

# Генерация .NET API клиента
echo "Генерация .NET API клиента..."
java -jar .vendor/$GENERATOR generate \
  -i .swagger/$FILE_NAME \
  -g go \
  -o new \
  --skip-validate-spec \
  -c $CONFIG

# Проверка успешности генерации
if [ ! -d "new" ]; then
    echo "Ошибка: Не удалось сгенерировать клиент!"
    exit 1
fi

# Удаление всех .go файлов в текущей директории (без рекурсии)
echo "Удаление всех .go файлов в текущей директории (без рекурсии)..."
rm -f *.go

# Копирование всех .go файлов из new/ в текущую директорию с заменой (кроме *_test.go)
echo "Копирование всех .go файлов из new/ в текущую директорию с заменой (кроме *_test.go)..."
find new -name "*.go" -not -name "*_test.go" -exec cp {} . \;

# Копирование документации если она была сгенерирована
if [ -d "new/docs" ]; then
    echo "Копирование документации..."
    rm -rf docs/* || true
    cp -r new/docs/* docs/ || true
fi

# Частичное обновление README.md
echo "Частичное обновление README.md..."
if [ -f "new/README.md" ]; then
    # Создаем копию нового README для обновления
    cp new/README.md README-NEW.md
    ./update-docs.sh
fi

echo "Генерация клиента завершена успешно!"

# Проверка и добавление импорта "time" в api_test_runs.go если отсутствует
echo "Проверка наличия импорта \"time\" в api_test_runs.go..."
if [ -f "api_test_runs.go" ]; then
    if ! grep -q "\"time\"" api_test_runs.go; then
        echo "Добавление импорта \"time\" в api_test_runs.go..."
        # Найдем строку с закрывающей скобкой блока import и вставим "time" перед ней
        awk '/import \($/ {imp=1; print; next} imp==1 && /\)/ {print "	\"time\""; imp=0} {print}' api_test_runs.go > temp_file && mv temp_file api_test_runs.go
        echo "Импорт \"time\" успешно добавлен."
    else
        echo "Импорт \"time\" уже присутствует."
    fi
else
    echo "Файл api_test_runs.go не найден."
fi

echo "Сборка проекта"
go build
