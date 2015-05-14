package helpers

import (
  "strings"
)

func IsGoGettable(path string) bool {
  return !strings.Contains(path, ":")
}
