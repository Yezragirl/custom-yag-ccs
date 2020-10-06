{{$id := .User.ID}}
{{$args := parseArgs 0 "correct syntax is `myv (who)`"
     (carg "userid" "Violator")}}
{{if $args.IsSet 0}}
{{$id = $args.Get 0}}
{{end}}
{{$color := randInt 111111 999999 }}
{{$r48 := (dbGet $id 48).Value}}
{{$r56 := (dbGet $id 56).Value}}
{{$r7 := (dbGet $id 7).Value}}
{{$r9 := (dbGet $id 9).Value}}
{{$r14 := (dbGet $id 14).Value}}
{{$r16 := (dbGet $id 16).Value}}
{{$r17 := (dbGet $id 17).Value}}
{{$r18 := (dbGet $id 18).Value}}
{{$r19 := (dbGet $id 19).Value}}
{{$r20 := (dbGet $id 20).Value}}
{{$embed := cembed 
"color" $color
"fields" (cslice 
    (sdict "name" "Current Violations for User" "value"  (joinStr "" "<@" $id ">") "inline" false)  
    (sdict "name" "Rules 4 & 8: Drama/Bad Attitude/Trolling/Kiting/Land grabbing/Stealing" "value" (toString (toInt $r48)) "inline" false) 
    (sdict "name" "Rules 5 & 6: Building in Render/Blocking Resource Spawns" "value" (toString (toInt $r56)) "inline" false) 
    (sdict "name" "Rule 7: Mass Breeding/Dinos Left on Aggressive" "value" (toString (toInt $r7)) "inline" false) 
    (sdict "name" "Rule 9: Public Teleporters" "value" (toString (toInt $r9)) "inline" false) 
    (sdict "name" "Rule 14: Traps" "value" (toString (toInt $r14)) "inline" false) 
	(sdict "name" "Rule 16: Base Size/Count" "value" (toString (toInt $r16)) "inline" false) 
	(sdict "name" "Rule 17: Base Registration" "value" (toString (toInt $r17)) "inline" false) 
	(sdict "name" "Rule 18: Piping" "value" (toString (toInt $r18)) "inline" false)
	(sdict "name" "Rule 19: Disrespecting of Admins" "value" (toString (toInt $r19)) "inline" false) 
    (sdict "name" "Rule 20: Direct Messaging of Admin Staff" "value" (toString (toInt $r20)) "inline" false) )
 }}
{{sendMessage nil $embed}}