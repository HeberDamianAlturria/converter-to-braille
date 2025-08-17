package cmd

import (
	"github.com/HeberDamianAlturria/converter-to-braille/internal/imageproc"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/loader"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

type GifOptions struct {
	Inverted   bool
	OutputPath string
}

var gifOpts GifOptions

var gifCmd = &cobra.Command{
	Use:   "gif [path]",
	Short: "Convert an gif to braille",
	Long:  "Convert an gif to braille ASCII art",
	Args:  cobra.ExactArgs(1),
	Run:   runGifCmd,
}

func runGifCmd(cmd *cobra.Command, args []string) {
	path := args[0]

	frames, err := loader.FromGifFileReconstructed(path)
	if err != nil {
		cmd.PrintErrln("Error loading image:", err)
		os.Exit(1)
	}

	for _, frame := range frames {
		imgProc := imageproc.New(frame.Image)

		var strBuilder strings.Builder

		imgProc.WriteToBrille(&strBuilder, gifOpts.Inverted)

		cmd.Print("\033[H\033[2J") // Clear terminal

		cmd.Println(strBuilder.String())

		delay := time.Duration(frame.Delay) * 10 * time.Millisecond

		time.Sleep(delay)
	}
}

func init() {
	gifCmd.Flags().BoolVarP(&gifOpts.Inverted, "inverted", "i", false, "Invert braille colors")

	rootCmd.AddCommand(gifCmd)
}
