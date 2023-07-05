package notehandler

import (
	"api-go-fiber/database"
	"api-go-fiber/internals/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetNotes func gets all existing notes
// @Description Get all existing notes
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @router /api/note [get]
func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}



// CreateNote func create a note
// @Description Create a Note
// @Tags Notes
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param sub_title body string true "SubTitle"
// @Param text body string true "Text"
// @Success 200 {object} model.Note
// @router /api/note [post]
func CreateNote(c *fiber.Ctx) error {
	db := database.DB

	note := new(models.Note)
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	note.ID = uuid.New()

	err = db.Create(&note).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}


// GetNote func one note by ID
// @Description Get one note by ID
// @Tags Note
// @Accept json
// @Produce json
// @Success 200 {object} model.Note
// @router /api/note/{id} [get]
func GetNote(c *fiber.Ctx) error {
	db := database.DB

	var note models.Note

	id := c.Params("noteId")
	db.Find(&note, "id=?", id)

	if note.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}



// UpdateNote update a note by ID
// @Description Update a Note by ID
// @Tags Notes
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param sub_title body string true "SubTitle"
// @Param text body string true "Text"
// @Success 200 {object} model.Note
// @router /api/note/{id} [post]
func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}

	db := database.DB
	var note models.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updateNoteData updateNote

	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.Subtitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes Found and Updated", "data": note})
}


// DeleteNote delete a Note by ID
// @Description Delete a note by ID
// @Tags Note
// @Accept json
// @Produce json
// @Success 200
// @router /api/note/{id} [delete]
func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	err := db.Delete(&note, "id=?", id).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}