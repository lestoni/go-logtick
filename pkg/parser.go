package logtick

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Type to Store parsed git log tokens
type LogTokens []Token

// Meta represents embedded info for JSONOutput
type Meta struct {
	Path    string `json:"path"`    // File Path
	Changes string `json:"changes"` // File Changes
	Raw     string `json:"raw"`     // Raw Diff Output
}

// JSONOutput represents JSON output for LogToken
type JSONOutput struct {
	Commit      string `json:"commit"`       // Commit Hash
	Author      string `json:"author"`       // Author Name
	AuthorDate  string `json:"author_date"`  // Authored Date
	Committer   string `json:"committer"`    // Committer Name
	CommitDate  string `json:"commit_date"`  // Committed Date
	Message     string `json:"message"`      // Commit Message
	ChangeStats string `json:"change_stats"` // Changed Files
	Files       []Meta `json:"files"`        // The Changed files
	Diff        []Meta `json:"diff"`         // Commit changes
}

// Parses a raw git log string
func Parse(log string) (LogTokens, error) {
	var tokens LogTokens

	slices := GLOBAL_SPLITTER.Split(log, -1)

	if len(slices) < 3 {
		return tokens, fmt.Errorf("Untrue Git log")
	}

	infoTokens := COMMON_SPLITTER.Split(slices[0], -1)
	info := PackInfo(infoTokens)
	tokens = append(tokens, info)

	commitTokens := COMMON_SPLITTER.Split(slices[1], -1)
	commit := PackCommit(commitTokens)
	tokens = append(tokens, commit)

	diffTokens := COMMON_SPLITTER.Split(slices[2], -1)
	diff := PackDiff(diffTokens)
	tokens = append(tokens, diff)

	return tokens, nil
}

// Encode logtokens to JSON output
func (lt LogTokens) ToJSON() (string, error) {
	var output []byte
	jsonOutput := &JSONOutput{}
	var paths []string

	for _, tkn := range lt {
		if tkn.Type == "INFO" {
			jsonOutput.setInfo(tkn)
		}

		if tkn.Type == "COMMIT" {
			paths = jsonOutput.setCommit(tkn)
		}

		if tkn.Type == "DIFF" {
			jsonOutput.setDiff(tkn, paths)
		}
	}

	output, err := json.Marshal(jsonOutput)
	if err != nil {
		return "", err
	}

	out := fmt.Sprintf("%s", output)

	return out, nil
}

func (j *JSONOutput) setInfo(tkn Token) {
	j.Commit = tkn.Index[0]
	j.Author = tkn.Index[1]
	j.AuthorDate = tkn.Index[2]
	j.Committer = tkn.Index[3]
	j.CommitDate = tkn.Index[4]
}

func (j *JSONOutput) setCommit(tkn Token) []string {
	var paths []string
	var filesMeta []Meta

	j.Message = strings.Join(tkn.Index, "\n")

	for index, change := range tkn.Changes {
		changeStr := strings.Join(change, "")

		if index == len(tkn.Changes)-1 {
			j.ChangeStats = changeStr

		} else {
			meta := CHANGE_SPLITTER.Split(changeStr, -1)
			changeMeta := Meta{
				Path:    meta[0],
				Changes: meta[1],
			}

			paths = append(paths, meta[0])

			filesMeta = append(filesMeta, changeMeta)
		}
	}

	j.Files = append(j.Files, filesMeta...)

	return paths
}

func (j *JSONOutput) setDiff(tkn Token, paths []string) {
	for index, path := range paths {
		diffMeta := Meta{
			Path: path,
			Raw:  strings.Join(tkn.Diff[index], "\n"),
		}

		j.Diff = append(j.Diff, diffMeta)
	}
}
