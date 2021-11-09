/*
	结巴分词
 */

package service

import (
	"github.com/yanyiwu/gojieba"
)

// [jieBaType]
// 		1: 全模式
// 		2: 精确模式
//		3: 搜索引擎模式
//		4: 词性标注
//      5: Tokenize 搜索引擎模式
//      6: Tokenize 默认模式
//      7: Extract
func JieBa(str string, jieBaType int) []interface{} {
	var (
		words []string
		rse []interface{}
		jieBaWord []gojieba.Word
		jieBaWordWeight []gojieba.WordWeight
	)
	x := gojieba.NewJieba()
	switch jieBaType {
	case 1:
		words = x.CutAll(str)
	case 2:
		words = x.Cut(str, true)
	case 3:
		words = x.CutForSearch(str, true)
	case 4:
		words = x.Tag(str)
	case 5:
		jieBaWord = x.Tokenize(str, gojieba.SearchMode, false)
	case 6:
		jieBaWord = x.Tokenize(str, gojieba.DefaultMode, false)
	case 7:
		jieBaWordWeight = x.ExtractWithWeight(str, 5)
	default:
		words = x.CutAll(str)
	}

	if len(words) > 0 {
		for _,v := range words {
			rse = append(rse, v)
		}
	}

	if len(jieBaWord) > 0 {
		for _,v := range jieBaWord {
			rse = append(rse, v)
		}
	}

	if len(jieBaWordWeight) > 0 {
		for _,v := range jieBaWordWeight {
			rse = append(rse, v)
		}
	}

	return rse
}