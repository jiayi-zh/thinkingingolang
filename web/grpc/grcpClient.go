package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

const (
	address = "192.168.9.27:21105"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewDeviceServerClient(conn)

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "123123")

	//info, err := c.PutChannelInfo(incomingContext, &ChannelInfoUpdateReq{
	//	ChannelUuid:    "ba19e6ae700c4eefa4cd3e1a86f9e003",
	//	ChannelName:    "读头2",
	//	ChannelType:    "readHead",
	//	ChannelNo:      "1",
	//	DeviceId:       "133220641294A193280",
	//	Direction:      "in",
	//	ExtInfo:        "{\"chn\":{\"k1\":\"v1\",\"k2\":\"v2\"}}",
	//	IotExtInfo:     "{\"chn\":{\"k1\":\"v1\",\"k2\":\"v2\"}}",
	//	Version:        0,
	//	GbCode:         "11111111",
	//	CaptureAbility: "face",
	//	Latitude:       "111.123",
	//	Longitude:      "111.321",
	//	Remarks:        "remarks",
	//})

	list, err := c.GetAreaInfoList(ctx, &AreaInfoQueryReq{ProjectId: "1111"})
	fmt.Printf("%v %v", list, err)
}
