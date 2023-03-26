# kbbi

Install Go KBBI:
```go 
go get github.com/dimassfeb-09/kbbi
```

Import:
```go 
import "github.com/dimassfeb-09/kbbi/kbbi"
```

How to Use:
1. Use ```kbbi```
2. Search method ```.Search("keyword here")```
```go
if result, err := kbbi.Search("Baju"); err != nil {
	log.Println(err)	
} else {
    fmt.Println(result)
}
```
