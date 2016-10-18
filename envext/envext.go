package envext

import (
  "os"
  "strings"
)

type EnvExtractor struct {
  EnvVars    []EnvVar
  Prefix     string
  SplitDelim string
}

func New(prefix string) *EnvExtractor {
  return &EnvExtractor{
    Prefix: prefix,
    SplitDelim: "",
  }
}

func (e *EnvExtractor) Run() error {
  var rawEnvs []string
  for _, rawEnv := range os.Environ() {
    if strings.HasPrefix(rawEnv, e.Prefix) {
      rawEnvs = append(rawEnvs, rawEnv)
    }
  }

  for _, rawEnv := range rawEnvs {
    env := EnvVar{
      Raw: rawEnv,
    }
    parts := strings.Split(rawEnv, "=")
    env.Key = parts[0]

    if len(parts) > 2 {
      parts = append(parts[:0], parts[1:]...)
      env.Value = strings.Join(parts, "=")
    } else {
      env.Value = parts[1]
    }

    e.EnvVars = append(e.EnvVars, env)
  }

  return nil
}
