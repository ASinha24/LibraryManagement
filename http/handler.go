package http

import (
	"encoding/json"
	"net/http"

	"github.com/ASinha24/LibraryManagementSystem"
	"github.com/ASinha24/LibraryManagementSystem/api"
	"github.com/ASinha24/LibraryManagementSystem/bookStore"
	"github.com/ASinha24/LibraryManagementSystem/http/utils"
	"github.com/gorilla/mux"
)

type bookHandler struct {
	libmanager LibraryManagementSystem.LibraryManager
	bookStore  bookStore.BookStore
}

func NewbookHandler(libmanager LibraryManagementSystem.LibraryManager, bookStore bookStore.BookStore) *bookHandler {
	return &bookHandler{
		libmanager: libmanager,
		bookStore:  bookStore,
	}
}

func (b *bookHandler) MuxInstaller(mux *mux.Router) {
	mux.Methods(http.MethodPost).Path("/bookAdd").HandlerFunc(b.AddNewBook)
	mux.Methods(http.MethodPut).Path("/bookUpdate/{bookID}").HandlerFunc(b.UpdateBookInfo)
	mux.Methods(http.MethodDelete).Path("/bookDelete/{bookID}").HandlerFunc(b.DeleteBookInfo)
	mux.Methods(http.MethodGet).Path("/find/{bookID}").HandlerFunc(b.FindBookByID)
	mux.Methods(http.MethodGet).Path("/findbyName/{bookName}").HandlerFunc(b.FindBookByName)
	mux.Methods(http.MethodGet).Path("/getAllBook").HandlerFunc(b.GetAllBooks)

}

func (b *bookHandler) AddNewBook(w http.ResponseWriter, req *http.Request) {
	createReq := &api.BookRequest{}
	if err := json.NewDecoder(req.Body).Decode(createReq); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}
	res, err := b.libmanager.AddBook(req.Context(), createReq)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusCreated, res, w)
}

func (b *bookHandler) UpdateBookInfo(w http.ResponseWriter, req *http.Request) {
	bookID := mux.Vars(req)["bookID"]
	updateReq := &api.BookRequest{}
	if err := json.NewDecoder(req.Body).Decode(updateReq); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}
	resp, err := b.libmanager.UpdateBook(req.Context(), bookID, updateReq)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusCreated, resp, w)
}

func (b *bookHandler) DeleteBookInfo(w http.ResponseWriter, req *http.Request) {
	bookID := mux.Vars(req)["bookID"]
	err := b.libmanager.DeleteBook(req.Context(), bookID)
	if err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
	}
}

func (b *bookHandler) FindBookByID(w http.ResponseWriter, req *http.Request) {
	bookID := mux.Vars(req)["bookID"]
	bookInfo, err := b.bookStore.FindBookByID(req.Context(), bookID)
	if err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
	}
	utils.WriteResponse(http.StatusOK, bookInfo, w)
}

func (b *bookHandler) FindBookByName(w http.ResponseWriter, req *http.Request) {
	bookName := mux.Vars(req)["bookName"]
	bookInfo, err := b.bookStore.FindBookByName(req.Context(), bookName)
	if err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
	}
	utils.WriteResponse(http.StatusOK, bookInfo, w)
}

func (b *bookHandler) GetAllBooks(w http.ResponseWriter, req *http.Request) {
	books, err := b.bookStore.GetAllBooks(req.Context())
	if err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
	}
	utils.WriteResponse(http.StatusOK, books, w)
}
