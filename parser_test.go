package logtick

import (
	"fmt"
	"github.com/tj/assert"
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) {
	var got CommitTokens

	t.Run("Parsing git log", func(t *testing.T) {
		log, err := ioutil.ReadFile("testdata/git.log")
		if err != nil {
			t.Fatal(err)
		}

		content := fmt.Sprintf("%s", log)

		got, err = Parse(content)
		if err != nil {
			t.Fatal(err)
		}

		expected := CommitTokens{
			Token{
				Index: []string{"57080f0e9ee97798f9ff1bfb2bc74bd7f98b5278", "Tony Mutai <tonnimutai@gmail.com>", "Wed Nov 21 10:36:00 2018 +0300", "Tony Mutai <tonnimutai@gmail.com>", "Wed Nov 21 10:36:00 2018 +0300"},
				Type:  "INFO",
			},
			Token{
				Index:   []string{"more test"},
				Changes: [][]string{{"index.js | 1 +"}, {"test.js  | 1 +"}, {"2 files changed, 2 insertions(+)"}},
				Type:    "COMMIT",
			},
			Token{
				Diff: [][]string{{"diff --git a/index.js b/index.js", "index e69de29..52f6164 100644", "--- a/index.js", "+++ b/index.js", "@@ -0,0 +1 @@", "+console.log('A');"}, {"diff --git a/test.js b/test.js", "index e69de29..e299cf1 100644", "--- a/test.js", "+++ b/test.js", "@@ -0,0 +1 @@", "+console.log('B');"}},
				Type: "DIFF",
			},
		}

		assert.Equal(t, expected, got)
	})

	t.Run("Commit Tokens to JSON", func(t *testing.T) {

		expected := `{"commit":"57080f0e9ee97798f9ff1bfb2bc74bd7f98b5278","author":"Tony Mutai \u003ctonnimutai@gmail.com\u003e","author_date":"Wed Nov 21 10:36:00 2018 +0300","committer":"Tony Mutai \u003ctonnimutai@gmail.com\u003e","commit_date":"Wed Nov 21 10:36:00 2018 +0300","message":"more test","change_stats":"2 files changed, 2 insertions(+)","files":[{"path":"index.js","changes":"1 +","raw":""},{"path":"test.js","changes":"1 +","raw":""}],"diff":[{"path":"index.js","changes":"","raw":"diff --git a/index.js b/index.js\nindex e69de29..52f6164 100644\n--- a/index.js\n+++ b/index.js\n@@ -0,0 +1 @@\n+console.log('A');"},{"path":"test.js","changes":"","raw":"diff --git a/test.js b/test.js\nindex e69de29..e299cf1 100644\n--- a/test.js\n+++ b/test.js\n@@ -0,0 +1 @@\n+console.log('B');"}]}`

		gotJSON, err := got.ToJSON()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expected, gotJSON)
	})

}
