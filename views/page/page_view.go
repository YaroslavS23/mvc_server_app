package page_view
//import
import (
	"fmt"
	"io"
	"text/template"

	"github.com/fops9311/mvc_server_app/tassets"
	layout "github.com/fops9311/mvc_server_app/views/layout"
)  //import
func Dummy(){
	fmt.Print("hi")
}

func renderTemplate(params map[string]interface{}, w io.Writer, templ string, templateName string) (err error) {
	tmpl, err := template.New(templateName).Parse(templ)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, params)
	if err != nil {
		return err
	}
	return nil
}


var Index func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, index, "page_Index")
}

var Edit func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, edit, "page_Edit")
}

var New func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, new, "page_New")
}

var Show func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, show, "page_Show")
}

var Create func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, create, "page_Create")
}

var Update func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, update, "page_Update")
}

var Delete func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, delete, "page_Delete")
}

var index string
var edit string
var new string
var show string
var create string
var update string
var delete string


func Init(){
	tassets.InitDir("./templates")
	index = tassets.GetAsset("templates/page/page_index.html")
	edit = tassets.GetAsset("templates/page/page_edit.html")
	new = tassets.GetAsset("templates/page/page_new.html")
	show = tassets.GetAsset("templates/page/page_show.html")
	create = tassets.GetAsset("templates/page/page_create.html")
	update = tassets.GetAsset("templates/page/page_update.html")
	delete = tassets.GetAsset("templates/page/page_delete.html")

	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

func init_continue() {
	edit = layout.Layout(edit)
}
