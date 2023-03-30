package run

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/promptc/promptc-go/cli/oper/cfg"
	"github.com/promptc/promptc-go/driver"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"strings"
)

func BlankHandler(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: promptc-cli blank [prompt]")
		return
	}
	providerDriver := driver.GetDefaultDriver(cfg.GetCfg().GetToken("openai"))
	toSend := models.PromptToSend{
		Items: []models.PromptItem{
			{
				Content: strings.Join(args, ""),
			},
		},
		Model: "gpt-3.5-turbo",
		Extra: nil,
	}
	runPrompt(providerDriver, toSend)
}

func runPrompt(providerDriver interfaces.ProviderDriver, toSend models.PromptToSend) {
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
