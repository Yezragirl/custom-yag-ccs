{{ $start := currentTime }}
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
 {{$ref := (dbGet $user.ID "ref").Value}} {{if not $ref}}{{$ref = "None"}}{{dbSet .User.ID "ref" $ref}}{{end}}
{{$parentguardian := (dbGet $user.ID "parentguardian").Value}}{{if not $parentguardian}}{{$parentguardian = "Not a Minor"}}{{dbSet $user.ID "parentguardian" $parentguardian}}{{end}}
{{$tribe := (dbGet $user.ID "tribe").Value}} {{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet $user.ID "tribe" $tribe}}{{end}}

			{{$embed := cembed 
				"title" "Registration Lookup"
				"description" "Here is the registration information as requested."
				"fields" (cslice 
				(sdict "name" "Discord Name" "value" (toString $user.Username) ) 
				(sdict "name" "In Game Name" "value" (toString $name) ) 
				(sdict "name" "Tribe Name" "value" (toString $tribe) ) 
				(sdict "name" "Gamer Tag" "value" (toString $gt) ) 
				(sdict "name" "Pin" "value" (toString $pin) )
				(sdict "name" "Age" "value" (toString (toInt $age)) )
				(sdict "name" "Referred By" "value" (toString $ref)))
				"thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
{{sendMessage nil $embed}}
{{ $dur := currentTime.Sub $start }}
{{ if gt $dur.Seconds 5.0 }}
    {{ sendMessage 587858995012698137 (print "CC #" .CCID " took longer than 5 seconds to run.\n**Duration:** " $dur "\n**Command executed:* `" .Message.Content "`") }}
{{ end }}