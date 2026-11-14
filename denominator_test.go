package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// readmePath is BentoPDF's own README, kept here as testdata rather than
// fetched.
//
// ⛔⛔ A COMMITTED COPY, AND THAT IS THE POINT. The denominator of this whole
// measurement is "how many tools does BentoPDF have", and it is read from
// somebody else's repository. Fetching it inside the test would make the gate
// fail when the network is down and pass quietly when the reference GREW —
// a reference that moves under a measurement turns a widening gap into an
// unchanged number. With the copy committed, refreshing it is a DIFF somebody
// has to look at:
//
//	gh api repos/alam00000/bentopdf/contents/README.md \
//	  -H "Accept: application/vnd.github.raw" > testdata/bentopdf-README.md
//
// and this test then names every row that appeared, vanished or was renamed.
const readmePath = "testdata/bentopdf-README.md"

// readmeTools is every tool BentoPDF's README names in a table whose first
// column is "Tool Name".
//
// ⛔ The header is what selects a table. The README holds other tables — the
// licence tiers, the AGPL components it loads from a CDN — and counting their
// rows as tools would inflate the denominator with things nobody could
// implement.
// ⛔ IT COUNTS ROWS, NOT NAMES. BentoPDF lists "Redact Content" in two of its
// own tables, so the table here judges it twice on purpose. A set keyed by name
// collapses that pair on both sides and agrees for the wrong reason — and the
// day a different tool is duplicated, the collapse would hide a row nobody
// judged.
func readmeTools(t *testing.T) map[string]int {
	t.Helper()
	f, err := os.Open(readmePath)
	if err != nil {
		t.Fatalf("the reference is not here: %v", err)
	}
	defer f.Close()

	out := map[string]int{}
	inTools := false
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "|") {
			inTools = false
			continue
		}
		cells := tableCells(line)
		if len(cells) == 0 {
			continue
		}
		switch {
		case strings.EqualFold(cells[0], "Tool Name"):
			inTools = true
		case isRule(cells[0]):
			// the ---- row under a header
		case inTools:
			if name := plainName(cells[0]); name != "" {
				out[name]++
			}
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatal("no tool rows were read, so this measures nothing")
	}
	return out
}

// tableCells splits a markdown table row into its cells.
func tableCells(line string) []string {
	line = strings.Trim(line, "|")
	parts := strings.Split(line, "|")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		out = append(out, strings.TrimSpace(p))
	}
	return out
}

// isRule says whether a cell is a table's dashed rule.
func isRule(cell string) bool {
	trimmed := strings.Trim(cell, "-: ")
	return cell != "" && trimmed == ""
}

// plainName strips the emphasis and links a README puts round a tool's name.
func plainName(cell string) string {
	name := strings.TrimSpace(cell)
	name = strings.Trim(name, "*`")
	// A linked name, [Thing](url), is the text inside the brackets.
	if strings.HasPrefix(name, "[") {
		if end := strings.Index(name, "]"); end > 1 {
			name = name[1:end]
		}
	}
	return strings.TrimSpace(strings.Trim(name, "*`"))
}

// ⛔⛔ THE DENOMINATOR IS A GATE, NOT A NUMBER IN A COMMENT. This table says it
// judges 111 of BentoPDF's tools; if BentoPDF has 115, then four of them are
// not missing and not present but UNCOUNTED, and the report would say the gap
// had not grown. Counting in one direction only is an argument, not a
// measurement, so this counts in both.
func TestTheTableJudgesExactlyTheReadmesTools(t *testing.T) {
	inReadme := readmeTools(t)
	judged := map[string]int{}
	for _, tl := range tools {
		judged[tl.name]++
	}

	rows := 0
	for _, n := range inReadme {
		rows += n
	}
	t.Logf("the reference holds %d tool rows under %d distinct names; "+
		"this table judges %d rows under %d",
		rows, len(inReadme), len(tools), len(judged))
	if rows != len(tools) {
		t.Errorf("the reference holds %d tool rows and this table judges %d",
			rows, len(tools))
	}
	for name, want := range inReadme {
		if got := judged[name]; got != want {
			t.Errorf("BentoPDF names %q %d time(s) and this table judges it %d: "+
				"the gap is not what the report says", name, want, got)
		}
	}
	for name, got := range judged {
		if inReadme[name] == 0 {
			t.Errorf("this table judges %q (%d row(s)) and BentoPDF's README no longer "+
				"names it: either it was renamed, or we are measuring against a tool "+
				"that is gone", name, got)
		}
	}
}
