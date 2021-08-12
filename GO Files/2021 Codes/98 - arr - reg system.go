{{/*arr Not sure what this does*/}}
{{$args := parseArgs 1 "no who given"
(carg "user" "user")}}
{{$user := $args.Get 0}}
 {{$db := dbGet $user.ID "registration"}}
 {{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}

 {{if (targetHasRoleID $user.ID 607041870966685706)}}{{/*Has Registered Role*/}}
      {{$reg.Set "name" (getMember $user.ID).Nick}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
             (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
             (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
             (sdict "name" "Tribe Name" "value" $reg.tribe "inline" false) )
             "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
         {{editMessage 658352884701724733 $reg.public $embed}}

   
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "Discord Name" "value" $user.Username "inline" false) 
            (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
            (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
            (sdict "name" "Referral" "value" $reg.ref "inline" false) 
            (sdict "name" "Pin" "value" $reg.pin "inline" false) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
         {{editMessage 658354135221141533 $reg.admin $embed2}}
       {{dbSet $user.ID "registration" $reg}}

	{{else if targetHasRoleID $user.ID 645404993863811072}}{{/*Has Kids Role*/}}
         {{$color := randInt 111111 999999 }}
      {{$reg.Set "name" (getMember $user.ID).Nick}}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "__***MINOR***__" "value" "_ _" "inline" false) 
			(sdict "name" "In Game Name" "value" $reg.name "inline" false)
			(sdict "name" "Tribe Name" "value" $reg.tribe "inline" false)  
            (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{editMessage 658352884701724733 $reg.public $embed}}
          {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
               (sdict "name" "Discord Name" "value" $user.Username "inline" false) 
               (sdict "name" "__***MINOR***__" "value" (toString $reg.age) "inline" false) 
            (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
			(sdict "name" "Gamer Tag" "value" $reg.gt "inline" false)  
			(sdict "name" "Tribe Name" "value" $reg.tribe "inline" false) 
            (sdict "name" "Parent/Guardian" "value" $reg.guardian "inline" false)
            (sdict "name" "Pin" "value" $reg.pin "inline" false))
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{editMessage 658354135221141533 $reg.admin $embed2}}
       {{dbSet $user.ID "registration" $reg}}

			{{end}}
