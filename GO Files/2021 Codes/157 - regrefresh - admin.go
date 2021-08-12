{{/*regrefresh used to refresh registration posts for players*/}}
{{$args := parseArgs 1 "regrefresh for who"
(carg "userid" "user")}}
{{$user := (userArg ($args.Get 0))}}
{{$Q := (dbGet $user.ID "Q").Value}}
{{$admin := (dbGet $user.ID "admin").Value}}
{{$age := (dbGet $user.ID "age").Value}}
{{$gt := (dbGet $user.ID "gt").Value}}
{{$name := (dbGet $user.ID "name").Value}}
{{$pin := (dbGet $user.ID "pin").Value}}
{{$public := (dbGet $user.ID "public").Value}}
{{$parentguardian := (dbGet $user.ID "parentguardian").Value}}
{{$ref := (dbGet $user.ID "ref").Value}}
{{$tribe := (dbGet $user.ID "tribe").Value}}
{{$nick := (getMember $user.ID).Nick}}
{{if not $nick}}
	{{if not $name}}
	{{$name = $user.Username}}{{dbSet $user.ID "name" $name}}
	{{end}}
{{else}}
{{$name = $nick}}{{dbSet $user.ID "name" $nick}}
{{end}}
{{if (targetHasRoleID $user.ID 607041870966685706)}}{{/*Has Registered Role*/}}
        {{$color := randInt 111111 999999 }}
        {{$embed := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "In Game Name" "value" (toString $name) "inline" false) 
            (sdict "name" "Gamer Tag" "value" (toString $gt) "inline" false) 
            (sdict "name" "Tribe Name" "value" (toString $tribe) "inline" false) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
		{{if $public}}
		{{editMessage 658352884701724733 $public $embed}}
Public Registration Updated.
		{{else}}
		{{$msg := sendMessageRetID 658352884701724733 $embed}}
No Public Registration to Update. New Public Registration Created.
		{{dbSet $user.ID "public" $msg}}
		{{end}}
        {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "Discord Name" "value" (toString $user.Username) "inline" false) 
            (sdict "name" "In Game Name" "value" (toString $name) "inline" false) 
            (sdict "name" "Gamer Tag" "value" (toString $gt) "inline" false) 
            (sdict "name" "Referral" "value" (toString $ref) "inline" false) 
            (sdict "name" "Pin" "value" (toString $pin) "inline" false) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{if $admin}}
			{{editMessage 658354135221141533 $admin $embed2}}
Admin Registration Updated.
			{{else}}
			{{$msg := sendMessageRetID 658354135221141533 $embed2}}
No Admin Registration to Update. New Admin Registration Created.
			{{dbSet $user.ID "admin" $msg}}
			{{end}}
{{else if targetHasRoleID $user.ID 645404993863811072}}{{/*Has Kids Role*/}}
        {{$color := randInt 111111 999999 }}
        {{$embed := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "__***MINOR***__" "value" "_ _" "inline" false) 
			(sdict "name" "In Game Name" "value" (toString $name) "inline" false)
			(sdict "name" "Tribe Name" "value" (toString $tribe) "inline" false)  
            (sdict "name" "Gamer Tag" "value" (toString $gt) "inline" false) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{if $public}}
		{{editMessage 658352884701724733 $public $embed}}
Public Registration Updated.
		{{else}}
		{{$msg := sendMessageRetID 658352884701724733 $embed}}
No Public Registration to Update. New Public Registration Created.
		{{dbSet $user.ID "public" $msg}}
		{{end}}
        {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "Discord Name" "value" $user.Username "inline" false) 
            (sdict "name" "__***MINOR***__" "value" (toString $age) "inline" false) 
            (sdict "name" "In Game Name" "value" (toString $name) "inline" false) 
			(sdict "name" "Gamer Tag" "value" (toString $gt) "inline" false)  
			(sdict "name" "Tribe Name" "value" (toString $tribe) "inline" false) 
            (sdict "name" "Parent/Guardian" "value" (toString $parentguardian) "inline" false)
            (sdict "name" "Pin" "value" (toString $pin) "inline" false))
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{if $admin}}
			{{editMessage 658354135221141533 $admin $embed2}}
Admin Registration Updated.
			{{else}}
			{{$msg := sendMessageRetID 658354135221141533 $embed2}}
No Admin Registration to Update. New Admin Registration Created.
			{{dbSet $user.ID "admin" $msg}}
			{{end}}

{{end}}
