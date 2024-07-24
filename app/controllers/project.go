package controllers

import (
	"context"
	"encoding/json"
	"go-mongodb/app/models"
	"go-mongodb/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var GetProjectListEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var projects = []*models.Projects{}

	collection := config.DB.Collection("projects")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		config.ErrorResponse(err.Error(), response)
		return
	}

	if err = cursor.All(context.TODO(), &projects); err != nil {
		config.ErrorResponse(err.Error(), response)
		return
	}

	json.NewEncoder(response).Encode(projects)
})

var GetProjectEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var project models.Projects
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	collection := config.DB.Collection("projects")
	err := collection.FindOne(context.TODO(), filter).Decode(&project)
	if err != nil {
		config.ErrorResponse("Project doesn't exist", response)
		return
	}
	json.NewEncoder(response).Encode(project)
})

var CreateProjectEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var req models.ProjectRequest
	collection := config.DB.Collection("projects")

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		config.ErrorResponse(err.Error(), response)
		return
	}

	project := models.Projects{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Link:        req.Link,
		Start:       req.Start,
		End:         req.End,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), project)

	if err != nil {
		config.ErrorResponse(err.Error(), response)
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(project)
})

var DeleteProjectEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	collection := config.DB.Collection("projects")
	result, _ := collection.DeleteOne(context.TODO(), filter)
	if result.DeletedCount == 0 {
		config.ErrorResponse("Project doesn't exist", response)
		return
	}
	config.SuccessResponse(`Project with ID `+params["id"]+` deleted successfully`, response)
})

var UpdateProjectEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var project models.Projects
	var req models.ProjectRequest
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := config.DB.Collection("projects")

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		config.ErrorResponse(err.Error(), response)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": req}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	err = collection.FindOne(context.TODO(), filter).Decode(&project)
	if err != nil {
		config.ErrorResponse(err.Error(), response)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(project)
})
