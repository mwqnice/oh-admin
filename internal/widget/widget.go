package widget

import (
	"github.com/mwqnice/oh-admin/pkg/utils"
	"html/template"
)

func Query(text string) template.HTML {
	return template.HTML("<button class=\"layui-btn\" lay-submit=\"\" lay-filter=\"searchForm\" id=\"search\"><i class=\"layui-icon\">\uE615</i>" + text + "</button>")
}
func Add(text string) template.HTML {
	return template.HTML("<a href=\"javascript:\" class=\"layui-btn btnOption  layui-btn-small btnadd\" id=\"add\" data-param=\"{}\" lay-event=\"add\"><i class=\"layui-icon layui-icon-add-1\"></i>" + text + "</a>")
}
func Edit(text string) template.HTML {
	return template.HTML("<a class=\"layui-btn layui-btn-xs btnEdit\" lay-event=\"edit\" title=\"" + text + "\"><i class=\"layui-icon\">\uE642</i>" + text + "</a>")
}
func Delete(text string) template.HTML {
	return template.HTML("<a class=\"layui-btn layui-btn-danger layui-btn-xs btnDel\" lay-event=\"del\" title=\"" + text + "\"><i class=\"layui-icon\">\uE640</i>" + text + "</a>")
}

func Expand(text string) template.HTML {
	return template.HTML("<a href=\"javascript:\" class=\"layui-btn btnOption layui-btn-normal layui-btn-small btnexpand\" id=\"expand\" data-param=\"{}\" lay-event=\"expand\"><i class=\"layui-icon layui-icon-shrink-right\"></i> " + text + "</a>")
}

func Collapse(text string) template.HTML {
	return template.HTML("<a href=\"javascript:\" class=\"layui-btn btnOption layui-btn-warm layui-btn-small btncollapse\" id=\"collapse\" data-param=\"{}\" lay-event=\"collapse\"><i class=\"layui-icon layui-icon-spread-left\"></i>" + text + "</a>")
}

func Addz(text string) template.HTML {
	return template.HTML("<a href=\"javascript:\" class=\"layui-btn btnOption layui-btn-normal layui-btn-xs btnaddz\" id=\"addz\" data-param=\"{}\" lay-event=\"addz\"><i class=\"layui-icon layui-icon-add-1\"></i> " + text + "</a>")
}
func In(target int, array []int) bool {
	return utils.InIntArray(target, array)
}
