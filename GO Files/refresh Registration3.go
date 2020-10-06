 {{$args := parseArgs 1 "regrefresh for who"
(carg "user" "user")}}
 {{$user := ($args.Get 0)}}
 
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


 {{$name = (getMember $user.ID).Nick}}{{dbSet $user "name" $name}}

 {{if (hasRoleID 607041870966685706)}}{{/*Has Registered Role*/}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
             (sdict "name" "In Game Name" "value" $name ) 
             (sdict "name" "Gamer Tag" "value" $gt ) 
             (sdict "name" "Tribe Name" "value" $tribe ) )
             "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
         {{editMessage 658352884701724733 $public $embed}}
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "Discord Name" "value" $user.Username ) 
            (sdict "name" "In Game Name" "value" $name ) 
            (sdict "name" "Gamer Tag" "value" $gt ) 
            (sdict "name" "Referral" "value" $ref ) 
            (sdict "name" "Pin" "value" $pin ) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
         {{editMessage 658354135221141533 $admin $embed2}}


	{{else if hasRoleID 645404993863811072}}{{/*Has Kids Role*/}}
         {{$color := randInt 111111 999999 }}
         {{$embed := cembed 
            "color" $color
            "fields" (cslice 
            (sdict "name" "__***MINOR***__" "value" "_ _" ) 
			(sdict "name" "In Game Name" "value" $name )
			(sdict "name" "Tribe Name" "value" $tribe )  
            (sdict "name" "Gamer Tag" "value" $gt ) )
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{editMessage 658352884701724733 $public $embed}}
         {{$embed2 := cembed 
            "color" $color
            "fields" (cslice 
               (sdict "name" "Discord Name" "value" $user.Username ) 
               (sdict "name" "__***MINOR***__" "value" (toString $age) ) 
            (sdict "name" "In Game Name" "value" $name ) 
			(sdict "name" "Gamer Tag" "value" $gt )  
			(sdict "name" "Tribe Name" "value" $tribe ) 
            (sdict "name" "Parent/Guardian" "value" $guardian )
            (sdict "name" "Pin" "value" $pin ))
            "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
			{{editMessage 658354135221141533 $admin $embed2}}


			{{end}}
