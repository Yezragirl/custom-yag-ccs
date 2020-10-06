{{if .ReactionAdded}}
    {{if and (not .User.Bot) (ne .ReactionMessage.Author.ID .User.ID) (eq .Reaction.Emoji.Name "✅") (not (dbGet .ReactionMessage.ID "limit"))}}
        {{$addPoint := dbIncr .ReactionMessage.Author.ID "reactionpoints" 1}}
        {{dbSet .ReactionMessage.ID "limit" true}}
    {{end}}
{{else}}
    {{$match := 0}}
    {{range .ReactionMessage.Reactions}}
        {{if eq .Reaction.Emoji.Name "✅"}}
            {{$match := 1}}
        {{end}}
    {{end}}
    {{if eq $match 0}}
        {{$addPoint := dbIncr .ReactionMessage.Author.ID "reactionpoints" -1}}
    {{end}}
{{end}}