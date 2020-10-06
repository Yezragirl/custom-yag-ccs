{{$color := 0}}


{{$args := parseArgs 1 ""
  (carg "string" "suggestion")}}

{{if (dbGet .User.ID "suggestCld")}}This command is on cooldown for {{humanizeDurationSeconds ((dbGet .User.ID "suggestCld").ExpiresAt.Sub currentTime)}} to avoid spam.{{else}}{{dbSetExpire .User.ID "suggestCld" "cooldown" 1200}}

{{ $embed := cembed
"description" (joinStr " " .CmdArgs)
"color" $color
"author" (sdict "name" (joinStr "" .User.Username "#" .User.Discriminator) "url" "" "icon_url" (.User.AvatarURL "512"))
"timestamp"  currentTime
"footer" (sdict "text" (joinStr "" "Submit your suggestion with -suggest - " .User.ID))
}}
{{ $id := (sendMessageRetID 745464594629132310 $embed) }}
{{ addMessageReactions 745464594629132310 $id "✅" "⛔" }}
{{ sendDM (joinStr "" "Suggestion submitted successfully. If you want to discuss this or other suggestions, use the <#745464630843015280> channel.") }}
{{addReactions "✅" }} {{deleteTrigger 20}}{{end}}
{{ $embed2 := cembed
"title" "New Suggestion Posted" 
"description" "A member of the community has posted a suggestion! Please go to <#745464594629132310> to vote on it! Voting on suggestions are great ways to help us improve the cluster and have your voice heard! Suggestions are reviewed after 1 week!"
"color" $color
"timestamp"  currentTime
"footer" (sdict "text" (joinStr "" "Submit your suggestion with -suggest Your-Idea-Here" ))
}}
{{sendMessage 744719451408302252 $embed2}}