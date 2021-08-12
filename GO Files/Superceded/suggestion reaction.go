{{/* CONFIGURATION SECTION,  CHANGE VALUES HERE */}}

{{/* Reaction needed to move to admin channel*/}}
{{$count := 10 }}
{{/* name of your upvote emote*/}}
{{$emoteName := "⭐"}}
{{/* ID of your upvote emote. Leave empty if its a default emote*/}}
{{$emotesID := ""}}
{{/* ID of your admin channel */}}
{{$adminID := 616184642302902282 }}


{{/* CODE DO NOT TOUCH BELOW */}}
{{/* Unless you know what you're doing :) */}}

{{$match := 0}}
{{/* check all reaction to see if upvote reaction is on it */}}
{{range .ReactionMessage.Reactions}}

{{/* check if emote is upvote emote */}}
{{if eq (toString .Emoji.Name) (toString $emoteName)}}

{{/* if upvote emote meets required count and bot has not reacted */}}
{{if eq .Count $count}}
{{if not .Me}}
{{$match = 1}}
{{end}}{{end}}{{end}}{{end}}

{{if eq $match 1}}

{{/* Creating the message link */}}
{{$tmpString := ""}}
{{if eq (len $emotesID ) 0}}
{{$tmpString = (printf "**%s [Starred Message](https://discordapp.com/channels/%d/%d/%d)** \n" $emoteName .ReactionMessage.GuildID .ReactionMessage.ChannelID .ReactionMessage.ID)}}
{{else}}
{{$tmpString = (printf "**%s [Starred Message](https://discordapp.com/channels/%d/%d/%d)** \n"  (joinStr "" "<:" $emoteName ":" $emotesID ">") .ReactionMessage.GuildID .ReactionMessage.ChannelID .ReactionMessage.ID)}}

{{end}}

{{$RawEmbed := sdict}}

{{/* If the message is a Embed we transpose the embed into a new one with Message Link*/}}

{{if .ReactionMessage.Embeds}}
{{$embed := (index .ReactionMessage.Embeds 0)}}

{{if $embed.Title}}
{{$RawEmbed.Set "title" ($embed.Title)}}
{{end}}

{{if $embed.Description}}
{{$RawEmbed.Set "description" (joinStr "" $tmpString ($embed.Description))}}
{{else}}
{{$RawEmbed.Set "description" ($tmpString)}}
{{end}}

{{if $embed.Image}}
{{$RawEmbed.Set "image" ($embed.Image)}}
{{end}}

{{if $embed.Thumbnail}}
{{$RawEmbed.Set "thumbnail" ($embed.Thumbnail)}}
{{end}}

{{if $embed.Fields}}
{{$RawEmbed.Set "fields" ($embed.Fields)}}
{{end}}

{{if $embed.URL}}
{{$RawEmbed.Set "url" ($embed.URL)}}
{{end}}


{{if $embed.Author}}
{{$RawEmbed.Set "author" (sdict "name" (toString $embed.Author.Name) "url" (toString ($embed.Author.URL) ) "icon_url" (toString ($embed.Author.IconURL))  )}}
{{else}}
{{$RawEmbed.Set "author" (sdict "name" "")}}
{{end}}

{{if $embed.Footer}}
{{$RawEmbed.Set "footer" (sdict "text" (toString $embed.Footer.Text) "icon_url" (toString ( $embed.Footer.IconURL) ) ) }}
{{else}}
{{$RawEmbed.Set "footer" (sdict "text" "")}}
{{end}}

{{/* Handling Embeds with Message Content */}}
{{if .ReactionMessage.Content}}
{{$temp:= $RawEmbed.Get "description"}}
{{$RawEmbed.Set "description" (joinStr "" $tmpString $temp "\n**Message Content**\n" .ReactionMessage.Content )}}
{{end}}


{{else}}
{{/* If message has an attachment image , the image is copied to the image field in Rawembed */}}
{{ if .ReactionMessage.Attachments}}
{{$Attachment := index .ReactionMessage.Attachments 0}}
{{$file := $Attachment.Filename}}

{{/* Making sure the Attachment is a image and not a audio file or some other filetype */}}
{{if (reFind "(.jpg|.jpeg|.png|.gif|.tif|.tiff)$" $file)  }}
{{$RawEmbed.Set "image" (sdict "url" ($Attachment.ProxyURL) ) }}
{{end}}
{{end}}

{{/* If message has textual content , the content is copied into the description of the RawEmbed */}}
{{if .ReactionMessage.Content}}
{{$RawEmbed.Set "description"  (joinStr "" $tmpString .ReactionMessage.Content)}}
{{else}}
{{$RawEmbed.Set "description" $tmpString}}
{{end}}

{{/* Create and Initialize EmbedRaw Author (author ->name) and footer(footer -> text) fields for later use */}} 
{{$RawEmbed.Set "author" (sdict "name" "")}}
{{$RawEmbed.Set "footer" (sdict "text" "")}}
{{end}}



{{/* Sending the embed to the appropriate channel and react to show completion */}}

{{sendMessage (toString $adminID) (cembed $RawEmbed) }}

{{/* If its a default emote, there is no ID */}}
{{if eq (len $emotesID) 0}}
{{ addReactions $emoteName}}
{{else}}
{{$fullEmote := (joinStr "" $emoteName ":" $emotesID)}}
{{ addReactions $fullEmote}}
{{end}}

{{end}}