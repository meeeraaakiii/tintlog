package logger

import (
	"bytes"
	"encoding/json"
)

func LogPrettyJSON(payload any, payloadName string) {
	// 1) Marshal the struct to raw JSON bytes
	raw, err := json.Marshal(payload)
	if err != nil {
		Log(Color, "%s (marshal error): %v\n'''\n%+v\n'''", payloadName, err, payload)
		return
	}

	// 2) Try to pretty-print; on failure, fall back to raw
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, raw, "", "    "); err != nil {
		// just print raw if can't prettify JSON
		Log(Color, "Raw %s (raw marshaled):\n'''\n%s\n'''", payloadName, string(raw))
		return
	}

	Log(Color, "%s (pretty JSON):\n'''\n%s\n'''", payloadName, pretty.String())
}

func LogPrettyJSONBytes(payload []byte, payloadName string) {
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, payload, "", "    "); err != nil {
		// just print raw if can't prettify JSON
		Log(Color, "%s (raw):\n'''\n%s\n'''", payloadName, string(payload))
		return
	}

	Log(Color, "%s (pretty JSON):\n'''\n%s\n'''", payloadName, pretty.String())
}
