package log

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Warning struct {
	Code         string
	Message      string
	Instructions string
	URL          string
	Fields       map[string]interface{}
}

func Warn(msg Warning) {
	fmt.Fprintf(os.Stderr, "%s %s\n", color.YellowString("WARNING: "+msg.Message), "")
	fmt.Fprintf(os.Stderr, "| %s\n", msg.Instructions)
	fmt.Fprintf(os.Stderr, "| %s\n\n", color.HiBlackString("See details at: "+msg.URL))
}
