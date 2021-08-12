{{/* promote Admin Promotion */}}
{{$args := parseArgs 1 "Promote Who"
	(carg "user" "User")}}
{{$target := $args.Get 0}}
{{dbSet 0 "Promote" (toString $target.ID)}}
{{sendMessage nil (joinStr "" $target.Mention " you have been approved for promotion. Type 'accept' to accept your promotion." )}}
