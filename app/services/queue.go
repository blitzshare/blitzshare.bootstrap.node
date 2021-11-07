package services

import (
	"context"
	"encoding/json"

	kubemq "github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type NodeJoinedEvent struct {
	NodeId string `json:"nodeId"`
}

func NewNodeJoinedEvent(nodeId string) *NodeJoinedEvent {
	return &NodeJoinedEvent{NodeId: nodeId}
}

const clientId = "bootstrap-node"
const channelName = "node-instances-channel"

func submitEvent(queueUrl string, event []byte, channelName string) (string, error) {
	ctx, _ := context.WithCancel(context.Background())
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, 50000),
		kubemq.WithClientId(clientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))

	if err != nil {
		log.Errorln("something is wrong", err)
		return "", err
	}
	defer client.Close()

	sendResult, err := client.NewQueueMessage().
		SetChannel(channelName).
		SetBody(event).
		Send(ctx)

	if err != nil {
		return "", err
	}
	log.Debugln("clientId", clientId)
	log.Debugln("uploadMsgEventChannelName", channelName)
	return sendResult.MessageID, nil
}

func EmitNodeJoinedEvent(queueUrl string, event *NodeJoinedEvent) (string, error) {
	log.Debug("SubmitNodeJoinedEvent: %v", event)
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	msgId, err := submitEvent(queueUrl, bEvent, channelName)
	if err != nil {
		return "", err
	}
	log.Debug("msgId: %s", msgId)
	return msgId, nil
}
