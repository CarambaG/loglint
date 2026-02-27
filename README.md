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
