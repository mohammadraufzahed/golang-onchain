package job

import "github.com/ario-team/glassnode-api/functions"

func RegisterEndpointJob() {
	Scheduler.Every("12h").Do(functions.GetEndPoints)
}
