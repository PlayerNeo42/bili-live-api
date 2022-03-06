---
inject: true
to: websocket/handler_event.go
before: \/\/ handler_map\(for hygen\)
---
	dto.Event<%=h.changeCase.pascal(name)%>: <%=h.changeCase.camel(name)%>Handler,