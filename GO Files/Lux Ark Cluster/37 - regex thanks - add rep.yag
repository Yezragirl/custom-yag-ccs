{{/*regex (?i)(?:\A(?:\+rep|(?:-|<@!?204255221017214977>)\s*(?:(?:(?:t(?:ake)?r(?:ep)?|-(?:rep)?)|(?:g(?:ive)?r(?:ep)?|\+(?:rep)?)|(?:(?:(?:rep(?:uptation)?)-?)(?:cooldown|cd)s?))(?:\s+|\z))(^ty |thanks|thank you|thx|Appreciate it|Thank you|Thank You|appreciate it|THX)*/}}
{{/*
	Regex:
	(?i)(?:\A(?:\+rep|(?:-|<@!?204255221017214977>)\s*(?:(?:t(?:ake)?r(?:ep)?|-(?:rep)?)|(?:g(?:ive)?r(?:ep)?|\+(?:rep)?)|(?:(?:(?:rep(?:uptation)?)-?)(?:cooldown|cd)s?)))|(?: |\n|^)(?:thanks?\pP*|danks|ty|thx)(?: |\n|$))
*/}}

{{/* Credits - Joe_#0001 (621907351267442698)  */}}
{{ $COOLDOWN := .TimeHour }}
{{ $NOREP := 653995701008728064 }}

{{ define "getCooldowns" }}
	{{ $cooldowns := dict }}
	{{ $time := currentTime.Unix }}
	{{ with dbGet .ID "rep_cooldowns" }}
		{{ range $user, $expiry := .Value }}
			{{- if gt $expiry $time }}
				{{- $cooldowns.Set $user $expiry }}
			{{- end -}}
		{{ end }}
	{{ end }}

	{{ .Set "Result" $cooldowns }}
{{ end }}

{{ $err := "" }}
{{ $user := 0 }}
{{ $page := 0 }}
{{ $cmd := "giverep" }}
{{ $isCmd := false }}
{{ with reFindAllSubmatches `(?i)\A(?:\+rep|(?:-|<@!?204255221017214977>)\s*(?:(t(?:ake)?r(?:ep)?|-(?:rep)?)|(?:g(?:ive)?r(?:ep)?|\+(?:rep)?)|((?:(?:rep(?:uptation)?)-?)(?:cooldown|cd)s?)))(\s+(?:<@!?(\d{17,19})>|(\d{17,19})|(\d+))\z)?` .Message.Content }}
	{{ $isCmd = true }}
	{{/* First match group matches the command if and only if the command takes rep rather than gives. Second is show cooldown command. */}}
	{{ if index . 0 1 }}
		{{ $cmd = "takerep" }}
	{{ else if index . 0 2 }}
		{{ $cmd = "rep-cooldowns" }}
	{{ end }}

	{{ if eq $cmd "rep-cooldowns" }}
		{{ if index . 0 3 }}
			{{ if $id := or (index . 0 4) (index . 0 5) }} {{ $user = or (userArg $id) 0 }}
			{{ else if $p := index . 0 5 }} {{ $page = toInt $p }}
			{{ end }}
		{{ end }}
	{{ else }}
		{{ if not (index . 0 3) }} {{ $err = print "No user provided - syntax is `-" $cmd " <user>` where user can be either a mention or an ID." }}
		{{ else if $found := or (index . 0 4) (index . 0 5) | userArg }} {{ $user = $found }}
		{{ else }} {{ $err = print "Invalid user provided - syntax is `-" $cmd " <user>` where user can be either a mention or an ID." }}
		{{ end }}
	{{ end }}
{{ end }}

{{ if and (not $isCmd) .Message.Mentions }}
	{{ $user = index .Message.Mentions 0 }}
{{ end }}

{{ if and (ne $cmd "rep-cooldowns") $user }}
	{{ if eq $user.ID .User.ID }}
		{{ $err = "You can't give rep to yourself, silly." }}
	{{ else if targetHasRoleID $user $NOREP }}
		{{ $err = print "**" $user "** has no-rep on, so you can't give them rep." }}
	{{ else }}
		{{ $time := currentTime.Unix }}
		{{ $query := sdict "ID" .User.ID }}
		{{ template "getCooldowns" $query }}
		{{ $cooldowns := $query.Result }}
		{{ with $cooldowns.Get $user.ID }}
			{{ $duration := sub . $time | mult $.TimeSecond | toDuration  }}
			{{ $err = print "You can't rep **" $user "** for " (humanizeDurationSeconds $duration) ". Try coming back later?" }}
		{{ else }}
			{{ exec $cmd $user }}
			{{ $cooldowns.Set $user.ID (add $time $COOLDOWN.Seconds) }}
		{{ end }}
		{{ dbSet .User.ID "rep_cooldowns" $cooldowns }}
	{{ end }}
{{ else if eq $cmd "rep-cooldowns" }}
	{{ $time := currentTime.Unix }}
	{{ $query := sdict "ID" .User.ID }}
	{{ template "getCooldowns" $query }}
	{{ $cooldowns := $query.Result }}

	{{ dbSet .User.ID "rep_cooldowns" $cooldowns }}

	{{ if $user }}
		{{ with $cooldowns.Get $user.ID }}
			{{ $duration := sub . $time | mult $.TimeSecond | toDuration }}
			You can rep **{{ $user.String }}** again in `{{ humanizeDurationSeconds $duration }}`.
		{{ else }}
			No cooldown for **{{ $user.String }}**!
		{{ end }}
	{{ else }}
		{{ $start := mult $page 10 }}
		{{ $end := add $start 10 }}
		{{ $desc := "" }}

		{{/* Sdicts aren't actually unordered, they're ordered alphabetically, so we can use that to our favor */}}
		{{ $index := 0 }}
		{{ range $user, $remaining := $cooldowns }}
			{{ if and (ge $index $start) (lt $index $end) }}
				{{ $tag := "unknown#????" }}
				{{ with getMember $user }} {{ $tag = .User.String }} {{ end }}
				{{ $duration := sub $remaining $time | mult $.TimeSecond | toDuration }}
				{{ $desc = print $desc "\n" "`" (add $index 1) ".` **" $tag "** • " (humanizeDurationSeconds $duration) }}
			{{ end }}
			{{ $index = add $index 1 }}
		{{ end }}

		{{ if $desc }}
			{{ sendMessage nil (cembed
				"title" "Rep cooldowns"
				"description" $desc
				"footer" (sdict "text" (print "Page " (add $page 1) " of " (fdiv (len $cooldowns) 10 | roundCeil)))
				"color" 0xACF2A9
			)}}
		{{ else }}
			There are no users on cooldown on page `{{ add $page 1 }}`.
		{{ end }}
	{{ end }}
{{ end }}

{{ if $err }} {{ .User.Mention }} **::** {{ $err }} {{ end }}

{{if .Message.Mentions}}
{{else}}
{{$x := sendMessageRetID nil "Did you mean to thank someone specific? Make sure you `thanks @mention` them so they get credit for being helpful! For more information on our Rep system, try `?rep`."}}
{{ deleteMessage nil $x 10}}
{{end}}