package main

import (
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/cli/oper/compile"
	"github.com/promptc/promptc-go/cli/oper/help"
	"github.com/promptc/promptc-go/cli/oper/run"
	"github.com/promptc/promptc-go/cli/oper/shared"
	"os"
)

func main() {
	//fmt.Println(shared.GetUserFolder())
	shared.InitPath()
	args := os.Args[1:]
	var handler func([]string)
	if len(args) == 0 {
		help.HelpHandler(args)
		return
	}
	verb := args[0]
	keepVerb := false
	switch verb {
	case "help":
		handler = help.HelpHandler
	case "set":
		handler = cfg.SetHandler
	case "show":
		handler = cfg.ShowHandler
	case "run":
		handler = run.RunHandler
	case "compile":
		handler = compile.CompileHandler
	case "analyse":
		handler = compile.AnalyseHandler
	case "blank":
		handler = run.BlankHandler
	default:
		handler = run.SimpleRunHandler
		keepVerb = true
	}
	if !keepVerb {
		args = args[1:]
	}
	if handler == nil {
		help.HelpHandler(args)
		return
	}

	handler(args)
}
