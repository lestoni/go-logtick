package logtick

import (
	"strings"
)

// Token Represents Structure for capturing tokens
type Token struct {
	Index   []string   `json:"index"`   // Commit Generic segments
	Changes [][]string `json:"changes"` // Commit Changed files segments
	Diff    [][]string `json:"diff"`    // Commit Diff Segments
	Type    string     `json:"type"`    // Token type
}

func (t *Token) addIndex(token string) {
	t.Index = append(t.Index, token)
}

func (t *Token) addChange(token []string) {
	t.Changes = append(t.Changes, token)
}

// Tokenize git log header info and pack it to Token
func PackInfo(slices []string) Token {
	token := Token{
		Type: "INFO",
	}

	for index, slice := range slices {
		if index == 0 {
			matches := COMMIT_HASH_SPLITTER.Split(slice, -1)
			token.addIndex(strings.TrimSpace(matches[1]))
		}

		if index == 1 {
			matches := AUTHOR_SPLITTER.Split(slice, -1)
			token.addIndex(strings.TrimSpace(matches[1]))
		}

		if index == 2 {
			matches := AUTHOR_DATE_SPLITTER.Split(slice, -1)
			token.addIndex(strings.TrimSpace(matches[1]))
		}

		if index == 3 {
			matches := COMMIT_SPLITTER.Split(slice, -1)
			token.addIndex(strings.TrimSpace(matches[1]))
		}

		if index == 4 {
			matches := COMMIT_DATE_SPLITTER.Split(slice, -1)
			token.addIndex(strings.TrimSpace(matches[1]))
		}
	}

	return token
}

// Tokenize git log Changed Files info and pack it to Token
func PackCommit(slices []string) Token {
	token := Token{
		Type: "COMMIT",
	}

	for index, slice := range slices {
		slice = strings.TrimSpace(slice)

		if index == 0 {
			token.addIndex(slice)

		} else if strings.Index(slice, "---") == 0 {
			continue

		} else {
			if index == len(slices) {
				token.addIndex(slice)
			} else {
				changes := COMMIT_SEPARATOR_SPLITTER.Split(slice, -1)
				token.addChange(changes)
			}
		}

	}

	return token
}

// Tokenize git log diff info and pack it to Token
func PackDiff(slices []string) Token {
	token := Token{
		Type: "DIFF",
	}
	var cache []string
	for index, slice := range slices {
		if strings.Index(slice, "diff --git") == 0 {
			if index != 0 {
				token.Diff = append(token.Diff, cache)
				cache = make([]string, 0)
			}

		}

		cache = append(cache, slice)

	}

	if len(cache) != 0 {
		token.Diff = append(token.Diff, cache)
	}

	return token
}
