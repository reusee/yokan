package main

import (
  "log"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/icccm"
)

type Window struct {
  *xwindow.Window
  set *Set
  forbidCommand bool
  currentFile *File
}

func NewWindow() *Window {
  win, err := xwindow.Generate(X)
  if err != nil {
    log.Fatal("cannot generate window %v\n", err)
    return nil
  }

  width, height := 800, 600
  win.Create(X.RootWin(), 0, 0, width, height, xproto.CwBackPixel, 0x0)

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
  win.Clear(0, 0, 0, 0)
  win.Map()

  self := &Window{
    win,
    nil,
    false,
    nil,
  }
  self.bindKeys()

  return self
}
