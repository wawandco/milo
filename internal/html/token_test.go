// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type tokenTest struct {
	// A short description of the test case.
	desc string
	// The HTML to parse.
	html string
	// The string representations of the expected tokens, joined by '$'.
	golden string
}

var tokenTests = []tokenTest{
	{
		"empty",
		"",
		"",
	},
	// A single text node. The tokenizer should not break text nodes on whitespace,
	// nor should it normalize whitespace within a text node.
	{
		"text",
		"foo  bar",
		"foo  bar",
	},
	// An entity.
	{
		"entity",
		"one &lt; two",
		"one &lt; two",
	},
	// A start, self-closing and end tag. The tokenizer does not care if the start
	// and end tokens don't match; that is the job of the parser.
	{
		"tags",
		"<a>b<c/>d</e>",
		"<a>$b$<c/>$d$</e>",
	},
	// Angle brackets that aren't a tag.
	{
		"not a tag #0",
		"<",
		"&lt;",
	},
	{
		"not a tag #1",
		"</",
		"&lt;/",
	},
	{
		"not a tag #2",
		"</>",
		"<!---->",
	},
	{
		"not a tag #3",
		"a</>b",
		"a$<!---->$b",
	},
	{
		"not a tag #4",
		"</ >",
		"<!-- -->",
	},
	{
		"not a tag #5",
		"</.",
		"<!--.-->",
	},
	{
		"not a tag #6",
		"</.>",
		"<!--.-->",
	},
	{
		"not a tag #7",
		"a < b",
		"a &lt; b",
	},
	{
		"not a tag #8",
		"<.>",
		"&lt;.&gt;",
	},
	{
		"not a tag #9",
		"a<<<b>>>c",
		"a&lt;&lt;$<b>$&gt;&gt;c",
	},
	{
		"not a tag #10",
		"if x<0 and y < 0 then x*y>0",
		"if x&lt;0 and y &lt; 0 then x*y&gt;0",
	},
	{
		"not a tag #11",
		"<<p>",
		"&lt;$<p>",
	},
	// EOF in a tag name.
	{
		"tag name eof #0",
		"<a",
		"",
	},
	{
		"tag name eof #1",
		"<a ",
		"",
	},
	{
		"tag name eof #2",
		"a<b",
		"a",
	},
	{
		"tag name eof #3",
		"<a><b",
		"<a>",
	},
	{
		"tag name eof #4",
		`<a x`,
		``,
	},
	// Some malformed tags that are missing a '>'.
	{
		"malformed tag #0",
		`<p</p>`,
		`<p< p="">`,
	},
	{
		"malformed tag #1",
		`<p </p>`,
		`<p <="" p="">`,
	},
	{
		"malformed tag #2",
		`<p id`,
		``,
	},
	{
		"malformed tag #3",
		`<p id=`,
		``,
	},
	{
		"malformed tag #4",
		`<p id=>`,
		`<p id="">`,
	},
	{
		"malformed tag #5",
		`<p id=0`,
		``,
	},
	{
		"malformed tag #6",
		`<p id=0</p>`,
		`<p id="0&lt;/p">`,
	},
	{
		"malformed tag #7",
		`<p id="0</p>`,
		``,
	},
	{
		"malformed tag #8",
		`<p id="0"</p>`,
		`<p id="0" <="" p="">`,
	},
	{
		"malformed tag #9",
		`<p></p id`,
		`<p>`,
	},
	// Raw text and RCDATA.
	{
		"basic raw text",
		"<script><a></b></script>",
		"<script>$&lt;a&gt;&lt;/b&gt;$</script>",
	},
	{
		"unfinished script end tag",
		"<SCRIPT>a</SCR",
		"<script>$a&lt;/SCR",
	},
	{
		"broken script end tag",
		"<SCRIPT>a</SCR ipt>",
		"<script>$a&lt;/SCR ipt&gt;",
	},
	{
		"EOF in script end tag",
		"<SCRIPT>a</SCRipt",
		"<script>$a&lt;/SCRipt",
	},
	{
		"scriptx end tag",
		"<SCRIPT>a</SCRiptx",
		"<script>$a&lt;/SCRiptx",
	},
	{
		"' ' completes script end tag",
		"<SCRIPT>a</SCRipt ",
		"<script>$a",
	},
	{
		"'>' completes script end tag",
		"<SCRIPT>a</SCRipt>",
		"<script>$a$</script>",
	},
	{
		"self-closing script end tag",
		"<SCRIPT>a</SCRipt/>",
		"<script>$a$</script>",
	},
	{
		"nested script tag",
		"<SCRIPT>a</SCRipt<script>",
		"<script>$a&lt;/SCRipt&lt;script&gt;",
	},
	{
		"script end tag after unfinished",
		"<SCRIPT>a</SCRipt</script>",
		"<script>$a&lt;/SCRipt$</script>",
	},
	{
		"script/style mismatched tags",
		"<script>a</style>",
		"<script>$a&lt;/style&gt;",
	},
	{
		"style element with entity",
		"<style>&apos;",
		"<style>$&amp;apos;",
	},
	{
		"textarea with tag",
		"<textarea><div></textarea>",
		"<textarea>$&lt;div&gt;$</textarea>",
	},
	{
		"title with tag and entity",
		"<title><b>K&amp;R C</b></title>",
		"<title>$&lt;b&gt;K&amp;R C&lt;/b&gt;$</title>",
	},
	{
		"title with trailing '&lt;' entity",
		"<title>foobar<</title>",
		"<title>$foobar&lt;$</title>",
	},
	// DOCTYPE tests.
	{
		"Proper DOCTYPE",
		"<!DOCTYPE html>",
		"<!DOCTYPE html>",
	},
	{
		"DOCTYPE with no space",
		"<!doctypehtml>",
		"<!DOCTYPE html>",
	},
	{
		"DOCTYPE with two spaces",
		"<!doctype  html>",
		"<!DOCTYPE html>",
	},
	{
		"looks like DOCTYPE but isn't",
		"<!DOCUMENT html>",
		"<!--DOCUMENT html-->",
	},
	{
		"DOCTYPE at EOF",
		"<!DOCtype",
		"<!DOCTYPE >",
	},
	// XML processing instructions.
	{
		"XML processing instruction",
		"<?xml?>",
		"<!--?xml?-->",
	},
	// Comments.
	{
		"comment0",
		"abc<b><!-- skipme --></b>def",
		"abc$<b>$<!-- skipme -->$</b>$def",
	},
	{
		"comment1",
		"a<!-->z",
		"a$<!---->$z",
	},
	{
		"comment2",
		"a<!--->z",
		"a$<!---->$z",
	},
	{
		"comment3",
		"a<!--x>-->z",
		"a$<!--x>-->$z",
	},
	{
		"comment4",
		"a<!--x->-->z",
		"a$<!--x->-->$z",
	},
	{
		"comment5",
		"a<!>z",
		"a$<!---->$z",
	},
	{
		"comment6",
		"a<!->z",
		"a$<!----->$z",
	},
	{
		"comment7",
		"a<!---<>z",
		"a$<!---<>z-->",
	},
	{
		"comment8",
		"a<!--z",
		"a$<!--z-->",
	},
	{
		"comment9",
		"a<!--z-",
		"a$<!--z-->",
	},
	{
		"comment10",
		"a<!--z--",
		"a$<!--z-->",
	},
	{
		"comment11",
		"a<!--z---",
		"a$<!--z--->",
	},
	{
		"comment12",
		"a<!--z----",
		"a$<!--z---->",
	},
	{
		"comment13",
		"a<!--x--!>z",
		"a$<!--x-->$z",
	},
	// An attribute with a backslash.
	{
		"backslash",
		`<p id="a\"b">`,
		`<p id="a\" b"="">`,
	},
	// Entities, tag name and attribute key lower-casing, and whitespace
	// normalization within a tag.
	{
		"tricky",
		"<p \t\n iD=\"a&quot;B\"  foo=\"bar\"><EM>te&lt;&amp;;xt</em></p>",
		`<p id="a&#34;B" foo="bar">$<em>$te&lt;&amp;;xt$</em>$</p>`,
	},
	// A nonexistent entity. Tokenizing and converting back to a string should
	// escape the "&" to become "&amp;".
	{
		"noSuchEntity",
		`<a b="c&noSuchEntity;d">&lt;&alsoDoesntExist;&`,
		`<a b="c&amp;noSuchEntity;d">$&lt;&amp;alsoDoesntExist;&amp;`,
	},
	{
		"entity without semicolon",
		`&notit;&notin;<a b="q=z&amp=5&notice=hello&not;=world">`,
		`¬it;∉$<a b="q=z&amp;amp=5&amp;notice=hello¬=world">`,
	},
	{
		"entity with digits",
		"&frac12;",
		"½",
	},
	// Attribute tests:
	// http://dev.w3.org/html5/pf-summary/Overview.html#attributes
	{
		"Empty attribute",
		`<input disabled FOO>`,
		`<input disabled="" foo="">`,
	},
	{
		"Empty attribute, whitespace",
		`<input disabled FOO >`,
		`<input disabled="" foo="">`,
	},
	{
		"Unquoted attribute value",
		`<input value=yes FOO=BAR>`,
		`<input value="yes" foo="BAR">`,
	},
	{
		"Unquoted attribute value, spaces",
		`<input value = yes FOO = BAR>`,
		`<input value="yes" foo="BAR">`,
	},
	{
		"Unquoted attribute value, trailing space",
		`<input value=yes FOO=BAR >`,
		`<input value="yes" foo="BAR">`,
	},
	{
		"Single-quoted attribute value",
		`<input value='yes' FOO='BAR'>`,
		`<input value="yes" foo="BAR">`,
	},
	{
		"Single-quoted attribute value, trailing space",
		`<input value='yes' FOO='BAR' >`,
		`<input value="yes" foo="BAR">`,
	},
	{
		"Double-quoted attribute value",
		`<input value="I'm an attribute" FOO="BAR">`,
		`<input value="I&#39;m an attribute" foo="BAR">`,
	},
	{
		"Attribute name characters",
		`<meta http-equiv="content-type">`,
		`<meta http-equiv="content-type">`,
	},
	{
		"Mixed attributes",
		`a<P V="0 1" w='2' X=3 y>z`,
		`a$<p v="0 1" w="2" x="3" y="">$z`,
	},
	{
		"Attributes with a solitary single quote",
		`<p id=can't><p id=won't>`,
		`<p id="can&#39;t">$<p id="won&#39;t">`,
	},
}

func TestTokenizer(t *testing.T) {
loop:
	for _, tt := range tokenTests {
		z := NewTokenizer(strings.NewReader(tt.html))
		if tt.golden != "" {
			for i, s := range strings.Split(tt.golden, "$") {
				if z.Next() == ErrorToken {
					t.Errorf("%s token %d: want %q got error %v", tt.desc, i, s, z.Err())
					continue loop
				}
				actual := z.Token().String()
				if s != actual {
					t.Errorf("%s token %d: want %q got %q", tt.desc, i, s, actual)
					continue loop
				}
			}
		}
		z.Next()
		if z.Err() != io.EOF {
			t.Errorf("%s: want EOF got %q", tt.desc, z.Err())
		}
	}
}

func TestMaxBuffer(t *testing.T) {
	// Exceeding the maximum buffer size generates ErrBufferExceeded.
	z := NewTokenizer(strings.NewReader("<" + strings.Repeat("t", 10)))
	z.SetMaxBuf(5)
	tt := z.Next()
	if got, want := tt, ErrorToken; got != want {
		t.Fatalf("token type: got: %v want: %v", got, want)
	}
	if got, want := z.Err(), ErrBufferExceeded; got != want {
		t.Errorf("error type: got: %v want: %v", got, want)
	}
	if got, want := string(z.Raw()), "<tttt"; got != want {
		t.Fatalf("buffered before overflow: got: %q want: %q", got, want)
	}
}

func TestMaxBufferReconstruction(t *testing.T) {
	// Exceeding the maximum buffer size at any point while tokenizing permits
	// reconstructing the original input.
tests:
	for _, test := range tokenTests {
		for maxBuf := 1; ; maxBuf++ {
			r := strings.NewReader(test.html)
			z := NewTokenizer(r)
			z.SetMaxBuf(maxBuf)
			var tokenized bytes.Buffer
			for {
				tt := z.Next()
				tokenized.Write(z.Raw())
				if tt == ErrorToken {
					if err := z.Err(); err != io.EOF && err != ErrBufferExceeded {
						t.Errorf("%s: unexpected error: %v", test.desc, err)
					}
					break
				}
			}
			// Anything tokenized along with untokenized input or data left in the reader.
			assembled, err := io.ReadAll(io.MultiReader(&tokenized, bytes.NewReader(z.Buffered()), r))
			if err != nil {
				t.Errorf("%s: ReadAll: %v", test.desc, err)
				continue tests
			}
			if got, want := string(assembled), test.html; got != want {
				t.Errorf("%s: reassembled html:\n got: %q\nwant: %q", test.desc, got, want)
				continue tests
			}
			// EOF indicates that we completed tokenization and hence found the max
			// maxBuf that generates ErrBufferExceeded, so continue to the next test.
			if z.Err() == io.EOF {
				break
			}
		} // buffer sizes
	} // tests
}

func TestPassthrough(t *testing.T) {
	// Accumulating the raw output for each parse event should reconstruct the
	// original input.
	for _, test := range tokenTests {
		z := NewTokenizer(strings.NewReader(test.html))
		var parsed bytes.Buffer
		for {
			tt := z.Next()
			parsed.Write(z.Raw())
			if tt == ErrorToken {
				break
			}
		}
		if got, want := parsed.String(), test.html; got != want {
			t.Errorf("%s: parsed output:\n got: %q\nwant: %q", test.desc, got, want)
		}
	}
}

func TestBufAPI(t *testing.T) {
	s := "0<a>1</a>2<b>3<a>4<a>5</a>6</b>7</a>8<a/>9"
	z := NewTokenizer(bytes.NewBufferString(s))
	var result bytes.Buffer
	depth := 0
loop:
	for {
		tt := z.Next()
		switch tt {
		case ErrorToken:
			if z.Err() != io.EOF {
				t.Error(z.Err())
			}
			break loop
		case TextToken:
			if depth > 0 {
				result.Write(z.Text())
			}
		case StartTagToken, EndTagToken:
			tn, _, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
	u := "14567"
	v := result.String()
	if u != v {
		t.Errorf("TestBufAPI: want %q got %q", u, v)
	}
}

func TestConvertNewlines(t *testing.T) {
	testCases := map[string]string{
		"Mac\rDOS\r\nUnix\n":    "Mac\nDOS\nUnix\n",
		"Unix\nMac\rDOS\r\n":    "Unix\nMac\nDOS\n",
		"DOS\r\nDOS\r\nDOS\r\n": "DOS\nDOS\nDOS\n",
		"":                      "",
		"\n":                    "\n",
		"\n\r":                  "\n\n",
		"\r":                    "\n",
		"\r\n":                  "\n",
		"\r\n\n":                "\n\n",
		"\r\n\r":                "\n\n",
		"\r\n\r\n":              "\n\n",
		"\r\r":                  "\n\n",
		"\r\r\n":                "\n\n",
		"\r\r\n\n":              "\n\n\n",
		"\r\r\r\n":              "\n\n\n",
		"\r \n":                 "\n \n",
		"xyz":                   "xyz",
	}
	for in, want := range testCases {
		if got := string(convertNewlines([]byte(in))); got != want {
			t.Errorf("input %q: got %q, want %q", in, got, want)
		}
	}
}

func TestReaderEdgeCases(t *testing.T) {
	const s = "<p>An io.Reader can return (0, nil) or (n, io.EOF).</p>"
	testCases := []io.Reader{
		&zeroOneByteReader{s: s},
		&eofStringsReader{s: s},
		&stuckReader{},
	}
	for i, tc := range testCases {
		got := []TokenType{}
		z := NewTokenizer(tc)
		for {
			tt := z.Next()
			if tt == ErrorToken {
				break
			}
			got = append(got, tt)
		}
		if err := z.Err(); err != nil && err != io.EOF {
			if err != io.ErrNoProgress {
				t.Errorf("i=%d: %v", i, err)
			}
			continue
		}
		want := []TokenType{
			StartTagToken,
			TextToken,
			EndTagToken,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("i=%d: got %v, want %v", i, got, want)
			continue
		}
	}
}

// zeroOneByteReader is like a strings.Reader that alternates between
// returning 0 bytes and 1 byte at a time.
type zeroOneByteReader struct {
	s string
	n int
}

func (r *zeroOneByteReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	r.n++
	if r.n%2 != 0 {
		return 0, nil
	}
	p[0], r.s = r.s[0], r.s[1:]
	return 1, nil
}

// eofStringsReader is like a strings.Reader but can return an (n, err) where
// n > 0 && err != nil.
type eofStringsReader struct {
	s string
}

func (r *eofStringsReader) Read(p []byte) (int, error) {
	n := copy(p, r.s)
	r.s = r.s[n:]
	if r.s != "" {
		return n, nil
	}
	return n, io.EOF
}

// stuckReader is an io.Reader that always returns no data and no error.
type stuckReader struct{}

func (*stuckReader) Read(p []byte) (int, error) {
	return 0, nil
}

const (
	rawLevel = iota
	lowLevel
	highLevel
)

func benchmarkTokenizer(b *testing.B, level int) {
	buf, err := os.ReadFile("testdata/go1.html")
	if err != nil {
		b.Fatalf("could not read testdata/go1.html: %v", err)
	}
	b.SetBytes(int64(len(buf)))
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z := NewTokenizer(bytes.NewBuffer(buf))
		for {
			tt := z.Next()
			if tt == ErrorToken {
				if err := z.Err(); err != nil && err != io.EOF {
					b.Fatalf("tokenizer error: %v", err)
				}
				break
			}
			switch level {
			case rawLevel:
				// Calling z.Raw just returns the raw bytes of the token. It does
				// not unescape &lt; to <, or lower-case tag names and attribute keys.
				z.Raw()
			case lowLevel:
				// Caling z.Text, z.TagName and z.TagAttr returns []byte values
				// whose contents may change on the next call to z.Next.
				switch tt {
				case TextToken, CommentToken, DoctypeToken:
					z.Text()
				case StartTagToken, SelfClosingTagToken:
					_, _, more := z.TagName()
					for more {
						_, _, _, _, more = z.TagAttr()
					}
				case EndTagToken:
					z.TagName()
				}
			case highLevel:
				// Calling z.Token converts []byte values to strings whose validity
				// extend beyond the next call to z.Next.
				z.Token()
			}
		}
	}
}

func TestAttrName(t *testing.T) {
	for _, tc := range []string{
		"HREF",
		"src",
		"Alt",
		"aLT",
		"ClAsS",
	} {
		z := NewTokenizer(strings.NewReader("<a " + tc + "='val' >"))
		z.Next()
		attr := z.Token().Attr
		if got, want := len(attr), 1; got != want {
			t.Errorf("attr len does not match got %d, and want %d", got, want)
		}

		if got, want := attr[0].Name, tc; got != want {
			t.Errorf("attr name does not match got %s, and want %s", got, want)
		}
	}

}

func TestColAndLineNumberForTags(t *testing.T) {
	for _, tc := range []struct {
		Name    string
		Content string
		Line    int
		Col     int
	}{
		{
			Name:    "col=1;line=1",
			Content: "<a>",
			Line:    1,
			Col:     1,
		},

		{
			Name:    "col=2;line1",
			Content: " <a>",
			Line:    1,
			Col:     2,
		},

		{
			Name:    "col=6;line1",
			Content: "blah <a>",
			Line:    1,
			Col:     6,
		},

		{
			Name:    "closed;col=11;line1",
			Content: "blah blah </a>",
			Line:    1,
			Col:     11,
		},

		{
			Name:    "selfclosed;col=11;line1",
			Content: "blah blah <a/>",
			Line:    1,
			Col:     11,
		},

		{
			Name: "selfclosed;col=1;line2",
			Content: `blah blah
<a/>`,
			Line: 2,
			Col:  1,
		},

		{
			Name: "selfclosed;col=11;line2",
			Content: `blah blah
blah blah <a/>`,
			Line: 2,
			Col:  11,
		},
	} {
		z := NewTokenizer(strings.NewReader(tc.Content))
		for {
			z.Next()
			if err := z.Err(); err != nil {
				if err != io.EOF {
					t.Errorf("%s: got error %s", tc.Name, err)
				}
				break
			}

			tok := z.Token()
			if tok.Type == TextToken {
				continue
			}

			if got, want := tok.Col, tc.Col; got != want {
				t.Errorf("%s: got %d, and want %d", tc.Name, got, want)
			}

			if got, want := tok.Line, tc.Line; got != want {
				t.Errorf("%s: got %d, and want %d", tc.Name, got, want)
			}
		}
	}

}

func TestAttrColNumber(t *testing.T) {
	for _, tc := range []struct {
		Name    string
		Content string
		Lines   []int
		Cols    []int
	}{
		{
			Name:    "attribute col = 4",
			Content: "<a href='#'>",
			Lines:   []int{1},
			Cols:    []int{4},
		},

		{
			Name:    "attributes with col = 4 and col = 14",
			Content: "<a href='#'  class=''>",
			Lines:   []int{1, 1},
			Cols:    []int{4, 14},
		},

		{
			Name:    "attributes with col = 10",
			Content: "<a       class=''>",
			Lines:   []int{1},
			Cols:    []int{10},
		},

		{
			Name: "attributes with line = 2 and col = 1",
			Content: `<a 
class=''>`,
			Lines: []int{2},
			Cols:  []int{1},
		},
	} {
		z := NewTokenizer(strings.NewReader(tc.Content))
		for {
			z.Next()
			if err := z.Err(); err != nil {
				if err != io.EOF {
					t.Errorf("%s: got error %s", tc.Name, err)
				}
				break
			}

			tok := z.Token()
			if tok.Type == TextToken {
				continue
			}

			for i := range tok.Attr {
				if got, want := tok.Attr[i].Col, tc.Cols[i]; got != want {
					t.Errorf("cols %s: got %d, and want %d", tc.Name, got, want)
				}

				if got, want := tok.Attr[i].Line, tc.Lines[i]; got != want {
					t.Errorf("lines %s: got %d, and want %d", tc.Name, got, want)
				}
			}
		}
	}

}

func TestTagNameWithERBTemplate(t *testing.T) {
	cases := []struct {
		Content          string
		ExpectedTagName  string
		ExpectedAttrs    []string
		ExpectedAttrsVal []string
		Name             string
	}{
		{
			Content:         "<%= if (true) { %><% } %><a>",
			ExpectedTagName: "a",
			Name:            "before tag definition",
		},
		{
			Content:         "<a<%= if (true) { %><% } %>>",
			ExpectedTagName: "a",
			Name:            "after tag name",
		},
		{
			Content:         "<<%= if (true) { %>a<% } %>>",
			ExpectedTagName: "a",
			Name:            "before tag name",
		},

		{
			Content:          "<a <%= if(true){ %>checked<% } %>>",
			ExpectedTagName:  "a",
			ExpectedAttrs:    []string{"checked"},
			ExpectedAttrsVal: []string{""},
			Name:             "attr checked",
		},

		{
			Content:          `<a <%= if(true){ %>checked<% } %> <%= if(Data == "yes"){ %>data-new<%= Get()%> ='<%= if(true) {%>hello <%}%>'>`,
			ExpectedTagName:  "a",
			ExpectedAttrs:    []string{"checked", "data-new"},
			ExpectedAttrsVal: []string{"", "<%= if(true) {%>hello <%}%>"},
			Name:             "attr checked and data-new",
		},

		{
			Content:          `<a <%= if(true){ %>checked<% } %> <%= if(Data == "yes"){ %> data-new<%= Get()%> ="<%= if(val == "hello") {%>hi<%}%>">`,
			ExpectedTagName:  "a",
			ExpectedAttrs:    []string{"checked", "data-new"},
			ExpectedAttrsVal: []string{"", `<%= if(val == "hello") {%>hi<%}%>`},
			Name:             "attr checked and data-new",
		},
	}

	for _, tc := range cases {
		z := NewTokenizer(strings.NewReader(tc.Content))
		for {
			z.Next()
			if err := z.Err(); err != nil {
				if err != io.EOF {
					t.Errorf("%s: got error %s", tc.Name, err)
				}
				break
			}

			tok := z.Token()
			if tok.Type != StartTagToken {
				continue
			}

			if got, want := tok.Data, tc.ExpectedTagName; got != want {
				t.Errorf("%s: got %s, and want %s", tc.Name, got, want)
			}

			for i := range tc.ExpectedAttrs {
				if got, want := tok.Attr[i].Key, tc.ExpectedAttrs[i]; got != want {
					t.Errorf("%s: got %s, and want %s", tc.Name, got, want)
				}

				if got, want := tok.Attr[i].Val, tc.ExpectedAttrsVal[i]; got != want {
					t.Errorf("%s: got %s, and want %s", tc.Name, got, want)
				}

			}
		}
	}
}

func BenchmarkRawLevelTokenizer(b *testing.B)  { benchmarkTokenizer(b, rawLevel) }
func BenchmarkLowLevelTokenizer(b *testing.B)  { benchmarkTokenizer(b, lowLevel) }
func BenchmarkHighLevelTokenizer(b *testing.B) { benchmarkTokenizer(b, highLevel) }
