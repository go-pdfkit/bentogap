package main

// tools is BentoPDF's own tables, read from its README on 2026-11-05, with a
// verdict against the go-pdfkit workbench (go-pdfkit/app, v0.19.1) and the
// fleet behind it.
//
// ⛔ EVERY ROW CARRIES A VERDICT. The zero value is `unjudged`, and the program
// exits non-zero if any row still has it: a list of gaps that silently drops
// the rows nobody looked at reads exactly like a list of gaps that is complete.
//
// ⛔ `elsewhere` means THE FLEET DOES IT AND THE RIBBON DOES NOT, and it is only
// used where a repository's own description names the capability. A library that
// reads a format is not a conversion to PDF: go-odf reads ODF and go-pdfkit
// writes PDF, and nothing joins them, so those rows are `missing` with the two
// halves named.
// asked is what a `missing` row cites: the sweep that looked for it and found
// nothing.
//
// ⛔ A `missing` row USED TO CITE NOTHING, on the reasoning that there is
// nothing to point at when a thing does not exist. That reasoning is how this
// table was wrong four times in one morning — an unwitnessed "missing" cannot
// be told from "I did not look", and eleven of them were `pdfops` verbs.
//
//	cd go-pdfkit-capability-sweep && fleetverbs --quiet --match <word> clones
//
// over the thirty repositories listed in that directory's fetch program, on
// 2026-11-05: the PDF stack, the richdoc model and every reader into it, the
// office formats, and the picture libraries underneath.
const asked = "asked across the 30 repositories of the capability sweep with fleetverbs; nothing answered"

// redaction is where the workbench stands against "Permanently remove
// sensitive content from your PDFs", which is BentoPDF's own wording.
//
// ⛔ STILL `partly`, AND THE REASON IS NARROWER AGAIN. The text in a region
// comes out of the content stream; since ops v0.35.0 the SAMPLES of a picture
// the region crosses are destroyed too, and since ops v0.38.0 (playground
// v0.58.0) that includes a JPEG — which is the whole of what a reader of
// scanned paper has, every scanner writing DCTDecode.
//
// ⛔⛔ THAT LAST CLAUSE IS WHY THIS ROW DID NOT BECOME `have` IN v0.35.0 AND
// SHOULD NOT HAVE READ AS GENEROUSLY AS IT DID. Between v0.35.0 and v0.38.0
// this table, the README, the docs and the landing all said a redaction
// handles the scanned page, and the one picture a scan actually contains was
// still refused. The sentence was true of a fixture.
//
// What is left refused now is a picture drawn TURNED or drawn TWICE, one
// carrying a mask or a palette, a CMYK JPEG (no faithful way back to RGB
// without the output profile the file does not carry), a form XObject, an
// annotation, and a half-covered operator. BentoPDF's "permanently remove"
// covers those, so `have` would retire a row a reader with a turned scan still
// cannot use.
const redaction = "`Redact` on the Mark up tab, over a selection (app v0.26.0, ops.Doc.Redact v0.28.1; " +
	"it sat on Read until v0.37.0, under a panel promising that nothing there changes the document): " +
	"the operator that drew the text is blanked out of the page's own content stream, so the " +
	"words are gone from the bytes rather than covered — and since ops v0.35.0 (app v0.51.0) a " +
	"PICTURE the region crosses has the samples under it DESTROYED and is written back, which " +
	"ISO 32000-1 12.5.6.23 requires in those words: \"clipping or image masks shall not be used " +
	"to hide that data\". ⛔ That is done for a plain or JPEG, 8-bit, DeviceGray or DeviceRGB " +
	"picture drawn UPRIGHT and drawn ONCE — extract v0.11.0's Image.Matrix was added to tell an " +
	"upright one from a turned one, because the box a picture covers cannot, and mapping a " +
	"region into a turned picture's samples by that box would destroy MORE than was asked for. " +
	"⛔⛔ THE JPEG CAME LAST AND IT WAS THE ONLY ONE A SCAN HAS (ops v0.38.0, closing ops#82): " +
	"every scanner writes DCTDecode, so for three releases \"a scanned page can be redacted\" " +
	"was true of a fixture and false of any file a reader owned. It is decoded and written back " +
	"as PLAIN SAMPLES, never re-encoded — re-encoding re-quantises block by block and would " +
	"degrade the whole photograph in order to destroy a corner of it — so what is written is " +
	"exactly what a reader of the original would have seen, and the gate compares the kept " +
	"samples BYTE FOR BYTE against its own decode of that JPEG. The file grows; that is the " +
	"trade, and it is paid only by a picture a region actually reaches. The rest are refused " +
	"BY NAME: a CMYK JPEG (no faithful way back to RGB without the output profile the file does " +
	"not carry), anything past 100 megapixels (read from the HEADER, before a decoder allocates " +
	"width x height x 3: 39 bytes can ask for 1.2 GB), a form XObject, an annotation, a " +
	"half-covered operator. ⛔ And the page's /Resources are " +
	"REBUILT rather than copied, or the original picture would sit in the file beside the " +
	"blanked one holding every pixel the redaction promised to remove — ops#74 one object along, " +
	"which is why the gate walks EVERY object of the written file and decodes it"

// toPDF is the half that did not exist until 2026-11-05 and now does:
// go-pdfkit/document.Write takes a richdoc to a laid-out PDF, through
// go-richdoc/html.Write and go-pdfkit/html2pdf.Export.
//
// ⛔ These rows are `elsewhere` and NOT `have`, by the same standard applied to
// go-pdfkit/convert: the fleet's libraries can do it in a few lines of Go, and
// nothing a person can run does. A library with no caller is not delivered.
// ⛔ "no command joins the two" WAS TRUE FOR ONE DAY. go-pdfkit/doc2pdf v0.1.0
// is that command, and this witness went stale the moment it was released —
// which is the hazard of a witness written in the present tense about an
// absence. The verdict does not change: `elsewhere` means the fleet does it and
// the workbench's ribbon does not, and a shell command is not a ribbon.
// ⛔ "NOT ON THE RIBBON" IS NOW A MEASURED REASON RATHER THAN AN OMISSION — and
// ⛔⛔ THE FIRST MEASUREMENT WAS TAKEN WITHOUT THE SEAMS THE ENGINE ALREADY HAS,
// which is the lesson worth keeping. An arm was added to the workbench's Open
// that reads .md/.csv/.tsv/.json/.xml/.txt and lays them out with
// document.Write; same commit, same toolchain, same flags:
//
//	plain                              24,011,429 -> 61,963,937 (+37,952,508, x2.58)
//	-tags webengine_nojs,webengine_nonet          -> 34,857,531 (+10,846,102, x1.45)
//	gzipped, plain                      7,780,496 -> 17,164,594 (+9,384,098)
//	gzipped, tagged                               -> 11,786,983 (+4,006,487)
//
// Plain, document.Write reaches a PDF through html2pdf, so it links
// go-webengine/engine and with it dop251/goja — A JAVASCRIPT INTERPRETER, 5.15
// MB — and refraction-networking/utls, A TLS STACK, 1.36 MB. ⛔ But goja is
// confined to engine/js (checked across the whole module: outside js/ it
// appears only in go.mod, go.sum, the README and one test), and the root
// package carries webengine_nojs / webengine_nonet build tags as file pairs
// (net_http.go/net_none.go, settle_scripts.go/settle_noscripts.go). With them,
// wasmsize --match reports goja and utls contributing ZERO bytes of code: both
// tables come back EMPTY. 71 % of the increment was a tag nobody had passed.
//
// What is left is +3.1 MB of code and +7.7 MB of DATA, which wasmsize does not
// attribute. Diffing the two `go list -deps` by module names it: four more
// embedded font families — dejavusans, gomono, inter, lora — on top of the four
// the workbench already carries. ⛔ dejavusans is the per-character GLYPH
// FALLBACK of html2pdf#22 (6 lost characters to 0 of 106) and is not in
// question; the other three duplicate roles arimo, tinos and cousine already
// fill metric-compatibly. go-pdfkit/html2pdf#50 asks for the seam.
// go-pdfkit/document#12 carries the whole measurement.
const toPDF = "go-pdfkit/document.Write takes richdoc to a laid-out PDF (v0.2.0); " +
	"go-pdfkit/doc2pdf v0.1.0 converts it from a shell; ⛔ NOT ON THE RIBBON, and MEASURED: an arm " +
	"that opened .md/.csv/.json/.xml/.txt in the workbench took main.wasm from 24,011,429 to " +
	"61,963,937 bytes (x2.58) — but to 34,857,531 (x1.45) with the engine's OWN build tags, " +
	"`-tags webengine_nojs,webengine_nonet`, which take dop251/goja (a JavaScript interpreter) and " +
	"refraction-networking/utls (a TLS stack) to ZERO bytes of code, measured with wasmsize " +
	"--match rather than inferred from the total. ⛔⛔ THE FIRST FIGURE WAS TAKEN WITHOUT SEAMS THE " +
	"ENGINE ALREADY HAD: goja lives only in engine/js and the root carries nojs/nonet file pairs. " +
	"What remains is +3.1 MB of code and +7.7 MB of DATA — four more embedded font families, of " +
	"which dejavusans is html2pdf#22's glyph fallback and stays, and inter/lora/gomono duplicate " +
	"roles arimo/tinos/cousine already fill metric-compatibly (go-pdfkit/html2pdf#50). " +
	"go-pdfkit/document#12 carries the measurement"

// toPDFnoSuffix is the same half-built shape for the three OpenDocument bodies
// doc2pdf does not offer a suffix for. go-odf's Parse reads all four bodies —
// its own documentation says a spreadsheet "is the same package as a text
// document ... so Parse reads one too" — so this is three lines of table, not a
// reader to write.
const toPDFnoSuffix = "go-pdfkit/document.Write takes richdoc to a laid-out PDF (v0.2.0); " +
	"odf.Parse reads this body already and doc2pdf offers no suffix for it, so nothing a person can run joins the two"

var tools = []tool{
	// ---- Organize & Manage PDFs (23) ----
	{"organize", "Merge PDFs", have, "`Add files after this one` takes AS MANY AS YOU CHOOSE, in the order you chose them (app v0.49.0): a second <input type=file> carrying `multiple` and a second method on the host, because every other verb that asks for a file wants exactly one and a shared picker allowing several would hand back a list the code took the first of in silence. ⛔ The order is the PICKER's, not the order the bytes arrive in — each file is read through a promise of its own, so each result goes to its own slot and the callback waits for the count; appending as they arrived would merge five scans in the order of their SIZES. ⛔ A file among them that cannot be read stops the whole merge and names it. ⛔ The browser check hands the sample in twice, because the picker and the host method live in the js/wasm half that `go test` never compiles. ⛔ An earlier verdict said \"the added file's bookmarks are not carried\", which was read off nothing: ops carries EVERY source's outline, and TestAddingAFileAfterThisOneKeepsBOTHTablesOfContents (v0.42.0) reads both tables of contents back"},
	{"organize", "Split PDFs", have, "`Split and hand them over`"},
	{"organize", "Organize Pages", have, "`Arrange the pages` on the Pages tab (app v0.45.0): every page as a grid of tiles where the page was — press one to pick it up, press the page it should go in front of. ⛔ TWO PRESSES AND NOT A DRAG: a drag over a grid inside a scroll view has to hit-test tiles against a moving viewport and decide what a press that moved four pixels meant, where two presses say the same thing and the tile and the status line both show what is held; it is also the gesture a touch screen gives for nothing. ⛔ The number asked of ops.Move is one LESS when a page moves forwards, because Move takes the page out first and a reader who pressed the fifth tile meant \"in front of this one\" — read back out of the written document by a fixture whose pages draw their own numbers. Rotate, delete, crop and the rest are on the same tab, acting on the page being read"},
	{"organize", "Extract Pages", have, "`Keep only these`"},
	{"organize", "Delete Pages", have, "`Delete these`, and `Delete` for the page on screen"},
	{"organize", "Rotate PDF", have, "`Turn them` for a range, `Rotate` for the page on screen"},
	{"organize", "Rotate by Custom Degrees", have, "`Turn them by an angle` on the Pages tab (app v0.52.0) and `pdfops turn -by <degrees>`, on ops.Doc.Turn, new in ops v0.36.0: any angle, clockwise, added to what the page already carries. ⛔ THE PAPER GROWS to the turned page's bounding box — w|cos|+h|sin| across, w|sin|+h|cos| down, the reference's own arithmetic — because a sheet that stayed the same size would have two corners of the page off the edge of it; that is the whole difference from Deskew, which keeps the paper because a fraction of a degree on a scan costs nothing. A multiple of ninety still goes through /Rotate and rewrites nothing. ⛔⛔ POSITIVE IS CLOCKWISE AND THE REFERENCE CONTRADICTS ITSELF ABOUT IT: rotate-custom.md says \"Positive values rotate clockwise\", its code writes a multiple of ninety as /Rotate (ISO 32000-1 Table 30: degrees CLOCKWISE) and sends every other angle through pdf-lib's drawPage, which concatenates [cos sin -sin cos] — counter-clockwise in a y-up space. The two arms turn opposite ways, and the seam is at exactly 90. Both are clockwise here, gated by a render across that seam at 89, 90 and 91 degrees"},
	{"organize", "N-Up PDF", have, "`Lay them out`"},
	{"organize", "View PDF", have, "the workbench is a viewer: minipages, contents, zoom to a page"},
	{"organize", "Alternate & Mix Pages", have, "`Mix another file in` on the Sheet tab (app v0.31.0), on ops.Doc.Interleave — three lines over `withAnother`, the seam `merge` and `overlay` already asked the file picker through"},
	{"organize", "Posterize PDF", have, "`Print it larger than the paper` on the Sheet tab (app v0.27.0), on ops.Doc.Poster, with its own tile numbers"},
	{"organize", "PDF Multi Tool", partly, "BentoPDF's row is one page grid that reorders, rotates, deletes and inserts without leaving it. `Arrange the pages` (app v0.45.0) is that grid: it reorders in it, and since v0.47.0 the page PICKED UP is the page every other verb acts on — so turning or deleting the page a reader can see in the grid is a press on the tile and a press on the ribbon, where before it meant finding the page again in the rail. ⛔ Any change to the document puts down what was held, in changeSaying, because a number is not a page: deleting page 3 while holding page 5 would leave the next press moving something nobody pointed at. And since v0.50.0 the PAGE MENU opens over a tile: a secondary press names that page and offers turn, drop, insert, hand over and read — the grid REPLACES the pages with tiles, so until then the one view showing every page at once was the one view where the menu did not open. What is left is a tile that carries its own buttons, which is a look rather than a capability"},
	{"organize", "PDF Booklet", have, "`Fold it into a booklet`"},
	{"organize", "Add Attachments", have, "`Put a file inside it` on the File tab (app v0.31.0), on ops.Doc.Attach — and ops refuses a name already used rather than replacing what is there (\"two files under one name is a document that has lost one of them\"), which reaches the status line unchanged"},
	{"organize", "Extract Attachments", have, "`Show the files inside` on the File tab (app v0.31.0) lists them with a `Take it out` beside each, on ops.Doc.Attachments — ⛔ the DOCUMENT's own list, not the saved file's, so a file attached a moment ago is in it"},
	{"organize", "Edit Attachments", have, "the same list's `Drop it`, on ops.Doc.Detach (app v0.31.0) — which reports whether it found anything, so dropping a name the document never carried says so rather than claiming success"},
	{"organize", "Divide Pages", have, "`Cut each page up` on the Sheet tab (app v0.27.0), on ops.Doc.Divide — ⛔ its own pair of numbers, not the poster's: \"2 × 2\" means halves there and quarters-enlarged here"},
	{"organize", "Combine to Single Page", have, "`Put them on one page` on the Sheet tab (app v0.27.0), on ops.Doc.OnePage"},
	{"organize", "Add Blank Page", have, "`Insert a blank page`"},
	{"organize", "Reverse Pages", have, "`Reverse the order`"},
	{"organize", "View Metadata", have, "`Show what it says` on the File tab (app v0.28.0): every entry of the information dictionary, the nine ISO 32000-1 Table 317 ones in the order the standard lists them and anything else the producer added — ⛔ the panel could WRITE a title and not show the title that was there, so anybody typing into those boxes was overwriting a value they could not see"},
	{"organize", "PDFs to ZIP", have, "`Split and hand them over` puts every piece in ONE zip (app v0.42.0): it handed over one download per piece and apologised for it in its own status line — \"a browser may ask before it takes more than one\" — which for a 200-page document split into twenties is ten prompts, each refusable on its own, leaving somebody with four pieces and no way of knowing which four. One piece is still handed over as the PDF itself. `Hand over every page, zipped` remains the one that zips page PICTURES"},
	// ⛔ `partly` AND NOT `have`: everything a library can answer is answered,
	// and the two VIEWS BentoPDF's row is mostly about are not here. A reader
	// who wants the two pages beside each other, or one over the other, gets a
	// list of changes instead. That is a real part of the tool and not a
	// garnish, so the row stays open.
	{"organize", "Compare PDFs", partly, "`Compare with a file` on the Read tab (app v0.53.0) and `pdfops compare`, " +
		"on ops.Compare, new in ops v0.37.0. ⛔ THE PAGES ARE PAIRED BY WHAT THEY SAY, NOT BY THEIR " +
		"NUMBERS — BentoPDF's own \"automatic page pairing based on content similarity\" — so a page " +
		"inserted in the middle moves nothing after it and is one `page added`; pairing by number " +
		"would report every page from there to the end as changed, which is both useless and false. A " +
		"pair must share HALF ITS WORDS or it is one page removed and another added, or a replaced " +
		"chapter reads as a rewrite of the one before it. Inside a pair: Myers, Algorithmica 1 (1986) " +
		"251-266, over words, with the words that changed together reported together. Of its six " +
		"categories, text, image and annotation are answered from the file (a picture by a HASH of its " +
		"samples, so a swapped logo of the same size is a change), and two more its pairing implies " +
		"and its list does not name — a page added and a page removed. ⛔ WHAT IS MISSING IS THE TWO " +
		"VIEWS: side by side and overlay, which are most of what its page offers, plus its category " +
		"filters, its OCR for scans with no text layer, and exporting the comparison as a PDF. " +
		"`formatting` and `background` are a RENDERING comparison, which is a different instrument and " +
		"a different cost"},

	// ---- Edit & Modify PDFs (23) ----
	{"edit", "PDF Editor", have, "BentoPDF's own row is \"Annotate, highlight, redact, comment, add shapes/images, search, and view\", and the workbench now has every one of them. Highlight, underline, strike out and SQUIGGLE (app v0.34.0 — ops has had four kinds since Mark existed and the ribbon offered three, found by counting the library's enum against the ribbon's), `Redact` (v0.26.0), IMAGES (`Draw a picture on them`, v0.34.0, on ops.Place, which the workbench had never called although its own documentation says \"a signature, a logo, a stamp that is an image rather than a line of text\") the viewer, and SEARCH (`Find` / `Find the next one`, app v0.35.0 — ⛔ a match becomes a SELECTION, so Copy, the four marks and Redact all act on what was found; it matches a phrase across words in the text the selection itself would give back, and never across a line break, because one selection cannot cover two lines) , NOTES and FREE TEXT (`Put a note on it` / `Write in the margin`, app v0.36.0 on ops.Doc.AddNotes and ops.Doc.AddFreeText, new in ops v0.29.0 — ⛔ both carry an /AP, because render draws an annotation ONLY from its appearance: a /Text with a /Name and a /FreeText with a /DA are correct annotations this fleet shows as NOTHING, so the ops tests count rendered pixels INSIDE the annotation's own rectangle) , and EDITING THE TEXT ITSELF (`Replace the words`, app v0.38.0 on ops.Doc.ReplaceText, new in ops v0.30.0 — ⛔ the hard part is not the string but everything after it: a replacement of another width moves the rest of the line, so the new text carries a TJ kern equal to the DIFFERENCE OF THE RAW WIDTH SUMS, which 9.4.4 makes independent of the font size because the size multiplies both the glyph width and the kern; asserted at 6/12/24/48pt, with a control showing the next word moving 18 points without it) , and SHAPES (`Box them` / `Ring them` / `Point at them`, app v0.39.0 on ops.Doc.AddShapes, new in ops v0.31.0: /Line with Table 176's ClosedArrow, /Square and /Circle — ⛔ an ellipse is FOUR CUBIC CURVES because a PDF has no ellipse operator, at 4(sqrt(2)-1)/3 computed rather than typed; the Rect is grown by HALF THE LINE WIDTH because 12.5.6.8 inscribes the shape in it and a stroke straddles its path; and each one goes ROUND the passage rather than over it, because a box on the glyphs hides what it draws attention to)"},
	// ⛔⛔ A FIFTH VERDICT OF MINE THAT HAD GONE STALE, and the worst of them: it
	// read "ops can Flatten, Links, RemoveAnnotations and Sanitize — read what
	// is on a page and destroy it, never author it", and cited a sweep for the
	// word "annotation", which is not what editing text is. ops.Doc.ReplaceText
	// has existed since v0.30.0 and the workbench has called it since v0.38.0.
	// Re-read the CODE, then a test in the repo described.
	{"edit", "Edit PDF Text", have, "BentoPDF's row is \"Edit existing text directly in your PDF. Reuses the document's " +
		"embedded fonts and keeps the rest of the page untouched\", which is ops.Doc.ReplaceText (v0.30.0) " +
		"word for word, on the ribbon as `Replace the words` (app v0.38.0): the replacement is written in " +
		"THE PAGE'S OWN FONT, in the codes that font understands — TestTheTwoRefusalsAboutTheFont refuses a " +
		"font the page's resources do not describe — and the rest of the line stays where it was, because " +
		"the new text carries a TJ kern equal to the DIFFERENCE OF THE RAW WIDTH SUMS, which 9.4.4 makes " +
		"independent of the font size (TestTheKernDoesNotDependOnTheFontSize, at 6/12/24/48pt; " +
		"TestWithoutTheKernTheNextWordWouldMove is the control, 18 points of drift). ⛔ ONE DIFFERENCE, " +
		"filed as go-pdfkit/ops#84: where the replacement needs a character the page's font cannot draw, " +
		"this REFUSES and BentoPDF fetches a fallback font by script and embeds it — so a Greek letter or a " +
		"Japanese word typed into a Latin-1 document is something it can do and this cannot. The refusal is " +
		"deliberate (a silent face change is a page that no longer looks like itself), which is why the " +
		"fallback has to be asked for rather than automatic"},
	{"edit", "Create Fillable Forms", have, "`Make a field for this` on the Mark up tab (app v0.48.0, on ops.Doc.AddFields, which `pdfops addfields -fields <description.json>` had been the only way to reach): a fillable field UNDER the words selected, named after them, as a box to type in or a tick. ⛔ The gesture is a selection because that is the question being asked — a blank printed form is a page of questions and what somebody wants is a box to answer each in — and the name is the words WITHOUT their punctuation, because a field's name travels into exported data and into whatever script fills the form. Each placement is CHECKED against the paper (under the words, else above them), never clamped; a tick is a square; a second field of the same name is refused by the library"},
	{"edit", "PDF Form Filler", have, "`Fill in`, on go-pdfkit/forms; go-pdfkit/xfa reads XFA"},
	{"edit", "Add Page Numbers", have, "`Number them`"},
	{"edit", "Bates Numbering", have, "`Stamp the numbers on`"},
	{"edit", "Add Watermark", have, "`Write it across them`"},
	{"edit", "Header & Footer", have, "`Head and foot them` on the Stamp tab (app v0.44.0): the header at the top and the footer at the foot in ONE change, each its own box, with {page}, {pages} and {n} filled in and an empty box writing nothing there. ⛔ The earlier verdict said \"there is no free header or footer text\", which was wrong -- `Stamp them` has written free text at any of ops.Stamp's nine positions since the verb existed, and ops.expand fills the tokens in for ANY stamp -- so what was missing was one press for a gesture that is one gesture, plus a hint that SAID the tokens work there (app v0.42.0). ⛔⛔ The two boxes took the Marks panel 109 pixels past the room it is given (21 of which it had left before them), so the stamp, the header/footer and the picture are a Stamp tab now"},
	{"edit", "Crop PDF", have, "`Crop them`"},
	{"edit", "Deskew PDF", have, "`Straighten the scan` on the Pages tab (app v0.27.0), on ops.Doc.Deskew — and it says what it MEASURED: \"measured 12 pages, straightened 3\", because a page already level is left alone on purpose"},
	{"edit", "Font to Outline", missing, asked + " for \"outline\": every hit is a path outline in gfx or a bookmark tree, none turns a font into paths in a page"},
	{"edit", "Invert Colors", have, "`Turn them inside out` in `Redraw the pixels` on the Redraw tab, which it was split onto for crowding the File tab (app v0.32.0), on convert.Redraw with Invert — ⛔ and the status line says `⚠ the pages are pictures now`, because every row in this group rasterises: after it nothing can select a word, search the document or read it out. ⛔ convert costs 443,265 bytes as an INCREMENT into the binary that was already there (22,489,931 → 22,933,196, by a control naming Redraw and ToArchive alone) and not the 18.18 MB it measures standalone, which is the measurement that gated this whole batch — contrast go-pdfkit/document's 26.9 MB increment (go-pdfkit/document#11). ⛔ The delivered binary nonetheless grew 1,126,671 bytes, so 683,406 of them are neither convert nor a dependency bump (the two go.mod files are byte-identical), which is an OPEN QUESTION (go-pdfkit/app#85) and not a rounding: the figure for anyone counting bytes is 1.07 MiB"},
	{"edit", "Change Background", have, "`Draw them on a colour` on the Redraw tab, which it was split onto for crowding the File tab (app v0.32.0), four grounds by name in the chooser beside it, on convert.Redraw with Background — ⛔ handed to the RENDERER rather than composited after, which convert learned the hard way: \"A first version painted the colour under the finished raster — which does nothing at all, because the renderer has already filled the page with opaque white\". ⚠ rasterises"},
	{"edit", "Change Text Color", have, "`Paint the text` on the Pages tab (app v0.29.0), on ops.Doc.RecolourText — ⛔ and it passes on UnreachedText: text drawn through a Form XObject is shared with every page that draws it, so it is left alone and the status line SAYS how much, because a document that came back looking unchanged with nothing said would be worse than a refusal"},
	{"edit", "Flatten PDF", have, "`Flatten the annotations in`"},
	{"edit", "Remove Annotations", have, "`Drop every annotation`"},
	{"edit", "Remove Blank Pages", have, "`Drop the blank pages`"},
	{"edit", "Edit Bookmarks", have, "app v0.40.0, Contents tab: the rail shows them, `Bookmark this page` adds one to the tree already there (titled with the selected passage, pointing at the page THAT is on), `Bookmarks from a file` writes a whole tree through ops.ParseOutline — the same file `pdfops outline -from` takes — and `Drop the bookmarks` removes them"},
	{"edit", "Add Stamps", have, "`Stamp them`"},
	{"edit", "Table of Contents", have, "`Write a contents page` on the Contents tab (app v0.43.0, on ops.Doc.ContentsPage new in ops v0.34.0): a page with a line per bookmark, its number at the right and dots leading to it, built from the entries the rail shows — and `pdfops contents -from toc.txt [-bookmarks]` writes it from the same text file `outline -from` takes. ⛔ The printed numbers are MOVED BY THE PAGES IT ADDS (an entry on page 7 is printed as 8), ShiftBookmarks is the same rule for the tree with its preface edge, and a title with no room at all is a REFUSAL naming the entry rather than a line showing a page number with nothing to read beside it — which it was until cmd/pdfops' own 100-point fixture drew six of them and reported success"},
	{"edit", "Redact Content", partly, redaction},
	{"edit", "Scanner Effect", have, "`Make them look scanned` on the Redraw tab, which it was split onto for crowding the File tab (app v0.32.0), on convert.Redraw with ScannerEffect — ⛔ at a FIXED seed, which is convert's own default and its own reason (\"two runs over the same file must produce the same bytes\"), and the test that gates it compares RENDERED PIXELS rather than doc.Bytes(), which carries the write's own date and so differs between two runs of the same verb. ⚠ rasterises"},
	{"edit", "Adjust Colors", have, "`Brighten or deepen them` on the Redraw tab, which it was split onto for crowding the File tab (app v0.32.0), on convert.Redraw — ⛔ the two spin buttons are read as PERCENTAGES, so 100 % contrast means \"leave it alone\" (convert's 1), and a zero never reaches convert, which treats it as 1 because 0 \"would turn every page into a flat grey rectangle\". ⚠ rasterises"},

	// ---- Automate (1) ----
	{"automate", "PDF Workflow Builder", missing, asked + " for \"workflow\""},

	// ---- Convert to PDF (33) ----
	// ⛔ THE WORKBENCH ALREADY DOES THESE, and I filed them as unreachable by
	// looking at the ribbon instead of at the code. `Open` falls back to
	// ops.Doc.Picture when the bytes are not a PDF (app scene.go, "Somebody who
	// hands a picture to a PDF workbench wants a PDF of it"), so handing it a
	// picture gives a one-page document and the status line says "this was a
	// picture, and is now a page". There is no verb for it because it needs
	// none.
	{"to pdf", "Image to PDF", have, "`Open` a picture: ops.Doc.Picture makes it a one-page document — ⛔⛔ and until app v0.33.0 THE FILE DIALOG WOULD NOT SHOW ONE. web/index.html carried `accept=\"application/pdf,.pdf\"` while this arm had been there since the workbench could open anything: the code was right, every test passed because a test hands bytes straight to open(), and a person standing at the dialog could not choose a PNG. This row and the five below it were counted as delivered for days behind a picker that hid their files. The list now lives in Go (openSuffixes) with a test that reads the HTML and compares the two IN BOTH DIRECTIONS"},
	{"to pdf", "JPG to PDF", have, "`Open` — codec decodes JPEG; reachable from the dialog since app v0.33.0, see Image to PDF"},
	{"to pdf", "PNG to PDF", have, "`Open` — codec decodes PNG; reachable from the dialog since app v0.33.0, see Image to PDF"},
	{"to pdf", "WebP to PDF", have, "`Open` — codec decodes WEBP; reachable from the dialog since app v0.33.0, see Image to PDF"},
	{"to pdf", "SVG to PDF", have, "`Open` a drawing (app v0.33.0), on convert.SVGToPDFVector — ⛔ it asks for the VECTOR form first although convert.SVGToPDF would fall back by itself, because the fallback is the right behaviour and convert still cannot tell the person at the workbench WHICH of the two they got: a page of paths scales and keeps its text, a page that was drawn is pixels. So the status line says either \"stays sharp at any size\" or warns that the page is a picture of the drawing"},
	{"to pdf", "BMP to PDF", have, "`Open` — codec decodes BMP; reachable from the dialog since app v0.33.0, see Image to PDF"},
	{"to pdf", "HEIC to PDF", missing, "codec names PNG, JPEG, GIF, WEBP, TIFF, BMP, ICO, ICNS, PNM, QOI, JP2 and JBIG2; HEIC is not among them"},
	{"to pdf", "TIFF to PDF", have, "`Open` — codec decodes TIFF; reachable from the dialog since app v0.33.0, see Image to PDF"},
	{"to pdf", "PSD to PDF", elsewhere, "`convert.Readers()` answers \"bmp gif jpeg png psd tiff webp\" and Open goes through ops.Picture, which uses codec and does not read PSD — ⛔ PSD is the ONLY one of convert's seven that codec lacks, so the arm would be four lines. It is held back for want of a WITNESS: nothing in the fleet WRITES a PSD, so the arm could only be tested through a faked seam, and a row claimed on a seam that never saw a real file of the format is the \"the filename is not the content\" defect with a test suite in front of it. The honest next step is a PSD encoder in go-images/psd, which the fleet wants anyway"},
	{"to pdf", "Word to PDF", elsewhere, "go-ooxml reads it into richdoc; " + toPDF},
	{"to pdf", "Excel to PDF", elsewhere, "go-ooxml reads it into richdoc; " + toPDF},
	{"to pdf", "PowerPoint to PDF", elsewhere, "go-ooxml reads it into richdoc; " + toPDF},
	{"to pdf", "ODT to PDF", elsewhere, "go-odf reads it into richdoc; " + toPDF},
	{"to pdf", "ODS to PDF", elsewhere, "go-odf reads it into richdoc; " + toPDFnoSuffix},
	{"to pdf", "ODP to PDF", elsewhere, "go-odf reads it into richdoc; " + toPDFnoSuffix},
	{"to pdf", "ODG to PDF", elsewhere, "go-odf reads it into richdoc; " + toPDFnoSuffix},
	{"to pdf", "RTF to PDF", partly, "`doc2pdf report.rtf` (go-pdfkit/doc2pdf v0.3.0, reading through go-richdoc/rtf v0.1.2) converts it, keeping headings, lists, tables, links, footnotes, PNG and JPEG pictures, bold, italic, strike and alignment. ⛔ The gap: BentoPDF goes through LibreOffice and keeps the LAYOUT, while the richdoc model (go-richdoc v0.10.0) has no place for font size, colour, underline, fonts, indents, headers or footers, so those are dropped. Input is bounded (16 MiB for RTF, `--max-input` to change it)"},
	{"to pdf", "CSV to PDF", elsewhere, "go-richdoc/data reads CSV into richdoc; " + toPDF},
	{"to pdf", "Markdown to PDF", elsewhere, "go-richdoc/markdown reads it into richdoc; " + toPDF},
	{"to pdf", "Text to PDF", elsewhere, "go-richdoc/data reads plain text into richdoc; " + toPDF},
	{"to pdf", "JSON to PDF", elsewhere, "go-richdoc/data reads JSON into richdoc; " + toPDF},
	{"to pdf", "XML to PDF", elsewhere, "go-richdoc/data reads XML into richdoc; " + toPDF},
	{"to pdf", "EPUB to PDF", elsewhere, "go-richdoc/epub reads it into richdoc; " + toPDF},
	{"to pdf", "MOBI to PDF", elsewhere, "go-richdoc/mobi reads it into richdoc; " + toPDF},
	{"to pdf", "FB2 to PDF", elsewhere, "go-richdoc/fb2 reads it into richdoc; " + toPDF},
	{"to pdf", "CBZ to PDF", have, "`Open` a comic archive (app v0.33.0), on convert.ArchiveToPDF, which orders the pages the way their NAMES do (page10 after page2) — ⛔ recognised by CONTENT (PK\\x03\\x04) and not by the name, and the archive arm is LAST because a .docx and an .odt are ZIPs too: what reaches the status line for one of those is convert's own \"no pictures in the archive: N entries held nothing readable\" rather than a guess at what was meant. convert's three ceilings are exercised by a test that LOWERS them and checks the refusal arrives naming the bound"},
	// ⛔⛔ THIS ROW SAID `elsewhere` AND NAMED A FUNCTION THAT DOES NOT EXIST.
	// `convert.XPSToPDF` is in no released version of go-pdfkit/convert — v0.16.0
	// is the latest and its API is ArchiveToPDF, Redraw, SVGToPDF,
	// SVGToPDFVector, ToArchive, ToPDF, ReadersToPDF and the raster helpers.
	// Checked 2026-11-08 with `go doc` against the module, and
	// `go list -m -versions` for the tags.
	//
	// What does exist: go-ooxml/opc reads the OPC container "behind docx, xlsx,
	// pptx and xps" — its own description — so the PACKAGE can be opened, and
	// the spec (ECMA-388) is in the bibliography at
	// ~/Documents/VCS/GIT/localhost/specs-openxps. Nothing turns XPS markup
	// into a page. A container reader is not a converter, the same way go-odf
	// reading ODF is not ODF to PDF.
	//
	// An `elsewhere` naming something that is not there is worse than a
	// `missing`: it understates the gap, and it understates it in the
	// direction that looks like progress.
	{"to pdf", "XPS to PDF", missing, "go-ooxml/opc reads the OPC container behind xps, and nothing turns the markup into a page; `convert.XPSToPDF` does not exist in v0.16.0 or any earlier tag, which is what this row used to cite; the spec is ECMA-388"},
	{"to pdf", "Email to PDF", elsewhere, "go-richdoc/email reads it into richdoc; " + toPDF},
	{"to pdf", "Pages to PDF", missing, asked + "; no reader for Apple's format"},
	{"to pdf", "WPD to PDF", missing, asked + " for \"wpd\": not one mention in thirty repositories"},
	{"to pdf", "WPS to PDF", missing, asked + "; no reader for Kingsoft's format"},
	{"to pdf", "PUB to PDF", missing, asked + "; no reader for Publisher's format"},
	{"to pdf", "VSD to PDF", missing, asked + " for \"vsd\" and \"visio\": every hit is a substring of MOVSD, revision or division"},

	// ---- Convert from PDF (15) ----
	{"from pdf", "PDF to Image", have, "`Hand over this page`, `Hand over every page, zipped`"},
	{"from pdf", "PDF to JPG", have, "the picture chooser offers it"},
	{"from pdf", "PDF to PNG", have, "the picture chooser's first offer"},
	{"from pdf", "PDF to WebP", have, "the picture chooser offers it"},
	{"from pdf", "PDF to BMP", have, "the picture chooser offers it"},
	{"from pdf", "PDF to TIFF", have, "the picture chooser offers it"},
	{"from pdf", "PDF to CBZ", have, "`Hand it over as a comic` in `Hand over` on the Read tab (app v0.32.0), on convert.ToArchive — ⛔ under .cbz and NOT the .zip the neighbouring `Hand over every page, zipped` writes: the same pictures under the name the format is known by, which is the whole difference between a file a comic reader opens and one somebody has to rename first"},
	{"from pdf", "PDF to SVG", elsewhere, "`pdfconv to-svg [-pages 1,3,5-9] [-password-file f] [-o dir] in.pdf` (go-pdfkit/convert v0.17.0; API StreamSVG, PDFToSVG, NewSVGArchive, SVGArchive), on render.PageSVG: paths, clips, images and transparency groups become SVG elements and text becomes outlines, so the file needs no font. Named the way BentoPDF names it: one page is `name.svg`, several are `name_svg.zip` holding `page_N.svg`. ⛔ NOT ON THE RIBBON: no workbench verb calls it, so this is `elsewhere` and not `have`. ⚠ Not the whole of BentoPDF's row: it takes SEVERAL PDFs at once into one `pdf_to_svg.zip`, and here one PDF is one run. ⚠ All or nothing, bounded before it is spent (1000 pages, 64 MiB a page, 256 MiB and 5 minutes a run), and the time limit is a request while go-pdfkit/render#142 is open; clips nested past 64 are refused until go-pdfkit/render#143 lands a raster fallback. This row used to be `missing` on a sweep made before v0.17.0"},
	{"from pdf", "PDF to Greyscale", have, "`Drop the colour` in `Redraw the pixels` on the Redraw tab, which it was split onto for crowding the File tab (app v0.32.0), on convert.Redraw with Greyscale — ⛔ and the test measures the GREYNESS of the rendered pixels (how far red, green and blue stand apart) rather than trusting the option's name. ⚠ rasterises"},
	{"from pdf", "PDF to Text", have, "`What this page says`, `Hand over what this page says`, on go-pdfkit/extract"},
	{"from pdf", "PDF to JSON", have, "`Hand over the page as data` on the Read tab (app v0.30.0): every piece of text WITH WHERE IT SITS, and the grids its words line up into — the positions being the whole reason a page is worth having as data rather than as the text file beside it"},
	{"from pdf", "PDF to CSV", have, "`Hand over the grids` on the Read tab (app v0.30.0), on extract.Tables and encoding/csv — ⛔ with a MARKER ROW between grids rather than a blank line: csv skips blank lines on read, so `pdfops tables -as csv` separates nothing while its comment says it does (go-pdfkit/ops#71)"},
	{"from pdf", "PDF to Excel", elsewhere, "`pdfops excel` — \"write the pages' tables out as a workbook\". ⛔ NOT on the ribbon ON PURPOSE, and the reason is measured rather than editorial: writing .xlsx means go-ooxml's writer, which cost 3.50 MB as an increment into this wasm binary when it was measured on 2026-11-08 — about 15 % more for every reader of the workbench, to serve one row. The same grids already come out as CSV and as JSON, which every spreadsheet opens, through a package the binary already links (extract). The trade is the row's whole content: it is reachable from the command line, where a megabyte costs nobody anything"},
	{"from pdf", "Extract Tables", have, "`What lines up into a grid` on the Read tab (app v0.30.0), on extract.Tables — ⛔ NOT through go-pdfkit/document, which the first reading of this row cited: that package costs about 33.5 MB in a wasm binary in EITHER direction, because it imports html2pdf and so the browser engine (measured 2026-11-08: 3.50 MB for ooxml's writer alone, 37.05 MB once document.Parse is named). extract has been linked since the Read tab could read text, so this row cost nothing"},
	{"from pdf", "OCR PDF", missing, "the fleet has no OCR engine"},

	// ---- Secure & Optimize PDFs (16) ----
	{"secure", "Compress PDF", have, "`Pack it smaller`"},
	{"secure", "Repair PDF", have, "`Show what it says` on the File tab (app v0.28.0) says whether the structure had to be rebuilt to read the file, and that Saving writes the rebuilt one — ⛔ repair is something a reader does on the way IN, so the honest verb is to report it and point at Save rather than to offer a button that has already happened"},
	{"secure", "Encrypt PDF", have, "`Protect it`"},
	{"secure", "Decrypt PDF", have, "`Take the protection off`"},
	{"secure", "Change Permissions", have, "`Protect it` allows what is ticked"},
	{"secure", "Sign PDF", have, "BentoPDF's row is \"Draw, type, or upload your signature\", and the workbench now does all three. DRAW: the Sign tab (playground v0.59.0, playground#118) takes a stroke from a mouse, pen or touch on toolkit.InkPad (go-widgets/toolkit v0.396.0), with undo and clear, then places it on the chosen pages (an empty list means all pages), at a corner and a width, through ops.Place as a transparent PNG. UPLOAD: `Draw a picture on them` (app v0.34.0) places an image through ops.Place too (also `pdfops place`). TYPE: the Stamp tab writes a line of text where you say. There is no flatten option: a placed picture is already page content, and flatten applies to annotations. Known limit: a stroke dragged sideways out of the panel stops at its edge (go-widgets/toolkit#616), which does not stop a signature being drawn. ⚠ This is a drawn signature, not a cryptographic one: that is the Digital Signature row"},
	{"secure", "Digital Signature", elsewhere, "go-pdfkit/ops v0.39.0 SIGNS: `ops.Sign` and `pdfops sign` add a digital signature as an incremental update, so every earlier signature still verifies, with a detached CMS from go-authn/cms (RSA, RSA-PSS, ECDSA, Ed25519) under `adbe.pkcs7.detached` or `ETSI.CAdES.detached`; the key and certificate come from PEM files. Since ops v0.40.0 the signature can be VISIBLE: `SignOptions.Appearance` and `pdfops sign -visible x0,y0,x1,y1 [-image file] [-text words]` give the widget a box, a picture and Latin text inside the signed bytes. ⚠ The text is WinAnsi only: a Cyrillic or CJK character is refused by name (ErrAppearanceText), not dropped. Judged by poppler pdfsig 26.10 (RSA and ECDSA valid, the last one \"Total document signed\"), reader/signature and `openssl cms -verify`. ⚠ pdfsig reports an Ed25519 signature as Invalid while the other two accept the same bytes. ⛔ NOT ON THE RIBBON: no workbench verb calls it. ⚠ Not the whole of BentoPDF's wording, \"X.509 certificates (PFX/PEM), private key never leaves browser\": ops reads PEM files, and it is a library and a command, not a browser page. Not claimed by ops: a PAdES baseline, timestamps or LTV, an encrypted file (refused with ErrEncrypted). ⚠ This row used to be `missing` on the finding that nothing assembled a CMS signature into a PDF; go-authn/cms and the /ByteRange plumbing of ops.Sign are that"},
	{"secure", "Validate Signature", elsewhere, "go-pdfkit/reader v0.20.0 READS and VERIFIES PDF digital signatures: `Document.Signatures` lists the signed fields, `Document.SignedBytes` returns the bytes a signature covers after validating its /ByteRange, and `reader/signature` (`Verify`, `VerifyAll`) checks `adbe.pkcs7.detached` and `ETSI.CAdES.detached` on top of go-authn/cms. Checked against poppler pdfsig 26.10.0 on OpenSSL-made fixtures: RSA, ECDSA and CAdES valid; an incremental update leaves the earlier signature valid and not whole-file; two stacked signatures; seven tamperings refused by both. One deliberate difference: pdfsig calls a SHA-1 signature valid and this package refuses it by name. ⛔ NOT ON THE RIBBON: no workbench verb calls it. ⚠ And not the whole of BentoPDF's wording, \"check certificate validity\": trust — the chain, the validity period at signing time, revocation — is the CALLER's, as in go-authn/cms; what is covered is the signature and the integrity of the signed bytes. ⚠ This row used to be `missing` on the finding that no CMS SignedData parser existed in the fleet (an external review of 2 584 clones); go-authn/cms is one, and the reader's tests run it against pdfsig. Verified per document at most 64 distinct signatures and 1 GiB of signed bytes"},
	{"secure", "Redact Content", partly, redaction + " — BentoPDF lists this tool twice, and so does this table, because its own tables do"},
	{"secure", "Edit Metadata", have, "`Show what it says` lists every entry and COUNTS whatever else the file carries (app v0.28.0 — not showing is what made writing destructive: the two boxes started empty whatever the file said, so opening a paper with a title and an author and pressing the button deleted both, silently); `Say so in the file` writes SIX entries since app v0.41.0 — title, author, subject, keywords, creator, producer — where it wrote two while displaying four more as read-only labels. ⛔ The boxes are filled from the UNCUT value: infoSays clips to 120 runes for a LABEL, and a box filled from that and written back would cut the file's own value down by way of showing it; an entry longer than that is shown rather than offered and written back untouched. The two dates stay shown: a PDF date is D:YYYYMMDDHHmmSSOHH'mm (ISO 32000-1, 7.9.4), nothing in the fleet parses one, and a box to type a date into is a box in which somebody writes what no reader will show"},
	{"secure", "Remove Metadata", have, "`Drop what it says about itself`"},
	{"secure", "Linearize PDF", partly, "`reader.Document.Linearize` WRITES one since reader v0.16.0: Annex F's eleven parts, both cross-reference sections, the parameter dictionary, and a primary hint stream holding Table F.3's page offset table and Table F.5's shared object hint table. `qpdf --check` reports nothing about the output it does not also report about the input, on 1, 2, 5 and 25 page documents AND on a one-page document inheriting an indirect /Resources — that last shape being the one that catches the inheritance regression, since /E and the per-page object count were wrong on it until reachFromPage followed Table 30. The test's hint-table exemption was DELETED rather than satisfied, so any new complaint fails the suite. The checker came first on purpose, which is why reader.CheckLinearized (v0.11.0) could judge every parameter against the document before a writer existed. ⭐ SINCE reader v0.18.0 THE SHARING REFUSAL IS GONE, and with it the reason this row carried from the day it was written. F.4.2's shared object hint table is filled: two runs, as Table F.5 describes — every first-page object counted POSITIONALLY from /O, then the objects shared only between later pages — so a document whose pages share anything is WRITTEN. That is the ordinary multi-page PDF, because Table 30 lets a page inherit /Resources, /MediaBox, /CropBox and /Rotate and one resource dictionary on the pages node is reached by every page. Measured over eight shapes against qpdf 12.4.2, each diffed against ITS OWN input: private 1/2/5/25, a shared content stream, 2 and 25 pages inheriting one /Resources, and pages 2 and 3 sharing an object page 1 never reaches — the only shape that fills part 8. ⛔⛔ STILL `partly`, AND FOR SIX NAMED REASONS RATHER THAN THE OLD ONE. Linearize refuses a document carrying /OpenAction (F.3.7 makes the page it opens at the FIRST page), /PageMode /UseOutlines with /Outlines (F.3.7 then puts the whole hierarchy in part 6), /Threads (F.3.5 part 4, F.3.10 their /I in part 9, beads with their pages), /AcroForm (F.3.5 takes the top-level dictionary, F.3.10 the field hierarchy), /Encrypt (F.3.5 requires the dictionary AND every value in it in part 4), and a page with /Thumb (F.3.10 gives thumbnails two ordered categories of their own). None of those placements is exercised by any file in the corpus, and a WRONG placement is silent — the reader does not fail, it reads the wrong bytes — so they are refused rather than guessed. ⚠ A form, an encrypted file and a scan that carries page thumbnails are each ordinary, so the row stays open. ⚠ And three values are deliberately coarse rather than absent: F.4.1's items 5 and 13 are written 0/1 (a reader fetches a shared object too early, never too late), every group describes ONE object where F.4.2 allows a run, and a page BORROWING a content stream gets 0 for items 6 and 7, since the stream lies before the page and the items are unsigned. ⚠ reader.CheckLinearized recomputes the page and group LENGTHS and not Table F.5's first shared object offset; qpdf does, and that is what caught a real defect in the writer while our own checker stayed silent. go-pdfkit/reader#68 (go-pdfkit/reader#40, #61, #66, #67)"},
	{"secure", "Sanitize PDF", have, "`Sanitize — drop what runs`"},
	{"secure", "Fix Page Size", have, "`Set the paper` on the Pages tab (app v0.29.0), on ops.Doc.Resize, with six sheets by name — ⛔ at the standard's numbers (A4 is 595.276 × 841.89, not 595 × 842: a tenth of a millimetre shows as a hairline against a trimmed edge)"},
	{"secure", "Page Dimensions", have, "`Show what it says` on the File tab (app v0.28.0) lists every page's size in points and how it is turned — ⛔ the size it will be SHOWN at, crop box over media box, because a readout quoting the paper of a cropped page would disagree with the picture beside it"},
	{"secure", "Remove Restrictions", have, "`Take the protection off` writes the file with no protection at all. ⛔ The earlier verdict said \"it needs the password first\", which is true of a file with a USER password and false of the case this tool is FOR: a file protected with an OWNER password only opens for anybody — the user password is empty — and what the owner password buys is a list of things a reader is asked not to do. Removing those needs no password, because the file was already open. TestTakingTheRestrictionsOffAFileThatOpensForAnybody (app v0.46.0) builds exactly that file, opens it with nothing, presses the verb and reads the saved bytes back unprotected. A file nobody can OPEN without a password is encryption rather than a restriction, and that one does need it"},
}

// ours is what the workbench has that BentoPDF's tables do not list. It is here
// so that "the gap" is not read as one-directional.
var ours = []string{
	"HTML to PDF, vector rather than rasterised (go-pdfkit/html2pdf)",
	"LaTeX source reconstructed from a PDF (go-pdfkit/latex)",
	"a PDF several people edit at once, as a CRDT (go-pdfkit/coedit)",
	"a page read back as a DOCUMENT — headings, paragraphs, lists, tables (go-pdfkit/document)",
	"a PDF read back as Markdown or as a Word document (`pdfops markdown`, `pdfops word`)",
	"the optional content a viewer calls layers, listed, shown and hidden (`pdfops layers`)",
	"page labels set as rules — i, ii, iii then 1, 2, 3 (`pdfops labels`)",
	"what a page carries: the pictures on it, listed and handed over one by one",
	"what the file is protected with, said in words before anything is changed",
	"the whole thing in one wasm binary with no network at all, not a page that loads libraries from a CDN",
}
