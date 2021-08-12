{{/*Register - begins registration process*/}}
{{$capture := exec "ticket open" "registration"}}

{{deleteTrigger 1}}

{{$response := sendMessageRetID nil (joinStr "" "↖️Please check the top of the channel list for your registration ticket.")}}

{{deleteMessage nil $response 300}}

