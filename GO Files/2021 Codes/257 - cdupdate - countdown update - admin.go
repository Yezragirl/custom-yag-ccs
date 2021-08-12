{{/*cdupdate for fixing the countdowns when they break*/}}
{{$args := parseArgs 2 ""
 (carg "duration" "countdown-duration")
 (carg "string" "message-ID")}}
{{$mID := ($args.Get 1)}}
{{$embedmsg := getMessage nil $mID}}
{{$embed := index $embedmsg.Embeds 0}}
{{$message := (reReplace "Countdown To " $embed.Title "")}}
{{$t := currentTime.Add ($args.Get 0)}}
{{$color := 32768 }}
	{{$embed := cembed 
	"Title" "Countdown"
	"color" $color
	"Description" (joinStr  "" "countdown starting..." $t.String)
	"footer" (sdict "text" "Yez's Ark Cluster")}}
{{editMessage nil $mID $embed}}
{{execCC 11 nil 0 (sdict "MessageID" $mID "T" $t "Message" $message) }}
{{deleteTrigger 1}} 