{{/*database checks how much of the database we are using*/}}
{{$dbcount := dbCount}}
{{$maxentries := mult .Server.MemberCount 50}}
{{if .IsPremium}}
     {{$maxentries = mult .Server.MemberCount 500}}
{{end}}
{{$perc := div (round (mult (fdiv ($dbcount) $maxentries) 10000)) 100}}
{{$dbcount}}/{{$maxentries}} ({{$perc}}% used)