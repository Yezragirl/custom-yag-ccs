{{$msgID := .ReactionMessage.ID}}
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
{{$color := randInt 111111 999999 }}
{{$tribeedit := .ExecData.tribe}}
{{$step := .ExecData.step}}
{{$tribes_db := (dbGet 0 "tribes").Value}}
{{$tribes := (cslice).AppendSlice $tribes_db}}
{{$members_new := cslice}}
{{$tribes_new := cslice}}


{{if and (eq $tribeedit "first") (eq $step "join")}}
	{{sendMessage nil (joinStr "" .User.Mention ", type the first few letters, or words of the tribe you'd like to join.")}}
	{{$Q = "First Tribe Join"}}
	{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "first") (eq $step "leave")}}
	{{$dbkey := print "tribe_" (lower $tribe)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{$owner := $tribeinfo.owner}}
	{{if $owner}}
		{{if eq .User.ID $owner}}
			{{$members_converted := (cslice).AppendSlice (or $tribeinfo.members cslice)}}
			{{if gt (len $members_converted) 0}}
				You are the Owner of this tribe. You can not leave this tribe until you have transferred ownership to another member. 
			{{else}}
				Tribe has no members...deleting tribe {{$tribe}}.
				{{dbDel 0 (print "tribe_" (lower $tribe))}}
				{{- range $tribes -}}
					{{- if ne . $tribe -}}
						{{- $tribes_new = $tribes_new.Append . -}}
					{{- end -}}
				{{end}}
				{{dbSet 0 "tribes" $tribes_new}}
				{{$tribe = "Unregistered"}}
				{{dbSet .User.ID "tribe" $tribe}}
				{{$Q = "First Tribe Leave"}}
				{{dbSet .User.ID "Q" $Q}}
			{{end}}
		{{else}}
			{{$members_converted := (cslice).AppendSlice $tribeinfo.members}}
			{{if lt (len $members_converted) 2}}
				Tribe has no members...deleting tribe {{$tribe}}.
				{{dbDel 0 (print "tribe_" (lower $tribe))}}
				{{- range $tribes -}}
					{{- if ne . $tribe -}}
						{{- $tribes_new = $tribes_new.Append . -}}
					{{- end -}}
				{{end}}
				{{dbSet 0 "tribes" $tribes_new}}
			{{else}}
				{{- range $members_converted -}}
					{{- if ne . $.User.ID -}}
						{{- $members_new = $members_new.Append . -}}
					{{- end -}}
					{{$tribeinfo.Set "members" $members_new}}
					{{dbSet 0 $dbkey $tribeinfo}}
				{{end}}
				{{sendMessage nil (joinStr "" .User.Mention ", ok, you have left " (toString $tribe) ".")}}
				{{$tribe = "Unregistered"}}
				{{dbSet .User.ID "tribe" $tribe}}
				{{$Q = "First Tribe Leave"}}
				{{dbSet .User.ID "Q" $Q}}
			{{end}}
		{{end}}
	{{else}}
		{{$members_converted := (cslice).AppendSlice $tribeinfo.members}}
		{{if gt (len $members_converted) 1}}
			{{- range $members_converted -}}
				{{- if ne . $.User.ID -}}
					{{- $members_new = $members_new.Append . -}}
				{{- end -}}
				{{$tribeinfo.Set "members" $members_new}}
				{{dbSet 0 $dbkey $tribeinfo}}
			{{end}}
		{{else}} 
		Tribe has no members...deleting tribe {{$tribe}}.
				{{dbDel 0 (print "tribe_" (lower $tribe))}}
				{{- range $tribes -}}
					{{- if ne . $tribe -}}
						{{- $tribes_new = $tribes_new.Append . -}}
					{{- end -}}
				{{end}}
				{{dbSet 0 "tribes" $tribes_new}}
		{{end}}
		{{sendMessage nil (joinStr "" .User.Mention ", ok, you have left " (toString $tribe) ".")}}
		{{$tribe = "Unregistered"}}
		{{dbSet .User.ID "tribe" $tribe}}
		{{$Q = "First Tribe Leave"}}
		{{dbSet .User.ID "Q" $Q}}
	{{end}}	
{{else if and (eq $tribeedit "first") (eq $step "create")}}
	{{sendMessage nil (joinStr "" .User.Mention ", ok, type what you would like the Tribe to be called.")}}
	{{$Q = "First Tribe Create"}}
	{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "second") (eq $step "join")}}
	{{sendMessage nil (joinStr "" .User.Mention ", type the first few letters, or words of the tribe you'd like to join.")}}
	{{$Q = "Second Tribe Join"}}
	{{dbSet .User.ID "Q" $Q}}
{{else if and (eq $tribeedit "second") (eq $step "leave")}}
	{{$dbkey := print "tribe_" (lower $tribe2)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{$owner := $tribeinfo.owner}}
	{{if $owner}}
		{{if eq .User.ID $owner}}
			{{$members_converted := (cslice).AppendSlice (or $tribeinfo.members cslice)}}
			{{if gt (len $members_converted) 0}}
			You are the Owner of this tribe. You can not leave this tribe until you have transferred ownership to another member. 
			{{else}}
				Tribe has no members...deleting tribe {{$tribe2}}.
				{{dbDel 0 (print "tribe_" (lower $tribe2))}}
				{{- range $tribes -}}
					{{- if ne . $tribe2 -}}
						{{- $tribes_new = $tribes_new.Append . -}}
					{{- end -}}
				{{end}}
				{{dbSet 0 "tribes" $tribes_new}}	
				{{$tribe2 = "Unregistered"}}
				{{dbSet .User.ID "tribe2" $tribe2}}
				{{$Q = "Second Tribe Leave"}}
				{{dbSet .User.ID "Q" $Q}}
			{{end}}
		{{else}}
			{{$members_converted := (cslice).AppendSlice $tribeinfo.members}}
			{{if lt (len $members_converted) 2}}
				Tribe has no members...deleting tribe {{$tribe2}}.
				{{dbDel 0 (print "tribe_" (lower $tribe2))}}
				{{- range $tribes -}}
					{{- if ne . $tribe2 -}}
						{{- $tribes_new = $tribes_new.Append . -}}
					{{- end -}}
				{{end}}
				{{dbSet 0 "tribes" $tribes_new}}
			{{else}} 
				{{- range $members_converted -}}
					{{- if ne . $.User.ID -}}
						{{- $members_new = $members_new.Append . -}}
					{{- end -}}
					{{$tribeinfo.Set "members" $members_new}}
					{{dbSet 0 $dbkey $tribeinfo}}
				{{end}}
				{{sendMessage nil (joinStr "" .User.Mention ", ok, you have left " (toString $tribe2) ".")}}
				{{$tribe2 = "Unregistered"}}
				{{dbSet .User.ID "tribe2" $tribe2}}
				{{$Q = "Second Tribe Leave"}}
				{{dbSet .User.ID "Q" $Q}}
			{{end}}
		{{end}}
	{{else}}
		{{$members_converted := (cslice).AppendSlice $tribeinfo.members}}
		{{if gt (len $members_converted) 1}}
			{{- range $members_converted -}}
				{{- if ne . $.User.ID -}}
					{{- $members_new = $members_new.Append . -}}
				{{- end -}}
				{{$tribeinfo.Set "members" $members_new}}
				{{dbSet 0 $dbkey $tribeinfo}}
			{{end}}
		{{else}} 
		Tribe has no members...deleting tribe {{$tribe2}}.
			{{dbDel 0 (print "tribe_" (lower $tribe2))}}
				{{- range $tribes -}}
					{{- if ne . $tribe2 -}}
						{{- $tribes_new = $tribes_new.Append . -}}
					{{- end -}}
				{{end}}
			{{dbSet 0 "tribes" $tribes_new}}
		{{end}}
		{{sendMessage nil (joinStr "" .User.Mention ", ok, you have left " (toString $tribe2) ".")}}
		{{$tribe2 = "Unregistered"}}
		{{dbSet .User.ID "tribe2" $tribe2}}
		{{$Q = "Second Tribe Leave"}}
		{{dbSet .User.ID "Q" $Q}}
	{{end}}	
{{else if and (eq $tribeedit "second") (eq $step "create")}}
	{{sendMessage nil (joinStr "" .User.Mention ", ok, type what you would like the Tribe to be called.")}}
	{{$Q = "Second Tribe Create"}}
	{{dbSet .User.ID "Q" $Q}}
{{end}}
