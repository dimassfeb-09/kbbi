# kbbi

Import:
```go 
import "github.com/dimassfeb-09/kbbi.git/kbbi"
```

How to Use:
1. Use ```kbbi```
2. User ```.Search("keyword here")```
```go
if result, err := kbbi.Search("keyword here"); err != nil {
	log.Println(err)	
} else {
    fmt.Println(result)
}
```