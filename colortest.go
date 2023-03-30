package main

import (
	"fmt"
)

var Color = struct {
	BLACK     string
	RED       string
	GREEN     string
	YELLOW    string
	BLUE      string
	PURPLE    string
	CYAN      string
	WHITE     string
	END       string
	BOLD      string
	UNDERLINE string
	INVISIBLE string
	REVERCE   string
}{
	BLACK:     "\033[30m",
	RED:       "\033[31m",
	GREEN:     "\033[32m",
	YELLOW:    "\033[33m",
	BLUE:      "\033[34m",
	PURPLE:    "\033[35m",
	CYAN:      "\033[36m",
	WHITE:     "\033[37m",
	END:       "\033[0m",
	BOLD:      "\033[1m",
	UNDERLINE: "\033[4m",
	INVISIBLE: "\033[08m",
	REVERCE:   "\033[07m",
}

func Decorate(text string, option string) string {
	var decorated string
	switch option {
	case "RED":
		decorated = Color.RED + text + Color.END
	case "GREEN":
		decorated = Color.GREEN + text + Color.END
	case "BOLD":
		decorated = Color.BOLD + text + Color.END
	case "UNDERLINE":
		decorated = Color.UNDERLINE + text + Color.END
	case "BLUE":
		decorated = Color.BLUE + text + Color.END
	case "YELLOW":
		decorated = Color.YELLOW + text + Color.END
	default:
		decorated = text
	}
	return decorated
}

func main() {
	red := Color.RED + "red?" + Color.END
	fmt.Println(red)
	fmt.Printf(Decorate("Hello", "UNDERLINE"))
}
