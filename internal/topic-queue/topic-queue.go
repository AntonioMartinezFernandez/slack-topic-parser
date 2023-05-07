package topic_queue

import "fmt"

var TopicName string = "⭐ Environment Queue\n  ⚙️  PROD: 🆓\n  🧪 PRE: 🆓"

func PreviousTopic() string {
	// TODO: Read topic from Slack channel
	fmt.Println("----- Previous channel topic -----")
	fmt.Println(TopicName)
	return TopicName
}

func StablishNewTopic(newTopic string) {
	TopicName = newTopic
	// TODO: Update Slack channel topic
	fmt.Println("----- Stablishing new channel topic -----")
	fmt.Println(newTopic)
}
