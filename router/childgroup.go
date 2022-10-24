package router

import (
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
)

var childGroupRoute fiber.Router = web.Server.Group("/api/childgroup")

func InitializeChildGroupRoute() {
	childGroupRoute.Post("/", createChildGroup)
	childGroupRoute.Patch("/:id", updateChildGroup)
	childGroupRoute.Delete("/:id", deleteChildGroup)
}

// Create child group
// @Summary Create the child group
// @Tags Groups Management
// @ID childgroup_create
// @Accept json
// @Produce json
// @Param data body types.ChildGroupCreate true "Data"
// @Success 200 {object} types.ChildGroupUpdate "Successfull"
// @Failure 400 {object} types.CreateTopGroupRes "Bad request"
// @Failure 500 {object} types.CreateTopGroupRes "Server Faild"
// @Router /api/childgroup [post]
func createChildGroup(c *fiber.Ctx) error {
	var body types.ChildGroupCreate
	c.BodyParser(&body)
	if body.Description == "" || body.Name == "" || body.MiddleGroupID == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Body is empty",
		})
	}
	var childGroup schema.ChildGroup
	dbResult := database.Connection.Where("name = ?", body.Name).First(&childGroup)
	var middlegroup schema.MiddleGroup
	if dbResult.RowsAffected == 1 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ChildGroup exists",
		})
	}
	dbResult = database.Connection.Where("id = ?", body.MiddleGroupID).First(&middlegroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "MiddleGroup not found",
		})
	}
	err := database.Connection.Model(&middlegroup).Association("ChildGroups").Append(&schema.ChildGroup{
		Name:        body.Name,
		Description: body.Description,
	})
	if err != nil {
		return c.Status(500).JSON(types.CreateTopGroupRes{
			Status:  500,
			Message: "Server faild",
		})
	}
	return c.Status(200).JSON(&types.CreateTopGroupRes{
		Status:  200,
		Message: "Created",
	})
}

// Update child group
// @Summary Update the child group
// @Tags Groups Management
// @ID childgroup_update
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param data body types.ChildGroupUpdate true "Data"
// @Success 200 {object} types.CreateTopGroupRes "Successfull"
// @Failure 400 {object} types.CreateTopGroupRes "Bad request"
// @Failure 500 {object} types.CreateTopGroupRes "Server Faild"
// @Router /api/childgroup/{id} [patch]
func updateChildGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	var body types.ChildGroupUpdate
	c.BodyParser(&body)
	var childgroup schema.ChildGroup
	dbResult := database.Connection.Where("id = ?", id).First(&childgroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(&types.CreateTopGroupRes{
			Status:  400,
			Message: "Group not found",
		})
	}
	if body.Name != "" {
		childgroup.Name = body.Name
	}
	if body.Description != "" {
		childgroup.Description = body.Description
	}
	dbResult = database.Connection.Save(&childgroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(500).JSON(&types.CreateTopGroupRes{
			Status:  500,
			Message: "Server fauild",
		})
	}
	return c.Status(200).JSON(&types.CreateTopGroupRes{
		Status:  200,
		Message: "Updated",
	})
}

// Delete child group
// @Summary Delete the child group
// @Tags Groups Management
// @ID childgroup_delete
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} types.CreateTopGroupRes "Successfull"
// @Failure 400 {object} types.CreateTopGroupRes "Bad request"
// @Failure 500 {object} types.CreateTopGroupRes "Server Faild"
// @Router /api/childgroup/{id} [delete]
func deleteChildGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	var childgroup schema.ChildGroup
	dbResult := database.Connection.Where("id = ?", id).First(&childgroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(&types.CreateTopGroupRes{
			Status:  400,
			Message: "ChildGroup not found",
		})
	}
	dbResult = database.Connection.Delete(&schema.ChildGroup{}, id)
	if dbResult.RowsAffected == 0 {
		return c.Status(500).JSON(&types.CreateTopGroupRes{
			Status:  500,
			Message: "Server faild",
		})
	}
	return c.Status(200).JSON(&types.CreateTopGroupRes{
		Status:  200,
		Message: "Deleted",
	})
}
