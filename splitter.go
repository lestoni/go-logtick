package logtick

import (
	"regexp"
)

var (
	GLOBAL_SPLITTER = regexp.MustCompile("(?m)\\s+\n")

	COMMON_SPLITTER = regexp.MustCompile("\n+")

	COMMIT_HASH_SPLITTER = regexp.MustCompile("commit\\s+")

	AUTHOR_SPLITTER = regexp.MustCompile("Author:\\s+")

	AUTHOR_DATE_SPLITTER = regexp.MustCompile("AuthorDate:\\s+")

	COMMIT_SPLITTER = regexp.MustCompile("Commit:\\s+")

	COMMIT_DATE_SPLITTER = regexp.MustCompile("CommitDate:\\s+")

	COMMIT_SEPARATOR_SPLITTER = regexp.MustCompile("CommitDate:\\s+")

	CHANGE_SPLITTER = regexp.MustCompile("\\s+\\|\\s+")
)
