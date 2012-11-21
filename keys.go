package main

import (
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"time"
)

const (
  CMD_FREQ = time.Millisecond * 200
  MOVE_STEP = 100
)

func (self *Window) bindKeys() {
  // navigate

  self.freqLimitedFunc(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    file := self.set.Next()
    if file != nil {
      self.currentFile = file
      file.Draw()
    }
  }).Connect(X, self.Id, "Space", false)

  self.freqLimitedFunc(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    file := self.set.Prev()
    if file != nil {
      self.currentFile = file
      file.Draw()
    }
  }).Connect(X, self.Id, "c", false)

  // image operations

  self.freqLimitedFunc(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    self.currentFile.ZoomIn()
    self.currentFile.Draw()
  }).Connect(X, self.Id, "z", false)

  self.freqLimitedFunc(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    self.currentFile.ZoomOut()
    self.currentFile.Draw()
  }).Connect(X, self.Id, "x", false)

  keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    self.currentFile.y += MOVE_STEP
    self.currentFile.Draw()
  }).Connect(X, self.Id, "w", false)

  keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    self.currentFile.y -= MOVE_STEP
    self.currentFile.Draw()
  }).Connect(X, self.Id, "s", false)

  keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    self.currentFile.x += MOVE_STEP
    self.currentFile.Draw()
  }).Connect(X, self.Id, "a", false)

  keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    self.currentFile.x -= MOVE_STEP
    self.currentFile.Draw()
  }).Connect(X, self.Id, "d", false)

}

func (self *Window) freqLimitedFunc(fun func(*xgbutil.XUtil, xevent.KeyPressEvent)) keybind.KeyPressFun {
  return keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    if self.forbidCommand { return }
    fun(X, ev)
    self.forbidCommand = true
    go func() {
      <-time.After(CMD_FREQ)
      self.forbidCommand = false
    }()
  })
}
