package core

import (
	"context"

	"gitlab.com/ruijc/app/core"
	"gitlab.com/ruijc/app/core/kernel"
	"gitlab.com/ruijc/app/core/notification"
	"gitlab.com/ruijc/app/protocol"
	"google.golang.org/protobuf/proto"
)

// Notification 通知
type Notification struct {
	client protocol.NotificationClient
}

func NewNotification(connection *Connection) *Notification {
	return &Notification{
		client: protocol.NewNotificationClient(connection),
	}
}

func (n *Notification) Notify(ctx context.Context, id core.Id, data proto.Message) (err error) {
	req := new(notification.NotifyReq)
	req.Id = int64(id)
	req.Type = kernel.NotificationType_NOTIFICATION_TYPE_MALEONN
	req.ContentType = "application/protobuf"
	if req.Data, err = proto.Marshal(data); nil == err {
		_, err = n.client.Notify(ctx, req)
	}

	return
}
