package main

import (
  "fmt"
  "flag"
  "os"
  "io/ioutil"
  "strings"
)

func InfoLine(message string, args... interface{}) {
  if ! config.Quiet {
    formattedMessage := fmt.Sprintf(message, args...)
    fmt.Fprintf(os.Stdout, "\033[34m+\033[00m  %s\n", formattedMessage)
  }
}

func SuccessLine(message string, args... interface{}) {
  if ! config.Quiet {
    formattedMessage := fmt.Sprintf(message, args...)
    fmt.Fprintf(os.Stdout, "\033[32m+\033[00m  \033[32m%s\033[00m\n", formattedMessage)
  }
}

func PrintUsage() {
  flag.Usage()
}

func PrintVersion() {
  fmt.Printf("quark version %s\n", VERSION)
  os.Exit(0)
}

func WarningLine(message string, args... interface{}) {
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Fprintf(os.Stderr, "\033[33m+\033[00m  \033[33mWARNING: %s\033[00m\n", formattedMessage)
}

func ErrorLine(message string, args... interface{}) {
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Fprintf(os.Stderr, "\033[31m!\033[00m  \033[31mERROR: %s\033[00m\n", formattedMessage)
  os.Exit(1)
}

func FileExists(filePath string) bool {
  if _, err := os.Stat(filePath); os.IsNotExist(err) {
    return false
  } else {
    return true
  }
}

func CheckForKey(filePath string, key string) bool {
  rawContents, err := ioutil.ReadFile(filePath)
  if err != nil {
    return false
  }

  contents := string(rawContents)

  if ! strings.Contains(contents, key) {
    return false
  }

  return true
}

func ReplaceInFile(filePath string, pairs map[string]string) error {
  rawContents, err := ioutil.ReadFile(filePath)
  if err != nil {
    return err
  }

  contents := string(rawContents)

  for key, value := range pairs {
    InfoLine("replacing '%s' in '%s' with '%s'...", key, filePath, value)
    contents = strings.Replace(contents, key, value, -1)
  }

  err = ioutil.WriteFile(filePath, []byte(contents), 0644)
  if err != nil {
    return err
  }

  return nil
}
