package main
import(
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/new", NewHandler)
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintf(w, "Hello from koyeb")
}

func NewHandler(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintf(w, "Hello from koyeb New Route")
}