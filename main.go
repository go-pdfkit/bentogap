// Command bentogap says, tool by tool, what BentoPDF offers that the go-pdfkit
// workbench does not.
//
// ⛔ IT REPORTS ITS DENOMINATOR. An inventory that prints only what is missing
// cannot be told from one that could not read half its input: the count of what
// was READ has to appear beside the count of what was found. So every one of
// BentoPDF's tools carries a verdict here, and a tool with no verdict is a
// failure of this program rather than a silent omission -- see TestEveryToolIsJudged.
//
// ⛔ AND THE JUDGEMENT IS DATA, NOT A STRING MATCH. "Rotate PDF" and "Turn them"
// are the same tool under two names, and "PDF Editor" and "Fill in" are not the
// same tool under two names. Matching on words would get both wrong, in opposite
// directions. Each row says which of our verbs answers it, or says plainly that
// none does.
//
// ⛔⛔ AND THE REASON THIS IS A REPOSITORY. go-pdfkit/playground#71 published
// "58 missing" and was still saying it after the instrument had been corrected
// five times in one morning: the issue was never re-rendered from the program.
// A count nobody can clone is a count nobody re-runs, and a count nobody re-runs
// goes stale silently while reading exactly like a measurement.
//
// Sources, read 2026-11-05:
//   - BentoPDF's tool tables, github.com/alam00000/bentopdf/blob/main/README.md
//   - the go-pdfkit workbench's own verb table, go-pdfkit/app ribbon.go
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// osExit is replaceable so a test can drive main without ending the process.
var osExit = os.Exit

func main() {
	if err := RootCmd().Execute(); err != nil {
		osExit(1)
	}
}

// RootCmd builds the command, so a test can drive it with SetArgs and SetOut
// rather than through the process.
func RootCmd() *cobra.Command {
	var only string
	cmd := &cobra.Command{
		Use:   "bentogap",
		Short: "What BentoPDF offers that the go-pdfkit workbench does not",
		Long: "Prints every tool in BentoPDF's tables with the verdict this fleet has\n" +
			"earned against it, and prints the DENOMINATOR beside the counts: how many\n" +
			"tools were read, how many judged, how many not.\n\n" +
			"A row with no verdict fails the program rather than being left out, because\n" +
			"a report that silently drops what nobody looked at reads exactly like a\n" +
			"report that covered everything.",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return report(cmd.OutOrStdout(), tools, only)
		},
	}
	cmd.Flags().StringVar(&only, "only", "",
		"print just one verdict (have, partly, missing, elsewhere, declined)")
	return cmd
}

// verdictNamed turns a --only value into the verdict it names.
//
// ⛔ A FILTER THAT MATCHES NOTHING MUST NOT LOOK LIKE A VERDICT WITH NO ROWS.
// The first version compared the flag against each verdict's name as it went,
// so `--only mising` printed the denominator line and nothing else -- which is
// precisely what `--only declined` prints when no row is declined. One is a
// typo and the other is an answer, and the output could not tell them apart.
func verdictNamed(s string) (verdict, bool) {
	for _, v := range []verdict{have, partly, missing, elsewhere, declined} {
		if strings.EqualFold(s, v.String()) {
			return v, true
		}
	}
	return unjudged, false
}

// report prints the partition and the denominator, and fails if anything was
// left unjudged.
func report(out io.Writer, all []tool, only string) error {
	want, ok := verdictNamed(only)
	if only != "" && !ok {
		return fmt.Errorf("--only %q is not a verdict: it is one of have, partly, missing, elsewhere, declined", only)
	}

	counts := map[verdict]int{}
	for _, t := range all {
		counts[t.v]++
	}
	order := []verdict{have, partly, elsewhere, declined, missing, unjudged}

	for _, v := range order {
		if counts[v] == 0 {
			continue
		}
		if only != "" && v != want {
			continue
		}
		fmt.Fprintf(out, "\n== %s (%d of %d) ==\n", v, counts[v], len(all))
		rows := []tool{}
		for _, t := range all {
			if t.v == v {
				rows = append(rows, t)
			}
		}
		sort.SliceStable(rows, func(i, j int) bool { return rows[i].section < rows[j].section })
		for _, t := range rows {
			if t.ours == "" {
				fmt.Fprintf(out, "  %-28s  %s\n", t.name, t.section)
				continue
			}
			fmt.Fprintf(out, "  %-28s  %-26s  %s\n", t.name, t.section, t.ours)
		}
	}

	fmt.Fprintf(out, "\n%d tools read from BentoPDF's tables; %d judged, %d not\n",
		len(all), len(all)-counts[unjudged], counts[unjudged])
	if counts[unjudged] > 0 {
		fmt.Fprintln(out, "⛔ a tool with no verdict is a gap nobody has looked at")
		return fmt.Errorf("%d tool(s) carry no verdict", counts[unjudged])
	}
	return nil
}

// A verdict says how the workbench stands against one of BentoPDF's tools.
type verdict int

const (
	// unjudged is the zero value on purpose: a row somebody forgot is a row
	// this program refuses to report on.
	unjudged verdict = iota
	// have: the workbench does this, under its own name.
	have
	// partly: it does some of this, and the gap is worth naming.
	partly
	// missing: it does not do this, and doing it would belong here.
	missing
	// elsewhere: the fleet does this, in another repository -- it is not a gap
	// in the workbench, it is a tool that has not been put on its ribbon.
	elsewhere
	// declined: it is not a gap. A browser PDF workbench is not an office suite
	// and not an OCR engine, and saying so is a decision, not an oversight.
	declined
)

func (v verdict) String() string {
	switch v {
	case have:
		return "have"
	case partly:
		return "partly"
	case missing:
		return "MISSING"
	case elsewhere:
		return "elsewhere"
	case declined:
		return "declined"
	}
	return "UNJUDGED"
}

// A tool is one row of BentoPDF's tables and what we make of it.
type tool struct {
	section string
	name    string
	v       verdict
	// ours is the workbench verb, repository or reason that justifies the
	// verdict. It is never empty, `missing` included: an unwitnessed "missing"
	// cannot be told from "I did not look".
	ours string
}
