package services

import (
	"context"
	"encoding/json"

	kubemq "github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type NodeJoinedEvent struct {
	NodeId string `json:"nodeId"`
	Addr   string `json:"addr"`
}

func NewNodeJoinedEvent(nodeId string, addr string) *NodeJoinedEvent {
	return &NodeJoinedEvent{NodeId: nodeId, Addr: addr}
}

const clientId = "bootstrap-node"
const channelName = "node-instances-channel"

func submitEvent(queueUrl string, event []byte, channelName string) string {
	ctx, _ := context.WithCancel(context.Background())
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, 50000),
		kubemq.WithClientId(clientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))

	if err != nil {
		log.Fatal("something is wrong", err)
	}
	defer client.Close()

	sendResult, err := client.NewQueueMessage().
		SetChannel(channelName).
		SetBody(event).
		Send(ctx)

	if err != nil {
		log.Fatal(err)
	}
	log.Debugln("clientId", clientId)
	log.Debugln("uploadMsgEventChannelName", channelName)
	return sendResult.MessageID
}

func EmitNodeJoinedEvent(queueUrl string, event *NodeJoinedEvent) string {
	log.Debug("SubmitNodeJoinedEvent: %v", event)
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	msgId := submitEvent(queueUrl, bEvent, channelName)
	log.Debug("msgId: %s", msgId)
	return msgId
}
