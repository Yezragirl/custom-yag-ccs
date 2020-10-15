{{/* RegEx, Trigger: `\A-(add|sugg(est)?).?topic` */}}
{{$tsc := 765811442237833216}}
{{$msg := sendMessageRetID $tsc .StrippedMsg}}{{/* Send suggested topic to the logs */}}
{{dbSet $msg "topicauthor" (toString .User.ID)}}
{{addMessageReactions $tsc $msg "✅" "☑️" "❌"}}{{/* Add admin reactions */}}


{{/* RegEx, Trigger: `\A-topic` */}}
{{$slice := (dbGet 0 "topics").Value}}
{{$topics := (cslice).AppendSlice $slice}}
{{sendMessage nil (joinStr "" (index (shuffle $topics) 0))}}


{{/* RegEx, Trigger: `\A-pandr.?topic` */}}
{{$slice := (dbGet 0 "pandrtopics").Value}}
{{$pandrtopics := (cslice).AppendSlice $slice}}
{{sendMessage nil (joinStr "" (index (shuffle $pandrtopics) 0))}}



{{/* Reaction, Trigger: Added Only */}}

{{$name := .User.Username}}
{{if .Member.Nick}}
    {{$name = .Member.Nick}}
{{end}}
 
{{$author := (userArg (dbGet .Message.ID "topicauthor").Value)}}
{{$authorname := $author.Username}}
{{if (getMember $author.ID).Nick}}
    {{$authorname = (getMember $author.ID).Nick}}
{{end}}
 
{{$col := 16777215}}{{$p := 0}}{{$r := (getMember $author).Roles}}{{range .Guild.Roles}}{{if and (in $r .ID) (.Color) (lt $p .Position)}}{{$p = .Position}}{{$col = .Color}}{{end}}{{end}}
{{ $status := "nil" }}
{{ $sugg := .Message.Content }}


{{ $slice := (dbGet 0 "topics").Value }}
{{ $topics := (cslice).AppendSlice $slice }}

{{ $slice2 := (dbGet 0 "pandrtopics").Value }}
{{ $pandrtopics := (cslice).AppendSlice $slice2 }}

{{ if eq .Reaction.Emoji.Name "✅" }}
    {{ $topics = $topics.Append $sugg }}
    {{ dbSet 0 "topics" $topics }}
    {{ deleteTrigger 0 }}
    {{ $status = "Added" }}
{{ else if eq .Reaction.Emoji.Name "☑️"}}
    {{ $pandrtopics = $pandrtopics.Append $sugg }}
    {{ dbSet 0 "pandrtopics" $pandrtopics }}
    {{ deleteTrigger 0 }}
    {{ $status = "Added (PandR)"}}
{{ else if eq .Reaction.Emoji.Name "❌" }}
    {{ deleteTrigger 0 }}
    {{ $status = "Denied" }}
{{ end }}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString $author.ID) "/" $author.Avatar ".png") }}
{{ $embed := cembed
    "author" (sdict "name" (joinStr "" "Topic Suggestion by " $authorname " — " $status) "url" "https://cdn.discordapp.com/attachments/710036018908102687/757033106250137680/Untitled.png" "icon_url" $avatar ) 
    "description" (joinStr "" "\"" $sugg "\"")
    "color" $col
    "footer" (sdict "text" (joinStr "" "Marked as " $status " by " $name " on "))
    "timestamp" currentTime  
}}
{{ sendMessage 765793026987393046 $embed }}
{{dbDel .Message.ID "topicauthor"}}