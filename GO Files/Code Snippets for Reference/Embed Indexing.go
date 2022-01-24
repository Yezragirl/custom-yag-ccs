{{ if .Message.Embeds }}
{{range .ReactionMessage.Reactions}}
{{.Count}} - {{.Emoji.Name}} 
{{end}}
{{ end }}




 {{ $embed := structToSdict (index .Message.Embeds 0) }}
 {{/* As structToSdict is not deep (only converts 1 layer) we need to range over */}}
 {{ range $k, $v := $embed }}
 {{- /* This checks whether the value is a struct, if so, we should convert it to an sdict. We use the `kindOf` function for this with the `indirect` bool set to true. See docs for more. */}}
 {{- if eq (kindOf $v true) "struct" }}
 {{- $embed.Set $k (structToSdict $v) }}
 {{- end -}}
 {{ end }}

 {{/* structToEmbed does not change the keys so there are some inconsistencies (causing some parts of the embed to be left out without the following code). */}}
 {{ if $embed.Author }} {{ $embed.Author.Set "Icon_URL" $embed.Author.IconURL }} {{ end }}
 {{ if $embed.Footer }} {{ $embed.Footer.Set "Icon_URL" $embed.Footer.IconURL }} {{ end }}
 {{$options := split $embed.Description "\n"}}
 {{/* You can now do whatever with $embed - for example, `$embed.Set "title" "test"` sets the embed title to test. */}}
{{ end }}