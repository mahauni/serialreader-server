package internal

import (
	"context"
	"log"

	pb "github.com/mahauni/serialreader-server/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SerialReaderServerImpl struct {
	arduinoReader *ArduinoReader
	pb.SerialReaderServiceServer
}

func (s *SerialReaderServerImpl) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SayHelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *SerialReaderServerImpl) GetSparkFunWeatherShieldData(ctx context.Context, in *pb.GetTimeSeriesData) (*pb.SparkFunWeatherShieldTimeSeriesData, error) {
	datum := s.arduinoReader.GetSparkFunWeatherShieldData()
	return &pb.SparkFunWeatherShieldTimeSeriesData{
		Status:                 true,
		Timestamp:              timestamppb.Now(),
		HumidityValue:          datum.HumidityValue,
		HumidityUnit:           datum.HumidityUnit,
		TemperatureValue:       datum.TemperatureValue,
		TemperatureUnit:        datum.TemperatureUnit,
		PressureValue:          datum.PressureValue,
		PressureUnit:           datum.PressureUnit,
		TemperatureBackupValue: datum.TemperatureBackupValue,
		TemperatureBackupUnit:  datum.TemperatureBackupUnit,
		AltitudeValue:          datum.AltitudeValue,
		AltitudeUnit:           datum.AltitudeUnit,
		IlluminanceValue:       datum.IlluminanceValue,
		IlluminanceUnit:        datum.IlluminanceUnit,
	}, nil
}
