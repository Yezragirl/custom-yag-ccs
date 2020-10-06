{{$args := parseArgs 2 "correct syntax is `addv rule# who (optional explanation)`"
    (carg "int" "Rule Number")
    (carg "user" "Violator")
    (carg "string" "Text")}}
{{$countofviolations := dbIncr ($args.Get 1).ID ($args.Get 0) 1}}
{{$fine := "0"}}
{{$warning := ""}}
{{ range seq 1 4 }}
{{$TierLevel := (joinStr "" "Tier" .)}}
{{$TierFine := (joinStr "" "Tier" . "Fine")}}
{{$TierWarning := (joinStr "" "Tier" . "Warn")}}
{{$TierViolationThreshold := (dbGet (toInt64  ($args.Get 0)) $TierLevel).Value}}
{{if ge (toInt $countofviolations) (toInt $TierViolationThreshold)}}
    {{$fine = (dbGet (toInt64  ($args.Get 0)) $TierFine).Value}}
    {{$warning = (dbGet (toInt64  ($args.Get 0)) $TierWarning).Value}}
  {{end}}
{{end}}
{{if and (eq (toString $fine) "0") (eq (toString $warning) "")}}  
{{sendMessage nil "error"}}
{{else}}
{{sendMessage nil (joinStr "" "<@" ($args.Get 1).ID ">, " $warning " " ($args.Get 2))}}
{{end}}
{{if or (eq (toString $fine) "0") (not $fine)}}
{{else}}
{{sendMessageNoEscape nil (joinStr "" (mentionRoleID 573210843840249889) ", " (getMember ($args.Get 1).ID).Nick " needs to be fined " (toString (toInt $fine)) " BBS for violations of Rule " ($args.Get 0) ".")}}
{{end}}
{{deleteTrigger 1}}

