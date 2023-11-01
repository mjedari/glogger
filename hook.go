package glogger

type SubscriptionList []ISubscriber

func Subscribe(subscribers ...ISubscriber) {
	for _, subscriber := range subscribers {
		std.subscribers = append(std.subscribers, subscriber)
	}
}

// publish based on method that subscriber provides
func (list SubscriptionList) publish(level LogLevel, message []byte) {
	for _, subscriber := range list {
		processFunc := subscriber.GetClosure()
		if processFunc == nil {
			return
		}

		err := processFunc(message)
		if err != nil {
			Errorf("logger subscriber error: %v \n", err)
		}
	}
}
