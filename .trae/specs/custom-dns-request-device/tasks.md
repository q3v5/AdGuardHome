# DNS 请求设备标识自定义功能实现任务

## 任务概述

实现 AdGuard Home 发送 DNS 请求时自定义请求设备标识的功能，允许用户配置自定义的 User-Agent 字符串，类似于 mosdns-x 的实现方式。

## 任务列表

### 阶段 1：后端配置结构实现

- [ ] **1.1 修改 dnsforward 配置结构体**
  - 在 `internal/dnsforward/config.go` 中添加 `DNSRequestDevice` 结构体
  - 在 `Config` 结构体中添加 `DNSRequestDevice` 字段
  - 添加 YAML 标签

- [ ] **1.2 实现配置验证逻辑**
  - 在 `internal/dnsforward/configvalidator.go` 中添加 `DNSRequestDevice` 验证函数
  - 验证 User-Agent 长度（最大 256 字符）
  - 验证 User-Agent 格式（允许的字符）
  - 当 `enabled=true` 时确保 User-Agent 不为空

### 阶段 2：Upstream 配置集成

- [ ] **2.1 传递配置到 upstream.Options**
  - 修改 `internal/dnsforward/upstreams.go` 中的 `newUpstreamConfig` 函数
  - 将 `DNSRequestDevice` 配置转换为 upstream.Options 中的字段

- [ ] **2.2 检查 upstream.Options 是否需要新增字段**
  - 检查 dnsproxy 库的 upstream.Options 是否已支持 UserAgent
  - 如需要，创建本地的配置传递方式

### 阶段 3：前端实现

- [ ] **3.1 更新 Redux 状态管理**
  - 在 `client/src/reducers/dnsConfig.ts` 中添加 `dns_request_device` 状态
  - 包含 `enabled` 和 `user_agent` 字段

- [ ] **3.2 更新 Redux Actions**
  - 在 `client/src/actions/dnsConfig.ts` 中确保正确处理新字段

- [ ] **3.3 创建配置表单组件**
  - 在 `client/src/components/Settings/Dns/Config/Form.tsx` 中添加配置区域
  - 添加启用复选框
  - 添加 User-Agent 文本输入框
  - 添加条件渲染逻辑

- [ ] **3.4 更新配置页面组件**
  - 在 `client/src/components/Settings/Dns/Config/index.tsx` 中传递初始值

### 阶段 4：国际化支持

- [ ] **4.1 添加中文翻译**
  - 在 `client/src/__locales/zh-cn.json` 中添加所有必要的翻译键

- [ ] **4.2 添加英文翻译**
  - 在 `client/src/__locales/en.json` 中添加所有必要的翻译键

### 阶段 5：测试与验证

- [ ] **5.1 配置验证测试**
  - 编写单元测试验证配置验证逻辑
  - 测试边界条件（空字符串、超长字符串、特殊字符）

- [ ] **5.2 E2E 测试**
  - 测试前端表单启用/禁用功能
  - 测试自定义 User-Agent 保存和加载

## 任务依赖关系

```
1.1 → 1.2 → 2.1 → 2.2
                  ↓
3.1 → 3.2 → 3.3 → 3.4
                  ↓
4.1 ← → 4.2
         ↓
5.1 ← → 5.2
```

## 注意事项

- 确保向后兼容性，默认禁用新功能
- User-Agent 格式应该符合 HTTP 规范
- 配置验证失败时应该给出清晰的错误信息
- 前端表单应该有合理的默认值和占位符提示
