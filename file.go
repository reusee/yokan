package main

import (
  "flag"
  "path/filepath"
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
  originXimage *xgraphics.Image
  loadMutex sync.Mutex
  scale float64
  window *Window
  x int
  y int
}

func loadFiles(window *Window) []*File {
  root := flag.Arg(0)
  if root == "" {
    root = "."
  }
  var files []*File
  filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
    if !f.IsDir() {
      files = append(files, NewFile(path, window))
    }
    return nil
  })
  return files
}

func NewFile(path string, window *Window) *File {
  return &File{
    path: path,
    scale: 1.0,
    window: window,
    x: 0,
    y: 0,
  }
}

func (self *File) load() {
  self.loadMutex.Lock()
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
    self.ximage.CreatePixmap()
    self.ximage.XDraw()
    self.originXimage = xgraphics.NewConvert(X, image)
    fmt.Printf("loaded %s\n", self.path)
  })
  self.loadMutex.Unlock()
}

func (self *File) ZoomIn() {
  self.scale *= 1.5
  self.Scale(self.scale)
}

func (self *File) ZoomOut() {
  self.scale *= 0.5
  self.Scale(self.scale)
}

func (self *File) Scale(ratio float64) {
  originRect := self.originXimage.Rect
  newWidth := int(float64(originRect.Dx()) * ratio)
  newHeight := int(float64(originRect.Dy()) * ratio)
  newImage := Resize(self.originXimage, originRect, newWidth, newHeight)
  self.ximage.Destroy()
  self.ximage = xgraphics.NewConvert(X, newImage)
  self.ximage.CreatePixmap()
  self.ximage.XDraw()
  self.x, self.y = 0, 0
  self.window.Clear(0, 0, 0, 0)
}

func (self *File) Move(dx int, dy int) {
  if dx > 0 {
    self.window.Clear(self.x, self.y, dx, self.ximage.Rect.Dy())
  } else if dx < 0 {
    self.window.Clear(self.x + self.ximage.Rect.Dx() + dx, self.y, -dx, self.ximage.Rect.Dy())
  }
  if dy > 0 {
    self.window.Clear(self.x, self.y, self.ximage.Rect.Dx(), dy)
  } else if dy < 0 {
    self.window.Clear(self.x, self.y + self.ximage.Rect.Dy() + dy, self.ximage.Rect.Dx(), -dy)
  }
  self.x += dx
  self.y += dy
}

func (self *File) Draw() {
  self.ximage.XExpPaint(self.window.Id, self.x, self.y)
}
