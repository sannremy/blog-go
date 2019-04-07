package main

import (
  "testing"
  "os"
  "net/http/httptest"
)

func TestIndexHandler(t *testing.T) {
  req := httptest.NewRequest("GET", "/", nil)
  w := httptest.NewRecorder()
  IndexHandler(w, req)

  resp := w.Result()

  if resp.StatusCode != 200 {
    t.Errorf("Status code was incorrect, got: %d, want: %d.", resp.StatusCode, 200)
  }
}

func TestMain(m *testing.M) {
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}
