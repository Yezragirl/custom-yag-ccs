{{if (eq .Channel.ParentID 635305499600093214)}} 
{{addRoleID 770393574734692363}}
{{$Q := (dbGet .User.ID "Q").Value}} 
{{$tribe := (dbGet .User.ID "tribe").Value}} {{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$tribe2 := (dbGet .User.ID "tribe2").Value}} {{if not $tribe2}}{{$tribe2 = "Unregistered"}}{{dbSet .User.ID "tribe2" $tribe2}}{{end}}
{{$color := randInt 111111 999999 }}
{{$rankemotes := sdict "1" "1️⃣" "2" "2️⃣" "3" "3️⃣" "4" "4️⃣" "5" "5️⃣" "6" "6️⃣" }}
{{$rank := 0}}
{{$emote := ""}}

{{$ownedtribes := cslice -}}
{{range (cslice $tribe $tribe2) -}}
	{{$tribeinfo := sdict (dbGet 0 (print "tribe_" (lower .))).Value -}}
		{{if eq $tribeinfo.owner $.User.ID -}}
			{{$ownedtribes = $ownedtribes.Append $tribeinfo.name -}}
		{{end -}}
	{{end -}}


	{{$fields := cslice}}
	{{$addEmotes := cslice}}
	{{- range $ownedtribes -}}
		{{- $rank = add $rank 1 -}}
		{{- $emote = $rankemotes.Get (toString $rank) -}}
		{{- $addEmotes = $addEmotes.Append $emote -}}
		{{- $fields = $fields.Append
		(sdict "name" (print $emote " " .) "value" "** **")}}
	{{- end}}
	{{$rank = add $rank 1}}
	{{$emote = $rankemotes.Get (toString $rank)}}
	{{$addEmotes = $addEmotes.Append $emote}}
	{{$fields = $fields.Append
	(sdict "name" (print $emote " Cancel") "value" "Cancel tribe management.")}}
	{{$embed := cembed 
	"title" "Which tribe would you like to manage?" 
	"description" "Use the reactions to indicate the tribe you'd like to manage." 
	"fields" $fields
	"color" 4645612 }}
	{{$msg := sendMessageRetID nil $embed}}
	{{- range $addEmotes -}}
	{{- addMessageReactions nil $msg . -}}
{{- end -}}

{{$Q = "TribeMan"}}
{{dbSet .User.ID "Q" $Q}}

{{- end -}}

