package redisopr

/*

订阅demo
ropr := redisopr.NewRedis()
defer ropr.Close()

done := make(chan error, 1)
err = listenPubSubChannels(ctx, ropr,
func() error {
	gConfig.setStatus(UP)
	logger.Infof("所有通道订阅成功，内存加载完成;")
	done <- nil
	return nil
},
func(channel string, data []byte) error {
	var err error
	logger.Infof("收到订阅通知[%s]:[%s]", channel, string(data))

	logger.Infof("通知[%s][%s]已加载成功;", channel, string(data))
	return nil
},
redisChannelChannel, //渠道配置
redisRouteChannel,   //路由配置
redisParamChannel)   //参数配置
if err != nil {
	logger.Errorf("listenPubSubChannels 注册失败:[%s]", err)
	done <- err
}

return <-done
*/
