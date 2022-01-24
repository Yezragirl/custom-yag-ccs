{{/*regex ^-(suggest|suggestion)\b*/}}
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
{{ $id := (sendMessageRetID 639554405590630400 $embed) }}
{{ addMessageReactions 639554405590630400 $id "upvote:524907425531428864" "downvote:524907425032175638" }}
{{ sendDM (joinStr "" "Suggestion submitted successfully. If you want to discuss this or other suggestions, use the <#588473094843400212> channel.") }}
{{addReactions "upvote:524907425531428864" }} {{deleteTrigger 20}}{{end}}
{{ $embed2 := cembed
"title" "New Suggestion Posted" 
"description" (joinStr "" "A member of the community has posted a suggestion! Please go to <#639554405590630400> to vote on it! Voting on suggestions, and participating in <#573245725228531712> are great ways to help us improve the cluster and have your voice heard! Suggestions are reviewed after 1 week!"  )
"color" $color
"thumbnail" (sdict "url" ("https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f0e6a60945624d1ea9b9cb9_Yez%20Tiny.jpg"))
"timestamp"  currentTime
"footer" (sdict "text" (joinStr "" "Submit your suggestion with -suggest Your-Idea-Here" ))
}}
{{sendMessage 573245421066125324 $embed2}}
{{execCC 167 573897276569813012 0 ""}}