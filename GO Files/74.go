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
{{if and $public $admin}}
   {{if (hasRoleID 607041870966685706)}}{{/*Has Registered Role*/}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
             (sdict "name" "In Game Name" "value" (toString $name) ) 
             (sdict "name" "Gamer Tag" "value" (toString $gt) ) 
             (sdict "name" "Tribe Name" "value" (toString $tribe) ) )
             "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
         {{editMessage 658352884701724733 $public $embed}}
         {{editNickname (toString $name)}}
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "Discord Name" "value" .User.Username ) 
            (sdict "name" "In Game Name" "value" (toString $name) ) 
            (sdict "name" "Gamer Tag" "value" (toString $gt) ) 
            (sdict "name" "Referral" "value" (toString $ref) ) 
            (sdict "name" "Pin" "value" (toString $pin) ) )
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
         {{editMessage 658354135221141533 $admin $embed2}}
	{{else if hasRoleID 645404993863811072}}{{/*Has Kids Role*/}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "__***MINOR***__" "value" "_ _" ) 
			(sdict "name" "In Game Name" "value" (toString $name) )
			(sdict "name" "Tribe Name" "value" (toString $tribe) )  
            (sdict "name" "Gamer Tag" "value" (toString $gt) ) )
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
			{{editMessage 658352884701724733 $public $embed}}
         {{editNickname (joinStr "-" $name $parentguardian)}}
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
               (sdict "name" "Discord Name" "value" .User.Username ) 
               (sdict "name" "__***MINOR***__" "value" (toString (toInt $age)) ) 
            (sdict "name" "In Game Name" "value" (toString $name) ) 
			(sdict "name" "Gamer Tag" "value" (toString $gt) )  
			(sdict "name" "Tribe Name" "value" (toString $tribe) ) 
            (sdict "name" "Parent/Guardian" "value" (toString $parentguardian) )
            (sdict "name" "Pin" "value" (toString $pin) ))
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
			{{editMessage 658354135221141533 $admin $embed2}}
   {{end}}
{{sendMessage 573897276569813012 (joinStr "" "<@514993305075843132> " .User.Mention " doesn't have an active registration.")}}
{{sendMessage nil (joinStr "" .User.Mention ", you don't have an active registration. Yez will have to fix it for you manually. Thanks for your patience.")}}
{{end}}