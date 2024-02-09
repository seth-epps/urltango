# urltango

Small tool to url encode/decode strings
```sh
Usage:
    urltango [command]

Available Commands:
    encode	url encode provided arguments
    decode	url decode provided arguments
```

Install with
```
❯ go install .
❯ urltango
Usage:
    urltango [command]

Available Commands:
    encode	url encode provided arguments
    decode	url decode provided arguments

❯ urltango decode my%20strings
my strings

❯ urltango encode my strings
my%20strings

❯ urltango decode `urltango encode my strings`
my strings
```