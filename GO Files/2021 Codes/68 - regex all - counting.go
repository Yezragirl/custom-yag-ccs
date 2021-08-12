{{/*regex | for counting*/}}
{{/* If you are not doing (no twice msg in a row)  or (role assignment for latest user)  you can remove counter_user and by extension everything to do with $lastUser*/}}

{{/* First time running command, set up initial values*/}}
{{$lastUser := dbGet 18 "counter_user"}}
{{if $lastUser}}
{{else}}
{{dbSet 18 "counter_user" 0}}
{{dbSet 18 "counter_count" "0"}}
{{end}}

{{/* OPTIONAL: this is just to prevent one person to type all the numbers themselves */}}
{{/* If current user ID matches the user who last successfully ran command */}}
{{if eq (toFloat $lastUser.Value) (toFloat .User.ID)}}
{{deleteTrigger 1}}
{{sendDM "You can not count twice in a row. Please give someone else a chance."}}
{{else}}


{{$next := dbGet 18 "counter_count"}}

{{/* If message is equal to the expected next number , update counter */}}
{{if eq (toInt .StrippedMsg) (toInt ($next.Value))}}
{{dbSet 18 "counter_count" (add (toInt ($next.Value)) 1)}}



{{/* OPTIONAL: If you don't want to give a role to the latest person delete everything but dbset */}}

{{dbSet 18 "counter_user" (toString .User.ID)}}
{{else}}

{{/* Message did not match expected next value */}}
{{deleteTrigger 1}}
{{giveRoleID .User.ID 584970327415848960}}
{{removeRoleID 584970327415848960 3600}}
{{end}}
{{end}}