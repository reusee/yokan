package main

import (
  "log"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

type Window struct {
  *xwindow.Window
  set *Set
}

func NewWindow() *Window {
  win, err := xwindow.Generate(X)
  if err != nil {
    log.Fatal("cannot generate window %v\n", err)
    return nil
  }

  width, height := 800, 600
  win.Create(X.RootWin(), 0, 0, width, height, 0)

  win.WMGracefulClose(func(w *xwindow.Window) {
    xevent.Detach(w.X, w.Id)
    keybind.Detach(w.X, w.Id)
    mousebind.Detach(w.X, w.Id)
    w.Destroy()
    xevent.Quit(w.X)
  })

  icccm.WmStateSet(X, win.Id, &icccm.WmState{
    State: icccm.StateNormal,
  })

  win.Listen(xproto.EventMaskKeyPress)

  self := &Window{win, nil}
  self.bindKeys()

  return self
}

func (self *Window) SetName(name string) {
  ewmh.WmNameSet(X, self.Id, name)
}

func (self *Window) DrawImage(image *xgraphics.Image) {
  image.XSurfaceSet(self.Id)
  image.XDraw()
  image.XPaint(self.Id)
  self.Map()
}
