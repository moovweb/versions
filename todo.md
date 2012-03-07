- Accept multiple pattern strings

- Fix edge cases:
  - A major minor version `0.3` will be zero'd out to `0.3.0`
  - This is fine, except with the pessimistic operator
  - Well, I should treat '0.0' unique from '0.0.0'

- Don't assume spaces are provided. Accept `>4.5.6` as a pattern

- Auto sort any arrays of results?

- Can I auto expand tilde on the command line: `versions ~/somedir` ?

- Searching by name is a bit too loose. If I search 'foo' it matches 'foobar'

- Support no names, just raw version strings '0.0.45'

- When searching by name I return the newest result. Add a "*" pattern to return all matching filepaths for that name

- Should I add "+" and "-" patterns for newest / oldest versions? How would this work on the command line? I guess the multiple patterns is a pre-req

- Add command line interface
  - Added, but consider how I'd match the library interface on the command line, esp:
  - get newest -- do I set a flag?