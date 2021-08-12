{{$capture := exec "ticket open" "Registration"}}
{{$ID := reFind `\d+` (reFind `<#\d+>` $capture)}}
{{$Q := dbGet .User.ID "Q"}}
{{$admin := dbGet .User.ID "admin"}}
{{$age := dbGet .User.ID "age"}}
{{$gt := dbGet .User.ID "gt"}}
{{$name := dbGet .User.ID "name"}}
{{$pin := dbGet .User.ID "pin"}}
{{$public := dbGet .User.ID "public"}}
{{$parentguardian := dbGet .User.ID "parentguardian"}}
{{$ref := dbGet .User.ID "ref"}}
{{$tribe := dbGet .User.ID "tribe"}}
{{$ticketnum := (dbGet 0 "ticketnum").Value}}
{{deleteTrigger 1}}

{{if not $name}}
{{sendMessage $ID (joinStr "" .User.Mention ", Thanks for joining Yez's Ark Cluster!\n\nJust a few questions, and you'll be registered and ready to play!\nInstructions first. **If emojis appear below the question, please use them to respond. Otherwise, type your answers.** These answers will be editable later, but be sure to check for typos before hitting enter, as your answers affect the direction these questions go.")}}
{{$Q = "Age"}}{{dbSet .User.ID "Q" $Q}}
{{removeRoleID 645404993863811072}}
{{removeRoleID 573210479753691170}}
{{removeRoleID 607041870966685706}}
{{giveRoleID .User.ID 606566506922377257}}
{{sendMessage $ID (joinStr "" .User.Mention ", how old are you?")}}
{{$ticketnum = add $ticketnum 1}}
{{dbSet 0 "ticketnum" $ticketnum}}

{{else}}
{{removeRoleID 645404993863811072}}
{{removeRoleID 573210479753691170}}
{{removeRoleID 607041870966685706}}
{{giveRoleID .User.ID 606566506922377257}}
{{sendMessage $ID (joinStr "" .User.Mention ", Thanks for joining Yez's Ark Cluster!\n\nJust a few questions, and you'll be registered and ready to play!\nInstructions first. **If emojis appear below the question, please use them to respond. Otherwise, type your answers.** These answers will be editable later, but be sure to check for typos before hitting enter, as your answers affect the direction these questions go.\n_ _")}}

{{sendMessage $ID (joinStr "" .User.Mention ", please be advised that by starting a new registration, any previous registration has been erased. You must complete THIS registration.")}}


{{if $admin}}
{{deleteMessage 658354135221141533 $admin}}
{{end}}
{{if $public}}
{{deleteMessage 658352884701724733 $public}}
{{end}}

{{$Q = "Age"}}{{dbSet .User.ID "Q" $Q}}
{{sendMessage $ID (joinStr "" .User.Mention ", how old are you?")}}
{{$ticketnum = add $ticketnum 1}}
{{dbSet 0 "ticketnum" $ticketnum}}


{{end}}









{{$admin := (dbGet .User.ID "admin").Value}}
{{$name := (dbGet .User.ID "name").Value}}
{{$public := (dbGet .User.ID "public").Value}}
{{$names_db := (dbGet 0 "names").Value}}
{{$names := (cslice).AppendSlice $names_db}}
{{$names_new := cslice}}
{{if ($admin)}}
{{deleteMessage 658354135221141533 $admin 86400}}
{{end}}
{{if ($public)}}
{{deleteMessage 658352884701724733 $public 86400}}
{{end}}
{{if $name}}
{{sendMessage 587858995012698137 (joinStr "" "Deleting registration information for " $name "...")}}
{{range $names}}
{{if ne . $name}}
{{$names_new = $names_new.Append .}}
{{end}}{{end}}
{{dbSet 0 "names" $names_new}}
{{dbDel .User.ID "name"}}{{dbDel .User.ID "gt"}}{{dbDel .User.ID "tribe"}}{{dbDel .User.ID "pin"}}{{dbDel .User.ID "parentguardian"}}{{dbDel .User.ID "admin"}}{{dbDel .User.ID "public"}}{{dbDel .User.ID "age"}}{{dbDel .User.ID "ref"}}{{dbDel .User.ID "tribe"}}{{dbDel .User.ID "tribe2"}}{{dbDel .User.ID "bday"}}
{{else}}
{{sendMessage 587858995012698137 (joinStr "" "Deleting registration information for " .User.Username "...")}}
{{dbDel .User.ID "name"}}{{dbDel .User.ID "gt"}}{{dbDel .User.ID "tribe"}}{{dbDel .User.ID "pin"}}{{dbDel .User.ID "parentguardian"}}{{dbDel .User.ID "admin"}}{{dbDel .User.ID "public"}}{{dbDel .User.ID "age"}}{{dbDel .User.ID "ref"}}{{dbDel .User.ID "tribe"}}{{dbDel .User.ID "tribe2"}}{{dbDel .User.ID "bday"}}
{{end}}
