# loglint (log/slog/zap)
`loglint` - это кастомный линтер для Go, интегрируемый в golangci-lint. Он проверяет лог-записи библиотек `slog` и `zap` на соблюдение корпоративных правил: стиль сообщений, язык, отсутствие спецсимволов и чувствительных данных.

## Способ интеграции `loglint` c `golangci-lint`
1. Скопируйте исходный репозиторий `golangci-lint`
2. Создайте папку `loglint` в `golangci-lint/pkg/golinters`
3. Скопируйте папку `analyzer` из `loglint` в `golangci-lint/pkg/golinters/loglint/`
4. Добавить в файл `golangci-lint/pkg/lint/lintersdb/builder_linter.go` недостающий код в `[]*linter.Config:`
```go
linter.NewConfig(loglint.New()).
  WithLoadForGoAnalysis(),
```
5. Готово. Можно собирать проект с новым линтером

## Правила линтера
`loglint` проверяет все лог-записи на соответствие следующим правилам:

**1. Лог-сообщения должны начинаться со строчной буквы**
❌ Неправильно
```go
log.Info("Starting server on port 8080")
slog.Error("Failed to connect to database")
```

✅ Правильно
```go
log.Info("starting server on port 8080")
slog.Error("failed to connect to database")
```
**2. Лог-сообщения должны быть только на английском языке**
❌ Неправильно
```go
log.Info("запуск сервера")
slog.Error("ошибка подключения к базе данных")
```

✅ Правильно
```go
log.Info("starting server")
slog.Error("failed to connect to database")
```

**3. Лог-сообщения не должны содержать спецсимволы или эмодзи**
❌ Неправильно
```go
log.Info("server started! 🎉🚀")
slog.Error("connection failed!!!")
slog.Warn("warning: something went wrong...")
```

✅ Правильно
```go
log.Info("server started")
slog.Error("connection failed")
slog.Warn("something went wrong")
```
**4. Лог-сообщения не должны содержать потенциально чувствительные данные**
❌ Неправильно
```go
log.Info("user password: " + password)
slog.Debug("api_key=" + apiKey)
slog.Info("token: " + token)
```

✅ Правильно
```go
log.Info("user authenticated successfully")
slog.Debug("api request completed")
slog.Info("token validated")
```
