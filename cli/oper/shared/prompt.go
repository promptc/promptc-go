package shared

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
	"strings"
)

func DefaultResponseBefore(id int) {
	console.Blue.AsForeground().WriteLine("Response #%d:", id)
}

func RunPrompt(providerDriver interfaces.ProviderDriver, toSend models.PromptToSend, responseBefore func(id int)) []string {
	if !providerDriver.StreamAvailable() {
		resp, err := providerDriver.GetResponse(toSend)
		if err != nil {
			panic(err)
		}
		for i, r := range resp {
			if responseBefore != nil {
				responseBefore(i)
			}
			fmt.Println(r)
		}
		return resp
	} else {
		streamer := providerDriver.ToStream()
		resp, err := streamer.GetStreamResponse(toSend)
		sb := strings.Builder{}
		if responseBefore != nil {
			responseBefore(0)
		}
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
			sb.WriteString(r[0])
			fmt.Print(r[0])
		}
		return []string{sb.String()}
	}
}
