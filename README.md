# rplpa

rplpa is an osu! replay parser/writer for golang.

## Examples:

## Reading the replay

```go
package main

import (
  "ioutil"

  "github.com/wieku/rplpa"
)

func main() {
  b, err := ioutil.ReadFile("path/to/replay.osr")
  if err != nil {
    panic(err)
  }
  replay, err := ParseReplay(b)
  if err != nil {
    panic(err)
  }
}
```

## Reading compressed input data

```go
package main

import (
  "ioutil"

  "github.com/wieku/rplpa"
)

func main() {
  RawData := []byte{} // Compressed LZMA stream of input events in delta1|x1|y1|keys1,delta2|x2|y2|keys2 format
  replaydata, err := ParseCompressed(RawData)
  if err != nil {
    panic(err)
  }
}
```
