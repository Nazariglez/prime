// Created by nazarigonzalez on 6/1/17.

package gor

import (
  "runtime"
  "bytes"
  "strconv"
)

//todo remove this package, it's just for debug purposes

func ID() uint64 {
  b := make([]byte, 64)
  b = b[:runtime.Stack(b, false)]
  b = bytes.TrimPrefix(b, []byte("goroutine "))
  b = b[:bytes.IndexByte(b, ' ')]
  n, _ := strconv.ParseUint(string(b), 10, 64)
  return n
}
