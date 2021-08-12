{{/*dk duration kick kick command with an automatic unkick?*/}}
{{$args := parseArgs 3 ""
 (carg "duration" "duration")
 (carg "user" "user")
 (carg "string" "message")}}

{{$t := ($args.Get 0)}}
{{$who := ($args.Get 1)}}
{{$msg := ($args.Get 2)}}

{{if .ExecData}}
	{{if targetHasRoleID $who.ID 648706778061602816}}
	{{ execAdmin (joinStr "" "Kick" (userArg $who).Mention $msg)}}
	{{end}}
{{end}}
{{execCC .CCID nil $t (sdict "msg" $msg  "who" (toInt64 $who.ID)) }}

