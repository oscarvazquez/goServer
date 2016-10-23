package models

import (
	"time"
)

const contentType = "application/octet-stream"
var acl string

type (
	Scene struct {
		uid 		string		`json:"uid"`
		Username	string 		`json:"username"`			
		Name		string		`json:"name"`
		token   	string  	`json:"token"`
		Public  	bool 		`json:"public"`
		Data		[]byte		`json:"data"`
		Length		int			`json:"seconds"`		
		Timestamp	time.Time 	`json:"created_at"`
	}

)

// func (s *Scene) UploadToS3() error {
// 	// There should be no name conflicts, usernames should be unique + time.Now()
// 	// I might have to consider users from different consoles using the same username
// 	storage_path := fmt.Sprintf("%v/%v", s.Username, time.Now())

// 	bucket := S3Bucket //You still need your s3 keys

// 	b := new(bytes.Buffer)
// 	encodeErr := json.NewEncoder(b).Encode(s)
// 	if encodeErr != nil {
// 		return encodeErr
// 	}

// 	if s.Public {
// 		acl = s3.Public
// 	} else {
// 		acl = s3.Private
// 	}

// 	return bucket.PutReader(storage_path, b, int64(b.Len()), contentType, acl, s3.Options{})
// }