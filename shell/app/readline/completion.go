package readline

import (
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

type Completion struct {
	trie   utils.Trie
	enable bool
}

func NewCompletion(enable bool) *Completion {
	return &Completion{
		enable: enable,
		trie:   *utils.NewTrie(),
	}
}

func (c *Completion) BulkInster(words []string) {
	for _, str := range words {
		c.trie.Insert(str)
	}
}
func (c *Completion) Insert(str string) {
	c.trie.Insert(str)
}
func (c *Completion) GetSuggestions(prefix string) []string {
	return c.trie.Suggest(prefix)
}
