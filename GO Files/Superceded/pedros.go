{{$gEmoji:=`🎟`}}

{{/*Code*/}}
{{$data:=sdict}}
{{$embed := sdict}}

{{$compareEmoji:=.Reaction.Emoji.Name}}
{{with reFindAllSubmatches `(\d+)>\z` $gEmoji}} {{$gEmoji =index . 0 1}} {{$compareEmoji =str $.Reaction.Emoji.ID}} {{end}}

{{if and (eq $compareEmoji $gEmoji) (not .User.Bot)}}
	{{with (dbGet 7777 "giveaway_active").Value}} {{$data =sdict .}} {{end}}
	{{$giveawayData := $data.Get (joinStr "" .Reaction.ChannelID .Reaction.MessageID)}}
	{{if $giveawayData}}
		{{$giveawayData =sdict $giveawayData}}
		{{$originalEmbed := index .ReactionMessage.Embeds 0}}
		{{$embed.Set "title" $originalEmbed.Title}}
		{{$embed.Set "description" $originalEmbed.Description}}
		{{$embed.Set "color" $originalEmbed.Color}}
		{{$embed.Set "footer" $originalEmbed.Footer}}
		{{$embed.Set "timestamp" $originalEmbed.Timestamp}}
		{{$part := "```"}}
		{{if .ReactionAdded}}
			{{$giveawayData.Set "listID" (print $giveawayData.listID  .User.ID "," )}}
			{{$giveawayData.Set "count" (add $giveawayData.count 1)}}
			{{$users := split (print $giveawayData.listID) ","}}
			{{range $k, $v := $users}}
				{{- if ne $k (sub (len $users) 1) -}}
					{{- if gt (len $users) 11 -}}
						{{- if  ge (sub $k (sub (len $users) 11)) 0 -}}
							{{- with getMember . -}}
								{{- $part = print $part "\n" .User.String -}}
							{{- end -}}
						{{- end -}}
					{{- else -}}
						{{- with getMember . -}}
							{{- $part = print $part "\n" .User.String -}}
						{{- end -}}
					{{- end -}}
				{{- end -}}
			{{end}}
			{{$part = print $part "```"}}
			{{$embed.Set "fields" (cslice (sdict "Name" (print "Participants (" $giveawayData.count ")") "Value" $part))}}
			{{editMessage nil .ReactionMessage.ID (cembed $embed)}}
		{{else}}
			{{$IDregex:= print .User.ID `,`}}
			{{$giveawayData.Set "listID" (reReplace $IDregex $giveawayData.listID "")}}
			{{$giveawayData.Set "count" (add $giveawayData.count -1)}}
			{{$users := split (print $giveawayData.listID) ","}}
			{{if gt (len $users) 1}}
				{{range $k, $v := $users}}
					{{- if ne $k (sub (len $users) 1) -}}
						{{- if gt (len $users) 11 -}}
							{{- if  ge (sub $k (sub (len $users) 11)) 0 -}}
								{{- with getMember . -}}
									{{- $part = print $part "\n" .User.String -}}
								{{- end -}}
							{{- end -}}
						{{- else -}}
							{{- with getMember . -}}
								{{- $part = print $part "\n" .User.String -}}
							{{- end -}}
						{{- end -}}
					{{- end -}}
				{{end}}
				{{$part = print $part "```"}}
			{{else}}
				{{$part = "```\nNone```\n"}}
			{{end}}
			{{$embed.Set "fields" (cslice (sdict "Name" (print "Participants (" $giveawayData.count ")") "Value" $part))}}
			{{editMessage nil .ReactionMessage.ID (cembed $embed)}}
		{{end}}
		{{$data.Set (joinStr "" .Reaction.ChannelID .Reaction.MessageID) $giveawayData}}{{dbSet 7777 "giveaway_active" $data}}
	{{end}}
{{end}}