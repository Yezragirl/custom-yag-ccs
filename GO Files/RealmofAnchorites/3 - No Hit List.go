{{$t := .ExecData.T}}
{{$mID := .ExecData.MessageID}}
{{$ts := .TimeSecond}}
{{$timeLeft := .ExecData.T.Sub currentTime}}
{{$message := (joinStr "" "**" .ExecData.Message "** Time Remaining on List: " $t.String)}}
{{$formattedTimeLeft := humanizeDurationSeconds $timeLeft}}

{{if lt $timeLeft (mult .TimeSecond 10)}}
{{deleteMessage  nil  $mID}}
{{else}}
    {{editMessage nil $mID (joinStr "" "**" .ExecData.Message "** Time Remaining on List: " $formattedTimeLeft)}}
    {{execCC 3 nil 30 .ExecData}}
{{end}}