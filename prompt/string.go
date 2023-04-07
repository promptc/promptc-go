package prompt

import (
	"github.com/promptc/promptc-go/prompt/provider"
	"github.com/promptc/promptc-go/utils"
	"github.com/promptc/promptc-go/variable"
	"strings"
)

func (f *File) Formatted() string {
	nf := File{
		FileInfo: f.FileInfo,
	}

	vars := make(map[string]string)
	for k, v := range f.VarConstraint {
		vars[k] = variable.ToPromptcString(v)
	}
	nf.Vars = vars
	for _, block := range f.ParsedPrompt {
		nf.Prompts = append(nf.Prompts, block.Formatted(f.RefProvider)...)
	}
	return utils.HjsonNoBrace(nf)

}

func TxTokens(tokens []BlockToken, varMoveTx, vatToLiterTx map[string]string) []BlockToken {
	newTokens := make([]BlockToken, 0)
	for _, token := range tokens {
		if token.Kind != BlockTokenKindVar {
			newTokens = append(newTokens, token)
			continue
		}
		if newV, ok := varMoveTx[token.Text]; ok {
			newTokens = append(newTokens, BlockToken{
				Kind: BlockTokenKindVar,
				Text: newV,
			})
			continue
		}
		if newV, ok := vatToLiterTx[token.Text]; ok {
			newTokens = append(newTokens, BlockToken{
				Kind: BlockTokenKindLiter,
				Text: newV,
			})
			continue
		}
		newTokens = append(newTokens, token)
	}
	return newTokens
}

func (r *ReferBlock) Formatted(prov provider.Privider) []string {
	varMoveTx := make(map[string]string)
	vatToLiterTx := make(map[string]string)
	for k, v := range r.VarMap {
		if strings.HasPrefix(v, "$") {
			newV := v[1:]
			if strings.HasPrefix(newV, "$") {
				vatToLiterTx[k] = newV
			} else {
				varMoveTx[k] = newV
			}
		} else {
			vatToLiterTx[k] = v
		}
		//fmt.Println(k, "->", v)
	}
	promptTxt, err := r.RefProvider.GetPrompt(r.RefTo)
	if err != nil {
		return nil
	}
	prompt := ParseUnstructuredFile(promptTxt)
	prompt.RefProvider = r.RefProvider
	for _, block := range prompt.ParsedPrompt {
		block.Tokens = TxTokens(block.Tokens, varMoveTx, vatToLiterTx)
	}
	var result []string
	for _, block := range prompt.ParsedPrompt {
		result = append(result, block.Formatted(prov)...)
	}
	return result
}

func (p *ParsedBlock) Formatted(prov provider.Privider) []string {
	if p.Type() == RefBlock {
		ref, err := p.ToReferBlock(prov)
		if err != nil {
			return nil
		}
		return ref.Formatted(prov)
	}
	meta := "{}"
	if len(p.Extra) > 0 {
		meta = utils.HjsonNoIdent(p.Extra)
	}
	sb := strings.Builder{}
	for _, token := range p.Tokens {
		switch token.Kind {
		case BlockTokenKindVar:
			sb.WriteString("{")
			sb.WriteString(token.Text)
			sb.WriteString("}")
		case BlockTokenKindLiter:
			replaced := strings.ReplaceAll(token.Text, "{", "{{")
			replaced = strings.ReplaceAll(replaced, "}", "}}")
			sb.WriteString(replaced)
		case BlockTokenKindReservedQuota:
			sb.WriteString("{%Q%}")
		case BlockTokenKindScript:
			sb.WriteString("{%\n")
			sb.WriteString(token.Text)
			sb.WriteString("\n%}")
		}
	}
	return []string{meta + "\n" + sb.String()}
}