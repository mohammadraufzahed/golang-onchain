package job

import "github.com/ario-team/glassnode-api/functions"

func RegisterEndpointJob() {
	Scheduler.Every("1h").Do(functions.GetEndPoints)
}
