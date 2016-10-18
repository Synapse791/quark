package envext

import "strings"

type EnvVar struct {
  Raw   string
  Key   string
  Value string
}

func (e EnvVar) SplitValue(delimeter string) []string {
  parts := strings.Split(e.Value, delimeter)
  return parts
}

func (e EnvVar) Search(search string) bool {
  return strings.Contains(e.Value, search)
}
