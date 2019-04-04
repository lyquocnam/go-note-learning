package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lyquocnam/go-note-learning/lib"
	"github.com/lyquocnam/go-note-learning/model"
	"github.com/lyquocnam/go-note-learning/repo"
	"net/http"
	"strconv"
)

type noteHandler struct {
	router   *gin.Engine
	noteRepo repo.NoteRepo
}

func NewNoteHandler(router *gin.Engine, noteRepo repo.NoteRepo) *noteHandler {
	handler := &noteHandler{
		router:   router,
		noteRepo: noteRepo,
	}

	notesGroup := handler.router.Group("/notes")
	notesGroup.GET("/", handler.GetList)
	notesGroup.GET("/:id", handler.Get)

	notesGroup.POST("/", handler.Add)
	notesGroup.PUT("/:id", handler.Update)
	notesGroup.DELETE("/:id", handler.Delete)

	return handler
}

type NoteHandler interface {
	Get(c *gin.Context)
	GetList(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func (h *noteHandler) Response(c *gin.Context, data interface{}, code int, err error) {
	var message string
	if err != nil {
		message = err.Error()
	}
	c.JSON(code, lib.NewResponse(code, message, data))
}

func (h *noteHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	note, err := h.noteRepo.Get(uint(id))
	if err != nil || note == nil {
		h.Response(c, nil, 404, err)
		return
	}

	h.Response(c, note, 200, nil)
}

func (h *noteHandler) GetList(c *gin.Context) {
	notes, err := h.noteRepo.GetList()

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *noteHandler) Add(c *gin.Context) {
	var note model.NoteRequest
	err := c.BindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, code, err := h.noteRepo.Insert(&note)
	h.Response(c, result, code, err)
}

func (h *noteHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var note model.NoteRequest
	err = c.BindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, code, err := h.noteRepo.Update(uint(id), &note)
	h.Response(c, result, code, err)
}

func (h *noteHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	note, code, err := h.noteRepo.Delete(uint(id))
	h.Response(c, note, code, err)
}
