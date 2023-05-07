package usecases

import (
	"fmt"
	custom_types "slack-topic-parser/internal/api/domain/custom-types"
	topic_parser "slack-topic-parser/internal/topic-parser"
	topic_queue "slack-topic-parser/internal/topic-queue"
)

func RemoveUser(data custom_types.PostRequest) custom_types.PostResponse {
	fmt.Println("Processing Remove User Usecase")

	previousTopic := topic_queue.PreviousTopic()
	lines := topic_parser.ParagraphToLines(previousTopic)

	topicTitle := lines[0]
	prodUsers := topic_parser.ExtractUsers(lines[1])
	preUsers := topic_parser.ExtractUsers(lines[2])

	// Remove user
	if data.Environment == "PROD" {
		prodUsers = topic_parser.RemoveUser(prodUsers, data.User)
	}
	if data.Environment == "PRE" {
		preUsers = topic_parser.RemoveUser(preUsers, data.User)
	}

	newTopic := topic_parser.MakeTopic(topicTitle, prodUsers, preUsers)

	topic_queue.StablishNewTopic(newTopic)

	res := custom_types.PostResponse{Data: newTopic}
	return res
}
