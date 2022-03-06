---
inject: true
to: websocket/handler_registration.go
before: \/\/ handler_type above\(for hygen\)
---
type <%=h.changeCase.pascal(name)%>Handler func(<%=h.changeCase.camel(name)%> *dto.<%=h.changeCase.pascal(name)%>)
