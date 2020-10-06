{{$args := parseArgs 1 "sd <user> <duration:00h00m00s>"
    (carg "user" "User")
    (carg "duration" "Duration")}}
{{$time := (toDuration ($args.Get 1).Seconds)}}
{{$distanced := ($args.Get 0).ID}}
{{$formattedTime:= humanizeDurationSeconds ($args.Get 1)}}
$time: {{$time}} Type: {{printf "%T" $time}}
$formattedTime: {{$formattedTime}} Type: {{printf "%T" $formattedTime}}

{{if not .ExecData}}
    {{giveRoleID $distanced 758126543276474399}}
{{sendMessage nil (print "<@" $distanced "> you have been ***socially distanced***, please do not resist.")}}
{{$embed := cembed 
    "title" (print "Someone was just Socially Distanced!") 
    "description" (print  .User.Mention " executed the <@&758126543276474399> command!\n<@" $distanced "> was given the role for a total of " $formattedTime)
    "color" 4645612 
    "timestamp"  currentTime
}}
{{sendMessage 713268227538419743 $embed }}
{{sendMessage 706969207983833189 (print "Really <@" $distanced ">!? That isn't allowed... You will be required to **Social Distance** for " $formattedTime )}}
{{execCC .CCID nil $time $distanced}}
{{else}}
    {{$distanced := .ExecData}}
    {{takeRoleID $distanced 758126543276474399}}
{{end}}
