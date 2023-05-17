package shared

import (
	"fmt"
	"strings"

	"github.com/KevinZonda/GoX/pkg/console"
	"github.com/promptc/promptc-go/driver/interfaces"
	"github.com/promptc/promptc-go/driver/models"
)

func DefaultResponseBefore(id int) {
	console.Blue.AsForeground().WriteLine("Response #%d:", id)
}

// RunPrompt runs a prompt in the given provider driver and sends the provided prompt to it.
// It takes in a ProviderDriver, a PromptToSend, and a function to execute before the response.
// It returns a slice of strings.
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
	}

	streamer := providerDriver.ToStream()
	resp, err := streamer.GetStreamResponse(toSend)
	if err != nil {
		panic(err)
	}
	defer resp.Close()

	var sb strings.Builder
	if responseBefore != nil {
		responseBefore(0)
	}

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
