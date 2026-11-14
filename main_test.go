package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// ⛔ THE TEST THAT MAKES THIS AN INVENTORY RATHER THAN A LIST OF OPINIONS. A row
// nobody judged keeps the zero value, and a report that quietly leaves those out
// reads exactly like a report that covered everything.
func TestEveryToolIsJudged(t *testing.T) {
	for _, tl := range tools {
		if tl.v == unjudged {
			t.Errorf("%q (%s) has no verdict", tl.name, tl.section)
		}
	}
}

// ⛔⛔ EVERY VERDICT NAMES ITS WITNESS, `missing` INCLUDED. It used to exempt
// `missing`, on the reasoning that there is nothing to cite when a thing does
// not exist. That reasoning is exactly how this table was wrong four times in
// one morning: an unwitnessed "missing" is indistinguishable from "I did not
// look", and eleven of them turned out to be `pdfops` verbs, seven more lived
// in go-pdfkit/convert, and six were already in the application.
//
// A `missing` row now has to say where the question was ASKED. The sweep that
// answers it is `fleetverbs --match <word> clones` over the thirty
// repositories listed in go-pdfkit-capability-sweep/fetch.
func TestEveryVerdictNamesItsWitness(t *testing.T) {
	for _, tl := range tools {
		if strings.TrimSpace(tl.ours) == "" {
			t.Errorf("%q is %s and names nothing — not what answers it, and not where it was asked",
				tl.name, tl.v)
		}
	}
}

// BentoPDF's six tables, and how many rows each held when they were read. ⛔ The
// counts are the DENOMINATOR: if a section is transcribed short, the gap list
// shrinks and looks like progress.
var sectionCounts = map[string]int{
	"organize": 23, "edit": 23, "automate": 1,
	"to pdf": 33, "from pdf": 15, "secure": 16,
}

func TestTheSectionsAreAsLongAsTheyWereRead(t *testing.T) {
	got := map[string]int{}
	for _, tl := range tools {
		got[tl.section]++
	}
	for s, want := range sectionCounts {
		if got[s] != want {
			t.Errorf("%s holds %d rows; BentoPDF's table had %d", s, got[s], want)
		}
	}
	for s := range got {
		if _, ok := sectionCounts[s]; !ok {
			t.Errorf("%s is a section nobody counted", s)
		}
	}
	total := 0
	for _, n := range sectionCounts {
		total += n
	}
	if len(tools) != total {
		t.Errorf("%d rows in all; the six tables held %d", len(tools), total)
	}
}

func TestNoToolIsListedTwiceInOneSection(t *testing.T) {
	seen := map[string]bool{}
	for _, tl := range tools {
		k := tl.section + "/" + tl.name
		if seen[k] {
			t.Errorf("%q appears twice under %s", tl.name, tl.section)
		}
		seen[k] = true
	}
}

// An unjudged row has to reach the exit status, not only the output.
func TestAnUnjudgedRowIsAFailure(t *testing.T) {
	if err := report(io.Discard, tools, ""); err != nil {
		t.Fatalf("the real table is complete and was reported as %v", err)
	}
	bad := append([]tool{}, tools...)
	bad = append(bad, tool{section: "organize", name: "something nobody looked at"})
	if err := report(io.Discard, bad, ""); err == nil {
		t.Fatal("a row with no verdict was reported as a clean inventory")
	}
}

// ⛔⛔ A TYPO IN --only USED TO BE INDISTINGUISHABLE FROM AN EMPTY VERDICT, and
// this pair is what separates them. `declined` is a real verdict that NO row
// currently carries, so it prints the denominator and nothing else — which is
// byte for byte what `mising` printed before the flag was validated. One is an
// answer and the other is a question that was never asked.
//
// ⛔ The pair is the test. Asserting only that the typo fails would pass against
// a version that rejected every --only value, including the correct ones.
func TestAVerdictNobodyCarriesIsNotTheSameAsAWordThatIsNotAVerdict(t *testing.T) {
	var real bytes.Buffer
	if err := report(&real, tools, "declined"); err != nil {
		t.Fatalf("declined is a verdict and was refused: %v", err)
	}
	if n := strings.Count(real.String(), "== "); n != 0 {
		t.Fatalf("no row is declined, so nothing should be printed under it; got %d section(s)", n)
	}
	if !strings.Contains(real.String(), "tools read from BentoPDF's tables") {
		t.Error("the denominator is what makes an empty answer readable, and it is missing")
	}

	err := report(io.Discard, tools, "mising")
	if err == nil {
		t.Fatal("a word that is not a verdict was accepted, and printed what an empty verdict prints")
	}
	if !strings.Contains(err.Error(), "is not a verdict") {
		t.Errorf("it was refused, but not for being an unknown verdict: %v", err)
	}
}

// Every verdict the flag advertises can actually be asked for, which a list in
// a help string cannot promise on its own.
func TestEveryVerdictTheFlagAdvertisesIsAccepted(t *testing.T) {
	for _, name := range []string{"have", "partly", "missing", "elsewhere", "declined"} {
		t.Run(name, func(t *testing.T) {
			if err := report(io.Discard, tools, name); err != nil {
				t.Errorf("--only %s was refused: %v", name, err)
			}
		})
	}
}

// ⛔ main() ITSELF, through the seam that exists for it. osExit is replaceable
// precisely so the one line that decides the process's exit status can be
// asserted rather than left as the only uncovered block and called "just
// main()". Both directions are the test: a failure must set the status, and a
// success must NOT — an exit that fires on both is indistinguishable from one
// that fires on neither when only the failing case is checked.
func TestMainSetsTheExitStatusFromTheCommand(t *testing.T) {
	oldExit, oldArgs := osExit, os.Args
	t.Cleanup(func() { osExit, os.Args = oldExit, oldArgs })

	code := -1
	osExit = func(c int) { code = c }

	os.Args = []string{"bentogap", "--only", "mising"}
	main()
	if code != 1 {
		t.Errorf("a word that is not a verdict left the exit status at %d", code)
	}

	code = -1
	os.Args = []string{"bentogap", "--only", "declined"}
	main()
	if code != -1 {
		t.Errorf("a complete report exited with %d", code)
	}
}

// The command writes through its own writer rather than to the process's
// stdout, which is what lets it be driven by a test at all.
func TestTheCommandReportsThroughTheWriterItIsGiven(t *testing.T) {
	var out bytes.Buffer
	cmd := RootCmd()
	cmd.SetArgs([]string{"--only", "missing"})
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("the command failed: %v", err)
	}
	if !strings.Contains(out.String(), "== MISSING") {
		t.Errorf("it did not print the verdict it was asked for:\n%s", out.String())
	}
	if strings.Contains(out.String(), "== have") {
		t.Error("--only missing printed another verdict as well")
	}
}
