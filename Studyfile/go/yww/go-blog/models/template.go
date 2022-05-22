package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index     TemplateBlog
	Category  TemplateBlog
	Custom    TemplateBlog
	Detail    TemplateBlog
	Login     TemplateBlog
	Pigenhole TemplateBlog
	Writing   TemplateBlog
}



func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(str []string, index int) string {
	return str[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	//return date.Format(time.RFC3339)
	return date.Format("2006-01-02 15:04:05")
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return 
		}
		//log.Println("error ", err)
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte(err.Error()))
		//log.Println("error ", err)
	}
}


func InitTemplate(templateDir string) (HtmlTemplate, error ){

	tp, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		//log.Println(err)
		return htmlTemplate,err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigenhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate,nil
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//访问首页模版由多个模板嵌套，解析文件时需要解析所有涉及到的
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		pagination := templateDir + "layout/pagination.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})

		t, err := t.ParseFiles(templateDir + viewName, home, header, footer, personal, post, pagination)

		if err != nil {
			log.Println("解析模版错误：", err)
			return nil, err
		}

		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil
}
