package handlers

import "github.com/m4tty/palaver/data/bundles"
import "github.com/m4tty/palaver/web/domain/bundles"
import "html/template"
import "appengine"
import "fmt"
import "net/http"
import "github.com/gorilla/mux"

type testtest struct {
	Whatever string
	Messages []string
}

var templates *template.Template

func ServeUserHome(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	dashHomeTemplatePath := "web/static/dist/home.html"
	c.Infof("test test -------- ---- ------- ")
	// if templates == nil {
	templates = template.New("homeTemplate")
	_, err := templates.ParseFiles(dashHomeTemplatePath)
	if err != nil {
		c.Infof("err", fmt.Sprintf("%v", err))
	}

	// } else {
	// 	// _, err := templates.New("homeTemplate").ParseFiles(dashHomeTemplatePath)
	// 	// c.Infof("err****", fmt.Sprintf("%v", err))
	// }

	// testObj := testtest{
	// 	Whatever: "woot",
	// 	Messages: []string{"hello", "name1", "name2"},
	// }
	vars := mux.Vars(r)
	userId := vars["userId"]

	dataManager := bundleDataMgr.GetDataManager(&c)
	dataMgr := bundlesDomain.NewBundlesMgr(dataManager)
	results, _ := dataMgr.GetBundlesByUserId(userId)
	// if dataErr != nil {
	// 	//execute the template w/ an error
	// }

	c.Infof("results... ", fmt.Sprintf("%v", results))
	if err := templates.ExecuteTemplate(w, "homeTemplate", results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//http.ServeFile(w, r, "web/static/dist/home.debug.html")
}
