package services

import (
	"context"
	"encoding/json"

	kubemq "github.com/kubemq-io/kubemq-go"
	log "github.com/sirupsen/logrus"
)

type NodeRegistryCmd struct {
	NodeId string `json:"nodeId"`
	Port   int    `json:"port"`
}

func NewNodeRegistryCmd(nodeId string, port int) *NodeRegistryCmd {
	return &NodeRegistryCmd{NodeId: nodeId, Port: port}
}

const (
	ClientId                         = "bootstrap-node"
	P2pBootstrapNodeRegistryCmdTopic = "p2p-bootstrap-node-registry-cmd"
)

func emitEvent(queueUrl string, event []byte, topic string) (string, error) {
	log.Infoln("SubmitNodeJoinedEvent, topic:", topic)
	ctx, _ := context.WithCancel(context.Background())
	client, err := kubemq.NewClient(ctx,
		kubemq.WithAddress(queueUrl, 50000),
		kubemq.WithClientId(ClientId),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))

	if err != nil {
		log.Errorln("cant connect to queue", queueUrl, err)
		return "", err
	}
	defer client.Close()

	sendResult, err := client.NewQueueMessage().
		SetChannel(topic).
		SetBody(event).
		Send(ctx)

	if err != nil {
		return "", err
	}
	log.Infoln("clientId", ClientId)
	log.Infoln("uploadMsgEventChannelName", topic)
	return sendResult.MessageID, nil
}

func EmitNodeRegistryCmd(queueUrl string, event *NodeRegistryCmd) (string, error) {
	bEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	msgId, err := emitEvent(queueUrl, bEvent, P2pBootstrapNodeRegistryCmdTopic)
	if err != nil {
		return "", err
	}
	log.Debugln("msgId", msgId)
	return msgId, nil
}
