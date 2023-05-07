package topic_parser

import (
	"fmt"
	"strings"
)

func ParagraphToLines(paragraph string) []string {
	return strings.Split(paragraph, "\n")
}

func ExtractUsers(envLine string) []string {
	var result []string

	lineValues := strings.Split(envLine, " ")
	for i := 0; i < len(lineValues); i++ {
		if strings.HasPrefix(lineValues[i], "@") {
			result = append(result, lineValues[i][1:])
		}
	}
	return result
}

func AddFirstUser(users []string, privilegedUser string) []string {
	users = RemoveUser(users, privilegedUser)
	return append([]string{privilegedUser}, users...)
}

func AddUser(users []string, newUser string) []string {
	users = RemoveUser(users, newUser)
	var result []string

	for i := 0; i < len(users); i++ {
		result = append(result, users[i])
	}
	result = append(result, newUser)

	return result
}

func RemoveUser(users []string, userToRemove string) []string {
	var result []string

	for i := 0; i < len(users); i++ {
		if users[i] != userToRemove {
			result = append(result, users[i])
		}
	}

	return result
}

func UsersToString(users []string) string {
	var result string
	for i := 0; i < len(users); i++ {
		result = result + "@" + users[i]
		if i < len(users)-1 && result != "" {
			result = result + " < "
		}
	}
	if len(users) == 0 {
		result = "🆓"
	}
	return result
}

func MakeTopic(topicTitle string, prodUsers []string, preUsers []string) string {
	prodQueue := "  ⚙️  PROD: " + UsersToString(prodUsers)
	preQueue := "  🧪 PRE: " + UsersToString(preUsers)

	return fmt.Sprintf("%s\n%s\n%s", topicTitle, prodQueue, preQueue)
}
