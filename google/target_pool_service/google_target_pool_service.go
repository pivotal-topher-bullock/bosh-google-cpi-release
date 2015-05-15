package targetpool

import (
	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	"github.com/frodenas/bosh-google-cpi/google/operation_service"
	"google.golang.org/api/compute/v1"
)

const googleTargetPoolServiceLogTag = "GoogleTargetPoolService"

type GoogleTargetPoolService struct {
	project          string
	computeService   *compute.Service
	operationService goperation.OperationService
	logger           boshlog.Logger
}

func NewGoogleTargetPoolService(
	project string,
	computeService *compute.Service,
	operationService goperation.OperationService,
	logger boshlog.Logger,
) GoogleTargetPoolService {
	return GoogleTargetPoolService{
		project:          project,
		computeService:   computeService,
		operationService: operationService,
		logger:           logger,
	}
}
