package server

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Server) imageUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(20 << 20)
		if err != nil {
			s.error(w, http.StatusBadRequest, fmt.Errorf("error parsing a formdata %v", err))
			return
		}

		// Get the image file from the form data
		fileHeaders := r.MultipartForm.File["images"]

		if len(r.MultipartForm.Value["path"]) == 0 {
			s.error(w, http.StatusBadRequest, errors.New("path variables are not provided"))
		}
		keys := r.MultipartForm.Value["path"][0]

		var paths []string
		for _, fileHeader := range fileHeaders {
			// Open uploaded file
			file, err := fileHeader.Open()
			if err != nil {
				s.error(w, http.StatusInternalServerError, fmt.Errorf("error while opening an uploaded file %v", err))
				return
			}
			defer file.Close()

			// Check file content type
			contentType := fileHeader.Header.Get("Content-Type")
			if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
				s.error(w, http.StatusBadRequest, errors.New("only JPG or PNG or GIF files are allowed"))
				return
			}

			err = uploadToS3(file, keys+fileHeader.Filename)
			if err != nil {

				fmt.Println(err)
				s.error(w, http.StatusBadRequest, fmt.Errorf("failed to upload file to S3 %v", err))
				return
			}
			paths = append(paths, "https://sn-media.s3.amazonaws.com/"+keys+fileHeader.Filename)

		}

		userID, ok := r.Context().Value(ctxUserID).(string)
		if !ok {
			s.error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		options := r.PathValue("option")
		switch options {
		case "avatar":
			fmt.Println("avatar changed")
			if err := s.store.User().SetAvatar(userID, paths[0]); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}
		case "cover":
			fmt.Println("cover changed")
			if err := s.store.User().SetCover(userID, paths[0]); err != nil {
				s.error(w, http.StatusUnprocessableEntity, err)
				return
			}
		default:
			fmt.Println("default changed")

			parent_id := r.PathValue("parent_id")

			err = s.store.Image().Add(parent_id, paths)
			if err != nil {
				s.error(w, http.StatusBadRequest, err)
				return
			}
		}

		s.respond(w, http.StatusOK, Response{
			Data: "Successfully uploaded to S3",
		})
	}
}

func uploadToS3(file multipart.File, fileName string) error {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(os.Getenv(awsAccessKey), os.Getenv(awsSecretKey), ""),
	})
	if err != nil {
		return err
	}
	svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return err
	}

	return nil
}
