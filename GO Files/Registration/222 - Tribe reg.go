{{$msgID := .ReactionMessage.ID}}
{{$Q := (dbGet .User.ID "Q").Value}}
{{$admin := (dbGet .User.ID "admin").Value}}
{{$age := (dbGet .User.ID "age").Value}}
{{$gt := (dbGet .User.ID "gt").Value}}
{{$name := (dbGet .User.ID "name").Value}}
{{$pin := (dbGet .User.ID "pin").Value}}
{{$public := (dbGet .User.ID "public").Value}}
{{$ref := (dbGet .User.ID "ref").Value}}
{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
{{$tribe := (dbGet .User.ID "tribe").Value}}{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$tribe2 := (dbGet .User.ID "tribe2").Value}}{{if not $tribe2}}{{$tribe2 = "Unregistered"}}{{dbSet .User.ID "tribe2" $tribe2}}{{end}}
{{$color := randInt 111111 999999 }}
{{$tribeedit := .ExecData.tribe}}
{{$step := .ExecData.step}}



{{if and (eq $tribeedit "first") (eq $step "join")}}
{{sendMessage nil (joinStr "" .User.Mention ", type the first few letters, or words of the tribe you'd like to join.")}}
{{$Q = "Edit First Tribe Join"}}
{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "first") (eq $step "leave")}}
{{sendMessage nil (joinStr "" .User.Mention ", ok, you have left " (toString $tribe) ".")}}
{{$tribe = "Unregistered"}}
{{dbSet .User.ID "tribe" $tribe}}
{{$Q = "Edit First Tribe Leave"}}
{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "first") (eq $step "create")}}
{{sendMessage nil (joinStr "" .User.Mention ", ok, type what you would like the Tribe to be called.")}}
{{$Q = "Edit First Tribe Create"}}
{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "second") (eq $step "join")}}
{{sendMessage nil (joinStr "" .User.Mention ", type the first few letters, or words of the tribe you'd like to join.")}}
{{$Q = "Edit Second Tribe Join"}}
{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "second") (eq $step "leave")}}
{{sendMessage nil (joinStr "" .User.Mention ", ok, you have left " (toString $tribe2) ".")}}
{{$tribe2 = "Unregistered"}}
{{dbSet .User.ID "tribe2" $tribe2}}
{{$Q = "Edit Second Tribe Leave"}}
{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "second") (eq $step "create")}}
{{sendMessage nil (joinStr "" .User.Mention ", ok, type what you would like the Tribe to be called.")}}
{{$Q = "Edit Second Tribe Create"}}
{{dbSet .User.ID "Q" $Q}}
{{end}}