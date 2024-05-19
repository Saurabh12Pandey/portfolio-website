package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
)

type Server struct {
	mongoClient *mongo.Client
}

type Home struct {
	ID                 string `json:"id" bson:"_id,omitempty"`
	AboutSectionHeader string `json:"about_section_header" bson:"about_section_header"`
	AboutMe            string `json:"about_me" bson:"about_me"`
}

type Experience struct {
	ID                   string `json:"id" bson:"_id,omitempty"`
	ExperienceTitle      string `json:"experienceTitle" bson:"experienceTitle"`
	ExperienceDuration   string `json:"experienceDuration" bson:"experienceDuration"`
	ExperienceDescription string `json:"experienceDescription" bson:"experienceDescription"`
}

type Project struct {
	ID              string   `json:"id" bson:"_id,omitempty"`
	ProjectName     string   `json:"projectName" bson:"projectName"`
	ProjectSummary  string   `json:"projectSummary" bson:"projectSummary"`
	TechnologiesUsed []string `json:"technologiesUsed" bson:"technologiesUsed"`
	ImageLink       string   `json:"imagelink" bson:"imagelink"`
}

type Skills struct {
	ID     string   `json:"id" bson:"_id,omitempty"`
	Skills []string `json:"skills" bson:"skills"`
}



func main() {
	// MongoDB connection URI
	uri := "mongodb://127.0.0.1:27017"

	// Create a new MongoDB client and connect to the server
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
			log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
			log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Create a new server instance
	server := &Server{mongoClient: client}

	// Create a new mux router
	router := mux.NewRouter()

	// Define the routes and handlers
	router.HandleFunc("/home", server.getHomeData)
	router.HandleFunc("/experience", server.getExperiences)
	router.HandleFunc("/projects", server.getProjects)
	router.HandleFunc("/skills", server.getSkills)

	// CORS middleware
	router.Use(corsMiddleware)

	// Start the HTTP server with the mux router
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}


func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
			if r.Method == "OPTIONS" {
					return
			}
			next.ServeHTTP(w, r)
	})
}


func (s *Server) getSkills(w http.ResponseWriter, r *http.Request) {
	// Select the MongoDB collection
	collection := s.mongoClient.Database("saurabh").Collection("skills")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find the skills document in the collection
	var skills Skills
	err := collection.FindOne(ctx, map[string]interface{}{}).Decode(&skills)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the skills as JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(skills); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (s *Server) getProjects(w http.ResponseWriter, r *http.Request) {
	// Select the MongoDB collection
	collection := s.mongoClient.Database("saurabh").Collection("projects")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find all projects in the collection
	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Decode the projects into a slice
	var projects []Project
	if err := cursor.All(ctx, &projects); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the projects as JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (s *Server) getHomeData(w http.ResponseWriter, r *http.Request) {
	// Select the MongoDB collection
	collection := s.mongoClient.Database("saurabh").Collection("portfolio")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find all items in the collection
	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	defer cursor.Close(ctx)

	// Decode the items into a slice
	var items []Home
	if err := cursor.All(ctx, &items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	// Encode the items as JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getExperiences(w http.ResponseWriter, r *http.Request) {
	// Select the MongoDB collection
	collection := s.mongoClient.Database("saurabh").Collection("experience")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find all experiences in the collection
	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Decode the experiences into a slice
	var experiences []Experience
	if err := cursor.All(ctx, &experiences); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the experiences as JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(experiences); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
