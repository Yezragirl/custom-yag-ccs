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
{{editMessage 660329324976799754 765226303569264680 (joinStr "" "There are currently " $ticketnum " open tickets.")}}

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
{{editMessage 660329324976799754 765226303569264680 (joinStr "" "There are currently " $ticketnum " open tickets.")}}


{{end}}