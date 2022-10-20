# Strings, bytes, runes and characters in Go

## [Source](https://go.dev/blog/strings)

## Intro

In go, using String well requires understanding not only how they work, but also the `difference between` a `byte, a character, and a rune`, the difference between `Unicode and UTF-8`, the difference between a `string and a string literal, and other even more` subtle distinctions.

An excellent introduction to some of these issues, independent of Go, is Joel Spolsky’s famous blog post, The [Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/). Many of the points he raises will be echoed here.

## What is a string?

In Go, a `string` is in effect a `read-only slice of bytes`. If you’re at all uncertain about what a slice of bytes is or how it works. , please read the previous [blog](https://go.dev/blog/slices) post; we’ll assume here that you have.

It’s important to state right up front that a `string holds arbitrary bytes`. It is not required to hold Unicode text, UTF-8 text, or any other predefined format.

As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

Here is a string literal (more about those soon) that uses the `\xNN` notation to define a string constant `holding some peculiar byte values`. (Of course, bytes range from `hexadecimal values` 00 through FF, inclusive.)

```text
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
```

## Printing strings

Because some of the bytes in our sample string are not valid ASCII, not even valid UTF-8, printing the string directly will produce ugly output. The simple print statement

```text
fmt.Println(sample)
```

produces this mess (whose exact appearance varies with the environment):

```text
��=� ⌘
```

`To find out what that string really holds`, we need to take it apart and examine the pieces. There are several ways to do this. The most obvious is to loop over its contents and pull out the bytes individually, as in this for loop:

```go
for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
```

As implied up front, `indexing a string accesses individual bytes, not characters`. We’ll return to that topic in detail below. For now, let’s stick with just the bytes. This is the output from the byte-by-byte loop:

```text
bd b2 3d bc 20 e2 8c 98
```

`Notice` how the individual bytes match the `hexadecimal escapes` that defined the string.

A shorter way to generate presentable output for a messy string is to use the `%x` (hexadecimal) format verb of `fmt.Printf`. It just `dumps out the sequential bytes of the string as hexadecimal digits, two per byte`.

```text
fmt.Printf("%x\n", sample)
```

Compare its output to that above:

```text
bdb23dbc20e28c98
```

A nice trick is to use the “space” flag in that format, putting a space between the % and the x. Compare the format string used here to the one above,

```text
fmt.Printf("% x\n", sample)
```

and notice how the bytes come out with spaces between, making the result a little less imposing:

```text
bd b2 3d bc 20 e2 8c 98
```

There’s more. The `%q` (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.

```text
    fmt.Printf("%q\n", sample)
```

This technique is handy when much of the string is intelligible as text but there are peculiarities to root out; it produces:

```text
    "\xbd\xb2=\xbc ⌘"
```

If we squint at that, we can see that buried in the noise is one ASCII equals sign, along with a regular space, and at the end appears the well-known Swedish “Place of Interest” symbol.

That symbol has Unicode value U+2318, encoded as UTF-8 by the bytes after the space (hex value 20): e2 8c 98.

If `we are unfamiliar or confused by strange values` in the string, we can use the `“plus” flag to the %q` verb. This flag causes the output to escape not only non-printable sequences, but also any non-ASCII bytes, all while interpreting UTF-8. The result is that it exposes the Unicode values of properly formatted UTF-8 that represents non-ASCII data in the string:

```text
fmt.Printf("%+q\n", sample)
```

With that format, the Unicode value of the Swedish symbol shows up as a `\u` escape:

```text
"\xbd\xb2=\xbc \u2318"
```

`These printing techniques are good to know when debugging the contents of strings`, and will be handy in the discussion that follows. It’s worth pointing out as well that all these methods behave exactly the same for byte slices as they do for strings.

Here’s the `full set of printing options` we’ve listed, presented as a complete program you can run (and edit) right in the browser:

```go
package main

import "fmt"

func main() {
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
    var Nsample []byte = []byte(sample)

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Byte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    for _, byte := range(Nsample){
        fmt.Printf("%x ", byte)
    }
    fmt.Printf("\n")

    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)
    fmt.Printf("%x\n", Nsample)

    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sample)
    fmt.Printf("% x\n", Nsample)

    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)
    fmt.Printf("%q\n", Nsample)

    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sample)
    fmt.Printf("%+q\n", Nsample)
}
```

[Exercise: Modify the examples above to use a slice of bytes instead of a string. Hint: Use a conversion to create the slice.]

[Exercise: Loop over the string using the %q format on each byte. What does the output tell you?]

## UTF-8 and string literals
