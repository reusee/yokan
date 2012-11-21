package main

import (
  "sync"
  "image"
  "os"
  "fmt"
  _ "image/gif"
  _ "image/png"
  _ "image/jpeg"
  "github.com/BurntSushi/xgbutil/xgraphics"
)

type File struct {
  path string
  loadOnce sync.Once
  invalid bool
  ximage *xgraphics.Image
  mutex sync.Mutex
}

func (self *File) load() {
  self.mutex.Lock()
  self.loadOnce.Do(func() {
    f, err := os.Open(self.path)
    if err != nil {
      self.invalid = true
      return
    }
    image, _, err := image.Decode(f)
    if err != nil {
      self.invalid = true
      return
    }
    self.ximage = xgraphics.NewConvert(X, image)
    fmt.Printf("loaded %s\n", self.path)
  })
  self.mutex.Unlock()
}
