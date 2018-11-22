# logtick

[![GoDoc](https://godoc.org/github.com/lestoni/go-logtick?status.svg)](https://godoc.org/github.com/lestoni/go-logtick)   [![Go Report Card](https://goreportcard.com/badge/github.com/lestoni/go-logtick)](https://goreportcard.com/report/github.com/lestoni/go-logtick)

Parse a `git log -1 -p --stat --pretty=fuller` .I Find it the log output simple and detailed enough.

Originally created for [go-logtick-http](https://github.com/lestoni/go-logtick-http)


## Install

```sh
  $ go get github.com/lestoni/go-logtick
```

## Usage


```
  func main(){
    log, err := ioutil.ReadFile("testdata/git.log")
    if err != nil {
      panic(err)
    }

    content := fmt.Sprintf("%s", log)

    output, err := logtick.Parse(content)
    if err != nil {
      panic(err)
    }

    out, err := output.ToJSON()
    if err != nil {
      panic(err)
    }

    fmt.Printf("%+v", output)
  }
```

## Testing

Test with Code Coverage

```
 $ go test -cover
```