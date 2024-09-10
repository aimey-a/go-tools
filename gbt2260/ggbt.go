package gbt2260

import (
	"fmt"

	"github.com/aimey-a/go-tools/gstring"
)

type BGT2260 struct{}

// 通过全局的方式创建trie树
var Trie = New()

func NewGBT2260() *BGT2260 {
	gbt2260Table := GetGbt2260Table()
	for _, cell := range gbt2260Table {
		createTrieTree(cell[0], cell[1], Trie)
	}
	return &BGT2260{}
}

// 向树中插入数据
func createTrieTree(code string, name string, trie *GTrie) {
	//检查传递参数
	if code == "" || len(code) == 0 {
		return
	}
	//过滤下数据构造插入lCode
	var lCode = stringParse(code)
	//创建trie树
	trieRoot := trie
	trieRoot.Add(lCode, name)
}

// 将传入的字符串解析成字符串数组
func stringParse(str string) []string {
	var lCode []string
	for i := 0; i < len(str)/2; i++ {
		if str[2*i:2*(i+1)] != "00" {
			lCode = append(lCode, str[2*i:2*(i+1)])
		}
	}
	return lCode
}

// 通过行政地域码获取地址
func (b *BGT2260) SearchGBT2260(code string) (newCode []string) {
	var lCode = stringParse(code)
	node := Trie.Root()
	for i := range lCode {
		r := lCode[i]
		if n, ok := node.children[r]; ok {
			newCode = append(newCode, n.value)
			node = n
		} else {
			fmt.Printf("对不起，您输入的地域码不在列表当中")
			return
		}
	}
	return
}

// 通过地址获取行政地域码
func (b *BGT2260) SearchCityGBT2260(citys ...string) (newCode string) {
	var province string
	var city string
	var prefecture string
	for i, v := range citys {
		switch i {
		case 0:
			province = v
		case 1:
			city = v
		case 2:
			prefecture = v
		}
	}
	if province == "" {
		newCode = "最少输入省份"
	} else {
		node := Trie.Root()

		for prcode, v := range node.children {
			if gstring.Contains(v.value, province) {
				newCode += prcode
				if city != "" {
					for cicode, u := range v.children {
						if gstring.Contains(u.value, city) {
							newCode += cicode
							if prefecture != "" {
								for precode, y := range u.children {
									if gstring.Contains(y.value, prefecture) {
										newCode += precode
									}
								}
							}
							break
						}
					}
				}
				return
			}
		}
	}

	// var lCode = stringParse(code)
	// node := Trie.Root()
	// for i := range lCode {
	// 	r := lCode[i]
	// 	if n, ok := node.children[r]; ok {
	// 		newCode = append(newCode, n.value)
	// 		node = n
	// 	} else {
	// 		fmt.Printf("对不起，您输入的地域码不在列表当中")
	// 		return
	// 	}
	// }
	newCode = "没有该地址"
	return
}
