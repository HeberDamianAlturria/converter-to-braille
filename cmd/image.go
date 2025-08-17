package cmd

import (
	"os"
	"strings"
	"github.com/spf13/cobra"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/loader"
	"github.com/HeberDamianAlturria/converter-to-braille/internal/imageproc"
)

var imageCmd = &cobra.Command{
	Use:   "image [path]",
	Short: "Convert an image to braille",
	Long:  "Convert an image to braille ASCII art",
	Args:  cobra.ExactArgs(1),
	Run:   runImageCmd,
}

func runImageCmd(cmd *cobra.Command, args []string) {
	path := args[0]

	img, err := loader.FromFile(path)
	if err != nil {
		cmd.PrintErrln("Error loading image:", err)
		os.Exit(1)
	}

	imgProc := imageproc.New(img)
	
	var strBuilder strings.Builder

	imgProc.WriteToBrille(&strBuilder)

	cmd.Println(strBuilder.String())
}


func init() {
	rootCmd.AddCommand(imageCmd)
}


