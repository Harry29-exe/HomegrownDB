package segparser

//
//import (
//	"HomegrownDB/backend/new/internal/parser/internal/sqlerr"
//	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
//	"HomegrownDB/backend/new/internal/pnode"
//	"fmt"
//)
//
//var InsertValues = insertValues{}
//
//type insertValues struct{}
//
//// Parse part of query starting with Values or Value token.
//// String to parse by this method can look like this:
////
//// VALUES (colName1, colName2 , colName4,colName6), ( colName1,colName2 ,colName4 , colName6 )
////
//// It will start parsing it when current token is Values and return
//// when source pointer is on closing parenthesis
//func (i insertValues) Parse(src tkSource, v tkValidator) (pnode.SelectStmt, error) {
//	src.Checkpoint()
//	err := v.CurrentIsAnd(token.Values).
//		SkipTokens().
//		TypeMax(token.SpaceBreak, 1).
//		SkipFromNext()
//	if err != nil {
//		return nil, err
//	}
//
//	selectStmt := pnode.NewSelectStmt()
//
//	src.Next()
//	for {
//		value, err = i.parseRow(src, v)
//		if err != nil {
//			src.Rollback()
//			return values, err
//		}
//
//		values = append(values, value)
//
//		err = v.SkipTokens().
//			TypeExactly(token.Comma, 1).
//			TypeMax(token.SpaceBreak, 2).
//			SkipFromNext()
//		if err != nil {
//			if len(values) > 0 {
//				src.Rollback()
//				return values, nil
//			}
//			src.Commit()
//			return values, err
//		}
//
//		err = v.NextIs(token.OpeningParenthesis)
//		if err != nil {
//			src.Rollback()
//			return values, err
//		}
//	}
//}
//
//func (i insertValues) parseRow(source tkSource, v tkValidator) (insertingRow []pnode.node, err error) {
//	if v.CurrentIs(token.OpeningParenthesis) != nil {
//		return
//	} else {
//		_ = v.SkipSpaceFromNext()
//	}
//
//	rowNode := pnode.NewInsertingRow()
//	err = i.parseInsertingFieldVal(&rowNode, source, v)
//	if err != nil {
//		return rowNode, err
//	}
//
//	for {
//		if v.SkipTokens().
//			TypeExactly(token.Comma, 1).
//			TypeMax(token.SpaceBreak, 2).
//			SkipFromNext() != nil {
//
//			_ = v.SkipSpaceFromNext()
//			nextTk := source.Next()
//			if nextTk.Code() != token.ClosingParenthesis {
//				return rowNode, sqlerr.NewSyntaxError(")", token.ToString(nextTk.Code()), source)
//			}
//
//			return rowNode, nil
//
//		} else {
//			err = i.parseInsertingFieldVal(&rowNode, source, v)
//			if err != nil {
//				return rowNode, err
//			}
//		}
//	}
//}
//
//func (i insertValues) parseInsertingFieldVal(rowNode *pnode.InsertingRow, source tkSource, v tkValidator) error {
//	if v.NextIs(token.OpeningParenthesis) == nil {
//		panic("select inside insert is not supported yet")
//	} else if v.CurrentIs(token.Identifier) == nil {
//		panic("expression inside insert is not supported yet")
//	} else if err := rowNode.AppendValue(source.Current(), source.CurrentTokenIndex()); err != nil {
//		return sqlerr.NewSyntaxError(
//			"value that can be used as column value",
//			fmt.Sprintf("got %s", token.ToString(source.Current().Code())),
//			source,
//		)
//	}
//
//	return nil
//}
