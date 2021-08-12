{{/* Adds the given task to the to do list for ALL the maps, individually. Command addtodo*/}}
{{$msg := .StrippedMsg}}
{{$lim := 8}}
{{if .ExecData}}
	{{$ub := $lim}}
	{{$msg = .ExecData.msg}}
	{{if lt (len .ExecData.todo) $lim}}
		{{$ub = len .ExecData.todo}}
	{{end}}
	{{$setCnt := 0}}{{$already := 0}}
	{{range seq 0 $ub}}
		{{$todo := index $.ExecData.todo .}}
		{{$map := sendMessageRetID 573860427721605120 (joinStr "" $msg " :" $todo)}}
		{{addMessageReactions 573860427721605120 $map "✔️" "❌"}}   
	{{end}}
	{{if gt (len .ExecData.todo) $lim -}}
		{{execCC .CCID nil 8 (sdict "todo" (slice .ExecData.todo $ub) "msg" $msg) -}}
	{{end -}}
{{else}}
	{{$maps := cslice "Refuge" "Excelsior" "Sphinx" "Venus" "Manticore" "Osiris" "Sanctuary" "Haven" "Central" "Beowulf" "Silvestria" "Merlin" "Hades" "Jupiter" "Phoenix" "Artemis" "Gryphon" "Titan" "Grayson" "Zelda" "Elysian" "Raiden" "Medusa" "Fenris"}}
	{{$todo := cslice}}
	{{range $maps}}
		{{$todo = $todo.Append .}}
	{{end}}
	{{if gt (len $todo) 0}}
		{{execCC .CCID nil 1 (sdict "todo" $todo "msg" $msg)}}
	{{else}}
	Done.
	{{end}}
{{end}}
