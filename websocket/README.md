# websocket部分

## 添加一个新的事件handler

### 使用hygen

```bash
hygen event new EVENT_NAME
```

### 手动

- `dto`目录下添加一个新的dto文件
- `dto/webscoket_event`中添加一个新的`event`
- `websocket/handler_registration`中添加事件`handler`类型
- `websocket/handler_registration`中的`DefaultHandlers`与`RegisterHandlers`添加事件注册
- `websocket/handler_event`中`eventPayloadHandlerMap`添加事件处理器

