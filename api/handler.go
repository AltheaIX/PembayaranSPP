package main

import (
	"database/sql"
	"github.com/AltheaIX/PembayaranSPP/user"
	"github.com/gofiber/fiber/v2"
)

func (env *Env) HandlerGetPetugas(c *fiber.Ctx) error {
	listPetugas, err := user.AllPetugas(env.db)
	result := &user.ResponseJson{}

	if err != nil {
		result = &user.ResponseJson{
			Success: false,
			Message: err.Error(),
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	}

	result = &user.ResponseJson{
		Success: true,
		Data:    listPetugas,
		Message: "Getting petugas list success.",
	}
	return c.Status(fiber.StatusServiceUnavailable).JSON(result)
}

func (env *Env) HandlerGetPetugasById(c *fiber.Ctx) error {
	listPetugas := []user.Petugas{}
	result := &user.ResponseJson{}

	id, err := c.ParamsInt("id")
	if err != nil {
		result = &user.ResponseJson{
			Success: false,
			Message: "Error on parsing parameter, please make sure you are inputing ID.",
		}
		return c.Status(fiber.StatusBadRequest).JSON(result)
	}

	petugas, err := user.PetugasById(env.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			result = &user.ResponseJson{
				Success: false,
				Message: "ID Petugas doesnt exist.",
			}
			return c.Status(fiber.StatusServiceUnavailable).JSON(result)
		}
		result = &user.ResponseJson{
			Success: false,
			Message: "Error on getting data, " + err.Error(),
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	}

	listPetugas = append(listPetugas, petugas)
	result = &user.ResponseJson{
		Success: true,
		Data:    listPetugas,
		Message: "Success on getting data.",
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (env *Env) HandlerCreatePetugas(c *fiber.Ctx) error {
	petugas := user.Petugas{}
	listPetugas := []user.Petugas{}
	response := &user.ResponseJson{}

	if err := c.BodyParser(&petugas); err != nil {
		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if petugas.Username == "" || petugas.Password == "" || petugas.NamaPetugas == "" || petugas.Level == "" {
		response = &user.ResponseJson{Success: false, Message: "Invalid payload json."}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newPetugas, err := user.CreatePetugas(env.db, petugas)
	if err != nil {
		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	listPetugas = append(listPetugas, newPetugas)
	response = &user.ResponseJson{Success: true, Data: listPetugas, Message: "Success on creating data."}
	return c.Status(fiber.StatusBadRequest).JSON(response)
}
