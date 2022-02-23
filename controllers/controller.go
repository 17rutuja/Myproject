package controllers

import (
	"context"
	"fmt"
	"golang_crudapp/database"
	"golang_crudapp/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "gopkg.in/mgo.v2/bson"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var validate = validator.New()

// Get all users
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		opts := options.Find().SetProjection(bson.M{
			"first_name": 1,
			"last_name":  1,
			"phone":      1,
			"email":      1,
		})

		res, err := userCollection.Find(ctx, bson.D{}, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var user []models.User
		e := res.All(ctx, &user)
		defer cancel()
		if e != nil {
			log.Fatal(e.Error())
			return
		}

		c.JSON(http.StatusOK, user)

	}
}

// Create user
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		//  json to struct i.e Binding
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validate the input
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		user.ID = primitive.NewObjectID()
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		result, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// Get single user
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		user_id := c.Param("user_id")

		objID, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			panic(err)
		}

		var user models.User
		errr := userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
		defer cancel()

		if errr != nil {
			fmt.Print(errr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// Delete user
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		user_id := c.Param("user_id")

		objID, objErr := primitive.ObjectIDFromHex(user_id)
		if objErr != nil {
			panic(objErr)
		}

		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objID})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

// Update user
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// get data and bind to struct
		user_id := c.Param("user_id")

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update object
		updateObj := primitive.D{}

		if user.First_name != "" {
			updateObj = append(updateObj, bson.E{"first_name", user.First_name})
		}
		if user.Last_name != "" {
			updateObj = append(updateObj, bson.E{"last_name", user.Last_name})
		}
		if user.Phone != "" {
			updateObj = append(updateObj, bson.E{"phone", user.Phone})
		}
		if user.Email != "" {
			updateObj = append(updateObj, bson.E{"email", user.Email})
		}

		updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", updated_at})

		// filter
		objID, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			panic(err)
		}
		filter := bson.M{"_id": objID}

		// opts like upsert
		upsert := true
		opts := options.UpdateOptions{
			Upsert: &upsert,
		}

		// actual collection call
		result, updateErr := userCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opts,
		)
		defer cancel()

		if updateErr != nil {
			fmt.Print(updateErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "There is some error while updating the user"})
			return
		}

		// response return
		c.JSON(http.StatusOK, result)
	}
}
