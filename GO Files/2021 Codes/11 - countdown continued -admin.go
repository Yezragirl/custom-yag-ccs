{{/*countdown part 2*/}}


{{$t := .ExecData.T}}
{{$tenpercent := toDuration (div $t 10)}}
{{$mID := .ExecData.MessageID}}
{{$cntDownMessageHeader := joinStr "" "Countdown To " .ExecData.Message}}
{{$color := 16711680}}
	{{$embed := cembed 
	"Title" $cntDownMessageHeader
	"color" $color
	"Description" (joinStr "" "This countdown ended <t:" $t.Unix ":R>.")
	"footer" (sdict "text" "Timer ended ") 
	"timestamp" $t}}
{{editMessage nil  .ExecData.MessageID $embed}}
{{sendMessage nil (joinStr "" .User.Mention ", your countdown is complete.")}}