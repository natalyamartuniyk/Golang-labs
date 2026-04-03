package main

import (
	"log"
	"strconv"

	"lab_07/model"

	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
)

var notes []model.Note
var nextId = 1

func main() {
	app := fiber.New()

	app.Get("/notes", getNotes)
	app.Get("/notes/:id", getNoteById)
	app.Post("/notes", createNote)
	app.Put("/notes/:id", updateNote)
	app.Delete("/notes/:id", deleteNote)

	log.Fatal(app.Listen(":8080"))
}

func getNotes(c fiber.Ctx) error {
	return c.JSON(notes)
}
func getNoteById(c fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id is invalid"})
	}

	for _, note := range notes {
		if note.Id == id {
			return c.Status(200).JSON(note)
		}
	}
	return c.Status(400).JSON(fiber.Map{"error": "note not found"})
}
func createNote(c fiber.Ctx) error {
	var note model.Note
	body := c.Body()

	if err := easyjson.Unmarshal(body, &note); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "body is invalid"})
	}

	if !note.IsPhoneValid() {
		return c.Status(400).JSON(fiber.Map{"error": "phone number is invalid"})
	}

	note.Id = nextId
	nextId++
	notes = append(notes, note)
	return c.Status(201).JSON(note)
}
func updateNote(c fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id is invalid"})
	}
	var updated model.Note

	body := c.Body()
	if err := easyjson.Unmarshal(body, &updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "body is invalid"})
	}

	if !updated.IsPhoneValid() {
		return c.Status(400).JSON(fiber.Map{"error": "phone number is invalid"})
	}

	for i, note := range notes {
		if note.Id == id {
			notes[i].Title = updated.Title
			notes[i].Content = updated.Content
			notes[i].Phone = updated.Phone

			return c.JSON(notes[i])
		}
	}
	return c.Status(400).JSON(fiber.Map{"error": "note not found"})

}
func deleteNote(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id is invalid"})
	}
	for i, note := range notes {
		if note.Id == id {
			notes = append(notes[:i], notes[i+1:]...)
			return c.Status(200).JSON(note)
		}
	}
	return c.Status(400).JSON(fiber.Map{"error": "note not found"})
}
