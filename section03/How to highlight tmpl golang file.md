# Highlighting .tmpl golang file

I found the way of highlighting .tmpl file type or golang literals inside go files in visual studio code. Also it allows emmet and format in .tmpl files exactly same as .html files.

1. Install [this vs code extension](https://marketplace.visualstudio.com/items?itemName=jinliming2.vscode-go-template) (Go Template Support by jinliming2). It highlights go literals and go templates
2. Add following lines in vs code settings.json. It makes vs code handle .tmpl files as .html files, which enables emmet and formatting .tmpl as same as html.

```
"files.associations": {
    "*.tmpl": "html"
  }
```

That's all. It will highlight your go template's own syntaxes like`{{define "some"}}` and `{{end}}`, and highlight all the html tags in .tmpl files. And also highlights the literals in the go files.

By the way, I'm learning go with the udemy and the lecture uses .tmpl files a lot. But is .tmpl files used in the reality? Or is it just for the education? I thought it's unpopular because the number of downloads of the extension i attached, or myabe it's because go itself is unpopular. Anyway, do you guys use .tmpl file?



Source:

https://www.reddit.com/r/golang/comments/11ffy5f/i_found_the_way_of_highlighting_tmpl_file_type_or/



**bubbzDotDev**

Thank you! I'm doing the Golang Udemy course by Trevor Sawler and he just started using `.tmpl` files so this is exactly what I needed.