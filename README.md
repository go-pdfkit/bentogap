# bentogap

What [BentoPDF](https://github.com/alam00000/bentopdf) offers that the go-pdfkit workbench does not — counted row by row, and printed **with its denominator**.

```
go run .                 # every verdict
go run . --only missing  # just one
```

## Why this is a repository

It was a directory on one workstation. `go-pdfkit/playground#71` published its numbers, the table was corrected five times in a single morning, **and the issue was never re-rendered from the program** — so it went on saying "58 missing" long after the instrument said otherwise.

That is the failure this tool was built to detect, happening to the tool itself. A count nobody can clone is a count nobody re-runs, and a count nobody re-runs goes stale silently *while still reading like a measurement*.

⛔ **So the numbers below are stamped, and they are not the authority — `go run .` is.** If this file and the program disagree, the program is right and this file is the bug.

## The count, 2026-11-14

| verdict | n | what it means |
|---|---:|---|
| **have** | 73 | the workbench does it, under its own name |
| **partly** | 4 | it does some of it, and the gap is worth naming |
| **elsewhere** | 18 | the fleet already does it and the ribbon does not |
| **declined** | 0 | not a gap: a decision, said out loud |
| **MISSING** | 16 | it does not exist anywhere in the fleet |
| | **111** | every row judged, none skipped |

## What makes it an inventory rather than a list of opinions

⛔ **It reports its denominator.** A report that prints only what is missing cannot be told from one that could not read half its input. `111 tools read; 111 judged, 0 not` is the sentence that makes the rest meaningful, and `TestEveryToolIsJudged` fails the build if any row still carries the zero value.

⛔ **Every verdict names its witness — `missing` included.** It used to exempt `missing`, on the reasoning that there is nothing to cite when a thing does not exist. That reasoning is exactly how the table was wrong four times in one morning: an unwitnessed *missing* cannot be told from *I did not look*. Eleven turned out to be `pdfops` verbs, seven more lived in `go-pdfkit/convert`, six were already in the application.

⛔ **The sections are held to the length they were read at.** If a table is transcribed short, the gap list shrinks and looks like progress. `TestTheSectionsAreAsLongAsTheyWereRead` pins all six.

⛔ **The judgement is data, not a string match.** "Rotate PDF" and "Turn them" are the same tool under two names; "PDF Editor" and "Fill in" are not. Matching on words gets both wrong, in opposite directions.

⛔ **`--only` refuses a word that is not a verdict.** `declined` is a real verdict that no row currently carries, so it prints the denominator and nothing else — byte for byte what a typo used to print. One is an answer and the other is a question nobody asked, and the output could not tell them apart.

## Reading a verdict

`elsewhere` is the one worth understanding. It does not mean *nearly done*; it means the capability exists in the fleet and cannot be reached from the workbench's ribbon. It is only used where a repository's **own description or command table** names the capability, and every such row cites it.

Several `elsewhere` rows are deliberate rather than pending — the conversions to PDF are reachable from `go-pdfkit/doc2pdf` at a shell, and putting them on the ribbon was measured at +10.8 MB of wasm, so the row records the trade instead of pretending it is free.

## Sources

- BentoPDF's own tool tables, read 2026-11-05 — `testdata/bentopdf-README.md` is the copy that was read, so a later edit upstream cannot silently change what this was judged against.
- The workbench's verb table, `go-pdfkit/app`.
