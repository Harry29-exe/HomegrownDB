package analyse

import (
	"HomegrownDB/backend/internal/analyser/anlctx"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"strings"
)

var ColumnDef = columnDef{}

type columnDef struct{}

func (c columnDef) Analyse(columnNode pnode.ColumnDef, currentCtx anlctx.QueryCtx) (tabdef.ColumnDefinition, error) {
	colType, err := ColumnType.Analyse(columnNode.Type, columnNode.Args)
	if err != nil {
		return nil, err
	}

	return tabdef.NewColumnDefinition(columnNode.Name, 0, 0, colType), nil
}

// -------------------------
//      ColumnType
// -------------------------

var ColumnType = columnType{}

type columnType struct{}

func (c columnType) Analyse(typeName string, constraint []pnode.TypeArg) (hgtype.ColType, error) {
	switch strings.ToUpper(typeName) {
	case "INT8":
		return c.analyseINT8(constraint)
	case "VARCHAR":
		return c.analyseVarchar(constraint)
	case "FLOAT8":
		//todo implement me
		panic("Not implemented")
	default:
		return nil, sqlerr.AnlsrErr.NewInvalidTypeErr(typeName)
	}
}

func (c columnType) analyseINT8(args []pnode.TypeArg) (hgtype.ColType, error) {
	err := c.validateTypeArgs(rawtype.TypeInt8, args, []pnode.TypeArgType{
		pnode.TypeArgTypeNullable,
	})
	if err != nil {
		return nil, err
	}

	nullable, ok := c.GetNullable(args)
	if !ok {
		nullable = true
	}

	return hgtype.NewInt8(hgtype.Args{
		Length:   8,
		Nullable: nullable,
	}), nil
}

func (c columnType) analyseVarchar(inputArgs []pnode.TypeArg) (hgtype.ColType, error) {
	err := c.validateTypeArgs(rawtype.TypeStr, inputArgs, []pnode.TypeArgType{
		pnode.TypeArgTypeNullable, pnode.TypeArgTypeLength,
	})
	if err != nil {
		return nil, err
	}
	args, ok := hgtype.Args{
		UTF8:   true,
		VarLen: true,
	}, true

	if args.Nullable, ok = c.GetNullable(inputArgs); !ok {
		args.Nullable = true
	}
	if args.Length, ok = c.GetLength(inputArgs); !ok {
		args.Length = 1
	}

	return hgtype.NewStr(args), nil
}

func (columnType) GetLength(args []pnode.TypeArg) (v int, ok bool) {
	for _, arg := range args {
		if arg.Arg == pnode.TypeArgTypeLength {
			v, ok = arg.Val.(int)
			return
		}
	}
	return
}

func (columnType) GetNullable(args []pnode.TypeArg) (v bool, ok bool) {
	for _, arg := range args {
		if arg.Arg == pnode.TypeArgTypeNullable {
			v, ok = arg.Val.(bool)
			return
		}
	}
	return
}

func (columnType) validateTypeArgs(typeTag rawtype.Tag, args []pnode.TypeArg, allowedTypes []pnode.TypeArgType) error {
	for _, arg := range args {
		allowed := false
		for _, allowedType := range allowedTypes {
			if arg.Arg == allowedType {
				allowed = true
				break
			}
		}
		if !allowed {
			return sqlerr.AnlsrErr.NewTypeArgErr(arg.Arg.ToString(), arg.Val, typeTag.ToStr())
		}
	}
	return nil
}
