package main

import (
	"github.com/absozero/vcheck/cmd/vcheck"
  _ "github.com/absozero/vcheck/cmd/vcheck/version"
  _ "github.com/absozero/vcheck/cmd/vcheck/get"
  _ "github.com/absozero/vcheck/cmd/vcheck/serve"

  )
  
  func main() {
	cmd.Execute()
  }