package trie

import (
	"bufio"
	"io"
	"os"
	"testing"
)

var (
	tr    = New()
	words []string
)

func init() {
	words = []string{
		"共产党",
		"法轮功",
		"快上车",
		"飙车",
		"老司机",
		"土共",
		"五毛党",
		"翻车",
		"qa",
		"asd",
		"asdf",
		"ascd",
		"csdf",
		"zxdss",
	}
	// load()
	for _, w := range words {
		tr.Add(w, 1)
	}
}
func load() {
	f, err := os.Open("filter.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		words = append(words, line[:len(line)-2])
	}
}

func TestPrefix(t *testing.T) {
	t.Log(tr.PrefixFind("as"))
}

func TestDel(t *testing.T) {
	t.Log(tr.Words())
	tr.Del("ascd")
	tr.Del("asdf")
	t.Log(tr.Words())
}

func TestFind(t *testing.T) {
	t.Log(tr.Find("qa"))
	t.Log(tr.Find("as"))
	t.Log(tr.Find("asd"))
}

func TestFilter(t *testing.T) {
	//tr.String(tr.root, "")
	ws := []string{
		"五毛五毛党，练练法法轮功夫,老司机带带我，老司机要飙车，飙车，翻翻车,自慰用,ssssss,作弊器sdwwfw,ssdfsdggw。身上的全球市场吃晚饭,老司机带带我，老司机要飙车，飙车，翻翻车,自慰用,ssssss,作弊器sdwwfw,ssdfsdggw。身上的全球市场吃晚饭",
		"老司机带带我，老司机要飙车，飙车，翻翻车,自慰用,ssssss,作弊器sdwwfw,ssdfsdggw。身上的全球市场吃晚饭",
	}
	for _, w := range ws {
		out, level := tr.Filter(w)
		t.Logf("input:%s\n  \toutput:%s", w, out, level)
	}
}

func BenchmarkFilter(b *testing.B) {
	ws := []string{
		"五毛五毛党，练练法法轮功夫",
		"老司机带带我，老司机要飙车，飙车，翻翻车,老司机带带我，老司机要飙车，飙车，翻翻车,自慰用,ssssss,作弊器sdwwfw,ssdfsdggw。身上的全球市场吃晚饭,老司机带带我，老司机要飙车，飙车，翻翻车,自慰用,ssssss,作弊器sdwwfw,ssdfsdggw。身上的全球市场吃晚饭,老司机带带我，老司机要飙车，飙车，翻翻车,自慰用,ssssss,作弊器sdwwfw,ssdfsdggw。身上的全球市场吃晚饭",
	}
	for i := 0; i < b.N; i++ {
		tr.Filter(ws[1])
	}
}
