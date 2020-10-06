 {{$db := dbGet .User.ID "registration"}}
 {{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}

 {{if (hasRoleID 607041870966685706)}}{{/*Has Registered Role*/}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
             (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
             (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
             (sdict "name" "Tribe Name" "value" $reg.tribe "inline" false) )
             "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
         {{editMessage 658352884701724733 $reg.public $embed}}
         {{editNickname $reg.name}}
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "Discord Name" "value" .User.Username "inline" false) 
            (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
            (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
            (sdict "name" "Referral" "value" $reg.ref "inline" false) 
            (sdict "name" "Pin" "value" $reg.pin "inline" false) )
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
         {{editMessage 658354135221141533 $reg.admin $embed2}}


	{{else if hasRoleID 645404993863811072}}{{/*Has Kids Role*/}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "__***MINOR***__" "value" "_ _" "inline" false) 
			(sdict "name" "In Game Name" "value" $reg.name "inline" false)
			(sdict "name" "Tribe Name" "value" $reg.tribe "inline" false)  
            (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) )
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
			{{editMessage 658352884701724733 $reg.public $embed}}
         {{editNickname (joinStr "-" $reg.name $reg.guardian)}}
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
               (sdict "name" "Discord Name" "value" .User.Username "inline" false) 
               (sdict "name" "__***MINOR***__" "value" (toString $reg.age) "inline" false) 
            (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
			(sdict "name" "Gamer Tag" "value" $reg.gt "inline" false)  
			(sdict "name" "Tribe Name" "value" $reg.tribe "inline" false) 
            (sdict "name" "Parent/Guardian" "value" $reg.guardian "inline" false)
            (sdict "name" "Pin" "value" $reg.pin "inline" false))
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
			{{editMessage 658354135221141533 $reg.admin $embed2}}


			{{end}}
