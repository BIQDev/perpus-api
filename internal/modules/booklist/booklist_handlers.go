package booklist

import (
	"context"
	"encoding/json"
	"github.com/BIQDev/perpus-api/internal/db"
	"github.com/BIQDev/perpus-api/internal/helper"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type IMongoBooklist struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Title    string             `json:"title" bson:"title"`
	Image    string             `json:"image" bson:"image"`
}

type booklistHandlers struct {
}

func (*booklistHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)

	username := params["username"]
	imgPathDir := "web/assets/" + username

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bookTotal, err := db.Mongo.DB().Collection("booklist").CountDocuments(ctx, bson.M{ "username": username })

	if err != nil {
		helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	if bookTotal >= 20 {
		helper.WriteResponse(w, http.StatusBadRequest, "error", "Max record has been reached", nil)
		return
	}

	if err = r.ParseMultipartForm(10 << 20); err != nil {
		log.Println(err)
		helper.WriteResponse(w, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}
	uploadedFile, handler, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		helper.WriteResponse(w, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}
	defer uploadedFile.Close()

	if _, err = os.Stat(imgPathDir); os.IsNotExist(err) {
		err = os.MkdirAll(imgPathDir, 0755)
		if err != nil {
			log.Println(err)
			helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
			return
		}
	}

	imgPathFile := imgPathDir + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + handler.Filename
	targetFile, err := os.OpenFile(imgPathFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		log.Println(err)
		helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	coll := db.Mongo.DB().Collection("booklist")
	var bookList = &IMongoBooklist{
		Username:  username,
		Title: r.FormValue("title"),
		Image: imgPathFile,
	}
	res, err := coll.InsertOne(ctx, bookList)

	if err != nil {
		log.Println(err)
		r := regexp.MustCompile(`(dup key: )(.*)(}]},)`)
		dupCheck := r.FindStringSubmatch(err.Error())
		errMsg := err.Error()
		if len(dupCheck) >= 2 {
			errMsg = "Error: unique field violation. Duplicate field: " + r.FindStringSubmatch(err.Error())[2]
		}
		helper.WriteResponse(w, http.StatusInternalServerError, "error", errMsg, nil)
		return
	}

	bookList.ID = res.InsertedID.(primitive.ObjectID)
	helper.WriteResponse(w, http.StatusOK, "success", "Success inserting data", bookList)
}

func (*booklistHandlers) Read(w http.ResponseWriter, r *http.Request) {
	var err error
	params := mux.Vars(r)

	username := params["username"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"_id", -1}})
	csr, err := db.Mongo.DB().Collection("booklist").Find(ctx, bson.M{ "username": username }, findOptions)

	if err != nil {
		log.Println(err)
		helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	var record []IMongoBooklist
	err = csr.All(ctx, &record)

	if err != nil {
		log.Println(err)
		helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	var resStr json.RawMessage
	resStr, err = json.Marshal(record)

	if err != nil {
		helper.WriteResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	helper.WriteResponse(w, http.StatusOK, "success", "", resStr)

}

var BooklistHandlers = &booklistHandlers{}
