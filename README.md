# spellchecker

## A simple spell checker written in Go


`go build`  
`./spellcheker`

```
$ curl -X POST localhost:8000/check -F "text=mistakus happon everywherre"
["mistakes","happen","everywhere"]
```
