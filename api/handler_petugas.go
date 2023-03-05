package main

import (
	"database/sql"
	"github.com/AltheaIX/PembayaranSPP/user"
	"github.com/gofiber/fiber/v2"
)

func (env *Env) HandlerGetPetugas(c *fiber.Ctx) error {
	listPetugas, err := user.GetAllPetugas(env.db)
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

	petugas, err := user.GetPetugasById(env.db, id)
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

	if petugas.Username == "" || petugas.Password == "" || petugas.NamaPetugas == "" || petugas.Level == "" {
		response = &user.ResponseJson{Success: false, Message: "Invalid payload json."}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := c.BodyParser(&petugas); err != nil {
		response = &user.ResponseJson{Success: false, Message: err.Error()}
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

func (env *Env) HandlerUpdatePetugas(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	petugas := user.Petugas{}
	listPetugas := []user.Petugas{}
	response := &user.ResponseJson{}

	if err != nil {
		response = &user.ResponseJson{Success: false, Message: "Error on parsing parameter, please make sure you are inputing ID."}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Validate ID
	_, err = user.GetPetugasById(env.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			response = &user.ResponseJson{Success: false, Message: "Petugas that you are trying to update doesnt exist."}
			return c.Status(fiber.StatusOK).JSON(response)
		}

		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusServiceUnavailable).JSON(response)
	}
	petugas.Id = id

	// Parse to Struct
	if err = c.BodyParser(&petugas); err != nil {
		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Validate Payload by Struct
	if petugas.Username == "" || petugas.Password == "" || petugas.NamaPetugas == "" || petugas.Level == "" {
		response = &user.ResponseJson{Success: false, Message: "Invalid payload json."}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Call update function
	updatedPetugas, err := user.UpdatePetugasById(env.db, petugas)
	if err != nil {
		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusServiceUnavailable).JSON(response)
	}
	listPetugas = append(listPetugas, updatedPetugas)

	response = &user.ResponseJson{Success: true, Data: listPetugas, Message: "Success on updating petugas."}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (env *Env) HandlerDeletePetugas(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	response := &user.ResponseJson{}

	if err != nil {
		response = &user.ResponseJson{Success: false, Message: "Error on parsing parameter, please make sure you are inputing ID."}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	_, err = user.GetPetugasById(env.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			response = &user.ResponseJson{Success: false, Message: "Petugas that you are trying to delete doesnt exist."}
			return c.Status(fiber.StatusOK).JSON(response)
		}

		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusServiceUnavailable).JSON(response)
	}

	err = user.DeletePetugasById(env.db, id)
	if err != nil {
		response = &user.ResponseJson{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusServiceUnavailable).JSON(response)
	}

	response = &user.ResponseJson{Success: true, Message: "Success on deleting petugas."}
	return c.Status(fiber.StatusOK).JSON(response)
}
