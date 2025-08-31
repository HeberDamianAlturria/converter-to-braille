package cmd

import (
	"github.com/HeberDamianAlturria/converter-to-braille/internal/imageproc"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/loader"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type ImageOptions struct {
	Inverted   bool
	OutputPath string
}

var imageOpts ImageOptions

var imageCmd = &cobra.Command{
	Use:   "image [path]",
	Short: "Convert an image to braille.",
	Long:  "Convert an image to braille ASCII art. The path must be a valid image file or URL.",
	Args:  cobra.ExactArgs(1),
	Run:   runImageCmd,
}

func runImageCmd(cmd *cobra.Command, args []string) {
	path := args[0]

	img, err := loader.GetImage(path)

	if err != nil {
		cmd.PrintErrln("Error loading image:", err)
		os.Exit(1)
	}

	imgProc := imageproc.New(img)

	var strBuilder strings.Builder

	imgProc.WriteToBrille(&strBuilder, imageOpts.Inverted)

	if imageOpts.OutputPath != "" {
		err = output.SaveTextFile(imageOpts.OutputPath, strBuilder.String())
		if err != nil {
			cmd.PrintErrln("Error writing to file:", err)
			os.Exit(1)
		}
	} else {
		cmd.Println(strBuilder.String())
	}
}

func init() {
	imageCmd.Flags().BoolVarP(&imageOpts.Inverted, "inverted", "i", false, "Invert braille colors")
	imageCmd.Flags().StringVarP(&imageOpts.OutputPath, "output", "o", "", "Output file path (default: stdout)")

	rootCmd.AddCommand(imageCmd)
}
