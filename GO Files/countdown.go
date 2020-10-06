
{{$timeleft := currentTime.Add $start}}
{{execCC 11 nil 0 (sdict "MessageID" $msgID "TimeLeft" $timeleft) }}





{{$timeLeft := .ExecData.T.Sub currentTime}}
{{$cntDownMessageHeader := joinStr "" "Countdown To " .ExecData.Message}}
{{$formattedTimeLeft := humanizeDurationSeconds $timeLeft}}

{{$t := .ExecData.T}}
{{$mID := .ExecData.MessageID}}
{{$ts := .TimeSecond}}

{{if lt $timeLeft (mult .TimeSecond 30)}}
  {{range seq 1 (toInt $timeLeft.Seconds) }}
    {{$timeLeft := $t.Sub currentTime}}
    {{$formattedTimeLeft := humanizeDurationSeconds $timeLeft}}

    {{editMessage nil $mID (joinStr "" $cntDownMessageHeader "\nTime left: " $formattedTimeLeft " seconds")}}
    {{if gt $timeLeft $ts}} {{sleep 1}} {{end}}
  {{end}}
  {{editMessage nil  .ExecData.MessageID (joinStr "" $cntDownMessageHeader "\nTime left: **ENDED**")}}
{{else}}
    {{editMessage nil .ExecData.MessageID (joinStr "" $cntDownMessageHeader "\nTime left: " $formattedTimeLeft)}}
    {{execCC 11 nil 10 .ExecData}}
{{end}}