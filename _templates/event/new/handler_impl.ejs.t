---
inject: true
to: websocket/handler_event.go
before: \/\/ handler_impl\(for hygen\)
---
func <%=h.changeCase.camel(name)%>Handler(payload *dto.WSPayload) {
	if DefaultEventHandlers.<%=h.changeCase.pascal(name)%> == nil {
		return
	}
	data := &dto.<%=h.changeCase.pascal(name)%>{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.<%=h.changeCase.pascal(name)%>(data)
}
