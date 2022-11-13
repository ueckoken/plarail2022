package internal

import (
	"ueckoken/plarail2022-external/pkg/envStore"
	"ueckoken/plarail2022-external/pkg/httphandler"
	"ueckoken/plarail2022-external/pkg/synccontroller"
	"ueckoken/plarail2022-external/spec"
)

// Run runs external server.
func Run() {
	synccontrollerInput := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	synccontrollerOutput := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	grpcHandlerInput := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	main2grpcHandler := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	httpInputKV := make(chan synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State])
	httpInput := make(chan *spec.Command2InternalRequest)
	httpOutput := make(chan *spec.Command2InternalRequest)

	go func() {
		for c := range synccontrollerOutput {
			main2grpcHandler <- c
			httpInputKV <- c
		}
	}()
	go func() {
		for c := range httpInputKV {
			httpInput <- &spec.Command2InternalRequest{Station: &spec.Stations{StationId: c.Key}, State: c.Value}
		}
	}()

	go func() {
		for c := range httpOutput {
			synccontrollerInput <- synccontroller.KV[spec.Stations_StationId, spec.Command2InternalRequest_State]{Key: c.GetStation().GetStationId(), Value: c.GetState()}
		}
	}()

	envVal := envStore.GetEnv()

	httpServer := httphandler.NewHTTPServer[*spec.Command2InternalRequest](
		httpOutput,
		httpInput,
		envVal,
	)
	StartStationSync(synccontrollerInput, synccontrollerOutput)
	grpcHandler := NewGrpcHandler(envVal, synccontrollerInput, grpcHandlerInput)

	go GRPCListenAndServe(uint(envVal.ClientSideServer.GrpcPort), grpcHandler)
	go httpServer.StartServer()
}
