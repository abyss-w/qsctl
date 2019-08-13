package main

import (
	"github.com/spf13/cobra"

	"github.com/yunify/qsctl/v2/action"
	"github.com/yunify/qsctl/v2/utils"
)

var (
	cpExpectSize int64
	cpMaxMemory  int64
)

// CpCommand will handle copy command.
var CpCommand = &cobra.Command{
	Use:   "cp <source-path> <dest-path>",
	Short: "copy from/to qingstor",
	Long:  "qsctl cp can copy file/folder/stdin to qingstor or copy qingstor objects to local/stdout",
	Example: utils.AlignPrintWithColon(
		"Copy file: qsctl cp /path/to/file qs://prefix/a",
		"Copy folder: qsctl cp qs://prefix/a /path/to/folder -r",
		"Read from stdin: cat /path/to/file | qsctl cp - qs://prefix/stdin",
		"Write to stdout: qsctl cp qs://prefix/b - > /path/to/file",
	),
	Args:    cobra.ExactArgs(2),
	RunE:    cpRun,
	PreRunE: validateCpFlag,
}

func cpRun(_ *cobra.Command, args []string) (err error) {
	// Package handler
	cpHandler := &action.CopyHandler{}
	return cpHandler.
		WithBench(bench).
		WithDest(args[1]).
		WithExpectSize(cpExpectSize).
		WithMaximumMemory(cpMaxMemory).
		WithSrc(args[0]).
		Copy()
}

func initCpFlag() {
	CpCommand.PersistentFlags().StringVar(&expectSize,
		"expect-size",
		"",
		"expected size of the input file"+
			"accept: 100MB, 1.8G\n"+
			"(only used and required for input from stdin)",
	)
	CpCommand.PersistentFlags().StringVar(&maximumMemoryContent,
		"maximum-memory-content",
		"",
		"maximum content loaded in memory\n"+
			"(only used for input from stdin)",
	)
}

func validateCpFlag(_ *cobra.Command, _ []string) (err error) {
	if expectSize != "" {
		cpExpectSize, err = utils.ParseByteSize(expectSize)
		if err != nil {
			return err
		}
	}

	if maximumMemoryContent != "" {
		cpMaxMemory, err = utils.ParseByteSize(maximumMemoryContent)
		if err != nil {
			return err
		}
	}
	return nil
}
