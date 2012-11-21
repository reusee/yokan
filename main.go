package main

import (
  "flag"
  "log"
  "math/rand"
  "time"

  "github.com/BurntSushi/xgbutil"
  "github.com/BurntSushi/xgbutil/xevent"
  "github.com/BurntSushi/xgbutil/keybind"
)

var X *xgbutil.XUtil

func init() {
  flag.Parse()

  rand.Seed(time.Now().UnixNano())

  var err error
  X, err = xgbutil.NewConn()
  if err != nil {
    log.Fatalf("error connect to X: %v\n", err)
  }
  keybind.Initialize(X)
}

func main() {
  window := NewWindow()
  if window == nil {
    return
  }

  set := NewSet(loadFiles(window))
  if len(set.files) == 0 {
    return
  }
  window.set = set

  file := set.Get(FORWARD)
  if file == nil {
    return
  }
  window.currentFile = file
  file.Draw()

  xevent.Main(X)
}
