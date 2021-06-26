{{$id := .User.ID}}
{{$args := parseArgs 0 "correct syntax is `myv (who)`"
     (carg "userid" "Violator")}}
{{if $args.IsSet 0}}
{{$id = $args.Get 0}}
{{end}}
{{$color := randInt 111111 999999 }}
{{$r1 := (dbGet $id "1").Value}}
{{$r2 := (dbGet $id "2").Value}}
{{$r3A := (dbGet $id "3A").Value}}
{{$r3B := (dbGet $id "3B").Value}}
{{$r3C := (dbGet $id "3C").Value}}
{{$r4 := (dbGet $id "4").Value}}
{{$r5 := (dbGet $id "5").Value}}
{{$r6 := (dbGet $id "6").Value}}
{{$r7 := (dbGet $id "7").Value}}
{{$r8 := (dbGet $id "8").Value}}
{{$r9 := (dbGet $id "9").Value}}
{{$r10 := (dbGet $id "10").Value}}
{{$r11 := (dbGet $id "11").Value}}
{{$r12 := (dbGet $id "12").Value}}
{{$r13 := (dbGet $id "13").Value}}
{{$r14 := (dbGet $id "14").Value}}
{{$r15 := (dbGet $id "15").Value}}
{{$embed := cembed 
"color" $color
"fields" (cslice 
    (sdict "name" "Current Violations for User" "value"  (joinStr "" "<@" $id ">") "inline" false)  
    (sdict "name" "Rule 1: Earn It" "value" (toString (toInt $r1)) "inline" false) 
    (sdict "name" "Rule 2: Age" "value" (toString (toInt $r2)) "inline" false) 
    (sdict "name" "Rule 3A: Player Registration" "value" (toString (toInt $r3A)) "inline" false) 
    (sdict "name" "Rule 3B: Tribe Registration" "value" (toString (toInt $r3B)) "inline" false) 
    (sdict "name" "Rule 3C: Base Registration" "value" (toString (toInt $r3C)) "inline" false) 
    (sdict "name" "Rule 4: Harassment" "value" (toString (toInt $r4)) "inline" false) 
    (sdict "name" "Rule 5: Drama/Attitude" "value" (toString (toInt $r5)) "inline" false) 
	(sdict "name" "Rule 6: Unsportsmanlike Conduct" "value" (toString (toInt $r6)) "inline" false) 
	(sdict "name" "Rule 7: Building In Render" "value" (toString (toInt $r7)) "inline" false) 
	(sdict "name" "Rule 8: Base Size/Count" "value" (toString (toInt $r8)) "inline" false)
	(sdict "name" "Rule 9: Piping" "value" (toString (toInt $r9)) "inline" false) 
    (sdict "name" "Rule 10: Mass Breeding" "value" (toString (toInt $r10)) "inline" false) 
    (sdict "name" "Rule 11: Teleporters" "value" (toString (toInt $r11)) "inline" false) 
    (sdict "name" "Rule 12: Traps" "value" (toString (toInt $r12)) "inline" false) 
    (sdict "name" "Rule 13: Dinos Out" "value" (toString (toInt $r13)) "inline" false) 
    (sdict "name" "Rule 14: Disrespecting an Admin" "value" (toString (toInt $r14)) "inline" false) 
    (sdict "name" "Rule 15: DMing Admins" "value" (toString (toInt $r15)) "inline" false) )
 }}
{{sendMessage nil $embed}}