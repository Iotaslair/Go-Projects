package alertevent

type AlertEvent struct {
	ProcessingId          int
	AlertId, AlertEventId string
}

func GetAlertEvent(processingId int) AlertEvent {
	//In our system this is a redis call so temp mocking it out
	//TODO fake a call to redis
	return AlertEvent{ProcessingId: processingId, AlertId: "alertId", AlertEventId: "alertEventId"}
}

// TODO build this out
// func IsAlertNew(alertEvent AlertEvent) bool {

// }
