package element

import "golang.org/x/net/html"

type IsForm interface {
	IsFormElement()
	Element() *HtmlElement
	Render() string
}

type form struct {
	el *HtmlElement
}

func (f *form) AcceptCharset(charset string) *form {
	f.el.AddAttribute("accept-charset", charset)
	return f
}

func (f *form) Action(url string) *form {
	f.el.AddAttribute("action", url)
	return f
}

type AutoComplete bool

func (m AutoComplete) String() string {
	if m {
		return "on"
	}
	return "off"
}

func (f *form) AutoComplete(option AutoComplete) *form {
	f.el.AddAttribute("autocomplete", option.String())
	return f
}

type EncType string

const (
	XWWWFormUrlEncoded EncType = "application/x-www-form-urlencoded" //	Default. All characters are encoded before sent (spaces are converted to "+" symbols, and special characters are converted to ASCII HEX values)
	MultiPartFormData  EncType = "multipart/form-data"               //	 This value is necessary if the user will upload a file through the form
	TextPlainEncoding  EncType = "text/plain"                        //	Sends data without any encoding at all. Not recommended
)

func (f *form) EncType(enctype EncType) *form {
	f.el.AddAttribute("enctype", string(enctype))
	return f
}

type FormMethod string

const (
	GetFormMethod    FormMethod = "get"    //	Default. Appends the form data to the URL in name/value pairs: URL?name=value&name=value
	PostFormMethod   FormMethod = "post"   //	Sends the form data as an HTTP post transaction
	DialogFormMethod FormMethod = "dialog" //	Sends the form data as a modal dialog box
)

func (f *form) Method(method FormMethod) *form {
	f.el.AddAttribute("method", string(method))
	return f
}

func (f *form) Name(text string) *form {
	f.el.AddAttribute("name", text)
	return f
}

func (f *form) NoValidate() *form {
	f.el.AddAttributeBool("novalidate")
	return f
}

type FormRel string

const (
	FormRelExternal   FormRel = "external"   //	Specifies that the referenced document is not a part of the current site
	FormRelHelp       FormRel = "help"       //	Links to a help document
	FormRelLicense    FormRel = "license"    //	Links to copyright information for the document
	FormRelNext       FormRel = "next"       //	The next document in a selection
	FormRelNofollow   FormRel = "nofollow"   //	Links to an unendorsed document, like a paid link. ("nofollow" is used by Google, to specify that the Google search spider should not follow that link)
	FormRelNoOpener   FormRel = "noopener"   //
	FormRelNoReferrer FormRel = "noreferrer" //	Specifies that the browser should not send a HTTP referrer header if the user follows the hyperlink
	FormRelOpener     FormRel = "opener"     //
	FormRelPrev       FormRel = "prev"       //	The previous document in a selection
	FormRelSearch     FormRel = "search"     //	Links to a search tool for the document
)

func (f *form) Rel(rel FormRel) *form {
	f.el.AddAttribute("rel", string(rel))
	return f
}

type FormTarget string

const (
	FormTargetBlank  FormTarget = "_blank"  //The response is displayed in a new window or tab
	FormTargetSelf   FormTarget = "_self"   //The response is displayed in the same frame (this is default)
	FormTargetParent FormTarget = "_parent" //The response is displayed in the parent frame
	FormTargetTop    FormTarget = "_top"    //The response is displayed in the full body of the window
)

func (f *form) Target(target FormTarget) *form {
	f.el.AddAttribute("target", string(target))
	return f
}

func (f *form) Element() *HtmlElement { return f.el }
func (f *form) Render() string        { return f.el.Render() }

func Form(children ...IsForm) *form {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "form",
		},
	}

	for _, c := range children {
		el.Children(IsElement(c))
	}

	f := &form{el: el}

	return f
}
