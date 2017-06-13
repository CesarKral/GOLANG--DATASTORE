package GAEb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func init() {

	http.Handle("/css/", http.FileServer(http.Dir("./xfklm")))
	http.Handle("/js/", http.FileServer(http.Dir("./xfklm")))
	http.Handle("/img/", http.FileServer(http.Dir("./xfklm")))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		templates.ExecuteTemplate(res, "main", nil)
	})

	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {

	})

	http.HandleFunc("/logout", func(res http.ResponseWriter, req *http.Request) {

	})
	//WORKS
	http.HandleFunc("/create", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			xa, _ := strconv.Atoi(req.FormValue("age"))
			user := &User{
				Name:       req.FormValue("name"),
				Age:        xa,
				Car:        req.FormValue("car"),
				Subscribed: time.Now(),
			}
			ctx := appengine.NewContext(req)
			key := datastore.NewKey(ctx, "User", req.FormValue("email"), 0, nil)
			_, err := datastore.Put(ctx, key, user)
			if err != nil {

			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "create", nil)
	})
	//WORKS
	http.HandleFunc("/get", func(res http.ResponseWriter, req *http.Request) {
		var user User
		if req.Method == "POST" {
			//var user User
			ctx := appengine.NewContext(req)
			key := datastore.NewKey(ctx, "User", req.FormValue("email"), 0, nil)

			err := datastore.Get(ctx, key, &user)
			if err != nil {

			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "get", &User{Name: user.Name, Age: user.Age, Car: user.Car})
	})
	//WORKS
	http.HandleFunc("/update", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			var user User
			ctx := appengine.NewContext(req)
			key := datastore.NewKey(ctx, "User", req.FormValue("email"), 0, nil)

			ea := datastore.Get(ctx, key, &user)
			if ea != nil {

			}
			var aa string
			var bb string
			var cc int

			if req.FormValue("name") == "" {
				aa = user.Name
			} else {
				aa = req.FormValue("name")
			}
			if req.FormValue("car") == "" {
				bb = user.Car
			} else {
				bb = req.FormValue("car")
			}
			if req.FormValue("age") == "" {
				cc = user.Age
			} else {
				cc, _ = strconv.Atoi(req.FormValue("age"))
			}
			userx := &User{
				Name: aa,
				Age:  cc,
				Car:  bb,
			}
			_, eb := datastore.Put(ctx, key, userx)
			if eb != nil {

			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "update", nil)
	})
	//WORKS
	http.HandleFunc("/delete", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {

			ctx := appengine.NewContext(req)
			key := datastore.NewKey(ctx, "User", req.FormValue("email"), 0, nil)

			err := datastore.Delete(ctx, key)
			if err != nil {

			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "delete", nil)
	})
	//WORKS
	http.HandleFunc("/query", func(res http.ResponseWriter, req *http.Request) {
		var user []User
		if req.Method == "POST" {
			ctx := appengine.NewContext(req)
			//var user []User
			query := datastore.NewQuery("User").Filter("Car =", req.FormValue("car")).Order("-Age").Limit(10)
			_, err := query.GetAll(ctx, &user)
			if err != nil {

			}
			//for _, p := range user {
			//	fmt.Println(p.Name)
			//}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "query", user)
	})
	//WORKS
	http.HandleFunc("/ancestor", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			ctx := appengine.NewContext(req)
			employee := &Employee{
				Name:    "Cesar",
				Age:     34,
				Company: "BMW",
				Country: "Spain",
			}
			empKey, ea := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Employee", nil), employee)
			if ea != nil {

			}

			address := &Address{
				Street:     "x street",
				Number:     1,
				Letter:     "C",
				Floor:      7,
				PostalCode: "50100",
			}
			addKey := datastore.NewIncompleteKey(ctx, "Address", empKey)
			_, eb := datastore.Put(ctx, addKey, address)
			if eb != nil {

			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "ancestor", nil)

	})
	//WORKS
	http.HandleFunc("/ancestorqueries", func(res http.ResponseWriter, req *http.Request) {
		var girl []Girlfriend
		ctx := appengine.NewContext(req)
		cesarKey := datastore.NewKey(ctx, "User", "cesar@a.com", 0, nil)
		query := datastore.NewQuery("Girlfriend").Ancestor(cesarKey)
		_, err := query.GetAll(ctx, &girl)
		if err != nil {

		}
		if req.Method == "POST" {

			natalia := &Girlfriend{
				Name: "Natalia",
				Age:  30,
				Eyes: "Brown",
			}
			nataliaKey := datastore.NewKey(ctx, "Girlfriend", "", 0, cesarKey)
			_, ea := datastore.Put(ctx, nataliaKey, natalia)
			if ea != nil {

			}

			nuria := &Girlfriend{
				Name: "Nuria",
				Age:  40,
				Eyes: "Blue",
			}
			nuriaKey := datastore.NewKey(ctx, "Girlfriend", "", 0, cesarKey)
			_, eb := datastore.Put(ctx, nuriaKey, nuria)
			if eb != nil {

			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "ancestorqueries", girl)
	})
	//WORKS
	http.HandleFunc("/iterator", func(res http.ResponseWriter, req *http.Request) {
		var users []User
		ctx := appengine.NewContext(req)
		query := datastore.NewQuery("User")
		iter := query.Run(ctx)
		for {
			var user User
			_, ea := iter.Next(&user) //Returns Key
			if ea == datastore.Done {
				break
			}
			users = append(users, user)
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "iterator", users)
	})
	//WORKS
	http.HandleFunc("/projection", func(res http.ResponseWriter, req *http.Request) {
		var ageCars []Xproj
		ctx := appengine.NewContext(req)
		query := datastore.NewQuery("User").Project("Age", "Car")
		iter := query.Run(ctx)
		for {
			var ageCar Xproj
			_, ea := iter.Next(&ageCar) //Returns Key
			if ea == datastore.Done {
				break
			}
			ageCars = append(ageCars, ageCar)
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "projection", ageCars)
	})
	//WORKS
	http.HandleFunc("/keysonly", func(res http.ResponseWriter, req *http.Request) {
		var xKeys []int64
		ctx := appengine.NewContext(req)
		query := datastore.NewQuery("Girlfriend").KeysOnly()
		iter := query.Run(ctx)
		for {
			key, ea := iter.Next(nil) //Returns Key
			if ea == datastore.Done {
				break
			}
			xKeys = append(xKeys, key.IntID())
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "keysonly", xKeys)

	})
	//WORKS
	http.HandleFunc("/keysonlystring", func(res http.ResponseWriter, req *http.Request) {
		var xKeys []string
		ctx := appengine.NewContext(req)
		query := datastore.NewQuery("User").KeysOnly()
		iter := query.Run(ctx)
		for {
			key, ea := iter.Next(nil) //Returns Key
			if ea == datastore.Done {
				break
			}
			xKeys = append(xKeys, key.StringID())
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "keysonlystring", xKeys)

	})
	//WORKS
	http.HandleFunc("/mail", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			const confirmMessage = `
		Thank you for creating an account!
		Please confirm your email address by clicking on the link below:

		%s
		`
			ctx := appengine.NewContext(req)
			addr := req.FormValue("email")
			url := `This should be a link to confirm account`
			msg := &mail.Message{
				Sender:  "xxxx@gmail.com",
				To:      []string{addr},
				Subject: "Confirm your registration",
				Body:    fmt.Sprintf(confirmMessage, url),
			}
			if err := mail.Send(ctx, msg); err != nil {
				log.Errorf(ctx, "Couldn't send email: %v", err)
			}
		}
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		templates.ExecuteTemplate(res, "mail", nil)
	})

	http.HandleFunc("/upload", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			ctx := appengine.NewContext(req)
			client, _ := storage.NewClient(ctx)
			file, fileHeader, _ := req.FormFile("filedata")
			defer file.Close()
			obj := client.Bucket("xxxx.appspot.com").Object(fileHeader.Filename)
			//obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader)
			sw := obj.NewWriter(ctx)
			//sw.ContentType = "image/jpeg"
			sw.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
			_, err := io.Copy(sw, file)
			if err != nil {

			}
			if err := sw.Close(); err != nil {
				io.WriteString(res, err.Error())
			}
			http.Redirect(res, req, "/", 301)
		}
		templates.ExecuteTemplate(res, "upload", nil)

	})

	http.HandleFunc("/getimagelink", func(res http.ResponseWriter, req *http.Request) {
		ctx := appengine.NewContext(req)
		client, _ := storage.NewClient(ctx)
		x := client.Bucket("xxxx.appspot.com").Object("banderas.jpg")
		y, _ := x.Attrs(ctx)
		htmp := `<!doctype html>

		<html lang="en">
		<head>
		  <meta charset="utf-8">

		  <title>The HTML5 Herald</title>
		  <meta name="description" content="Image Test">
		  <meta name="author" content="Cesar">
		</head>

		<body>
		  <img src = ` + y.MediaLink + `>
		</body>
		</html>`
		io.WriteString(res, htmp)
	})

}
