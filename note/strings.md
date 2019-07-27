# Package & Method

``` [strings / type {interface / struct} ]


```

``` [strings {strings /  replacer / reader / compare / builder}]

// Package strings implements simple functions to manipulate UTF-8 encoded strings.
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(s, substr string) int

// Contains reports whether substr is within s.
func Contains(s, substr string) bool

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func IndexRune(s string, r rune) int

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny(s, chars string) int

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexAny(s, chars string) int

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s string, c byte) int

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitN(s, sep string, n int) []string

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for SplitAfter.
func SplitAfterN(s, sep string, n int) []string

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
func Split(s, sep string) []string

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(s, sep string) []string

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func Fields(s string) []string

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
// FieldsFunc makes no guarantees about the order in which it calls f(c).
// If f does not return consistent results for a given c, FieldsFunc may crash.
func FieldsFunc(s string, f func(rune) bool) []string

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string) string

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func Repeat(s string, count int) string

// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string

// ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.
func ToLower(s string) string

// ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.
func ToTitle(s string) string

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecial(c unicode.SpecialCase, s string) string

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecial(c unicode.SpecialCase, s string) string

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// title case, giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase, s string) string

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their title case.
//
// BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
func Title(s string) string

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s string, f func(rune) bool) string

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s string, f func(rune) bool) string

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s string, f func(rune) bool) string

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s string, f func(rune) bool) int

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc(s string, f func(rune) bool) int

// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
type asciiSet [8]uint32

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(s string, cutset string) string

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use TrimPrefix instead.
func TrimLeft(s string, cutset string) string

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRight(s string, cutset string) string

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string 

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll(s, old, new string) string

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func EqualFold(s, t string) bool

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int

```

``` [strings / replacer]

// Replacer replaces a list of strings with replacements.
// It is safe for concurrent use by multiple goroutines.
type Replacer struct {
	once   sync.Once // guards buildOnce method
	r      replacer
	oldnew []string
}

// replacer is the interface that a replacement algorithm needs to implement.
type replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

// NewReplacer returns a new Replacer from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches.
func NewReplacer(oldnew ...string) *Replacer

// Replace returns a copy of s with all replacements performed.
func (r *Replacer) Replace(s string) string

// WriteString writes s to w with all replacements performed.
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

// trieNode is a node in a lookup trie for prioritized key/value pairs. Keys
// and values may be empty. For example, the trie containing keys "ax", "ay",
// "bcbc", "x" and "xy" could have eight nodes:
//
//  n0  -
//  n1  a-
//  n2  .x+
//  n3  .y+
//  n4  b-
//  n5  .cbc+
//  n6  x+
//  n7  .y+
//
// n0 is the root node, and its children are n1, n4 and n6; n1's children are
// n2 and n3; n4's child is n5; n6's child is n7. Nodes n0, n1 and n4 (marked
// with a trailing "-") are partial keys, and nodes n2, n3, n5, n6 and n7
// (marked with a trailing "+") are complete keys.
type trieNode struct {
	// value is the value of the trie node's key/value pair. It is empty if
	// this node is not a complete key.
	value string
	// priority is the priority (higher is more important) of the trie node's
	// key/value pair; keys are not necessarily matched shortest- or longest-
	// first. Priority is positive if this node is a complete key, and zero
	// otherwise. In the example above, positive/zero priorities are marked
	// with a trailing "+" or "-".
	priority int

	// A trie node may have zero, one or more child nodes:
	//  * if the remaining fields are zero, there are no children.
	//  * if prefix and next are non-zero, there is one child in next.
	//  * if table is non-zero, it defines all the children.
	//
	// Prefixes are preferred over tables when there is one child, but the
	// root node always uses a table for lookup efficiency.

	// prefix is the difference in keys between this trie node and the next.
	// In the example above, node n4 has prefix "cbc" and n4's next node is n5.
	// Node n5 has no children and so has zero prefix, next and table fields.
	prefix string
	next   *trieNode

	// table is a lookup table indexed by the next byte in the key, after
	// remapping that byte through genericReplacer.mapping to create a dense
	// index. In the example above, the keys only use 'a', 'b', 'c', 'x' and
	// 'y', which remap to 0, 1, 2, 3 and 4. All other bytes remap to 5, and
	// genericReplacer.tableSize will be 5. Node n0's table will be
	// []*trieNode{ 0:n1, 1:n4, 3:n6 }, where the 0, 1 and 3 are the remapped
	// 'a', 'b' and 'x'.
	table []*trieNode
}

// genericReplacer is the fully generic algorithm.
// It's used as a fallback when nothing faster can be used.
type genericReplacer struct {
	root trieNode
	// tableSize is the size of a trie node's lookup table. It is the number
	// of unique key bytes.
	tableSize int
	// mapping maps from key bytes to a dense index for trieNode.table.
	mapping [256]byte
}

type appendSliceWriter []byte

// Write writes to the buffer to satisfy io.Writer.
func (w *appendSliceWriter) Write(p []byte) (int, error)

// WriteString writes to the buffer without string->[]byte->string allocations.
func (w *appendSliceWriter) WriteString(s string) (int, error)

type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (int, error)

func (r *genericReplacer) Replace(s string) string

func (r *genericReplacer) WriteString(w io.Writer, s string) (n int, err error)

// singleStringReplacer is the implementation that's used when there is only
// one string to replace (and that string has more than one byte).
type singleStringReplacer struct {
	finder *stringFinder
	// value is the new string that replaces that pattern when it's found.
	value string
}

func (r *singleStringReplacer) Replace(s string) string

func (r *singleStringReplacer) WriteString(w io.Writer, s string) (n int, err error)

// byteReplacer is the implementation that's used when all the "old"
// and "new" values are single ASCII bytes.
// The array contains replacement bytes indexed by old byte.
type byteReplacer [256]byte

func (r *byteReplacer) Replace(s string) string

func (r *byteReplacer) WriteString(w io.Writer, s string) (n int, err error)

// byteStringReplacer is the implementation that's used when all the
// "old" values are single ASCII bytes but the "new" values vary in size.
type byteStringReplacer struct {
	// replacements contains replacement byte slices indexed by old byte.
	// A nil []byte means that the old byte should not be replaced.
	replacements [256][]byte
	// toReplace keeps a list of bytes to replace. Depending on length of toReplace
	// and length of target string it may be faster to use Count, or a plain loop.
	// We store single byte as a string, because Count takes a string.
	toReplace []string
}

// countCutOff controls the ratio of a string length to a number of replacements
// at which (*byteStringReplacer).Replace switches algorithms.
// For strings with higher ration of length to replacements than that value,
// we call Count, for each replacement from toReplace.
// For strings, with a lower ratio we use simple loop, because of Count overhead.
// countCutOff is an empirically determined overhead multiplier.
// TODO(tocarip) revisit once we have register-based abi/mid-stack inlining.
const countCutOff = 8

func (r *byteStringReplacer) Replace(s string) string

func (r *byteStringReplacer) WriteString(w io.Writer, s string) (n int, err error)

```

``` [strings / reader]

// A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
// io.ByteScanner, and io.RuneScanner interfaces by reading
// from a string.
// The zero value for Reader operates like a Reader of an empty string.
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

// Len returns the number of bytes of the unread portion of the
// string.
func (r *Reader) Len() int

// Size returns the original length of the underlying string.
// Size is the number of bytes available for reading via ReadAt.
// The returned value is always the same and is not affected by calls
// to any other method.
func (r *Reader) Size() int64

func (r *Reader) Read(b []byte) (n int, err error)

func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

func (r *Reader) ReadByte() (byte, error)

func (r *Reader) UnreadByte() error

func (r *Reader) ReadRune() (ch rune, size int, err error)

func (r *Reader) UnreadRune() error

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Reset resets the Reader to be reading from s.
func (r *Reader) Reset(s string)}

// NewReader returns a new Reader reading from s.
// It is similar to bytes.NewBufferString but more efficient and read-only.
func NewReader(s string) *Reader

```

``` [strings / compare]

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func Compare(a, b string) int

```

``` [strings / builder]

// A Builder is used to efficiently build a string using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}

// String returns the accumulated string.
func (b *Builder) String() string

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *Builder) Len() int

// Cap returns the capacity of the builder's underlying byte slice. It is the
// total space allocated for the string being built and includes any bytes
// already written.
func (b *Builder) Cap() int

// Reset resets the Builder to be empty.
func (b *Builder) Reset()

// Grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After Grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, Grow panics.
func (b *Builder) Grow(n int)

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *Builder) Write(p []byte) (int, error)

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (b *Builder) WriteByte(c byte) error

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
func (b *Builder) WriteRune(r rune) (int, error)

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *Builder) WriteString(s string) (int, error)

```