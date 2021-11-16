// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pdf

import (
	"fmt"
	"io"
	"log"
	"sort"
)

// A Stack represents a stack of values.
type Stack struct {
	stack []Value
}

func (stk *Stack) Len() int {
	return len(stk.stack)
}

func (stk *Stack) Push(v Value) {
	stk.stack = append(stk.stack, v)
}

func (stk *Stack) Pop() Value {
	n := len(stk.stack)
	if n == 0 {
		return Value{}
	}
	v := stk.stack[n-1]
	stk.stack[n-1] = Value{}
	stk.stack = stk.stack[:n-1]
	return v
}

func newDict() Value {
	return Value{nil, objptr{}, make(dict)}
}

// Interpret interprets the content in a stream as a basic PostScript program,
// pushing values onto a stack and then calling the do function to execute
// operators. The do function may push or pop values from the stack as needed
// to implement op.
//
// Interpret handles the operators "dict", "currentdict", "begin", "end", "def", and "pop" itself.
//
// Interpret is not a full-blown PostScript interpreter. Its job is to handle the
// very limited PostScript found in certain supporting file formats embedded
// in PDF files, such as cmap files that describe the mapping from font code
// points to Unicode code points.
//
// There is no support for executable blocks, among other limitations.
//
func Interpret(strm Value, do func(stk *Stack, op string)) {
	rd := strm.Reader()
	b := newBuffer(rd, 0)
	b.allowEOF = true
	b.allowObjptr = false
	b.allowStream = false
	var stk Stack
	var dicts []dict
Reading:
	for {
		tok := b.readToken()
		if tok == io.EOF {
			break
		}
		if kw, ok := tok.(keyword); ok {
			switch kw {
			case "null", "[", "]", "<<", ">>":
				break
			default:
				for i := len(dicts) - 1; i >= 0; i-- {
					if v, ok := dicts[i][name(kw)]; ok {
						stk.Push(Value{nil, objptr{}, v})
						continue Reading
					}
				}
				do(&stk, string(kw))
				continue
			case "dict":
				stk.Pop()
				stk.Push(Value{nil, objptr{}, make(dict)})
				continue
			case "currentdict":
				if len(dicts) == 0 {
					panic("no current dictionary")
				}
				stk.Push(Value{nil, objptr{}, dicts[len(dicts)-1]})
				continue
			case "begin":
				d := stk.Pop()
				if d.Kind() != Dict {
					panic("cannot begin non-dict")
				}
				dicts = append(dicts, d.data.(dict))
				continue
			case "end":
				if len(dicts) <= 0 {
					panic("mismatched begin/end")
				}
				dicts = dicts[:len(dicts)-1]
				continue
			case "def":
				if len(dicts) <= 0 {
					panic("def without open dict")
				}
				val := stk.Pop()
				key, ok := stk.Pop().data.(name)
				if !ok {
					panic("def of non-name")
				}
				dicts[len(dicts)-1][key] = val.data
				continue
			case "pop":
				stk.Pop()
				continue
			}
		}
		b.unreadToken(tok)
		obj := b.readObject()
		stk.Push(Value{nil, objptr{}, obj})
	}
}

type seqReader struct {
	rd     io.Reader
	offset int64
}

func (r *seqReader) ReadAt(buf []byte, offset int64) (int, error) {
	if offset != r.offset {
		return 0, fmt.Errorf("non-sequential read of stream")
	}
	n, err := io.ReadFull(r.rd, buf)
	r.offset += int64(n)
	return n, err
}

// 按照行读取数据
func (p *Page) GetRow() ([]string, error){
	rse := make([]string, 0)
	row, err := p.GetTextByRow()
	if err != nil {
		return rse, err
	}
	for _, v := range row {
		t := ""
		for _,txt := range v.Content {
			t+=txt.S
		}
		rse = append(rse, t)
	}
	return rse, nil
}


// 获取为存文本
func (p *Page) GetTxt() (string, error) {
	fonts := make(map[string]*Font)
	for _, name := range p.Fonts() { // cache fonts so we don't continually parse charmap
		if _, ok := fonts[name]; !ok {
			f := p.Font(name)
			fonts[name] = &f
		}
	}
	return p.GetPlainText(fonts)
}

/*
	提取表格
*/


type fw struct {
	YMin int
	YMax int
	X map[int][]int
	Str map[int]string
}

func (p *Page) GetTable() []map[int]string {

	rse := make([]map[int]string,0)

	yMap := make(map[int][]int) // y轴对应的x的每个点坐标
	for _, v := range p.Content().Rect {
		if _, ok := yMap[int(v.Max.Y)]; !ok {
			yMap[int(v.Max.Y)] = []int{}
		}
		yMap[int(v.Max.Y)] = append(yMap[int(v.Max.Y)], int(v.Max.X))
		if _, ok := yMap[int(v.Min.Y)]; !ok {
			yMap[int(v.Min.Y)] = []int{}
		}
		yMap[int(v.Min.Y)] = append(yMap[int(v.Min.Y)], int(v.Min.X))
	}

	var ySlice []int // y轴每个坐标
	for k, _ := range yMap {
		ySlice = append(ySlice, k)
	}
	sort.Ints(ySlice)

	// 横轴线降噪, 因为有重叠线的存在
	for i:=0; i<len(ySlice)-1; i++ {
		// 2表示两根线的间距
		if ySlice[i+1] - ySlice[i] < 2 {
			delete(yMap, ySlice[i])
		}
	}

	// 整理出线坐标, 划定每个单元格范围
	fws := make([]*fw,0)
	for i:=0; i<len(ySlice)-1; i++ {
		if ySlice == nil {
			// y 轴没有线
			continue
		}
		qcData := qc(yMap[ySlice[i]])
		x := rewei(qcData)
		if len(x) > 0 {
			n := 0
			for _,v := range x {
				f := &fw{
					YMin: ySlice[i],
					YMax: ySlice[i+1],
					X: make(map[int][]int),
					Str: make(map[int]string),
				}
				f.X[n] = v
				f.Str[n] = ""
				fws = append(fws, f)
				n++
			}
		}
	}

	// 以行的方式读取数据
	row, err := p.GetTextByRow()
	if err != nil {
		log.Println("读取数据失败 : ", err)
		return rse
	}
	for i:=0; i<len(row); i++ {
		rowData := row[i]
		// 将出现在单元格范围内的数据加入单元格
		for _,v := range fws {
			if int(rowData.Position) >= v.YMin && int(rowData.Position) <= v.YMax {
				for _, txtData := range rowData.Content {
					for k, x := range v.X {
						if int(txtData.X) >= x[0] && int(txtData.X) <= x[1] {
							v.Str[k]+=txtData.S
							break
						}
					}
				}
			}
		}
	}

	// 整理数据
	rd := map[int]string{}
	for _, v := range fws {
		if _,ok := v.Str[0]; ok {
			rse = append(rse, rd)
			rd = map[int]string{}
		}
		for k,v := range v.Str {
			rd[k] = v
		}
	}
	rse = append(rse, rd)
	rse = reverse(rse)

	//// 打印结果
	//for k, jg := range rse {
	//	log.Println("第",k+1,"行数据 = ", jg)
	//}
	return rse

}

func qc(a []int) []int {
	temp := map[int]struct{}{}
	for _,v := range a {
		temp[v]= struct{}{}
	}
	rse := []int{}
	for k,_ := range temp {
		rse = append(rse, k)
	}
	sort.Ints(rse)
	return rse
}

func reverse(s []map[int]string) []map[int]string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func rewei(a []int) [][]int {
	rse := [][]int{}
	for i:=0; i< len(a)-1; i++ {
		if a[i+1] - a[i] > 2 {
			rse = append(rse, []int{a[i], a[i+1]})
		}
	}
	return rse
}

