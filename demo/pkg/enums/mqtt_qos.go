package enums

//go:generate toolkit gen enum MqttQOS
type MqttQOS int8

const (
	MQTT_QOS_UNKNOWN        MqttQOS = iota - 1
	MQTT_QOS__ONCE                  // 0
	MQTT_QOS__AT_LEAST_ONCE         // 1
	MQTT_QOS__ONLY_ONCE             // 2
)

// func (QoS) Offset() int { return -1 }
