# gen-elm-wrappers

You are writing an Elm program. You have a custom type, which you would
like to use for the keys of a `Dict`, or the elements of a `Set`.
Unfortunately, for both of these cases, your type has to be `comparable`,
and custom types cannot be `comparable`. What to do?

Solution: add a `gen-elm-wrappers.json`, run `gen-elm-wrappers`,
and it will generate one or more Elm modules that wrap `Dict` or `Set`,
so that you can use them with your custom type.

## Installing

If you’re using node anyway, you can download it from NPM: `npm install
--save-dev gen-elm-wrappers` or install it globally or using yarn/pnpm/whatever.
This won’t install it on your `$PATH` unless your `$PATH` already includes
`./node_modules/.bin`. You might only want to run it via a script in
your `package.json` anyway.

Otherwise, you can [download prebuilt binaries](https://github.com/dave4420/gen-elm-wrappers/releases).
Each file is a self-contained executable for the appropriate platform.
Rename it to `gen-elm-wrappers` and move it to somewhere on your `$PATH`.
`chmod +x` it if you’re not on Windows.

Or to install from source:

- `brew install go` (or whatever the best way of installing Go on your laptop is)
  - don’t worry, you don’t need to know Go to _use_ this
- `npm ci`
- `npm test`

Assuming the tests all pass, this builds a `gen-elm-wrappers` executable.
Copy it to somewhere on your `$PATH`.

## Using

Create a `gen-elm-wrappers.json` file. This needs to contain an
object which (will eventually contain more config data, but currently
only) contains a `generate` key holding an array of module definitions.

To wrap `Dict`, the module definition is an object containing these keys:

- `underlying-type`
  - Must be `"Dict"`
- `wrapper-type`
  - The fully-qualified name of the type to generate. The generated
    code will be stored in the module named here. e.g. to generate
    a `Foo.Bar` module containing a `MyDict` type, you would put
    `"Foo.Bar.MyDict"` here
- `public-key-type`
  - The fully-qualified name of your custom type that you want to use
    as keys
- `private-key-type`
  - The type of keys to use in the underlying `Dict`. This will
    typically be `Int` or `String`, but can be any concrete
    `comparable` type
- `public-key-to-private-key`
  - The fully-qualified name of a function that converts values from
    `public-key-type` to `private-key-type`. i.e. it has a type
    like `PublicKey -> PrivateKey`
- `private-key-to-public-key`
  - The fully-qualified name of a function that converts values from
    `private-key-type` to `public-key-type`. It has a type
    like `PrivateKey -> Maybe PublicKey`. You can’t use a function
    with a type like `PrivateKey -> PublicKey` here — you may need
    to write a wrapper function in your code with the correct type

To wrap `Set`, the module definition is an object containing these keys:

- `underlying-type`
  - Must be `"Set"`
- `wrapper-type`
  - The fully-qualified name of the type to generate. The generated
    code will be stored in the module named here. e.g. to generate
    a `Foo.Bar` module containing a `MySet` type, you would put
    `"Foo.Bar.MySet"` here
- `public-key-type`
  - The fully-qualified name of your custom type that you want to use
    as elements in the set
- `private-key-type`
  - The type of elements to use in the underlying `Set`. This will
    typically be `Int` or `String`, but can be any concrete
    `comparable` type
- `public-key-to-private-key`
  - The fully-qualified name of a function that converts values from
    `public-key-type` to `private-key-type`. i.e. it has a type
    like `PublicKey -> PrivateKey`
- `private-key-to-public-key`
  - The fully-qualified name of a function that converts values from
    `private-key-type` to `public-key-type`. It has a type
    like `PrivateKey -> Maybe PublicKey`. You can’t use a function
    with a type like `PrivateKey -> PublicKey` here — you may need
    to write a wrapper function in your code with the correct type

Then, run `gen-elm-wrappers`. It expects `elm.json` to be in the
current directory. It writes the generated code to the appropriate
location inside your `src` directory.

For `Dict`s, the generated code wraps all functions from the core `Dict`
module. If your program also has `elm-community/dict-extra` as a direct
dependency, it will also wrap several functions from `Dict.Extra`.

For `Set`s, the generated code wraps all functions from the core `Set`
module. If your program also has `stoeffel/set-extra` as a direct
dependency, it will also wrap some functions from `Set.Extra`.

If `elm-format` is on your PATH (and not a relative path, i.e. not
starting with `.` or `..`) then the generated code will be beautifully
formatted. (This is the case, for example, when `elm-format` and
`gen-elm-wrappers` were both installed locally using npm, and you’re
running `gen-elm-wrappers` via npm.)

## Example

If you put this in `src/Helpers.elm`:

```
module Helpers exposing (..)

import Time

maybePosixFromMillis : Int -> Maybe Time.Posix
maybePosixFromMillis millis =
    Just <| Time.millisToPosix millis
```

and then you put this in `gen-elm-wrappers.json`:

```
{
    "generate": [
        {
            "underlying-type": "Dict",
            "wrapper-type": "Type.DictTimePosix.DictTimePosix",
            "public-key-type": "Time.Posix",
            "private-key-type": "Int",
            "private-key-to-public-key": "Helpers.maybePosixFromMillis",
            "public-key-to-private-key": "Time.posixToMillis"
        },
        {
            "underlying-type": "Set",
            "wrapper-type": "Type.SetTimePosix.SetTimePosix",
            "public-key-type": "Time.Posix",
            "private-key-type": "Int",
            "private-key-to-public-key": "Helpers.maybePosixFromMillis",
            "public-key-to-private-key": "Time.posixToMillis"
        }
    ]
}
```

then `gen-elm-wrappers` will produce

- a `Type.DictTimePosix` module in `src/Type/DictTimePosix.elm`,
  containing a `DictTimePosix v` type that acts like a `Dict` with
  `Time.Posix` keys and `v` values
- a `Type.SetTimePosix` module in `src/Type/SetTimePosix.elm`,
  containing a `SetTimePosix` type that acts like a `Set` with
  `Time.Posix` elements

## Limitations

- It won’t work if you try to wrap more than one type into the same module
- `Set.map` is not wrapped
- `Dict.Extra`:
  - `removeMany` and `keepOnly` are not wrapped, even when `Set` is also
    being wrapped for the key type
  - `mapKeys` and `invert` are not wrapped
- `Set.Extra`:
  - `concatMap` and `filterMap` are not wrapped

## Portability

I’ve only tested this on my Mac. But it’s written in Go, and I hear
Go’s really portable, so presumably it should also run on \*BSD, various
Linuxes, Windows, smart toasters, ZX-81, etc.

Actually… the test script won’t run on Windows. (Unless you use WSL?)

## Roadmap

This isn’t in priority order yet and I’ve probably forgotten something.

- Support type variables in dict key types and set element types
- Support versions of `elm-community/dict-extra` before 2.4.0
- Support versions of `stoeffel/set-extra` before 1.2.0
- Wrap more functions from `elm-community/dict-extra` and
  `stoeffel/set-extra`
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

- Senior / tech lead roles
- Full stack or backend
- Permanent only (no contracting)
- IC only (no line management)
- Remote (UK timezone ± an hour or two) or on-site/hybrid (London/Medway); not willing to relocate
- I prefer to work with statically typed languages (e.g. Typescript, not plain Javascript)
- not blockchain (except for catching crims), not ad tech (unless it’s surveillance-free),
  don’t really want to work for a hedge fund
- Ich kann ein bisschen Deutsch sprechen, aber mein Deutsch ist schlect
  und Englisch is meine Muttersprache.

Please [connect to me on LinkedIn](https://www.linkedin.com/in/dave-hinton-7507b4ab)
and mention this repo in your invitation (and don’t bury the lede if you want to
talk to me about an Elm or climate tech job).
