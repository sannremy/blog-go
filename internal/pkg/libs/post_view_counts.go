package libs

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// PostViewCounts contains post view counts
var PostViewCounts = make(map[string]int64)

type postType struct {
	Slug  string `firestore:"slug"`
	Views int64  `firestore:"views"`
}

func getClient() (*firestore.Client, context.Context, error) {
	// Create Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "srchea-com")
	return client, ctx, err
}

// InitPostViewCounts Initialize post view counts
func InitPostViewCounts() {
	PostViewCounts = getPostViewCounts()
}

func getPostViewCounts() map[string]int64 {
	// Create Firestore client
	client, ctx, err := getClient()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	// Set all post view counts from DB
	postViewCounts := make(map[string]int64)
	var postData postType
	iter := client.Collection("posts").Documents(ctx)
	for {
		doc, errDoc := iter.Next()
		if errDoc == iterator.Done {
			break
		}

		if errCast := doc.DataTo(&postData); errCast != nil {
			fmt.Println(errCast)
		}

		if postData.Slug != "" && postData.Views > 0 {
			postViewCounts[postData.Slug] = postData.Views
		}
	}

	return postViewCounts
}

// GetPostViewCount Get view count of a specific post
func GetPostViewCount(slug string) int64 {
	return PostViewCounts[slug]
}

// SetPostViewCount Set view count for a specific post
func SetPostViewCount(slug string, viewCount int64) {
	PostViewCounts[slug] = viewCount
}

// IncrementPostViewCount Increment by 1 view count
func IncrementPostViewCount(slug string) {
	if _, ok := PostViewCounts[slug]; ok {
		PostViewCounts[slug]++
	} else {
		PostViewCounts[slug] = 1
	}
}

// UpdateAllPosts Update all posts in DB
func UpdateAllPosts() {
	// Create Firestore client
	client, ctx, err := getClient()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	// Get post view counts from DB
	postViewCounts := getPostViewCounts()

	// Batch instance
	batch := client.Batch()
	hasUpdates := false

	// Populate all posts into the batch
	for slug, views := range PostViewCounts {
		// Get update from DB
		if viewsFromDb, postExists := postViewCounts[slug]; postExists {
			if viewsFromDb > views {
				views = viewsFromDb
			}
		}

		post := postType{
			Slug:  slug,
			Views: views,
		}

		ref := client.Collection("posts").Doc(slug)
		batch = batch.Set(ref, post)

		hasUpdates = true
	}

	// Commit batch
	if hasUpdates {
		_, errBatch := batch.Commit(ctx)
		if errBatch != nil {
			log.Printf("Cannot batch write (UpdateAllPosts): %s", errBatch)
		}
	}
}
