{{$capture := exec "ticket open" "Tribe Registration"}}
{{$ID := reFind `\d+` (reFind `<#\d+>` $capture)}}
{{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}
{{deleteTrigger 1}}

{{if not $db}}
{{sendMessage $ID (joinStr "" .User.Mention ", Thanks for joining Yez's Ark Cluster!\n\nJust a few questions, and you'll be registered and ready to play!\nInstructions first. **If emojis appear below the question, please use them to respond. Otherwise, type your answers.** These answers will be editable later, but be sure to check for typos before hitting enter, as your answers affect the direction these questions go.")}}
{{$reg.Set "Q" "A"}}
{{dbSet .User.ID "registration" $reg}}
{{removeRoleID 645404993863811072}}
{{removeRoleID 573210479753691170}}
{{removeRoleID 607041870966685706}}
{{giveRoleID .User.ID 606566506922377257}}
{{sendMessage $ID (joinStr "" .User.Mention ", how old are you?")}}
{{else}}
{{removeRoleID 645404993863811072}}
{{removeRoleID 573210479753691170}}
{{removeRoleID 607041870966685706}}
{{giveRoleID .User.ID 606566506922377257}}
{{sendMessage $ID (joinStr "" .User.Mention ", Thanks for joining Yez's Ark Cluster!\n\nJust a few questions, and you'll be registered and ready to play!\nInstructions first. **If emojis appear below the question, please use them to respond. Otherwise, type your answers.** These answers will be editable later, but be sure to check for typos before hitting enter, as your answers affect the direction these questions go.\n_ _")}}

{{sendMessage $ID (joinStr "" .User.Mention ", please be advised that by starting a new registration, any previous registration has been erased. You must complete THIS registration.")}}


{{if ($reg.admin)}}
{{deleteMessage 658354135221141533 $reg.admin}}
{{end}}
{{if ($reg.public)}}
{{deleteMessage 658352884701724733 $reg.public}}
{{end}}

{{$reg.Set "Q" "A"}}
{{dbSet .User.ID "registration" $reg}}
{{sendMessage $ID (joinStr "" .User.Mention ", how old are you?")}}


{{end}}