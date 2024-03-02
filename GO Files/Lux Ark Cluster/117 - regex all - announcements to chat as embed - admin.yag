{{/* Regex .* Announcements to General Chat as embed */}}

{{$msg := .Message}}
{{if not $msg}}

{{else}}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString $msg.Author.ID) "/" $msg.Author ".png") }}

{{$embedRaw := sdict 
"description" (joinStr "" "**<#" $msg.ChannelID ">**\n" $msg.Content)
"color" 4645612 
"author" (sdict "name"  $msg.Author.Username "icon_url" ($msg.Author.AvatarURL "64")) 
"footer" (sdict "text" (joinStr "" "Req. by "  .Message.Author.Username ". Quote from ")) 
"timestamp" $msg.Timestamp  }}

{{if $msg.Attachments}}
{{$embedRaw.Set "image" (sdict "url" (index $msg.Attachments 0).URL) }}
{{end}}

{{ sendMessage 573245421066125324 (cembed $embedRaw) }}
{{ sendMessage 646033978721173523 (cembed $embedRaw) }}

{{end}}