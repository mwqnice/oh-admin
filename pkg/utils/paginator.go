package utils

import (
	"bytes"
	"strconv"
	"strings"
)

type Page struct {
	PageNo     int    //当前页
	PageSize   int    //每页多少数据
	TotalPage  int    //总共多少页
	TotalCount int    //总共多少条数据
	FirstPage  int    //第一页
	LastPage   int    //最后一页
	Url        string //链接
}

func NewPage(PageNo int, PageSize int, TotalCount int, Url string) Page {
	return Page{PageNo: PageNo, PageSize: PageSize, TotalCount: TotalCount, Url: Url}
}

//计算总页数
func (this *Page) getPageCount() {
	var tp float32 = float32(this.TotalCount) / float32(this.PageSize)
	if tp < 1 {
		this.TotalPage = 1
	}
	var tpint float32 = float32(int(tp))

	if tp > tpint {
		tpint += 1
	}
	this.TotalPage = int(tpint)
	this.LastPage = int(tpint)
	this.FirstPage = 1
	this.execUrl()
}

//格式化URL地址
func (this *Page) execUrl() {
	if strings.Contains(this.Url, "?") {
		this.Url = strings.Join([]string{this.Url, "&page="}, "")
	} else {
		this.Url = strings.Join([]string{this.Url, "?page="}, "")
	}
}

//获取URL组织
func (this *Page) getUrl(page int) string {
	return strings.Join([]string{this.Url, strconv.Itoa(page)}, "")
}

func (this *Page) getNum(page int, num int) (int, int) {
	var k = 1
	var v = 10
	if page > 5 {
		k = page - 4
		v = page + 5
		if v > num {
			k = num - 9
			v = num
		}
	} else {
		if v > num {
			v = num
		}
	}
	return k, v
}

func (this *Page) Show() string {
	this.getPageCount()
	var buf bytes.Buffer
	buf.WriteString(" <a class = \"layui-laypage-prev\" href=\"")
	buf.WriteString(this.getUrl(1))
	buf.WriteString("\">首页</a> ")

	if this.PageNo > 1 {
		buf.WriteString("<a class = \"layui-laypage-prev\" href=\"")
		buf.WriteString(this.getUrl(this.PageNo - 1))
	} else {
		buf.WriteString("<a class = \"layui-laypage-prev layui-disabled\" href=\"")
	}
	buf.WriteString("\">上一页</a>")

	k, v := this.getNum(this.PageNo, this.TotalPage)
	for i := k; i <= v; i++ {
		if i == this.PageNo {
			buf.WriteString("<span class=\"layui-laypage-curr\"><em class=\"layui-laypage-em\"></em><em>")
			buf.WriteString(strconv.Itoa(i))
			buf.WriteString("</em></span>")
		} else {
			buf.WriteString("<a href=\"")
			buf.WriteString(this.getUrl(i))
			buf.WriteString("\">")
			buf.WriteString(strconv.Itoa(i))
			buf.WriteString("</a>")
		}
	}

	if this.PageNo < this.TotalPage {
		buf.WriteString("<a class=\"layui-laypage-next\" href=\"")
		var nextPage int = this.PageNo + 1
		buf.WriteString(this.getUrl(nextPage))
	} else {
		buf.WriteString("<a class=\"layui-laypage-next layui-disabled\"")
	}
	buf.WriteString("\">下一页</a>")
	buf.WriteString("<a href=\"")
	buf.WriteString(this.getUrl(this.TotalPage))
	buf.WriteString("\">尾页</a>")
	buf.WriteString("<span class=\"layui-laypage-count\">共")
	buf.WriteString(strconv.Itoa(this.TotalCount))
	buf.WriteString("条</span>")
	return buf.String()
}
