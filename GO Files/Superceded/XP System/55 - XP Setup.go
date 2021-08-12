{{/*command xpsetup*/}}
{{$minimum := toInt (dbGet 0 "minimum").Value}}
{{$maximum := toInt (dbGet 0 "maximum").Value}}
{{$cooldown := toInt (dbGet 0 "cooldown").Value}}
{{$entries := toInt (dbGet 0 "entries").Value}}

{{if not $minimum}}
{{$minimum = 15}}
{{dbSet 0 "minimum" $minimum}}
{{end}}

{{if not $maximum}}
{{$maximum = 25}}
{{dbSet 0 "maximum" $maximum}}
{{end}}

{{if not $cooldown}}
{{$cooldown = 60}}
{{dbSet 0 "cooldown" $cooldown}}
{{end}}

{{if not $entries}}
{{$entries = 25}}
{{dbSet 0 "entries" $entries}}
{{end}}

{{$embed := cembed 
"title" "XP Setup Menu" 
"description" "**This XP System is using the following settings.**" 
"color" 4645612 
"fields" (cslice 
       (sdict "name" "**[Minimum] XP Per Message**" "value" (toString $minimum) "inline" false) 
       (sdict "name" "**[Maximum] XP Per Message**" "value" (toString $maximum) "inline" false) 
       (sdict "name" "**[Cooldown] Length in Seconds**" "value" (toString $cooldown) "inline" false)
       (sdict "name" "**[Entries] Per Leaderboard Page**" "value" (toString $entries) "inline" false) ) 
"footer" (sdict "text" "To change any of these settings, use -set [setting] [new value].")}}
{{sendMessage nil $embed}}


