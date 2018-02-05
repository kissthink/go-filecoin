package commands

import (
	"context"

	cmds "gx/ipfs/QmWGgKRz5S24SqaAapF5PPCfYfLT7MexJZewN5M82CQTzs/go-ipfs-cmds"

	"github.com/filecoin-project/go-filecoin/node"
)

// Env is the environment passed to commands. Implements cmds.Environment.
type Env struct {
	ctx  context.Context
	Node *node.Node
}

var _ cmds.Environment = (*Env)(nil)

// Context returns the context of the environment.
func (ce *Env) Context() context.Context {
	return ce.ctx
}

// GetNode returns the Filecoin node of the environment.
func GetNode(env cmds.Environment) *node.Node {
	ce := env.(*Env)
	return ce.Node
}