package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RakaiSeto/golangWebDevTodd/171_Go-mongo/config"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Isbn   string  // `json:"isbn" bson:"isbn"`
	Title  string  // `json:"title" bson:"title"`
	Author string  // `json:"author" bson:"author"`
	Price  float32 // `json:"price" bson:"price"`
}

func AllBooks() ([]Book, error) {
	bks := []Book{}
	err := config.Books.Find(bson.M{}).All(bks)
	if err != nil {
		return nil, err
	}
	return bks, nil
}

func OneBook(req *http.Request) (Book, error) {
	bk := Book{}
	isbn := req.FormValue("isbn")
	err := config.Books.Find(bson.M{"isbn": isbn}).One(bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func InputBook(req *http.Request) (Book, error) {
	// get form Values
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	price := req.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || price == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Invalid price.")
	}
	bk.Price = float32(f64)

	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func UpdateBook(req *http.Request) (Book, error) {
	// get form Values
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	price := req.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || price == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	f64, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Invalid price.")
	}
	bk.Price = float32(f64)

	err = config.Books.Update(bson.M{"isbn" : bk.Isbn}, &bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func DeleteBook(req *http.Request) error {
	isbn := req.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}