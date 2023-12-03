package main

import (
    "net/http"
    "log"
    "html/template"
)

type PageData struct {
    Title string
}

func main() {

    // Serve the pages
    http.HandleFunc("/", servePage("Home", "./src/index.html"))
    http.HandleFunc("/contact", servePage("Contact", "./src/contact.html"))
    http.HandleFunc("/aboutme", servePage("About Me", "./src/aboutme.html"))
    http.HandleFunc("/projects", servePage("Projects", "./src/projects.html"))

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./src/css"))))
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./src/js"))))

    // Serve images from the "assets/img/" directory.
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./assets/img/"))))

    // Start the server on port 8081 and log if there's an error.
    log.Println("Server starting on port 8081...")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}

func servePage(title, filePath string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
       // Parse the main template
       tmpl, err := template.New("").ParseFiles(filePath)
       if err != nil {
           log.Println("Error parsing main template:", err)
           http.Error(w, "Internal Server Error", http.StatusInternalServerError)
           return
       }

       // Parse the partials
       tmpl, err = tmpl.ParseGlob("src/partials/*.html")
       if err != nil {
           log.Println("Error parsing partials:", err)
           http.Error(w, "Internal Server Error", http.StatusInternalServerError)
           return
       }
    
       data := PageData{Title: title}
       err = tmpl.ExecuteTemplate(w, "main", data)
       if err != nil {
           log.Println("Error executing template:", err)
           http.Error(w, "Internal Server Error", http.StatusInternalServerError)
       }
    }
}