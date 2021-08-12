{{$message :=  (getMessage "356486960417734666" (index  .CmdArgs 1))}}
{{$args := (joinStr " " .CmdArgs)}}

{{if (reFind "(?i)^(edit|dupe|markdupe)" $args)}}
{{$edited := cembed "author" (sdict "name"  (index $message.Embeds 0).Author.Name "icon_url" ((userArg (slice (index $message.Embeds 0).Footer.Text 39)).AvatarURL "512")) "footer" (sdict "text" (index $message.Embeds 0).Footer.Text) "description" (joinStr "" (index $message.Embeds 0).Description "\n\n**This message has been marked as a dupe of:\n**https://discordapp.com/channels/166207328570441728/356486960417734666/" (slice .CmdArgs 2))}}
{{editMessage 356486960417734666 $message.ID $edited}}
Doneso

{{else if (reFind "(?i)^deny" $args)}}
{{deleteMessage "356486960417734666" $message.ID}}{{sleep 3}}{{sendMessage 525535451839463424 (joinStr "" ((userArg (slice (index $message.Embeds 0).Footer.Text 39)).Mention) " | The suggestion below has been deleted for reason: `" (joinStr " " (slice (.CmdArgs) 2)) "`")}}{{sleep 1}}{{sendMessage 525535451839463424 (index $message.Embeds 0)}}{{sleep 3}}
doneso

{{else if (reFind "(?i)^(implement|implemented|archive)" $args)}}
{{deleteTrigger 1}}{{deleteResponse 3}}
{{$suggestionsChannel := 356486960417734666}}
{{$archiveChannel := 532918109036740608}}
{{if (index .CmdArgs 1)}}{{$message := getMessage $suggestionsChannel (index .CmdArgs 1)}}
{{$comment := ""}}
{{if gt (len .CmdArgs) 2}}
{{$comment =  joinStr "" "\n\n**__Comment__**```" (joinStr " " (slice .CmdArgs 2)) "```" }}
{{end}}
{{if $message}}
{{deleteMessage $suggestionsChannel (index .CmdArgs 1)}}{{sleep 3}}
{{$embed := cembed "title" "Successfully Implemented Suggestion:" "description" (joinStr "" (index $message.Embeds 0).Description $comment) "author" (sdict "name" (index $message.Embeds 0).Author.Name) "footer" (sdict "text" (joinStr "" "Implemented by: " .User.Username " - " .User.ID))
"timestamp" $message.Timestamp}}
{{sendMessage $archiveChannel $embed}}
doneso
{{else}}No message ID provided!{{end}}
{{else}}Unknown message{{end}}

{{end}}