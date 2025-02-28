package mqtt

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var PrivateMqttConfig Config
var LocalMqttConfig Config

func MqttInit() error {
	// 初始化配置
	err := loadPrivateConfig()
	if err != nil {
		return err
	}
	err = loadLocalConfig()
	if err != nil {
		return err
	}
	return nil
}

// 连接到路由器内部MQTT服务器的配置
func loadLocalConfig() error {
	var configMap map[string]interface{}
	// 将 map 映射到 PrivateMqttConfig 结构体
	// 注意！！！yml文件中带_的key，是无法通过UnmarshalKey解析的
	err := viper.Unmarshal(&configMap)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %s", err)
	}
	// 将 map 转换为 json
	jsonStr, err := json.Marshal(configMap["mqttLocal"])
	if err != nil {
		return fmt.Errorf("unable to marshal config, %s", err)
	}
	// 将 json 转换为结构体
	err = json.Unmarshal(jsonStr, &LocalMqttConfig)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config, %s", err)
	}

	// 打印配置
	logrus.Debug("local mqtt config:", LocalMqttConfig)

	// 单独获取 broker 配置
	broker := viper.GetString("mqttLocal.broker")
	if broker == "" {
		broker = "localhost:1883"
		logrus.Println("Using default broker:", broker)
	}
	LocalMqttConfig.Broker = broker

	// 单独获取 user 配置
	user := viper.GetString("mqttLocal.user")
	if user == "" {
		user = "root"
		logrus.Println("Using default user:", user)
	}
	LocalMqttConfig.User = user

	// 单独获取 pass 配置
	pass := viper.GetString("mqttLocal.pass")
	if pass == "" {
		pass = "root"
		logrus.Println("Using default pass:", pass)
	}
	LocalMqttConfig.Pass = pass

	// 单独获取 channel_buffer_size 配置
	channelBufferSize := viper.GetInt("mqttLocal.channel_buffer_size")
	if channelBufferSize == 0 {
		channelBufferSize = 10000
		logrus.Println("Using default channel_buffer_size:", channelBufferSize)
	}
	LocalMqttConfig.ChannelBufferSize = channelBufferSize

	// 单独获取 write_workers 配置
	writeWorkers := viper.GetInt("mqttLocal.write_workers")
	if writeWorkers == 0 {
		writeWorkers = 10
		logrus.Println("Using default write_workers:", writeWorkers)
	}
	LocalMqttConfig.WriteWorkers = writeWorkers

	// 单独获取 pool_size 配置
	poolSize := viper.GetInt("mqttLocal.telemetry.pool_size")
	if poolSize == 0 {
		poolSize = 100
		logrus.Println("Using default pool_size:", poolSize)
	}
	LocalMqttConfig.Telemetry.PoolSize = poolSize

	// 单独获取 batch_size 配置
	batchSize := viper.GetInt("mqttLocal.telemetry.batch_size")
	if batchSize == 0 {
		batchSize = 100
		logrus.Println("Using default batch_size:", batchSize)
	}
	return nil
}

// 连接到内网监控站的配置
func loadPrivateConfig() error {
	var configMap map[string]interface{}
	// 将 map 映射到 PrivateMqttConfig 结构体
	// 注意！！！yml文件中带_的key，是无法通过UnmarshalKey解析的
	err := viper.Unmarshal(&configMap)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %s", err)
	}
	// 将 map 转换为 json
	jsonStr, err := json.Marshal(configMap["mqtt"])
	if err != nil {
		return fmt.Errorf("unable to marshal config, %s", err)
	}
	// 将 json 转换为结构体
	err = json.Unmarshal(jsonStr, &PrivateMqttConfig)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config, %s", err)
	}

	// 打印配置
	logrus.Debug("mqtt config:", PrivateMqttConfig)

	// 单独获取 broker 配置
	broker := viper.GetString("mqtt.broker")
	if broker == "" {
		broker = "localhost:1883"
		logrus.Println("Using default broker:", broker)
	}
	PrivateMqttConfig.Broker = broker

	// 单独获取 user 配置
	user := viper.GetString("mqtt.user")
	if user == "" {
		user = "root"
		logrus.Println("Using default user:", user)
	}
	PrivateMqttConfig.User = user

	// 单独获取 pass 配置
	pass := viper.GetString("mqtt.pass")
	if pass == "" {
		pass = "root"
		logrus.Println("Using default pass:", pass)
	}
	PrivateMqttConfig.Pass = pass

	// 单独获取 channel_buffer_size 配置
	channelBufferSize := viper.GetInt("mqtt.channel_buffer_size")
	if channelBufferSize == 0 {
		channelBufferSize = 10000
		logrus.Println("Using default channel_buffer_size:", channelBufferSize)
	}
	PrivateMqttConfig.ChannelBufferSize = channelBufferSize

	// 单独获取 write_workers 配置
	writeWorkers := viper.GetInt("mqtt.write_workers")
	if writeWorkers == 0 {
		writeWorkers = 10
		logrus.Println("Using default write_workers:", writeWorkers)
	}
	PrivateMqttConfig.WriteWorkers = writeWorkers

	// 单独获取 pool_size 配置
	poolSize := viper.GetInt("mqtt.telemetry.pool_size")
	if poolSize == 0 {
		poolSize = 100
		logrus.Println("Using default pool_size:", poolSize)
	}
	PrivateMqttConfig.Telemetry.PoolSize = poolSize

	// 单独获取 batch_size 配置
	batchSize := viper.GetInt("mqtt.telemetry.batch_size")
	if batchSize == 0 {
		batchSize = 100
		logrus.Println("Using default batch_size:", batchSize)
	}
	return nil
}
