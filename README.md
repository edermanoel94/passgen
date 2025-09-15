# passgen

A tiny Linux-only CLI that generates random passwords straight from the kernel’s entropy pool.

> Version: **0.0.1**  
> Platform: **Linux** (enforced via build tag)  
> Language: **Go**

---

## Why passgen?

- Uses bytes read from `/dev/random` (high-quality entropy).
- Easy-read mode to avoid ambiguous characters: `l`, `1`, `O`, and `0`.
- Zero external dependencies; a single static binary.

---

## Installation

### From source

Requires Go **1.21+**.

```bash
# clone your repo path or drop this main.go into a folder named passgen
go build -tags=linux -o passgen
sudo mv passgen /usr/local/bin/
```

#### Via go install
GOOS=linux go install github.com/you/passgen@latest

## Behavior & Notes

Entropy source: reads from /dev/random. May block if entropy is low.
Character range: printable ASCII from ! (33) to ~ (126).
Easy-read mode (-e): skips 0, 1, O, l.
Output: password is printed to stdout followed by a newline.


## Exit codes

| Code | Meaning                          |
| ---- | -------------------------------- |
| 0    | Success                          |
| 1    | Length ≤ 0                       |
| 255  | Other runtime errors (I/O, etc.) |
