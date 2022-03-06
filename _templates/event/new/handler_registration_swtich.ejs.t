---
inject: true
to: websocket/handler_registration.go
before: default
---
		case <%=h.changeCase.pascal(name)%>Handler:
			DefaultEventHandlers.<%=h.changeCase.pascal(name)%> = handler