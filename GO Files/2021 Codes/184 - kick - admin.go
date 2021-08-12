{{/* Kick*/}}
{{$args := parseArgs 2 "kick who why"
(carg "userid" "user")
(carg "string" "reason")}}
{{$user := (userArg ($args.Get 0))}}
{{$reason := ($args.Get 1)}}
{{execAdmin (joinStr "" "kick " $user.ID $reason) }}

