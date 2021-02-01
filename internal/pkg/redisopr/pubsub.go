package redisopr

import (
	"context"
	"time"

	"uncresys/modules/log"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func SubChannels(ctx context.Context,
	onStart func() error,
	onMessage func(channel string, data []byte) error,
	channels ...string) error {
	// A ping is set to the server with this period to test for the health of
	// the connection and server.
	const healthCheckPeriod = 20 * time.Second

	logger, err := log.ContextEntry(ctx)
	if err != nil {
		return err
	}

	psc := redis.PubSubConn{Conn: NewRedis().Conn}

	if err := psc.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
		logger.Errorf("订阅[%v]失败;[%s]", channels, err)
		return err
	}

	done := make(chan error, 1)

	// Start a goroutine to receive notifications from the server.
	go func() {
		defer psc.Close()
		for {
			switch n := psc.Receive().(type) {
			case error:
				logger.Errorf("订阅收到错误[%s]", n)
				done <- n
				return
			case redis.Message:
				logger.Infof("订阅收到消息[%s]:[%s]", n.Channel, string(n.Data))
				if err := onMessage(n.Channel, n.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				logger.Infof("订阅返回订阅数[%d]", n.Count)
				switch n.Count {
				case len(channels):
					// Notify application when all channels are subscribed.
					if err := onStart(); err != nil {
						done <- err
						return
					}
				case 0:
					// Return from the goroutine when all channels are unsubscribed.
					done <- nil
					return
				}
			case redis.Pong:
				continue
			}

		}
	}()

	ticker := time.NewTicker(healthCheckPeriod)
	defer ticker.Stop()
loop:
	for {
		select {
		case <-ticker.C:
			// Send ping to test health of connection and server. If
			// corresponding pong is not received, then receive on the
			// connection will timeout and the receive goroutine will exit.
			if err = psc.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			break loop
		case err = <-done:
			// Return error from the receive goroutine.
			return err
		}
	}

	// Signal the receiving goroutine to exit by unsubscribing from all channels.
	psc.Unsubscribe()

	// Wait for goroutine to complete.
	return <-done
}

func Publish(ctx context.Context, channel string, data []byte) error {
	logger, _ := log.ContextEntry(ctx)
	cli := NewRedis()
	defer cli.Close()

	v, err := cli.Do("PUBLISH", channel, data)
	if err != nil {
		logger.Errorf("通知[%s]失败;[%s]", channel, err)
		return errors.Wrapf(err, "PUBLISH [%s] error[%s]", channel, err)
	}
	logger.Infof("推送通知[%s]成功;订阅数[%v]", channel, v)
	return nil
}
