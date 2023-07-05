package noteroutes

import (
	notehandler "api-go-fiber/internals/handlers/noteHandler"
	"github.com/gofiber/fiber/v2"
)

func SetUpNoteRotes(router fiber.Router) {
	note := router.Group("/note")

	note.Post("/", notehandler.CreateNote)

	note.Get("/", notehandler.GetNotes)

	note.Get("/:noteId", notehandler.GetNote)

	note.Put("/:noteId", notehandler.UpdateNote)

	note.Delete("/:noteId", notehandler.DeleteNote)
}



