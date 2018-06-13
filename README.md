# rplpa [![Build Status](https://travis-ci.org/Mempler/rplpa.svg?branch=master)](https://travis-ci.org/Mempler/rplpa)

rplpa is an Replay Parser for golang. \
IDK for what you'll use it, but i'll use it for an anticheat.

Example:
```go
package main

import (
  "ioutil"

  "github.com/Mempler/rplpa"
)

func main() {
  buf, err := ioutil.ReadFile("path/to/replay.osr")
  if err != nil {
    panic(err)
  }
  replay, err := ParseReplay(b)
  if err != nil {
    panic(err)
  }
}
```

or compressed replays

```go
package main

import (
  "ioutil"

  "github.com/Mempler/rplpa"
)

func main() {
  RawData := []byte{} // IDK, some data.
  replaydata, err := ParseCompressed(RawData)
  if err != nil {
    panic(err)
  }
}
```
