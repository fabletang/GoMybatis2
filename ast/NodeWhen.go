package ast

import (
	"github.com/fabletang/GoMybatis2/stmt"
	"github.com/fabletang/GoMybatis2/utils"
)

type NodeWhen struct {
	childs []Node
	test   string
	t      NodeType

	holder *NodeConfigHolder
}

func (it *NodeWhen) Type() NodeType {
	return NWhen
}

func (it *NodeWhen) Eval(env map[string]interface{}, arg_array *[]interface{}, stmtConvert stmt.StmtIndexConvert) ([]byte, error) {
	if it.holder == nil {
		return nil, nil
	}
	var result, err = it.holder.GetExpressionEngineProxy().LexerAndEval(it.test, env)
	if err != nil {
		err = utils.NewError("GoMybatisSqlBuilder", "[GoMybatis] <test `", it.test, `> fail,`, err.Error())
	}
	if result.(bool) {
		return DoChildNodes(it.childs, env, arg_array, stmtConvert)
	}
	return nil, nil
}
