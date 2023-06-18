package turnip

import "context"

type baseCmd struct {
	ctx context.Context
	err error
}

// Err returns error of cmd
func (cmd *baseCmd) Err() error {
	return cmd.err
}

// ------------------------------------------------------------------------------------

// StringCmd is used to commands which returns string
type StringCmd struct {
	baseCmd
	val string
}

func NewStringCmd(ctx context.Context, value string, err error) *StringCmd {
	return &StringCmd{
		baseCmd: baseCmd{
			err: err,
			ctx: ctx,
		},
		val: value,
	}
}

// Result returns value and error of string cmd
func (cmd *StringCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

// ------------------------------------------------------------------------------------

// BoolCmd is used to commands which returns bool
type BoolCmd struct {
	baseCmd
	val bool
}

func NewBoolCmd(ctx context.Context, value bool, err error) *BoolCmd {
	return &BoolCmd{
		baseCmd: baseCmd{
			err: err,
			ctx: ctx,
		},
		val: value,
	}
}

// Result returns value and error of string cmd
func (cmd *BoolCmd) Result() (bool, error) {
	return cmd.val, cmd.err
}
