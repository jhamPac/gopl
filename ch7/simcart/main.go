package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var listHTML = template.Must(template.New("list").Parse(`
<html>
<body>
<table>
	<tr>
		<th>item</th>
		<th>price</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td>{{$k}}</td>
		<td>{{$v}}</td>
	</tr>
{{end}}
</table>
</body>
</html>
`))

// PriceDB is the main database that houses all the items and prices
type PriceDB struct {
	sync.Mutex
	db map[string]int
}

// Create an item in the database
func (p *PriceDB) Create(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "no integer price given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; ok {
		http.Error(w, fmt.Sprintf("%s already exists", item), http.StatusBadRequest)
		return
	}

	p.Lock()
	if p.db == nil {
		p.db = make(map[string]int, 0)
	}
	p.db[item] = price
	p.Unlock()
	w.Write([]byte("successfully created"))
	time.AfterFunc(2*time.Second, func() {
		http.Redirect(w, r, "http://localhost:9000/", http.StatusSeeOther)
	})
}

// Update an item in the database
func (p *PriceDB) Update(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "no integer price given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s does not exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	p.db[item] = price
	p.Unlock()
}

// Delete an item from the database
func (p *PriceDB) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s does not exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	delete(p.db, item)
	p.Unlock()
}

// Read all the products from the simcart inventory
func (p *PriceDB) Read(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s does not exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	fmt.Fprintf(w, "%s: %d\n", item, p.db[item])
	p.Unlock()
}

// List generates an HTML output of the simcart inventory
func (p *PriceDB) List(w http.ResponseWriter, r *http.Request) {
	p.Lock()
	if err := listHTML.Execute(w, p.db); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
	p.Unlock()
}

func main() {
	p := &PriceDB{}
	p.db = make(map[string]int, 0)
	p.db["shirt"] = 40
	http.HandleFunc("/create", p.Create)
	http.HandleFunc("/read", p.Read)
	http.HandleFunc("/update", p.Update)
	http.HandleFunc("/delete", p.Delete)
	http.HandleFunc("/list", p.List)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to SimCart 🛒")
	})
	fmt.Println("SimCart is running on port 9000")
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
