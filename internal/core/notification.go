package core

import (
	"context"

	"gitlab.com/ruijc/app/core"
	"gitlab.com/ruijc/app/notification"
	"google.golang.org/protobuf/proto"
)

// Notification 通知
type Notification struct {
	client notification.RpcClient
}

func NewNotification(connection *Connection) *Notification {
	return &Notification{
		client: notification.NewRpcClient(connection),
	}
}

func (n *Notification) Notify(ctx context.Context, id Id, data proto.Message) (err error) {
	req := new(notification.NotifyReq)
	req.Id = int64(id)
	req.Type = core.NotificationType_NOTIFICATION_TYPE_MALEONN
	req.ContentType = "application/protobuf"
	if req.Data, err = proto.Marshal(data); nil == err {
		_, err = n.client.Notify(ctx, req)
	}

	return
}
