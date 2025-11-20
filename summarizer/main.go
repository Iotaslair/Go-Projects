package main

import (
	"summarizer/alertevent"
	"summarizer/assets"
)

func main() {
	/* Make a controller class:
	This should get the alertEvent from the processingId
	Use the alertEventId to call the asset module to get the current assets
	If it's a new alert then write the summary
	else get previous alert event
		Get assets from previous alert event
		Do the diff
		Write the summary
	*/
	processingId := 123

	currentAlertEvent := alertevent.GetAlertEvent(processingId)

	currentAssets := assets.GetAssetIds(currentAlertEvent.AlertId)

}
