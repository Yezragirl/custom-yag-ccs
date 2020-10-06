{{if hasRoleID 640220712560492544}}
{{deleteMessageReaction nil .Message.ID .User.ID "upvote:524907425531428864" "downvote:524907425032175638"}}
{{sendDM "So sorry, but we only allow players who have been on the server for at least 30 days to vote on suggestions. We hope you understand!"}}
{{end}}