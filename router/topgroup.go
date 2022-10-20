package router

import (
	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
)

var TopGroupRouter fiber.Router = web.Server.Group("/api/topgroup")

func InitializeTopGroupRouter() {
	TopGroupRouter.Post("/", createTopGroup)
	TopGroupRouter.Get("/", getTopGroups)
	TopGroupRouter.Patch("/:id", updateTopGroups)
	TopGroupRouter.Delete("/:id", deleteTopGroups)
}

// Create Top Group
// @Summary Add a new Top group
// @Tags    Groups Management
// @id      topgroup_create
// @Accept  json
// @Produce json
// @Param   name body     types.CreateTopGroupReq true "TopGroup name"
// @Success 200  {object} types.CreateTopGroupRes "TopGroup created"
// @Failure 400  {object} types.CreateTopGroupRes "Bad request"
// @Failure 500  {object} types.CreateTopGroupRes "Creating faild"
// @Router  /api/topgroup [post]
func createTopGroup(c *fiber.Ctx) error {
	var data schema.TopGroup
	c.BodyParser(&data)
	if data.Name == "" {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Body is not correct",
		})
	}
	var foundedGroup schema.TopGroup
	res := database.Connection.Where("name = ?", data.Name).Find(&foundedGroup)
	if res.RowsAffected != 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{Status: 400, Message: "Group exists"})
	}
	res = database.Connection.Create(&data)
	if res.RowsAffected == 0 {
		return c.Status(500).JSON(types.CreateTopGroupRes{Status: 500, Message: "Faild to create"})
	}
	return c.Status(200).JSON(types.CreateTopGroupRes{Status: 200, Message: "Created"})
}

// Update TopGroup
// @Summary Update TopGroups
// @Tags    Groups Management
// @ID      topgroups_update
// @Produce json
// @Param   id   path     int                     true "TopGroup ID"
// @Param   data body     types.CreateTopGroupReq true "Request body"
// @Success 200  {object} types.CreateTopGroupRes "TopGroup updated"
// @Failure 400  {object} types.CreateTopGroupRes "Bad request"
// @Failure 500  {object} types.CreateTopGroupRes "Updating faild"
// @Router  /api/topgroup/{id} [patch]
func updateTopGroups(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Please pass id",
		})
	}
	var topgroup schema.TopGroup
	res := database.Connection.Where("id = ?", id).First(&topgroup)
	if res.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "TopGroup Not Found",
		})
	}
	var body types.CreateTopGroupReq
	c.BodyParser(&body)
	if body.Name == "" {
		return c.Status(400).JSON(types.CreateTopGroupRes{Status: 400, Message: "Body is not correct"})
	}
	topgroup.Name = body.Name
	res = database.Connection.Save(&topgroup)
	if res.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "TopGroup Not Found",
		})
	}
	return c.JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "TopGroup Updated",
	})
}

// Get TopGroups
// @Summary Get TopGroups
// @Tags    Groups Management
// @ID      topgroups_get
// @Produce json
// @Success 200 {object} []types.GetTopGroups "TopGroup"
// @Router  /api/topgroup [get]
func getTopGroups(c *fiber.Ctx) error {
	var data []schema.TopGroup
	var res []types.GetTopGroups
	dbRes := database.Connection.Preload("MiddleGroups.ChildGroups.Endpoints").Find(&data)
	if dbRes.RowsAffected != 0 {
		for _, top := range data {
			var middles []types.MiddleGroup
			for _, middle := range top.MiddleGroups {
				var childs []types.ChildGroups
				for _, child := range middle.ChildGroups {
					childTemp := types.ChildGroups{
						ID:   child.ID,
						Name: child.Name,
					}
					for _, endpoint := range child.Endpoints {
						childTemp.EndpointID = append(childTemp.EndpointID, endpoint.ID)
					}
					childs = append(childs, childTemp)
				}
				middles = append(middles, types.MiddleGroup{
					ID:          middle.ID,
					Name:        middle.Name,
					ChildGroups: childs,
				})

			}
			res = append(res, types.GetTopGroups{
				ID:           top.ID,
				Name:         top.Name,
				MiddleGroups: middles,
			})
		}
		return c.JSON(res)
	} else {
		c.Response().Header.Add("Content-Type", "application/json")
		return c.SendString("[]")
	}
}

// Delete TopGroups
// @Summary Delete TopGroup
// @Tags    Groups Management
// @ID      topgroups_delete
// @Produce json
// @Param   id  path     int                     true "TopGroup ID"
// @Success 200 {object} types.CreateTopGroupRes "Result"
// @Failure 500 {object} types.CreateTopGroupRes "Error"
// @Router  /api/topgroup/{id} [delete]
func deleteTopGroups(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Pass the id",
		})
	}
	res := database.Connection.Delete(&schema.TopGroup{}, id)
	if res.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	return c.JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "TopGroup deleted",
	})
}
