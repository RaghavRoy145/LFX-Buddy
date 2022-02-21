package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func handler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}
func healthHandler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
}
func getIpListHandler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("buddy-namespace").List(context.TODO(), metav1.ListOptions{})
	var buddyList []string
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		buddyList = append(buddyList, pod.Status.PodIP)
	}
	for _, buddy := range buddyList {
		w.Write([]byte(buddy))
		w.Write([]byte("\n"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	http.HandleFunc("/buddy/list", getIpListHandler)
	fmt.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
