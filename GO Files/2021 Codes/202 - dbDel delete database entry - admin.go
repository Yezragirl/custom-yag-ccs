{{/*dbDel Delete Database, for testing*/}}
{{dbDel (toInt64 (index .CmdArgs 0)) (index .CmdArgs 1)}}
**Deleted database {{(index .CmdArgs 1)}} for {{(index .CmdArgs 0)}}.**