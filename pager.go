//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-22

package webr

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"

	"github.com/xanygo/anygo/store/xdb"
	"github.com/xanygo/anygo/xhtml"
)

type Pager1 struct {
	// Info 页面分页基础信息，必填
	Info xdb.Pagination

	// Request 当前页面的请求信息，必填，读取 RequestURI 用于生成分页链接
	Request *http.Request

	// Near 可选，默认值为3，当前页前后显示几页，总页数不超过 2 * Near +1
	Near int

	// PageLink 可选，自定义生成链接的方法
	PageLink func(r *http.Request, page int) string

	// DisableStat 不显示记录条数等统计信息
	DisableStat bool
}

func (p *Pager1) getPageLink(id int) string {
	if p.PageLink != nil {
		return p.PageLink(p.Request, id)
	}
	u, err := url.ParseRequestURI(p.Request.RequestURI)
	if err != nil {
		return err.Error()
	}
	qs := u.Query()
	qs.Set("page", strconv.Itoa(id))
	return u.Path + "?" + qs.Encode()
}

func (p *Pager1) pageLi(page int, txt string) xhtml.Element {
	href := p.getPageLink(page)
	str := fmt.Sprintf(`<li class="page-item"><a class="page-link" href=%q>%s</a></li>`, href, txt)
	return xhtml.HTMLString(str)
}

func (p *Pager1) disabled(txt string) xhtml.Element {
	str := `<li class="page-item disabled"><a class="page-link">` + txt + `</a></li>`
	return xhtml.HTMLString(str)
}

func (p *Pager1) getNear() int {
	if p.Near > 0 {
		return p.Near
	}
	return 3
}

func (p *Pager1) HTML() template.HTML {
	totalPage := p.Info.TotalPages()
	ul := xhtml.NewAny("ul")
	xhtml.SetClass(ul, "pagination", "justify-content-center")

	page := p.Info.Page

	if !p.DisableStat {
		tp := totalPage
		if p.Info.Total == 0 {
			tp = 0
		}
		ul.Add(p.disabled(fmt.Sprintf("共 %d 条记录 %d 页", p.Info.Total, tp)))
	}

	if page == 1 {
		ul.Add(p.disabled("首页"))
		ul.Add(p.disabled("上一页"))
	} else {
		ul.Add(p.pageLi(1, "首页"))
		ul.Add(p.pageLi(page-1, "上一页"))
	}

	ns, ne := p.Info.NearPages(p.getNear())
	for i := ns; i <= ne; i++ {
		if i == page {
			str := `<li class="page-item active"><a class="page-link">` + strconv.Itoa(i) + `</a></li>`
			li := xhtml.HTMLString(str)
			ul.Add(li)
		} else {
			ul.Add(p.pageLi(i, strconv.Itoa(i)))
		}
	}

	if page >= totalPage {
		ul.Add(p.disabled("下一页"))
		ul.Add(p.disabled("尾页"))
	} else {
		ul.Add(p.pageLi(page+1, "下一页"))
		ul.Add(p.pageLi(totalPage, "尾页"))
	}

	return ul.TplHTML()
}
