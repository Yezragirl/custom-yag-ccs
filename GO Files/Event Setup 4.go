{{$channel := 573897276569813012}}
{{$t := .ExecData.TimeLeft}}
{{$msgID := .ExecData.MessageID}}
{{$ts := .TimeSecond}}
{{$message := getMessage $channel (toString $msgID)}}
{{ if $message.Embeds }}
	{{ $embed := structToSdict (index $message.Embeds 0) }}
	{{/* As structToSdict is not deep (only converts 1 layer) we need to range over */}}
	{{ range $k, $v := $embed }}
 		{{- if eq (kindOf $v true) "struct" }}
 		{{- $embed.Set $k (structToSdict $v) }}
		 {{- end -}}
		 {{index $embed "Time Until"}}
 	{{ end }}
	{{/* structToEmbed does not change the keys so there are some inconsistencies (causing some parts of the embed to be left out without the following code). */}}
	{{ if $embed.Author }} {{ $embed.Author.Set "Icon_URL" $embed.Author.IconURL }} {{ end }}
 	{{ if $embed.Footer }} {{ $embed.Footer.Set "Icon_URL" $embed.Footer.IconURL }} {{ end }}
	{{$timeLeft := $t.Sub currentTime}}
	{{$formattedTimeLeft := humanizeDurationSeconds $timeLeft}}
 		{{if lt $timeLeft (mult .TimeSecond 30)}}
   			{{range seq 1 (toInt $timeLeft.Seconds) }}
				{{$timeLeft := .ExecData.TimeLeft.Sub currentTime}}
				{{$formattedTimeLeft := humanizeDurationSeconds $timeLeft}}
				{{$embed.Set "Time Until" (joinStr "" $formattedTimeLeft " seconds")}}
				{{editMessage $channel (toString $msgID) (cembed $embed)}}
	 				{{if gt $timeLeft $ts}} {{sleep 1}} {{end}}
   			{{end}}
			{{/* You can now do whatever with $embed - for example, `$embed.Set "title" "test"` sets the embed title to test. */}}
		 {{$embed.Set "Time Until" (joinStr "" "**Event has Begun**")}}
		 {{editMessage $channel (toString $msgID) (cembed $embed)}}
	{{else}}
		{{$embed.Set "Time Until" (joinStr "" $formattedTimeLeft " seconds")}}
		{{editMessage $channel (toString $msgID) (cembed $embed)}}
		{{execCC 194 nil 10 .ExecData}}
{{end}}{{end}}

