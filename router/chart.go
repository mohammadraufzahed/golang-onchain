package router

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ario-team/glassnode-api/database"
	"github.com/ario-team/glassnode-api/functions"
	"github.com/ario-team/glassnode-api/schema"
	"github.com/ario-team/glassnode-api/types"
	"github.com/ario-team/glassnode-api/web"
	"github.com/gofiber/fiber/v2"
)

var chartRouter fiber.Router = web.Server.Group("/api/chart")

func InitializeChartRouter() {
	chartRouter.Get("/", getChart)
}

// Get Charts
// @Summary Get the chart data
// @Tags    Charts
// @ID      get_chart
// @Accept  json
// @Produce json
// @Param   id  query    int                      true "Endpoint id"
// @Param   a   query    string                   true "Asset name"
// @Param   r   query    string                   true "Resolution"
// @Param   s   query    int                      true "Start"
// @Param   e   query    int                      true "End"
// @Success 200 {object} []types.ChartGetResponse "Data"
// @Router  /api/chart/ [get]
func getChart(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("id", "0"))
	asset := c.Query("a", "")
	resolution := c.Query("r", "")
	end, _ := strconv.Atoi(c.Query("e", "0"))
	start, _ := strconv.Atoi(c.Query("s", "0"))
	if id == 0 || asset == "" || resolution == "" || end == 0 || start == 0 {
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
			Message: "Endpoint not found",
		})
	}
	if !endpoint.Initialized {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "This endpoint is not initialized",
		})
	}
	if !strings.Contains(string(endpoint.Assets), fmt.Sprintf(`"symbol": "%v"`, strings.ToUpper(asset))) {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "This endpoint does not support this asset",
		})
	}
	if !strings.Contains(string(endpoint.Resolutions), fmt.Sprintf(`%v`, strings.ToUpper(resolution))) {
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "This endpoint does not support this resolution",
		})
	}
	chart, err := functions.GetChart(endpoint.Path, asset, resolution, start, end)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(types.CreateTopGroupRes{
			Status:  400,
			Message: "Server faild or bad filters",
		})
	}
	return c.SendString(chart)
}
