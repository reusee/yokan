package main

import (
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
)

func (self *Window) bindKeys() {
  // navigate
  keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    file := self.set.Next()
    if file != nil {
      self.DrawImage(file.ximage)
    }
  }).Connect(X, self.Id, "Space", false)
  keybind.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
    file := self.set.Prev()
    if file != nil {
      self.DrawImage(file.ximage)
    }
  }).Connect(X, self.Id, "c", false)
}
