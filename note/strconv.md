# Package & Method

``` [strconv / type {interface / struct} ]


```

``` [strconv {atob / atoi / decimal / itoa / quote}]

// ParseBool returns the boolean value represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func ParseBool(str string) (bool, error)

// FormatBool returns "true" or "false" according to the value of b.
func FormatBool(b bool) string

// AppendBool appends "true" or "false", according to the value of b,
// to dst and returns the extended buffer.
func AppendBool(dst []byte, b bool) []byte

```

```` [strconv / atoi]

// ErrRange indicates that a value is out of range for the target type.
var ErrRange = errors.New("value out of range")

// ErrSyntax indicates that a value does not have the right syntax for the target type.
var ErrSyntax = errors.New("invalid syntax")

// A NumError records a failed conversion.
type NumError struct {
	Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
	Num  string // the input
	Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *NumError) Error() string

const intSize = 32 << (^uint(0) >> 63)

// IntSize is the size in bits of an int or uint value.
const IntSize = intSize

const maxUint64 = 1<<64 - 1

// ParseUint is like ParseInt but for unsigned numbers.
func ParseUint(s string, base int, bitSize int) (uint64, error)

// ParseInt interprets a string s in the given base (0, 2 to 36) and
// bit size (0 to 64) and returns the corresponding value i.
//
// If base == 0, the base is implied by the string's prefix:
// base 16 for "0x", base 8 for "0", and base 10 otherwise.
// For bases 1, below 0 or above 36 an error is returned.
//
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
// For a bitSize below 0 or above 64 an error is returned.
//
// The errors that ParseInt returns have concrete type *NumError
// and include err.Num = s. If s is empty or contains invalid
// digits, err.Err = ErrSyntax and the returned value is 0;
// if the value corresponding to s cannot be represented by a
// signed integer of the given size, err.Err = ErrRange and the
// returned value is the maximum magnitude integer of the
// appropriate bitSize and sign.
func ParseInt(s string, base int, bitSize int) (i int64, err error)

// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
func Atoi(s string) (int, error)

````

``` [strconv / decimal]

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

type decimal struct {
	d     [800]byte // digits, big-endian representation
	nd    int       // number of digits used
	dp    int       // decimal point
	neg   bool      // negative flag
	trunc bool      // discarded nonzero digits beyond d[:nd]
}

func (a *decimal) String() string

// Assign v to a.
func (a *decimal) Assign(v uint64)

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const uintSize = 32 << (^uint(0) >> 63)
const maxShift = uintSize - 4

// Cheat sheet for left shift: table indexed by shift count giving
// number of new digits that will be introduced by that shift.
//
// For example, leftcheats[4] = {2, "625"}.  That means that
// if we are shifting by 4 (multiplying by 16), it will add 2 digits
// when the string prefix is "625" through "999", and one fewer digit
// if the string prefix is "000" through "624".
//
// Credit for this trick goes to Ken.

type leftCheat struct {
	delta  int    // number of new digits
	cutoff string // minus one digit if original < a.
}

// Binary shift left (k > 0) or right (k < 0).
func (a *decimal) Shift(k int)

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (a *decimal) Round(nd int)

// Round a down to nd digits (or fewer).
func (a *decimal) RoundDown(nd int)

// Round a up to nd digits (or fewer).
func (a *decimal) RoundUp(nd int)

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (a *decimal) RoundedInteger() uint64

```

``` [strconv / itoa]
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// FormatUint returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatUint(i uint64, base int) string

// FormatInt returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func FormatInt(i int64, base int) string

// Itoa is equivalent to FormatInt(int64(i), 10).
func Itoa(i int) string

// AppendInt appends the string form of the integer i,
// as generated by FormatInt, to dst and returns the extended buffer.
func AppendInt(dst []byte, i int64, base int) []byte

// AppendUint appends the string form of the unsigned integer i,
// as generated by FormatUint, to dst and returns the extended buffer.
func AppendUint(dst []byte, i uint64, base int) []byte

// small returns the string for an i with 0 <= i < nSmalls.
func small(i int) string

```

``` [strconv / quote]

//go:generate go run makeisprint.go -output isprint.go

// Quote returns a double-quoted Go string literal representing s. The
// returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// control characters and non-printable characters as defined by
// IsPrint.
func Quote(s string) string {
	return quoteWith(s, '"', false, false)
}

// AppendQuote appends a double-quoted Go string literal representing s,
// as generated by Quote, to dst and returns the extended buffer.
func AppendQuote(dst []byte, s string) []byte

// QuoteToASCII returns a double-quoted Go string literal representing s.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// non-ASCII characters and non-printable characters as defined by IsPrint.
func QuoteToASCII(s string) string

// AppendQuoteToASCII appends a double-quoted Go string literal representing s,
// as generated by QuoteToASCII, to dst and returns the extended buffer.
func AppendQuoteToASCII(dst []byte, s string) []byte

// QuoteToGraphic returns a double-quoted Go string literal representing s.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// non-ASCII characters and non-printable characters as defined by IsGraphic.
func QuoteToGraphic(s string) string

// AppendQuoteToGraphic appends a double-quoted Go string literal representing s,
// as generated by QuoteToGraphic, to dst and returns the extended buffer.
func AppendQuoteToGraphic(dst []byte, s string) []byte

// QuoteRune returns a single-quoted Go character literal representing the
// rune. The returned string uses Go escape sequences (\t, \n, \xFF, \u0100)
// for control characters and non-printable characters as defined by IsPrint.
func QuoteRune(r rune) string

// AppendQuoteRune appends a single-quoted Go character literal representing the rune,
// as generated by QuoteRune, to dst and returns the extended buffer.
func AppendQuoteRune(dst []byte, r rune) []byte

// QuoteRuneToASCII returns a single-quoted Go character literal representing
// the rune. The returned string uses Go escape sequences (\t, \n, \xFF,
// \u0100) for non-ASCII characters and non-printable characters as defined
// by IsPrint.
func QuoteRuneToASCII(r rune) string

// AppendQuoteRuneToASCII appends a single-quoted Go character literal representing the rune,
// as generated by QuoteRuneToASCII, to dst and returns the extended buffer.
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte

// QuoteRuneToGraphic returns a single-quoted Go character literal representing
// the rune. The returned string uses Go escape sequences (\t, \n, \xFF,
// \u0100) for non-ASCII characters and non-printable characters as defined
// by IsGraphic.
func QuoteRuneToGraphic(r rune) string

// AppendQuoteRuneToGraphic appends a single-quoted Go character literal representing the rune,
// as generated by QuoteRuneToGraphic, to dst and returns the extended buffer.
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte

// CanBackquote reports whether the string s can be represented
// unchanged as a single-line backquoted string without control
// characters other than tab.
func CanBackquote(s string) bool

// UnquoteChar decodes the first character or byte in the escaped string
// or character literal represented by the string s.
// It returns four values:
//
//	1) value, the decoded Unicode code point or byte value;
//	2) multibyte, a boolean indicating whether the decoded character requires a multibyte UTF-8 representation;
//	3) tail, the remainder of the string after the character; and
//	4) an error that will be nil if the character is syntactically valid.
//
// The second argument, quote, specifies the type of literal being parsed
// and therefore which escaped quote character is permitted.
// If set to a single quote, it permits the sequence \' and disallows unescaped '.
// If set to a double quote, it permits \" and disallows unescaped ".
// If set to zero, it does not permit either escape and allows both quote characters to appear unescaped.
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

// Unquote interprets s as a single-quoted, double-quoted,
// or backquoted Go string literal, returning the string value
// that s quotes.  (If s is single-quoted, it would be a Go
// character literal; Unquote returns the corresponding
// one-character string.)
func Unquote(s string) (string, error)

// TODO: IsPrint is a local implementation of unicode.IsPrint, verified by the tests
// to give the same answer. It allows this package not to depend on unicode,
// and therefore not pull in all the Unicode tables. If the linker were better
// at tossing unused tables, we could get rid of this implementation.
// That would be nice.

// IsPrint reports whether the rune is defined as printable by Go, with
// the same definition as unicode.IsPrint: letters, numbers, punctuation,
// symbols and ASCII space.
func IsPrint(r rune) bool

// IsGraphic reports whether the rune is defined as a Graphic by Unicode. Such
// characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories L, M, N, P, S, and Zs.
func IsGraphic(r rune) bool

```