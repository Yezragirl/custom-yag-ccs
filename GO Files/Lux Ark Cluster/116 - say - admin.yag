{{/* say Admin Command to make yag say something specific in a specified channel */}}
{{$args := parseArgs 2 "correct syntax is `say channel message`"
     (carg "channel" "channel")
(carg "string" "message")}}
{{$channel := ($args.Get 0).ID}}
{{$message := (toString ($args.Get 1))}}
{{sendMessage $channel $message}}
{{deleteTrigger 0}}