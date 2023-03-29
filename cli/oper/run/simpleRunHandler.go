package run

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/models"
	"github.com/promptc/promptc-go/prompt"
	"strings"
)

func SimpleRunHandler(args []string) {
	if len(args) < 1 {
		panic("Usage: promptc-cli [prompt-file] [input?]")
	}
	path := args[0]
	input := ""
	if len(args) == 2 {
		input = args[1]
	}
	txt, structured, err := fetchFile(path)
	if err != nil {
		panic(err)
	}
	var file *prompt.File
	if structured {
		file = prompt.ParseFile(txt)
	} else {
		file = prompt.ParseUnstructuredFile(txt)
	}
	if len(file.Vars) > 1 {
		fmt.Println("Required following vars:")
		for k, v := range file.Vars {
			fmt.Println("-", k, "->", v)
		}
		panic("Too many vars")
	}
	varMap := make(map[string]string)
	if len(file.Vars) == 1 {
		for k, _ := range file.Vars {
			varMap[k] = input
		}
	}
	printSep()
	printInfo(file.FileInfo)
	printSep()
	compiled := file.Compile(varMap)
	fmt.Println("Compiled To: ")
	for _, c := range compiled.Prompts {
		fmt.Println(c.Prompt)
	}

	provider := strings.ToLower(strings.TrimSpace(file.Conf.Provider))
	model := strings.ToLower(strings.TrimSpace(file.Conf.Model))
	providerDriver, err := driver.GetDriver(provider, model, cfg.GetCfg().GetToken(provider))

	var items []models.PromptItem
	for _, c := range compiled.Prompts {
		items = append(items, convCompiledToSend(c))
	}
	toSend := models.PromptToSend{
		Items: items,
		Model: model,
		Extra: nil,
	}
	printSep()
	if !providerDriver.StreamAvailable() {
		resp, err := providerDriver.GetResponse(toSend)
		if err != nil {
			panic(err)
		}
		for i, r := range resp {
			console.Blue.AsForeground().WriteLine("Response #%d:", i)
			fmt.Println(r)
		}
	} else {
		console.Blue.AsForeground().WriteLine("Response #%d:", 0)
		streamer := providerDriver.ToStream()
		resp, err := streamer.GetStreamResponse(toSend)
		if err != nil {
			panic(err)
		}
		defer resp.Close()
		for {
			r, err, eof := resp.Receive()
			if eof {
				fmt.Println()
				break
			}
			if err != nil {
				panic(err)
			}
			lenOfR := len(r)
			if lenOfR == 0 {
				continue
			}
			fmt.Print(r[0])
		}
	}

}

func printSep() {
	fmt.Println("================")
}

func printInfo(f prompt.FileInfo) {
	sb := strings.Builder{}
	if f.Project != "" {
		sb.WriteString("Project: ")
		sb.WriteString(f.Project)
		sb.WriteString("\n")
	}
	if f.Version != "" {
		sb.WriteString("Version: ")
		sb.WriteString(f.Version)
		sb.WriteString("\n")
	}
	if f.Author != "" {
		sb.WriteString("Author: ")
		sb.WriteString(f.Author)
		sb.WriteString("\n")
	}
	if f.License != "" {
		sb.WriteString("License: ")
		sb.WriteString(f.License)
		sb.WriteString("\n")
	}
	if sb.Len() > 0 {
		fmt.Println(sb.String())
	} else {
		fmt.Println("No info provided by prompt file")
	}
}

func fetchFile(file string) (txt string, structured bool, err error) {
	structured = true
	txt, err = iox.ReadAllText(file)
	if err == nil {
		return
	}
	txt, err = iox.ReadAllText(file + ".promptc")
	if err == nil {
		return
	}
	txt, err = iox.ReadAllText(file + ".ptc")
	if err == nil {
		return
	}
	structured = false
	txt, err = iox.ReadAllText(file + ".ptc")
	return
}
