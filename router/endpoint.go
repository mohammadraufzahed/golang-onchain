package router

import (
	"encoding/json"
	"time"

	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/redis"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
)

var endpointRoute fiber.Router = web.Server.Group("/api/endpoint")

func InitializeEndpointRouter() {
	endpointRoute.Get("/all", getEndpoints)
	endpointRoute.Get("/:id", getEndpoint)
	endpointRoute.Patch("/:id", updateEndpoint)
}

// Get supported endpoints
// @Summary Get the supported endpoints
// @Tags    Endpoints
// @ID      endpoints_get
// @Produce json
// @Success 200 {object} []types.EndpointGetAll "Successfull"
// @Router  /api/endpoint/all [get]
func getEndpoints(c *fiber.Ctx) error {

	var endpoints []schema.Endpoint
	database.Connection.Find(&endpoints)
	if len(endpoints) == 0 {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.SendString("[]")
	}
	endpointsResult := []types.EndpointGetAll{}
	for _, endpoint := range endpoints {
		endpointsResult = append(endpointsResult, types.EndpointGetAll{
			ID:          endpoint.ID,
			Path:        endpoint.Path,
			Tier:        endpoint.Tier,
			Assets:      endpoint.Assets,
			Name:        endpoint.Name,
			Description: endpoint.Description,
			Initialized: endpoint.Initialized,
			Currencies:  endpoint.Currencies,
			Resolutions: endpoint.Resolutions,
			Formats:     endpoint.Formats,
		})
	}
	endpointsJson, err := json.Marshal(endpointsResult)
	if err != nil {
		return err
	}
	redis.Set("endpoints", string(endpointsJson), time.Hour*5)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.SendString(string(endpointsJson))
}

// Get supported endpoint
// @Summary Get the supported endpoint
// @Tags    Endpoints
// @ID      endpoints_get_one
// @Produce json
// @Param   id  path     int                     true "Endpoint id"
// @Success 200 {object} types.EndpointGetAll    "Successfull"
// @Failure 400 {object} types.CreateTopGroupRes "Bad request"
// @Router  /api/endpoint/{id} [get]
func getEndpoint(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	var endpoint schema.Endpoint
	dbResult := database.Connection.Where("id = ?", id).First(&endpoint)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Endpoint not found",
		})
	}
	return c.Status(200).JSON(types.EndpointGetAll{
		ID:          endpoint.ID,
		Path:        endpoint.Path,
		Tier:        endpoint.Tier,
		Name:        endpoint.Name,
		Description: endpoint.Description,
		Initialized: endpoint.Initialized,
		Assets:      endpoint.Assets,
		Currencies:  endpoint.Currencies,
		Resolutions: endpoint.Resolutions,
		Formats:     endpoint.Formats,
	})
}

// Update the endpoint
// @Summary Update the endpoint information
// @Tags    Endpoints
// @ID      endpoints_update
// @Accept  json
// @Produce json
// @Param   id   path     int                     true "Endpoint ID"
// @Param   data body     types.EndpointUpdate    true "Data"
// @Success 200  {object} types.CreateTopGroupRes "Updated Successfully"
// @Success 400  {object} types.CreateTopGroupRes "Bad request"
// @Success 500  {object} types.CreateTopGroupRes "Server faild"
// @Router  /api/endpoint/{id} [patch]
func updateEndpoint(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "ID not found",
		})
	}
	var body types.EndpointUpdate
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Bad body",
		})
	}
	var endpoint schema.Endpoint
	dbResult := database.Connection.Where("id = ?", id).First(&endpoint)
	if dbResult.RowsAffected == 0 {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Endpoint with this id not found",
		})
	}
	if body.Name != "" {
		database.Connection.Model(&endpoint).Update("name", body.Name)
	}
	if body.Description != "" {
		database.Connection.Model(&endpoint).Update("description", body.Description)
	}
	return c.Status(200).JSON(types.CreateTopGroupRes{
		Status:  200,
		Message: "Endpoint Updated",
	})
}
