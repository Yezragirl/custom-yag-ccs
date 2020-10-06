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
{{$tribe := (dbGet .User.ID "tribe").Value}}
{{$names_db := (dbGet 0 "names").Value}}
{{$names := (cslice).AppendSlice $names_db}}
  {{if not $name}}
You can't edit your registration. You don't have one yet. 
{{else}}
      {{if eq (toString $Q) "Edit GT"}}
         {{$gt = (toString .Message.Content)}}{{dbSet .User.ID "gt" $gt}}
         {{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
		 {{else if eq (toString $Q) "Edit Name"}}
		 {{if gt (len (toString .Message.Content)) 25}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's too long for a name. I'll ask again, What would you like to use as your In Game Name?")}}
		 {{else if (in $names (toString .Message.Content))}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that name is already taken. Since names need to be unique, I'll ask again, what would you like to use as your In Game Name?")}}
		 {{else}}
		 {{$name = (toString .Message.Content)}}
		 {{dbSet .User.ID "name" $name}}
		 {{dbSet 0 "names" ($names.AppendSlice (cslice $name))}}
		 {{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
		 {{end}}
		 {{else if eq (toString $Q) "Edit Tribe"}}
	     {{$name = (toString .Message.Content)}}{{dbSet .User.ID "tribe" $name}}
         {{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
         {{else if eq (toString $Q) "Edit Pin"}}
         {{$pin = (toString .Message.Content)}}{{dbSet .User.ID "pin" $pin}}
         {{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
      {{else if eq (toString $Q) "Edit ParentGuardian"}}
	  {{$parentguardian = (toString .Message.Content)}}{{dbSet .User.ID "parentguardian" $parentguardian}}
	  {{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
	 {{else if eq (toString $Q) "Edit Age"}}
		 {{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", did you have a birthday?")}}
		 {{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
		 {{end}}
	{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", got it. Do you want to edit something else?")}}
	{{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}} 
  {{end}}
{{end}}