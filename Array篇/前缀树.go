package main

import "fmt"

/*
Trie树的根结点是不保存数据的，所有的数据都保存在它的孩子节点中。
Trie的核心思想是空间换时间。利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的。
基本性质：
1.根节点不包含字符，除根节点外每一个节点都只包含一个字符。
2.从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串。
3.每个节点的所有子节点包含的字符都不相同。

*/
//前缀树结构体
type Trie struct {
	next [26]*Trie	//26长度的数组用来指向下一层的a-z
	isEnd bool//是否是句尾
}
//初始化
func Constructor()Trie{
	return Trie{}
}
//增加节点
func (this *Trie)Insert(word string)  {
	for _,v:=range word{
		// v - ‘a’ 的意思是 ascii码相减 减去a的话 a-z 等于 0-25 刚好得到对应相对的index
		if this.next[v-'a']==nil{
			this.next[v-'a']=&Trie{}
		}
		this=this.next[v-'a']
	}
	this.isEnd=true
}
//查找一个单词是不是在前缀树
func (this *Trie)Search(word string)bool{
	// 搜索也类似 往下走，走到哪一层没有的话 就直接返回false
	for _,v:=range word{
		if this.next[v-'a']==nil{
			return false
		}else{
			this=this.next[v-'a']
		}
	}
	return this.isEnd
}
func (this *Trie)StartWith(prefix string)bool  {
	for _,v:=range prefix{
		if this.next[v-'a']==nil{
			return false
		}else {
			this=this.next[v-'a']
		}
	}
	return true
}
func main()  {
	t:=Constructor()
	t.Insert("hello")
	fmt.Println(t.Search("hello"))
	fmt.Println(t.StartWith("he"))
	fmt.Println(t.StartWith("fdsaf"))
}
