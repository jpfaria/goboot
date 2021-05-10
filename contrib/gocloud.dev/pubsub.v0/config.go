package pubsub

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root     = "gi.gocloud"
	resource = root + ".resource"
	tp       = root + ".type"
	region   = root + ".region"
)

func init() {
	config.Add(tp, "memory", "define queue type")
	config.Add(resource, "topicA", "define queue resource")
	config.Add(region, "", "define queue region")
}
