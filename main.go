/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package main

import (
	"context"
	"fmt"
	"log"

	// "testing/quick"
	"time"

	// "github.com/go-delve/delve/pkg/proc/test"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/gitpod/mycli/cmd"
	// "google.golang.org/api/ids/v1"
	// "google.golang.org/genproto/googleapis/cloud/ids/v1"
)


func main() {
	// cmd.Execute()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/test?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	keployDatabase := client.Database("keploy-prod")
	keployCollection := keployDatabase.Collection("tests")
	testcaseCollection := keployDatabase.Collection("testcase")
	t := time.Unix(1645531701, 0).UTC()
	fmt.Println(t)
	y := time.Date(2022, 2, 23, 0, 0, 0, 0, time.UTC)
	fmt.Println(y)
	z := y.Unix()
	fmt.Println(z)
	// opts2 := options.Delete()
	// result, err := keployCollection.DeleteMany(ctx, bson.D{{"completed", bson.D{{"$lte", 1645574400}}}},opts2)
	// result, err := keployCollection.DeleteMany(ctx, bson.D{{"run_id", bson.D{{"$lte", "77265b38-3763-4719-9b4c-957192a27278"}}}})
	// result, err := keployCollection.DeleteMany(ctx, bson.M{"uri":"/:*"})
	// ids,err:=testcaseCollection.Find(ctx,bson.M{"app_id":"TestKeploy"})
	// println(ids.All(ctx,ids))
	cursor, err := testcaseCollection.Find(ctx, bson.M{"app_id": "TestAPP"})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var ids []string
	// ids = make(string, 0)
	for cursor.Next(ctx) {
		var testcase bson.M
		if err = cursor.Decode(&testcase); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(testcase["_id"].(string))
		ids = append(ids, testcase["_id"].(string))

	}

	println(len(ids))
	cnt:=0
	for i := 0; i < len(ids); i++ {
		println(ids[i])
		for j := 0; j < 1; j++ {
			
			result, err := keployCollection.DeleteOne(ctx, bson.M{"test_case_id": ids[i]})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
			if result.DeletedCount==1 {
				cnt++
			}else{
				break
			}
		}
		println("Testcases of ",ids[i]," deleted")
	}
	println("TOTAL TESTCASES DELETED ",cnt)

}