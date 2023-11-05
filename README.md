# gen-elm-wrappers

You are writing an Elm program.  You have a custom type, which you would
like to use for the keys of a `Dict`, or the elements of a `Set`.
Unfortunately, for both of these cases, your type has to be `comparable`,
and custom types cannot be `comparable`.  What to do?

Solution: add some config to your `elm.json`, run `gen-elm-wrappers`,
and it will generate one or more Elm modules that wrap `Dict` or `Set`,
so that you can use them with your custom type.


## Installing

I haven’t wrapped this all up nicely in an NPM package yet.

For now, you need to run:
- `brew install go`
    - don’t worry, you don’t need to know Go to _use_ this
- `npm ci`
- `npm test`

Assuming the tests all pass, this builds a `gen-elm-wrappers` executable.
Copy it to somewhere on your `$PATH`.


## Using

Create a `gen-elm-wrappers.json` file.  This needs to contain an
object which (will eventually contain more config data, but currently
only) contains a `generate` key holding an array of module definitions.

To wrap `Dict`, the module definition is an object containing these keys:
- `underlying-type`
    - Must be `"Dict"`
- `wrapper-type`
    - The fully-qualified name of the type to generate.  The generated
        code will be stored in the module named here.  e.g. to generate
        a `Foo.Bar` module containing a `MyDict` type, you would put
        `"Foo.Bar.MyDict"` here
- `public-key-type`
    - The fully-qualified name of your custom type that you want to use
        as keys
- `private-key-type`
    - The type of keys to use in the underlying `Dict`.  This will
        typically be `Int` or `String`, but can be any concrete
        `comparable` type
- `public-key-to-private-key`
    - The fully-qualified name of a function that converts values from
        `public-key-type` to `private-key-type`.  i.e. it has a type
        like `PublicKey -> PrivateKey`
- `private-key-to-public-key`
    - The fully-qualified name of a function that converts values from
        `private-key-type` to `public-key-type`.  It has a type
        like `PrivateKey -> Maybe PublicKey`.  You can’t use a function
        with a type like `PrivateKey -> PublicKey` here — you may need
        to write a wrapper function in your code with the correct type

(Currently only `Dict` is supported.)

Then, run `gen-elm-wrappers`.  It expects `elm.json` to be in the
current directory.  It writes the generated code to the appropriate
location inside your `src` directory.

For `Dict`s, the generated code wraps all functions from the core `Dict`
module.  If your program also has `elm-community/dict-extra` as a direct
dependency, it will also wrap several functions from `Dict.Extra`.

If `elm-format` is on your PATH (and not a relative path, i.e. not
starting with `.` or `..`) then the generated code will be beautifully
formatted.  (This is the case, for example, when `elm-format` and
`gen-elm-wrappers` were both installed locally using npm, and you’re
running `gen-elm-wrappers` via npm.)


## Examples

See the [`bin/test.sh`](bin/test.sh) script.


## Portability

I’ve only tested this on my Mac.  But it’s written in Go, and I hear
Go’s really portable, so presumably it should also run on \*BSD, various
Linuxes, Windows, smart toasters, ZX-81, etc.

Actually… the test script won’t run on Windows.  (Unless you use WSL?)


## Roadmap

This isn’t in priority order yet and I’ve probably forgotten something.

- Improve error message when module name is missing — it might actually
    be the type name that’s missing
- Wrap it all up nicely in an NPM package that includes/downloads
    prebuilt binaries (like the `elm` NPM package does)
- Support `Set`
- Support type variables in key types
- Support versions of `elm-community/dict-extra` before 2.4.0
- Wrap more functions from `elm-community/dict-extra`
- Support writing the generated code to a directory other than `src`;
    optionally wipe it first
- Write more unit tests around reading the config from `elm.json`
- Improve error messages when something’s wrong in `elm.json`
- Validate identifiers in the config (instead of blindly writing them
    out and letting Elm complain about them)


## Development

If you don’t know Go: this is written in Go, sorry.

If you do know Go: I’m learning Go, this is my first Go program, the
code’s probably highly non-idiomatic, sorry.

`npm run test` runs the unit tests; on success it then runs the
component tests.

If you `brew install fswatch` then you can `npm run test:go:watch`.
This runs the unit tests whenever the source code changes.


## If you’re a hiring manager or a recruiter

I’m looking for a job, yes.

- Senior/lead roles
- Full stack or backend
- Permanent only (no contracting)
- Remote (UK timezone ± 1 hour) or on-site/hybrid (London/Medway)
- Ich kann ein bisschen Deutsch sprechen, aber mein Deutsch ist schlect
    und Englisch is meine Muttersprache.

Please [connect to me on LinkedIn](https://www.linkedin.com/in/dave-hinton-7507b4ab)
and mention this repo in your invitation.
