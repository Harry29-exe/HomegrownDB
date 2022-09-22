package tx

import "strings"

type QueryCtx struct {
	QueryTokens []string
}

func (c *QueryCtx) Reconstruct(startToken, endToken uint32) string {
	strBuilder := strings.Builder{}
	for i := startToken; i < endToken; i++ {
		strBuilder.WriteString(c.QueryTokens[i])
	}
	return strBuilder.String()
}
