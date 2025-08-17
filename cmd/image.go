package cmd

import (
	"github.com/HeberDamianAlturria/converter-to-braille/internal/imageproc"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/loader"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/output"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var inverted bool
var outputPath string

var imageCmd = &cobra.Command{
	Use:   "image [path]",
	Short: "Convert an image to braille",
	Long:  "Convert an image to braille ASCII art",
	Args:  cobra.ExactArgs(1),
	Run:   runImageCmd,
}

func runImageCmd(cmd *cobra.Command, args []string) {
	path := args[0]

	img, err := loader.FromImageFile(path)
	if err != nil {
		cmd.PrintErrln("Error loading image:", err)
		os.Exit(1)
	}

	imgProc := imageproc.New(img)

	var strBuilder strings.Builder

	imgProc.WriteToBrille(&strBuilder, inverted)

	if outputPath != "" {
		err = output.SaveTextFile(outputPath, strBuilder.String())
		if err != nil {
			cmd.PrintErrln("Error writing to file:", err)
			os.Exit(1)
		}
	} else {
		cmd.Println(strBuilder.String())
	}
}

func init() {
	imageCmd.Flags().BoolVarP(&inverted, "inverted", "i", false, "Invert braille colors")
	imageCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output file path (default: stdout)")

	rootCmd.AddCommand(imageCmd)
}
