// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func FormPage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>URL Shortener</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"></head><body class=\"bg-light\"><div class=\"container mt-5\"><div class=\"row justify-content-center\"><div class=\"col-md-6\"><div class=\"card shadow\"><div class=\"card-body\"><h1 class=\"card-title text-center mb-4\">URL Shortener</h1><form method=\"POST\" action=\"/shorten\"><div class=\"mb-3\"><label for=\"name\" class=\"form-label\">Enter URL to shorten</label> <input type=\"url\" class=\"form-control\" id=\"name\" name=\"name\" required placeholder=\"https://example.com\"></div><div class=\"d-grid\"><button type=\"submit\" class=\"btn btn-primary\">Shorten URL</button></div></form></div></div></div></div></div><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ResultPage(newUrl string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Shortened URL</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\"></head><body class=\"bg-light\"><div class=\"container mt-5\"><div class=\"row justify-content-center\"><div class=\"col-md-6\"><div class=\"card shadow\"><div class=\"card-body\"><h1 class=\"card-title text-center mb-4\">Shortened URL</h1><div class=\"alert alert-success\" role=\"alert\">Your shortened URL is ready!</div><div class=\"mb-3\"><label class=\"form-label\">Your shortened URL:</label><div class=\"input-group\"><input type=\"text\" class=\"form-control\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("https://url-shortner-0vq3.onrender.com" + newUrl)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/main.templ`, Line: 58, Col: 136}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" readonly> <button class=\"btn btn-outline-secondary\" type=\"button\" onclick=\"copyToClipboard()\">Copy</button></div></div><div class=\"d-grid gap-2\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL("https://url-shortner-0vq3.onrender.com" + newUrl)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-primary\" target=\"_blank\">Visit URL</a> <a href=\"/\" class=\"btn btn-secondary\">Shorten Another URL</a></div></div></div></div></div></div><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\"></script><script>\n                function copyToClipboard() {\n                    var copyText = document.querySelector(\"input[readonly]\");\n                    copyText.select();\n                    document.execCommand(\"copy\");\n                    alert(\"Copied to clipboard: \" + copyText.value);\n                }\n            </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
