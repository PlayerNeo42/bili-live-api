---
inject: true
to: websocket/handler_registration.go
before: \/\/ handler_struct above\(for hygen\)
---
	<%=h.changeCase.pascal(name)%> <%=h.changeCase.pascal(name)%>Handler