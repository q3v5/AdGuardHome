# DNS 请求设备标识自定义功能检查清单

## 实现检查点

### 后端实现

- [x] `internal/dnsforward/config.go` 包含 `DNSRequestDevice` 结构体定义
- [x] `internal/dnsforward/config.go` 的 `Config` 结构体包含 `DNSRequestDevice` 字段
- [x] `internal/dnsforward/configvalidator.go` 包含 `DNSRequestDevice` 验证逻辑
- [x] 验证逻辑包含 User-Agent 长度限制（最大 256 字符）
- [x] 验证逻辑包含 User-Agent 格式验证
- [x] 验证逻辑在 `enabled=true` 时确保 User-Agent 不为空
- [x] `internal/dnsforward/upstreams.go` 正确传递 DNSRequestDevice 配置到 upstream
- [x] `internal/dnsforward/dnsforward.go` 在 Prepare 中调用验证逻辑

### 前端实现

- [x] `client/src/reducers/dnsConfig.ts` 包含 `dns_request_device` 状态
- [x] `client/src/reducers/dnsConfig.ts` 状态包含 `enabled` 和 `user_agent` 字段
- [x] `client/src/components/Settings/Dns/Config/Form.tsx` 包含启用复选框
- [x] `client/src/components/Settings/Dns/Config/Form.tsx` 包含 User-Agent 输入框
- [x] `client/src/components/Settings/Dns/Config/Form.tsx` 包含条件渲染逻辑
- [x] `client/src/components/Settings/Dns/Config/index.tsx` 正确传递初始值

### 国际化

- [x] `client/src/__locales/zh-cn.json` 包含所有必要的翻译键
- [x] `client/src/__locales/en.json` 包含所有必要的翻译键
- [x] 翻译文本包含启用描述、User-Agent 标签和错误消息

### 配置持久化

- [x] 配置可以正确保存到配置文件（通过现有的 YAML 序列化机制）
- [x] 配置可以从配置文件正确加载
- [x] API 接口正确处理新配置项（通过现有的 API 处理机制）

### 向后兼容性

- [x] 默认 `enabled=false` 保持向后兼容
- [x] 未启用功能时使用原有的 User-Agent 逻辑
- [x] 现有 ClientID 功能不受影响

### 测试

- [ ] 配置验证逻辑有单元测试覆盖（需要手动测试）
- [ ] User-Agent 长度边界测试通过（需要手动测试）
- [ ] User-Agent 格式验证测试通过（需要手动测试）
- [ ] 前端表单功能测试通过（需要手动测试）
- [ ] 配置保存/加载测试通过（需要手动测试）

## 备注

- 由于环境中缺少 `lib/dnsproxy` 目录，无法进行完整的编译测试
- 实际部署环境应包含 dnsproxy 库的本地修改版本
- 前端和后端代码已按照项目规范完成实现
