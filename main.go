package main

import (
  "log"
  "flag"
  "path/filepath"
  "os"
  "fmt"

  "github.com/BurntSushi/xgbutil"
  "github.com/BurntSushi/xgbutil/xevent"
  "github.com/BurntSushi/xgbutil/keybind"
)

var X *xgbutil.XUtil

func init() {
  flag.Parse()

  var err error
  X, err = xgbutil.NewConn()
  if err != nil {
    log.Fatalf("error connect to X: %v\n", err)
  }
  keybind.Initialize(X)
}

func main() {
  set := NewSet(loadFiles())
  if len(set.files) == 0 {
    return
  }
  fmt.Printf("files loaded\n")

  window := NewWindow()
  if window == nil {
    return
  }
  window.set = set

  file := set.Get(FORWARD)
  if file == nil {
    return
  }
  window.DrawImage(file.ximage)

  xevent.Main(X)
}

func loadFiles() []*File {
  root := flag.Arg(0)
  if root == "" {
    root = "."
  }
  var files []*File
  filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
    if !f.IsDir() {
      files = append(files, &File{path: path})
    }
    return nil
  })
  return files
}
