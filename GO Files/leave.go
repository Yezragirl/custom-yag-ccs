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
{{$message := sendMessageRetID 573860427721605120 (joinStr "" $name " has left the discord. Pod their dinos for auction and clear their bases.")}}
{{addMessageReactions 573860427721605120 $message "✔️" "❌"}}
{{sendMessage 587858995012698137 (joinStr "" "Deleting registration information for " $name "...")}}
{{range $names}}
{{if eq . $name}}
{{else}}
{{$names_new = (cslice).AppendSlice .}}
{{end}}{{end}}
{{dbSet 0 "names" $names_new}}
{{dbDel .User.ID "name"}}{{dbDel .User.ID "gt"}}{{dbDel .User.ID "tribe"}}{{dbDel .User.ID "pin"}}{{dbDel .User.ID "parentguardian"}}{{dbDel .User.ID "admin"}}{{dbDel .User.ID "public"}}{{dbDel .User.ID "age"}}{{dbDel .User.ID "ref"}}{{dbDel .User.ID "tribe"}}
{{else}}
{{$message := sendMessageRetID 573860427721605120 (joinStr "" .User.Username " has left the discord. Pod their dinos for auction and clear their bases.")}}
{{addMessageReactions 573860427721605120 $message "✔️" "❌"}}
{{sendMessage 587858995012698137 (joinStr "" "Deleting registration information for " .User.Username "...")}}
{{dbDel .User.ID "name"}}{{dbDel .User.ID "gt"}}{{dbDel .User.ID "tribe"}}{{dbDel .User.ID "pin"}}{{dbDel .User.ID "parentguardian"}}{{dbDel .User.ID "admin"}}{{dbDel .User.ID "public"}}{{dbDel .User.ID "age"}}{{dbDel .User.ID "ref"}}{{dbDel .User.ID "tribe"}}
{{end}}
{{dbDel .User.ID "XP"}}
{{dbDel .User.ID "inactive"}}
{{exec "delrep" .User.ID}}
