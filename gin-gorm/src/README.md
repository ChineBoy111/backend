# src

signal.Notify 函数：将 os 信号转发 (relay) 到一个通道

```go
// sigChan 接收 os 信号的通道
// signals 指定转发的 os 信号，signals 为空时转发所有 os 信号
// 例 ctrl+c 时，转发 os.Interrupt (syscall.SIGINT) 信号
func Notify(sigChan chan<- os.Signal, signals ...os.Signal)
```

context.Background 函数：创建一个不可取消的空上下文 rootContext，是所有上下文的根

```go
var rootCtx context.Context = context.Background()
```

signal.NotifyContext 函数：收到任一 os 信号，或调用 cancelFunc 函数后，父上下文的 Done 通道关闭

```go
// parentCtx: 父上下文
// signals: 指定接收的 os 信号，signals 为空时接收所有 os 信号
// ctx: 新上下文，可取消
// cancelFunc: 调用 cancelFunc 函数取消 ctx
// 即停止接收 os 信号，关闭 ctx 的 Done 通道，释放资源
func NotifyContext(parentCtx context.Context, signals ...os.Signal) (ctx context.Context, cancelFunc context.CancelFunc)
```
