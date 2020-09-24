package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomwright/dasel/internal/oflag"
	"strings"
)

type putObjectOpts struct {
	File        string
	Out         string
	Parser      string
	Selector    string
	InputTypes  []string
	InputValues []string
}

func getMapFromTypesValues(inputTypes []string, inputValues []string) (map[string]interface{}, error) {
	if len(inputTypes) != len(inputValues) {
		return nil, fmt.Errorf("exactly %d types are required, got %d", len(inputValues), len(inputTypes))
	}

	updateValue := map[string]interface{}{}

	for k, arg := range inputValues {
		splitArg := strings.Split(arg, "=")
		name := splitArg[0]
		value := strings.Join(splitArg[1:], "=")
		parsedValue, err := parseValue(value, inputTypes[k])
		if err != nil {
			return nil, fmt.Errorf("could not parse value [%s]: %w", name, err)
		}
		updateValue[name] = parsedValue
	}

	return updateValue, nil
}

func runPutObjectCommand(opts putObjectOpts) error {
	parser, err := getParser(opts.File, opts.Parser)
	if err != nil {
		return err
	}
	rootNode, err := getRootNode(opts.File, parser)
	if err != nil {
		return err
	}

	updateValue, err := getMapFromTypesValues(opts.InputTypes, opts.InputValues)
	if err != nil {
		return err
	}

	if err := rootNode.Put(opts.Selector, updateValue); err != nil {
		return fmt.Errorf("could not put value: %w", err)
	}

	if err := writeNodeToOutput(rootNode, parser, opts.File, opts.Out); err != nil {
		return fmt.Errorf("could not write output: %w", err)
	}

	return nil
}

func putObjectCommand() *cobra.Command {
	typeList := oflag.NewStringList()

	cmd := &cobra.Command{
		Use:   "object -f <file> -s <selector> <value>",
		Short: "Update a string property in the given file.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPutObjectCommand(putObjectOpts{
				File:        cmd.Flag("file").Value.String(),
				Out:         cmd.Flag("out").Value.String(),
				Parser:      cmd.Flag("parser").Value.String(),
				Selector:    cmd.Flag("selector").Value.String(),
				InputTypes:  typeList.Strings,
				InputValues: args,
			})
		},
	}

	cmd.Flags().VarP(typeList, "type", "t", "Types of the variables in the object.")
	_ = cmd.MarkFlagRequired("type")

	return cmd
}
