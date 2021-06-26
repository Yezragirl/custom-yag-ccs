{{if (eq .Channel.ParentID 635305499600093214)}}
{{$Q := (dbGet .User.ID "Q").Value}} 
{{$admin := (dbGet .User.ID "admin").Value}}
{{$age := (dbGet .User.ID "age").Value}}
{{$gt := (dbGet .User.ID "gt").Value}}
{{$name := (dbGet .User.ID "name").Value}}
{{$pin := (dbGet .User.ID "pin").Value}}
{{$public := (dbGet .User.ID "public").Value}}
{{$ref := (dbGet .User.ID "ref").Value}}
{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
{{$tribe := (dbGet .User.ID "tribe").Value}}{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$tribe2 := (dbGet .User.ID "tribe2").Value}}{{if not $tribe2}}{{$tribe2 = "Unregistered"}}{{dbSet .User.ID "tribe2" $tribe2}}{{end}}
{{$MsgID := .Reaction.MessageID}}
{{$tribeselect := ""}}


{{range (index .Message.Embeds 0).Fields -}}
	{{$parts := reFindAllSubmatches `([^ ]+) (.+)` .Name -}}
	{{$emojipart := index $parts 0 1 -}}
	{{if (eq $.Reaction.Emoji.Name $emojipart)}}
		{{$tribeselect = index $parts 0 2}}
	{{end}}
{{end}}

{{$tribeBitsLookup := sdict
	"Edit First Tribe Join"
		(sdict "ordinal" "first" "key" "tribe")
	"Edit Second Tribe Join"
		(sdict "ordinal" "second" "key" "tribe2")
	}}
	
	{{if $tribeBits := $tribeBitsLookup.Get (toString $Q)}}
		{{if eq $tribeselect "Try Again"}}
			{{execCC 222 nil 0 (sdict "tribe" $tribeBits.ordinal "step" "join")}}
		{{else}}
			{{$dbkey := print "tribe_" (lower $tribeselect)}}
			{{dbSet .User.ID $tribeBits.key $tribeselect}}
			{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
			{{$tribeinfo.Set "members" (((cslice).AppendSlice (or $tribeinfo.members cslice)).Append .User.ID)}}
			{{dbSet 0 $dbkey $tribeinfo}}
			Added Tribe {{$tribeselect}} to your registration.
			Added {{.User.Mention}} to {{$tribeselect}} member roster.
		{{end}}
	{{end}}
{{end}}