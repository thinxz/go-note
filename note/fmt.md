# Package & Method

``` [fmt / type {interface / struct} ]


```

``` [fmt {format / print / scan}]

// flags placed in a separate struct for easy clearing.
type fmtFlags struct {
	widPresent  bool
	precPresent bool
	minus       bool
	plus        bool
	sharp       bool
	space       bool
	zero        bool

	// For the formats %+v %#v, we set the plusV/sharpV flags
	// and clear the plus/sharp flags since %+v and %#v are in effect
	// different, flagless formats set at the top level.
	plusV  bool
	sharpV bool
}

// A fmt is the raw formatter used by Printf etc.
// It prints into a buffer that must be set up separately.
type fmt struct {
	buf *buffer

	fmtFlags

	wid  int // width
	prec int // precision

	// intbuf is large enough to store %b of an int64 with a sign and
	// avoids padding at the end of the struct on 32 bit architectures.
	intbuf [68]byte
}

func (f *fmt) clearflags()

func (f *fmt) init(buf *buffer)

```

``` [fmt / print]

// State represents the printer state passed to custom formatters.
// It provides access to the io.Writer interface plus information about
// the flags and options for the operand's format specifier.
type State interface {
	// Write is the function to call to emit formatted output to be printed.
	Write(b []byte) (n int, err error)
	// Width returns the value of the width option and whether it has been set.
	Width() (wid int, ok bool)
	// Precision returns the value of the precision option and whether it has been set.
	Precision() (prec int, ok bool)

	// Flag reports whether the flag c, a character, has been set.
	Flag(c int) bool
}

// Formatter is the interface implemented by values with a custom formatter.
// The implementation of Format may call Sprint(f) or Fprint(f) etc.
// to generate its output.
type Formatter interface {
	Format(f State, c rune)
}

// Stringer is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
type Stringer interface {
	String() string
}

// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand
// to a %#v format.
type GoStringer interface {
	GoString() string
}

// Use simple []byte instead of bytes.Buffer to avoid large dependency.
type buffer []byte

func (b *buffer) Write(p []byte)

func (b *buffer) WriteString(s string)

func (b *buffer) WriteByte(c byte)

func (bp *buffer) WriteRune(r rune)

// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
type pp struct {
	buf buffer

	// arg holds the current item, as an interface{}.
	arg interface{}

	// value is used instead of arg for reflect values.
	value reflect.Value

	// fmt is used to format basic items such as integers or strings.
	fmt fmt

	// reordered records whether the format string used argument reordering.
	reordered bool
	// goodArgNum records whether the most recent reordering directive was valid.
	goodArgNum bool
	// panicking is set by catchPanic to avoid infinite panic, recover, panic, ... recursion.
	panicking bool
	// erroring is set when printing an error string to guard against calling handleMethods.
	erroring bool
}

var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

func (p *pp) Width() (wid int, ok bool)

func (p *pp) Precision() (prec int, ok bool)

func (p *pp) Flag(b int) bool

// Implement Write so we can call Fprintf on a pp (through State), for
// recursive use in custom verbs.
func (p *pp) Write(b []byte) (ret int, err error)

// Implement WriteString so that we can call io.WriteString
// on a pp (through state), for efficiency.
func (p *pp) WriteString(s string) (ret int, err error)

// These routines end in 'f' and take a format string.

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error)

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error

// These routines do not take a format string

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error)

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) (n int, err error)

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...interface{}) string

// These routines end in 'ln', do not take a format string,
// always add spaces between operands, and add a newline
// after the last operand.

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error)

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintln(a ...interface{}) string

```

``` [fmt / scan]

// ScanState represents the scanner state passed to custom scanners.
// Scanners may do rune-at-a-time scanning or ask the ScanState
// to discover the next space-delimited token.
type ScanState interface {
	// ReadRune reads the next rune (Unicode code point) from the input.
	// If invoked during Scanln, Fscanln, or Sscanln, ReadRune() will
	// return EOF after returning the first '\n' or when reading beyond
	// the specified width.
	ReadRune() (r rune, size int, err error)
	// UnreadRune causes the next call to ReadRune to return the same rune.
	UnreadRune() error
	// SkipSpace skips space in the input. Newlines are treated appropriately
	// for the operation being performed; see the package documentation
	// for more information.
	SkipSpace()
	// Token skips space in the input if skipSpace is true, then returns the
	// run of Unicode code points c satisfying f(c).  If f is nil,
	// !unicode.IsSpace(c) is used; that is, the token will hold non-space
	// characters. Newlines are treated appropriately for the operation being
	// performed; see the package documentation for more information.
	// The returned slice points to shared data that may be overwritten
	// by the next call to Token, a call to a Scan function using the ScanState
	// as input, or when the calling Scan method returns.
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Width returns the value of the width option and whether it has been set.
	// The unit is Unicode code points.
	Width() (wid int, ok bool)
	// Because ReadRune is implemented by the interface, Read should never be
	// called by the scanning routines and a valid implementation of
	// ScanState may choose always to return an error from Read.
	Read(buf []byte) (n int, err error)
}

// Scanner is implemented by any value that has a Scan method, which scans
// the input for the representation of a value and stores the result in the
// receiver, which must be a pointer to be useful. The Scan method is called
// for any argument to Scan, Scanf, or Scanln that implements it.
type Scanner interface {
	Scan(state ScanState, verb rune) error
}

// Scan scans text read from standard input, storing successive
// space-separated values into successive arguments. Newlines count
// as space. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
}

// Scanln is similar to Scan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Scanln(a ...interface{}) (n int, err error) {
	return Fscanln(os.Stdin, a...)
}

// Scanf scans text read from standard input, storing successive
// space-separated values into successive arguments as determined by
// the format. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
// Newlines in the input must match newlines in the format.
// The one exception: the verb %c always scans the next rune in the
// input, even if it is a space (or tab etc.) or newline.
func Scanf(format string, a ...interface{}) (n int, err error) {
	return Fscanf(os.Stdin, format, a...)
}

type stringReader string

func (r *stringReader) Read(b []byte) (n int, err error)

// Sscan scans the argument string, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
func Sscan(str string, a ...interface{}) (n int, err error)

// Sscanln is similar to Sscan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Sscanln(str string, a ...interface{}) (n int, err error)

// Sscanf scans the argument string, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
func Sscanf(str string, format string, a ...interface{}) (n int, err error)

// Fscan scans text read from r, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
func Fscan(r io.Reader, a ...interface{}) (n int, err error)

// Fscanln is similar to Fscan, but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)

// Fscanf scans text read from r, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

// scanError represents an error generated by the scanning software.
// It's used as a unique signature to identify such errors when recovering.
type scanError struct {
	err error
}

// ss is the internal implementation of ScanState.
type ss struct {
	rs    io.RuneScanner // where to read input
	buf   buffer         // token accumulator
	count int            // runes consumed so far.
	atEOF bool           // already read EOF
	ssave
}

// ssave holds the parts of ss that need to be
// saved and restored on recursive scans.
type ssave struct {
	validSave bool // is or was a part of an actual ss.
	nlIsEnd   bool // whether newline terminates scan
	nlIsSpace bool // whether newline counts as white space
	argLimit  int  // max value of ss.count for this arg; argLimit <= limit
	limit     int  // max value of ss.count.
	maxWid    int  // width of this arg.
}

// The Read method is only in ScanState so that ScanState
// satisfies io.Reader. It will never be called when used as
// intended, so there is no need to make it actually work.
func (s *ss) Read(buf []byte) (n int, err error)

func (s *ss) ReadRune() (r rune, size int, err error)

func (s *ss) Width() (wid int, ok bool)

func (s *ss) UnreadRune() error

func (s *ss) Token(skipSpace bool, f func(rune) bool) (tok []byte, err error)

// space is a copy of the unicode.White_Space ranges,
// to avoid depending on package unicode.
var space = [][2]uint16{
	{0x0009, 0x000d},
	{0x0020, 0x0020},
	{0x0085, 0x0085},
	{0x00a0, 0x00a0},
	{0x1680, 0x1680},
	{0x2000, 0x200a},
	{0x2028, 0x2029},
	{0x202f, 0x202f},
	{0x205f, 0x205f},
	{0x3000, 0x3000},
}

// readRune is a structure to enable reading UTF-8 encoded code points
// from an io.Reader. It is used if the Reader given to the scanner does
// not already implement io.RuneScanner.
type readRune struct {
	reader   io.Reader
	buf      [utf8.UTFMax]byte // used only inside ReadRune
	pending  int               // number of bytes in pendBuf; only >0 for bad UTF-8
	pendBuf  [utf8.UTFMax]byte // bytes left over
	peekRune rune              // if >=0 next rune; when <0 is ^(previous Rune)
}

// ReadRune returns the next UTF-8 encoded code point from the
// io.Reader inside r.
func (r *readRune) ReadRune() (rr rune, size int, err error)

func (r *readRune) UnreadRune() error

var ssFree = sync.Pool{
	New: func() interface{} { return new(ss) },
}

// SkipSpace provides Scan methods the ability to skip space and newline
// characters in keeping with the current scanning mode set by format strings
// and Scan/Scanln.
func (s *ss) SkipSpace()

```