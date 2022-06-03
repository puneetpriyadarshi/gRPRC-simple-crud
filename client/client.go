package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"root/prototype"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Helllo i am  a client")
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	c := prototype.NewUserProfilesClient(conn)
	router := gin.Default()

	router.POST("/user", func(ctx *gin.Context) {
		req := &prototype.UserProfile{}
		ctx.ShouldBindJSON(req)

		req2 := &prototype.CreateUserProfileRequest{
			UserProfile: req,
		}
		res, err := c.CreateUserProfile(context.Background(), req2)

		if err != nil {
			log.Fatalf("Error while creating user profile %v", err)
		}
		fmt.Println(res)
		ctx.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User Created",
		})
	})
	router.GET("/user/:id", func(ctx *gin.Context) {
		s := ctx.Param("id")
		req := &prototype.GetUserProfileRequest{
			Id: s,
		}
		res, err := c.GetUserProfile(context.Background(), req)
		if err != nil {
			log.Fatalf("Error while getting user profile %v", err)
		}
		fmt.Println(res)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User by ID",
			"data":    res,
		})
	})
	router.GET("/users", func(ctx *gin.Context) {
		req := &prototype.ListUsersProfilesRequest{}
		res, err := c.ListUsersProfiles(context.Background(), req)
		if err != nil {
			log.Fatalf("Error while getting user profile %v", err)
		}
		fmt.Println(res)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "All the Users",
			"data":    res,
		})
	})
	router.PUT("user/:id", func(ctx *gin.Context) {
		s := ctx.Param("id")
		
		req1 := &prototype.UserProfile{}
		ctx.ShouldBindJSON(req1)
		req2 := &prototype.UpdateUserProfileRequest{
			UserProfile: req1,
		}
		req2.UserProfile.Id = s

		fmt.Println("req------",req2.UserProfile.Id)
		// // req1 := &prototype.UpdateUserProfileRequest{}
		res, err := c.UpdateUserProfile(context.Background(), req2)

		if err != nil {
			log.Fatalf("Error while creating user profile %v", err)
		}
		fmt.Println(res)
	})
	router.DELETE("user/:id", func(ctx *gin.Context) {
		s := ctx.Param("id")
		req := &prototype.DeleteUserProfileRequest{Id: s}
		res, err := c.DeleteUserProfile(context.Background(), req)

		if err != nil {
			log.Fatalf("Error while creating user profile %v", err)
		}
		fmt.Println(res)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User Deleted",
		})
	})
	if err := router.Run(":5000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
