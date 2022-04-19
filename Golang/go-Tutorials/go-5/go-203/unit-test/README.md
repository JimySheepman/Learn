# Unit Test

```Bash
# run test
go test main_test.go
# belirli test funsiyonunu çalıştırmak için
go test -timeout 30s -run ^TestPercentges$ test/basket
# coverage
go test -cover
# satır satır coverage
go test -cover -coverprofile=cover.out
# görselleştirmek için
go tool cover -html=cover.out -o coverage.html
```
