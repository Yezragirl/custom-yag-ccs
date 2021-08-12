{{/*%apply for applying for admin*/}}
{{giveRoleID .User.ID 641285945584517131}}
{{sendDM "Please fill out the application at this address: https://docs.google.com/forms/d/e/1FAIpQLSenTRt36e8qNRu54zC6SZb7L-x7KPHsugoCWz4ajVao_3F_bQ/viewform"}}
{{deleteTrigger 1}}
{{sendMessage 573897276569813012 (joinStr "" "__***FYI, " .User.Mention " is applying for Admin. Expect their application shortly.***__")}}