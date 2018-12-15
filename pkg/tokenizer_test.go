package logtick

import (
	"github.com/tj/assert"
	"testing"
)

func TestPackInfo(t *testing.T) {
	t.Parallel()

	slices := []string{
		"commit 57080f0e9ee97798f9ff1bfb2bc74bd7f98b5278",
		"Author:     Tony Mutai <tonnimutai@gmail.com>",
		"AuthorDate: Wed Nov 21 10:36:00 2018 +0300",
		"Commit:     Tony Mutai <tonnimutai@gmail.com>",
		"CommitDate: Wed Nov 21 10:36:00 2018 +0300",
	}

	expected := Token{
		Index: []string{"57080f0e9ee97798f9ff1bfb2bc74bd7f98b5278", "Tony Mutai <tonnimutai@gmail.com>", "Wed Nov 21 10:36:00 2018 +0300", "Tony Mutai <tonnimutai@gmail.com>", " Wed Nov 21 10:36:00 2018 +0300"},
		Type:  "INFO",
	}

	got := PackInfo(slices)

	assert.Equal(t, len(expected.Index), len(got.Index))
	assert.Equal(t, expected.Type, got.Type)
}

func TestPackCommit(t *testing.T) {
	t.Parallel()

	slices := []string{
		"more test",
		"---",
		"index.js | 1 +",
		"test.js  | 1 +",
		"2 files changed, 2 insertions(+)",
	}

	expected := Token{
		Index:   []string{"more test"},
		Changes: [][]string{{"index.js | 1 +"}, {"test.js  | 1 +"}, {"2 files changed, 2 insertions(+)"}},
		Type:    "COMMIT",
	}

	got := PackCommit(slices)

	assert.Equal(t, len(expected.Index), len(got.Index))
	assert.Equal(t, len(expected.Changes), len(got.Changes))
	assert.Equal(t, expected.Type, got.Type)
}

func TestPackDiff(t *testing.T) {
	t.Parallel()

	slices := []string{
		"diff --git a/index.js b/index.js",
		"index e69de29..52f6164 100644",
		"--- a/index.js",
		"+++ b/index.js",
		"@@ -0,0 +1 @@",
		"+console.log('A');",
		"diff --git a/test.js b/test.js",
		"index e69de29..e299cf1 100644",
		"--- a/test.js",
		"+++ b/test.js",
		"@@ -0,0 +1 @@",
		"+console.log('B');",
	}

	expected := Token{
		Diff: [][]string{{"diff --git a/index.js b/index.js index e69de29..52f6164 100644 --- a/index.js +++ b/index.js @@ -0,0 +1 @@ +console.log('A');"}, {"diff --git a/test.js b/test.js index e69de29..e299cf1 100644 --- a/test.js +++ b/test.js @@ -0,0 +1 @@ +console.log('B');"}},
		Type: "DIFF",
	}

	got := PackDiff(slices)

	assert.Equal(t, len(expected.Diff), len(got.Diff))
	assert.Equal(t, expected.Type, got.Type)
}
