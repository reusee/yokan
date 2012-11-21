package main

const (
  FORWARD = iota
  BACKWARD
)

type Set struct {
  files []*File
  index int
  invalidFileCount int
}

func NewSet(files []*File) *Set {
  self := &Set{
    files: files,
    index: 0,
  }
  return self
}

func (self *Set) Get(direction int) *File {
  self.files[self.index].load()
  for self.files[self.index].invalid {
    if direction == FORWARD {
      self.nextIndex()
    } else {
      self.prevIndex()
    }
    self.files[self.index].load()
  }
  return self.files[self.index]
}

func (self *Set) Next() *File {
  self.nextIndex()
  return self.Get(FORWARD)
}

func (self *Set) Prev() *File {
  self.prevIndex()
  return self.Get(BACKWARD)
}

func (self *Set) nextIndex() {
  if self.index == len(self.files) - 1 {
    self.index = 0
  } else {
    self.index++
  }
}

func (self *Set) prevIndex() {
  if self.index == 0 {
    self.index = len(self.files) - 1
  } else {
    self.index--
  }
}
