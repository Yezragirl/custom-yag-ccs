-se -color "#388E3C" -title "Arkmart" -desc "If you want to place an order with the ArkMart, react to this panel with 1️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#A54BD2" -title "Base Issues" -desc "If you need help with any Base Issues, react to this panel with the 3️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#3ADDA6" -title "Auction" -desc "If you want to ask a question about, or put in a proxy bid for either our Discord Auction or Live Auction, react to this panel with the 2️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#F3003A" -title "Dino Issues" -desc "If you need help with any Dino Issues, react to this panel with the 4️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#DDFB28" -title "Discord" -desc "If you need help with any Discord Issues, react to this panel with the 5️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#0018FF" -title "Quests" -desc "If you need help with Quests, or want to turn in a New Player Quest, react to this panel with the 6️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#9EABAC" -title "Traps" -desc "If you have found a trap that needs to be removed, react to this panel with the 7️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#000000" -title "Tribe Issues" -desc "If you need help with any Tribe Issues, react to this panel with the 8️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#E902DE" -title "Change Registration Info" -desc "To edit your Registration Information, including your Pin, Tribe Name, In Game Name, or Gamer Tag, please reactive with the 9️⃣ below." -footer "Yez's Ark Cluster"
-se -color "#FDFFFD" -title "General" -desc "If your question/issue isn't covered by any other panel, react to this panel with the 🔟 below." -footer "Yez's Ark Cluster"

-se -color "#FDFFFD" -title "You leveled up!" -desc (joinStr "" "Congratulations " .User.Username ",\n you leveled up to **level " (toInt $newlevel) "**!\n\nFor more information on our XP system, type `?xp` in any channel.") 

1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣🔟


For restarting counting at 1.


{{if eq .Channel.ID 584957111872782366}}{{dbSet 18 "counter_count" 1}}{{end}}


{{$matches := reFindAll   "(\\p{So})|(<(a?):((\\w|\\d|~|_)+):(\\d{18})>)|(\\x{20e3})"  .Message.Content}}
{{if gt (len $matches) 5}}
Emoji Spam Detected
{{end}} 





-se -color "#388E3C" -title "Resource Roundup Winners" -desc "Here are you winners!
`1st Place - McLovin
2nd Place - Bug
3rd Place - Mrs. K
4th Place - Wiz
5th Place - Fire
6th Place - Mr. Pita
7th Place - Miller
8th Place - King
9th Place - Starbright
10th Place - Draygoon Rider`" -footer "Yez's Ark Cluster"




First command

{{$name := .User.Username}}
{{if .Member.Nick}}
  {{$name = .Member.Nick}}
{{end}}
{{editNickname (joinStr " " "[AFK]" $name)}}


second command
{{ $matches := reFindAllSubmatches `(AFK)` .Member.Nick }}
{{ if $matches }}
{{$matches}}
{{editNickname (reReplace "(\\[AFK\\] )" .Member.Nick "")}}
{{end}}


-se -color "#388E3C" -title "Random Tribes Base Registration" -desc "Tribe Owner: Owner's Name
Base Locations:
Gryphon
coords (blank if no base on map)
Manticore
coords (blank if no base on map)
Artemis
coords (blank if no base on map)
Zelda
coords (blank if no base on map)
Phoenix
coords (blank if no base on map)
Elysian
coords (blank if no base on map)
Titan
coords (blank if no base on map)
Raiden
coords (blank if no base on map)
Medusa
coords (blank if no base on map)
Beowulf
coords (blank if no base on map)
Water Pen 
Map and coords (blank if no water Pen)
Buiding Challenge Build
Map and coords (blank if no Building Challenge Build)" -footer "Yez's Ark Cluster"




{{$skip = add $skip 20}}
{{execCC 169 nil 2 (sdict "skip" $skip)}}


{{if reFind `\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d` .Key}} Matches Pattern
{{if getChannel .Key}} DO Nothing {{else}} {{dbDel .UserID .Key}} Deleted {{end}}{{end}}



WEEKLY ANNOUNCEMENT SCHEDULE AT A TIME YOU SPECIFY


{{$Weekday := 1}} {{$Hour := 22}} {{$Minute := 30}} {{$Second := 0}}

{{/* code starts */}}
{{$now := currentTime.Add .TimeSecond}} {{/*Add one second buffer because scheduler schedules a bit early at times*/}}

{{$date := add $now.Day (mod (sub $Weekday (toInt (json $now.Weekday))) 7)}}
{{$next := newDate $now.Year $now.Month $date $Hour $Minute $Second}}

{{if not ($now.Before $next)}}
    {{$next = $next.Add (toDuration "7d")}}
{{end}}

{{if .ExecData}}
Hello World
{{end}}

Next Run - {{$next}}

{{scheduleUniqueCC .CCID nil ($next.Sub currentTime).Seconds "scheduled-command-1" "dummy-data"}}






-poll "Should we turn off building in Mission Zones, if so when?" "Leave Building in Mission Zones ON (oh well about the boss fight)" "Turn Building in Mission Zones Off - Immediately (oh well about the bases)" "Turn Building in Mission Zones Off - in 1 Week" "Turn Building in Mission Zones Off - in 2 Weeks" "Other - Discuss in Suggestion Discussion"

-se -color "#FDFFFD" -title "Beowulf Mission Zones" -desc "PER COMMUNITY VOTE: Building WITHIN Mission Zones will be turned off! If you are currently built in a Mission Zone, you have until Next Sunday (August 9th) to move your base or it will disappear. Admin's will be happy to assist, but once it's gone, it's gone. To get Admin Assistance moving, please put in a base ticket." -footer "Yez's Ark Cluster"




{{$startOfCmd := currentTime}}
{{/* code here */}}
{{sendMessage 587858995012698137 (print "Steal Command took " (currentTime.Sub $startOfCmd) " to run.")}}



{{ $start := currentTime }}
{{/* code */}}

{{ $dur := currentTime.Sub $start }}
{{ if gt $dur.Seconds 5.0 }}
    {{ sendMessage 587858995012698137 (print "CC #" .CCID " took longer than 5 seconds to run.\n**Duration:** " $dur "\n**Command executed:* `" .Message.Content "`") }}
{{ end }}	




reFindAll `\d+` $str
$embed.Footer.Text


len (toRune $x)  - for length of nickname


{{$skip := 0}}
{{ if .ExecData }}
	{{$skip = .ExecData.skip}}
{{end}}
	{{$lb := dbBottomEntries "registration" 2 $skip}}
{{range $lb}}{{.UserID}}**$**{{.Key}}**$**{{.Value}}
{{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}

{{if reFind `\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d\d` .Key}} Matches Pattern
{{if getChannel .Key}}
{{dbSet (toInt64 .Key) "ticket_control" .Value}}{{dbDel .UserID .Key}} New Database Entry Created. 
 {{else}} {{dbDel .UserID .Key}} Deleted {{end}}{{end}}
{{end}}
{{$skip = add $skip 2}}
{{execCC 169 nil 20 (sdict "skip" $skip)}}
{{dbDel .UserID "registration"}}


{{if not $reg.admin}}{{$reg.Set "admin" "undefined"}}{{else}} {{dbSet .UserID "admin" $reg.admin}}{{end}}
{{if not $reg.public}}{{$reg.Set "public" "undefined"}}{{else}} {{dbSet .UserID "public" $reg.public}}{{end}}
{{if not $reg.name}}{{$reg.Set "name" "undefined"}}{{else}} {{dbSet .UserID "name" $reg.name}}{{end}}
{{if not $reg.age}}{{$reg.Set "age" "undefined"}}{{else}} {{dbSet .UserID "age" $reg.age}}{{end}}
{{if not $reg.pin}}{{$reg.Set "pin" "undefined"}}{{else}} {{dbSet .UserID "pin" $reg.pin}}{{end}}
{{if not $reg.gt}}{{$reg.Set "gt" "undefined"}}{{else}} {{dbSet .UserID "gt" $reg.gt}}{{end}}
{{if not $reg.Q}}{{$reg.Set "Q" "undefined"}}{{else}} {{dbSet .UserID "Q" $reg.Q}}{{end}}
{{if not $reg.ref}}{{$reg.Set "ref" "undefined"}}{{else}} {{dbSet .UserID "ref" $reg.ref}}{{end}}
{{if not $reg.tribe}}{{$reg.Set "tribe" "undefined"}}{{else}} {{dbSet .UserID "tribe" $reg.tribe}}{{end}}



{{$skip := 0}}
{{ if .ExecData }}
	{{$skip = .ExecData.skip}}
{{end}}
	{{$lb := dbBottomEntries "registration" 2 $skip}}
{{range $lb}}
{{$reg := sdict}}{{range $k,$v := .Value}}{{$reg.Set $k $v}}{{end}}
{{.UserID}}**$**{{.Key}}**$**{{.Value}}
{{if not $reg.admin}}{{$reg.Set "admin" "undefined"}}{{else}}{{end}} {{dbSet .UserID "admin" $reg.admin}}
{{if not $reg.public}}{{$reg.Set "public" "undefined"}}{{else}}{{end}} {{dbSet .UserID "public" $reg.public}}
{{if not $reg.name}}{{$reg.Set "name" "undefined"}}{{else}}{{end}}{{dbSet .UserID "name" $reg.name}}
{{if not $reg.age}}{{$reg.Set "age" "undefined"}}{{else}}{{end}} {{dbSet .UserID "age" $reg.age}}
{{if not $reg.pin}}{{$reg.Set "pin" "undefined"}}{{else}} {{end}}{{dbSet .UserID "pin" $reg.pin}}
{{if not $reg.gt}}{{$reg.Set "gt" "undefined"}}{{else}}{{end}} {{dbSet .UserID "gt" $reg.gt}}
{{if not $reg.Q}}{{$reg.Set "Q" "undefined"}}{{else}}{{end}} {{dbSet .UserID "Q" $reg.Q}}
{{if not $reg.ref}}{{$reg.Set "ref" "undefined"}}{{else}} {{end}}{{dbSet .UserID "ref" $reg.ref}}
{{if not $reg.tribe}}{{$reg.Set "tribe" "undefined"}}{{else}}{{end}} {{dbSet .UserID "tribe" $reg.tribe}}
{{end}}
{{$skip = add $skip 2}}
{{execCC 169 nil 2 (sdict "skip" $skip)}}







{{$string1 := "GILY"}}
{{$string2 := "GEELY"}}


{{$out := dict 1 $string1 2 $string2}}
{{template "levenshtein" $out}}

{{with $out.err}} `error:` {{.}}
{{else}} `distance:` {{$out.dist}}
{{end}}

{{define "levenshtein"}}
    {{$a := .Get 1}} {{$b := .Get 2}} {{$m := cslice 0}} {{$last := 0}} {{$x := 1}} {{$y := 1}}
    {{if and $a $b}}
        {{if ne $a $b}}
            {{$a = toRune $a}} {{$b = toRune $b}}
            {{$lena := len $a}} {{$lenb := len $b}}
            {{range seq 1 (add $lena 1)}} {{- $m = $m.Append . -}} {{end}}
            {{range $v := $b}}
                {{- $m.Set 0 $x -}}
                {{- $last = sub $x 1 -}}
                {{- $y = 1 -}}
                {{- range $a -}}
                    {{- $third := 1 -}}
                    {{- if eq . $v -}} {{- $third = 0 -}} {{- end -}}
                    {{- $min3 := cslice (add (index $m $y) 1) (add (index $m (sub $y 1)) 1) (add $last $third) -}}
                    {{- $last = index $m $y}}
                    {{- template "min3" $min3 -}}
                    {{- $m.Set $y (index $min3 0) -}}
                    {{- $y = add $y 1 -}}
                {{- end -}}
                {{- $x = add $x 1 -}}
            {{end}}
            {{.Set "dist" (index $m $lena)}}
        {{else}}
            {{.Set "dist" 0}}
        {{end}}
    {{else}}
        {{.Set "err" "You dint provide 2 strings to compare."}}
    {{end}}
{{end}}

{{define "min"}}
    {{if gt (index . 0) (index . 1)}}
        {{.Set 0 (index . 1)}}
    {{end}}
{{end}}

{{define "min3"}}
    {{$inner := cslice (index . 1) (index . 2)}}
    {{template "min" $inner}}
    {{$outter := cslice (index . 0) (index $inner 0)}}
    {{template "min" $outter}}
    {{.Set 0 (index $outter 0)}}
{{end}}