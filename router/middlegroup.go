package router

import (
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
)

var middleGroupRouter fiber.Router = web.Server.Group("/api/middlegroup")

func InitializeMiddleGroupRouter() {
	middleGroupRouter.Post("/", createMiddleGroup)
	middleGroupRouter.Post("/endpoint/:id", appendEndpoint)
	middleGroupRouter.Patch("/:id", updateMiddleGroup)
	middleGroupRouter.Delete("/:id", deleteMiddleGroup)
}

// Create MiddleGroup
// @Summary Create the middle group
// @Tags    Groups Management
// @ID      middlegroup_create
// @Accpet  json
// @Produce json
// @Param   data body     types.MiddleGroupCreate true "Data"
// @Success 200  {object} types.CreateTopGroupRes "Created"
// @Failure 400  {object} types.CreateTopGroupRes "Bad request"
// @Failure 500  {object} types.CreateTopGroupRes "Server faild"
// @Router  /api/middlegroup [post]
func createMiddleGroup(c *fiber.Ctx) error {
	var body types.MiddleGroupCreate
	c.BodyParser(&body)
	if body.Name == "" || body.TopGroupID == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Body is not correct",
		})
	}
	var dbMiddleGroup schema.MiddleGroup
	dbResult := database.Connection.Where("name = ?", body.Name).Find(&dbMiddleGroup)
	if dbResult.RowsAffected != 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Group Exists",
		})
	}
	var topgroup schema.TopGroup
	dbResult = database.Connection.Where("id = ?", body.TopGroupID).First(&topgroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "TopGroup not found",
		})
	}
	err := database.Connection.Model(&topgroup).Association("MiddleGroups").Append(&schema.MiddleGroup{
		Name: body.Name,
	})
	if err != nil {
		return c.Status(500).JSON(types.CreateTopGroupRes{
			Status:  500,
			Message: "Server faild",
		})
	}
	return c.Status(200).JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "Created",
	})

}

// Append Endpoint
// @Summary Append a endpoint to middle group
// @Tags Groups Management
// @ID middlegroup_append_endpoint
// @Accept json
// @Produce json
// @Param id path int true "MiddleGroup id"
// @Param endpoint_id body types.MiddleGroupAppendEndpoint true "Endpoint id"
// @Success 200  {object} types.CreateTopGroupRes "Created"
// @Failure 400  {object} types.CreateTopGroupRes "Bad request"
// @Failure 500  {object} types.CreateTopGroupRes "Server faild"
// @Router /api/middlegroup/endpoint/{id} [post]
func appendEndpoint(c *fiber.Ctx) error {
	middlegroup_id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	var body types.MiddleGroupAppendEndpoint
	err = c.BodyParser(&body)
	if err != nil || body.EndpointID == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Wrong body",
		})
	}
	var middlegroup schema.MiddleGroup
	dbResult := database.Connection.Where("id = ?", middlegroup_id).First(&middlegroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "MiddleGroup not found",
		})
	}
	var endpoint schema.Endpoint
	dbResult = database.Connection.Where("id = ?", body.EndpointID).First(&endpoint)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Endpoint not found",
		})
	}
	if endpoint.MiddleGroupID != 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Endpoint have a middlegroup",
		})
	}
	err = database.Connection.Model(&middlegroup).Association("Endpoints").Append(&endpoint)
	if err != nil {
		return c.Status(500).JSON(types.CreateTopGroupRes{
			Status:  500,
			Message: "Server faild",
		})
	}
	return c.Status(200).JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "Endpoint added",
	})
}

// Update MiddleGroup
// @Summary Update the middle group
// @Tags    Groups Management
// @ID      middlegroup_update
// @Accept  json
// @Produce json
// @Param   id   path     int                     true "Group id"
// @Param   data body     types.MiddleGroupUpdate true "Updating data"
// @Success 200  {object} types.CreateTopGroupRes "Updated"
// @Failure 400  {object} types.CreateTopGroupRes "Bad request"
// @Failure 500  {object} types.CreateTopGroupRes "Server faild"
// @Router  /api/middlegroup/{id} [patch]
func updateMiddleGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	var body types.MiddleGroupUpdate
	c.BodyParser(&body)
	if body.Name == "" {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Body isn't correct",
		})
	}
	var middleGroup schema.MiddleGroup
	dbResult := database.Connection.Where("id = ?", id).First(&middleGroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Group not found",
		})
	}
	middleGroup.Name = body.Name
	dbResult = database.Connection.Save(&middleGroup)
	if dbResult.RowsAffected == 0 {
		return c.Status(500).JSON(types.CreateTopGroupRes{Status: 500, Message: "Server faild"})
	}
	return c.Status(200).JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "Updated",
	})
}

// Delete MiddleGroup
// @Summary Delete the middle group
// @Tags    Groups Management
// @ID      middlegroup_delete
// @Accept  json
// @Produce json
// @Param   id  path     int                     true "Group ID"
// @Success 200 {object} types.CreateTopGroupRes "Deleted"
// @Failure 400 {object} types.CreateTopGroupRes "Bad request"
// @Router  /api/middlegroup/{id} [delete]
func deleteMiddleGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	dbResult := database.Connection.Delete(&schema.MiddleGroup{}, id)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Group not found",
		})
	}
	return c.Status(200).JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "Deleted",
	})
}
