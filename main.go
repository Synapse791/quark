package main

import(
  "flag"
  "github.com/Synapse791/quark/envext"
)

type Flags struct {
  Delimeter  string
  Prefix     string
  Quiet      bool
  SkipChecks bool
}

var config Flags

func init() {
  flag.StringVar(&config.Delimeter, "d", ":", "delimeter used to split environment variable values")
  flag.StringVar(&config.Prefix, "p", "QUARK_", "prefix to search for in environment variables")
  flag.BoolVar(&config.Quiet, "q", false, "supress stdout messages")
  flag.BoolVar(&config.SkipChecks, "s", false, "skip checks for keys in file")
}

func main() {
  flag.Parse()

  InfoLine("starting quark...")

  if config.Prefix == "" {
    PrintUsage()
    ErrorLine("prefix cannot be empty")
  } else if config.Delimeter == "" {
    PrintUsage()
    ErrorLine("delimeter cannot be empty")
  }

  if config.SkipChecks {
    InfoLine("skipping key checks")
  }

  InfoLine("searching for environment variables starting with '%s'...", config.Prefix)

  extractor := envext.New(config.Prefix)
  extractor.Run()

  if len(extractor.EnvVars) == 0 {
    ErrorLine("no environment varables found")
  } else if len(extractor.EnvVars) == 1 {
    InfoLine("found %d environment variable!", len(extractor.EnvVars))
  } else {
    InfoLine("found %d environment variables!", len(extractor.EnvVars))
  }

  replacements := map[string]map[string]string{}

  for _, env := range extractor.EnvVars {
    if ! env.Search(config.Delimeter) {
      ErrorLine("malformed environment variable: delimeter '%s' not found in %s", config.Delimeter, env.Raw)
    }

    key := env.Key
    parts := env.SplitValue(config.Delimeter)

    if len(parts) != 2 {
      ErrorLine("malformed environment variable: %s", config.Delimeter, env.Raw)
    }

    filePath := parts[0]
    value := parts[1]

    if ! FileExists(filePath) {
      ErrorLine("file '%s' does not exist", filePath)
    }

    if ! config.SkipChecks {
      if ! CheckForKey(filePath, key) {
        ErrorLine("key '%s' not found in '%s'", key, filePath)
      }
    }

    if _, ok := replacements[filePath]; !ok {
      replacements[filePath] = map[string]string{}
    }

    replacements[filePath][key] = value
  }

  for filePath, rs := range replacements {
    if ! FileExists(filePath) {
      ErrorLine("file '%s' not found", filePath)
    }

    InfoLine("processing %d unique keys in '%s'...", len(rs), filePath)

    if rErr := ReplaceInFile(filePath, rs); rErr != nil {
      ErrorLine(rErr.Error())
    }

  }

  SuccessLine("successfully processed %d environment variable in %d files!", len(extractor.EnvVars), len(replacements))
}
