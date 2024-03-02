{{/*deletesuggestion*/}}
{{if hasRoleName "Admin"}} You might have meant to use denysuggestion. :){{end}}
{{sleep 1}}{{$args := parseArgs 1 "" (carg "int" "message id")}}
{{$suggestionsChannel := 639554405590630400}}
{{$message := getMessage $suggestionsChannel ($args.Get 0)}}
{{if $message}}
	{{if $message.Embeds}}
		{{if gt (len (index $message.Embeds 0).Footer.Text) 37}}
			{{$suggestionAuthorID := (slice (index $message.Embeds 0).Footer.Text 39)}}
			{{if eq $suggestionAuthorID (toString .User.ID)}}
				{{ deleteMessage $suggestionsChannel $message.ID 1 }}
				Suggestion deleted.
			{{else}}
				You can only delete your own suggestions.
			{{end}}
		{{else}}
			You can only delete suggestions made after the 15th of January 2019.
		{{end}}
	{{else}}
		Please provide a valid message ID (with embed).
	{{end}}
{{else}}
	Please provide a valid message ID.
{{end}}
{{deleteTrigger 5}} {{deleteResponse 5}}