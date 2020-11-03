package main

import (
"context"
"flag"
"fmt"
"server/pb"
"github.com/gin-gonic/gin"
"google.golang.org/grpc"
)

var addr string
var port string
var apiLocation string

func init() {
	flag.StringVar(&addr, "server", "127.0.0.1:9999", "gRPC server address")
	flag.StringVar(&port, "port", "8080", "API Port")
	flag.Parse()
	apiLocation = fmt.Sprintf(":%s", port)
}

type Document struct {
	Text string `json:"text"`
}

func handleCategories(c *gin.Context) {
	var doc Document
	if c.ShouldBind(&doc) == nil {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		defer conn.Close()
		client := pb.NewBstoneClient(conn)
		req := pb.BstoneRequest{
			Text: doc.Text,
		}
		resp, err := client.Categories(context.Background(), &req)
		if err != nil {
			c.JSON(406, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": resp.Categories,
		})
	} else {
		c.String(406, "Malformed request")
	}
}

func handleEntities(c *gin.Context) {
	var doc Document
	if c.ShouldBind(&doc) == nil {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		defer conn.Close()
		client := pb.NewBstoneClient(conn)
		req := pb.BstoneRequest{
			Text: doc.Text,
		}
		resp, err := client.Entities(context.Background(), &req)
		if err != nil {
			c.JSON(406, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": resp.Entities,
		})
	} else {
		c.String(406, "Malformed request")
	}
}

func handleAbbreviations(c *gin.Context) {
	var doc Document
	if c.ShouldBind(&doc) == nil {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		defer conn.Close()
		client := pb.NewBstoneClient(conn)
		req := pb.BstoneRequest{
			Text: doc.Text,
		}
		resp, err := client.Abbreviations(context.Background(), &req)
		if err != nil {
			c.JSON(406, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": resp.Abbreviations,
		})
	} else {
		c.String(406, "Malformed request")
	}
}

func handleCompoundReferences(c *gin.Context) {
	var doc Document
	if c.ShouldBind(&doc) == nil {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		defer conn.Close()
		client := pb.NewBstoneClient(conn)
		req := pb.BstoneRequest{
			Text: doc.Text,
		}
		resp, err := client.CompoundReferences(context.Background(), &req)
		if err != nil {
			c.JSON(406, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": resp.References,
		})
	} else {
		c.String(406, "Malformed request")
	}
}



func main() {

	//gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.POST("/categories", handleCategories)
	route.POST("/entities", handleEntities)
	route.POST("/abbreviations", handleAbbreviations)
	route.POST("/compound-references", handleCompoundReferences)
	route.Run(apiLocation)
}
