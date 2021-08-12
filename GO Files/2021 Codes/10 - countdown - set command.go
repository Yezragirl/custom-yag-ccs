{{/*countdown part 1 */}}
{{$args := parseArgs 2 ""
 (carg "duration" "countdown-duration")
 (carg "string" "countdown-message")}}
 {{$cntDownMessageHeader := joinStr "" "Countdown To " ($args.Get 1)}}
{{$t := currentTime.Add ($args.Get 0)}}
{{$color := 32768 }}
	{{$embed := cembed 
	"Title" $cntDownMessageHeader
	"color" $color
	"Description" (joinStr  "" ($args.Get 1) " is <t:" $t.Unix ":R>.")
	"footer" (sdict "text" "Timer ends ") 
	"timestamp" $t}}
{{$mID := sendMessageRetID nil $embed}}
{{$timeLeft := $t.Sub currentTime}}
{{$timetrigger := toInt ($args.Get 0).Seconds}}
{{execCC 11 nil $timetrigger (sdict "MessageID" $mID "T" $t "Message" ($args.Get 1)) }}