练习7.7：请解释为什么默认值 20.0 没写 ℃，而帮助消息中却包含 ℃。

答：因为 Celsius 实现了 fmt.Stringer 接口 `func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }`。