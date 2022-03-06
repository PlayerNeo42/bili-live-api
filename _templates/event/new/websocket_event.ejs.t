---
inject: true
to: dto/websocket_event.go
before: \)
skip_if: Event<%=h.changeCase.pascal(name)%>
---
	Event<%=h.changeCase.pascal(name)%> EventType = "<%=name%>"