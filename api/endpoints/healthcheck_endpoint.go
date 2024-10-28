package endpoints

import (
	"context"
	infraestructure_services "my_wallet/api/services/healtcheck"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

type HealtcheckDbRequest struct {
}

type HealtcheckDbResponse struct {
	Database string `json:"database,omitempty"`
	Service  string `json:"service,omitempty"`
	Err      string `json:"error,omitempty"`
}

func MakeGetHealthCheckEndpoint(s infraestructure_services.HealtcheckService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.GetHealtcheck(ctx)
		if err != nil || !ok {
			logger.Errorln("Layer: Healthcheck endoint ", "Method:MakeGetHealthCheckEndpoint", "Error:", err)
			return HealtcheckDbResponse{}, err
		}
		return HealtcheckDbResponse{Database: "ok", Service: "ok"}, nil
	}
}
