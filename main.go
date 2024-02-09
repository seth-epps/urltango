package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	ENCODE      = "encode"
	ENCODE_HELP = "url encode provided arguments"
	DECODE      = "decode"
	DECODE_HELP = "url decode provided arguments"
)

func main() {
	help := fmt.Sprintf("Usage:\n    urltango [command]\n\nAvailable Commands:\n    %s\t%s\n    %s\t%s", ENCODE, ENCODE_HELP, DECODE, DECODE_HELP)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(help)
		os.Exit(1)
	}

	command, rest := args[0], args[1:]

	switch command {
	case ENCODE, DECODE:
		toDecode := strings.Join(rest, " ")
		if err := run(command, toDecode); err != nil {
			fmt.Printf("failed to decode `%s`\n", toDecode)
			os.Exit(1)
		}
	default:
		fmt.Printf("invalid command `%s`\n%s", command, help)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(command string, toDecode string) error {
	if command == ENCODE {
		fmt.Println(url.PathEscape(toDecode))
	} else {
		if decoded, err := url.PathUnescape(toDecode); err != nil {
			return err
		} else {
			fmt.Println(decoded)
		}
	}
	return nil
}
