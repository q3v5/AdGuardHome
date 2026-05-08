# DNS 请求设备标识自定义功能规格说明

## 变更原因

当前 AdGuard Home 在向上游 DNS 服务器发送请求时，HTTP 请求中的 User-Agent 标识是固定的或仅支持简单的 ClientID。用户希望能够像 mosdns-x 那样自定义发送 DNS 请求时附带的设备标识信息（如软件名称、版本号等），以便上游 DNS 服务能够识别请求来源。

## 需求分析

### 功能对比

| 特性 | 当前实现 | 目标实现 |
|------|----------|----------|
| HTTP User-Agent | 仅支持简单 ClientID | 支持完整的自定义 User-Agent 字符串 |
| DNS 请求标识 | 固定格式 | 支持自定义格式和内容 |
| 前端配置 | 基础文本输入 | 支持完整配置界面 |

### mosdns-x 参考实现

mosdns-x 在发送 DNS 请求时会附带类似 `MOSDNS/版本号` 的标识信息，AdGuard Home 应支持类似的自定义功能。

## 实现方案

### 1. 配置结构变更

#### 后端配置 (Go)

在 `internal/dnsforward/config.go` 的 `Config` 结构体中新增字段：

```go
// DNSRequestDevice 包含 DNS 请求时发送的设备标识信息
type DNSRequestDevice struct {
    // Enabled 启用自定义请求设备标识
    Enabled bool `yaml:"enabled"`
    
    // UserAgent 自定义的 User-Agent 字符串
    // 格式示例: "AdGuardHome/1.0" 或 "CustomClient/2.0"
    UserAgent string `yaml:"user_agent"`
}
```

更新 `Config` 结构体：

```go
// DNS请求设备标识配置
DNSRequestDevice *DNSRequestDevice `yaml:"dns_request_device"`
```

#### 前端配置 (TypeScript)

在 `client/src/reducers/dnsConfig.ts` 中添加状态：

```typescript
interface DNSConfigState {
    // ... 现有字段
    dns_request_device?: {
        enabled: boolean;
        user_agent: string;
    };
}
```

### 2. 配置验证

在 `internal/dnsforward/configvalidator.go` 中添加验证逻辑：

- `UserAgent` 最大长度限制（如 256 字符）
- `UserAgent` 格式验证（允许的字符）
- 至少在 `Enabled=true` 时 `UserAgent` 不能为空

### 3. Upstream 选项传递

在创建 upstream 连接时，将 `DNSRequestDevice` 配置转换为 `upstream.Options`：

```go
// upstream.Options 新增字段
type Options struct {
    // ... 现有字段
    // UserAgent 自定义 User-Agent
    UserAgent string
}
```

### 4. HTTP Header 设置

在 HTTP 请求构造时，使用配置的 `UserAgent`：

- 对于 DoH、DoQ 请求，设置 HTTP Header `User-Agent`
- 保持向后兼容，如果未配置则使用默认值

### 5. 前端实现

#### 表单组件 (`client/src/components/Settings/Dns/Config/Form.tsx`)

新增配置区域：

```tsx
<div className="form__group form__group--settings">
    <Checkbox
        name="dns_request_device_enabled"
        title={t('dns_request_device_enable')}
        desc={t('dns_request_device_enable_desc')}
    />
</div>

{dns_request_device_enabled && (
    <div className="form__group form__group--settings">
        <Input
            name="dns_request_device_user_agent"
            type="text"
            label={t('dns_request_device_user_agent')}
            desc={t('dns_request_device_user_agent_desc')}
            placeholder="AdGuardHome/1.0"
        />
    </div>
)}
```

#### 国际化文本 (`client/src/__locales/zh-cn.json`)

```json
{
  "dns_request_device": "DNS 请求设备标识",
  "dns_request_device_enable": "启用自定义设备标识",
  "dns_request_device_enable_desc": "向上游 DNS 服务器发送请求时，附带自定义的 User-Agent 信息",
  "dns_request_device_user_agent": "自定义 User-Agent",
  "dns_request_device_user_agent_desc": "设置向上游 DNS 服务器发送 HTTP 请求时使用的 User-Agent 标识，例如：AdGuardHome/1.0 或 CustomClient/2.0",
  "dns_request_device_user_agent_placeholder": "例如：AdGuardHome/1.0",
  "form_error_user_agent_format": "User-Agent 格式无效",
  "form_error_user_agent_too_long": "User-Agent 长度不能超过 256 个字符"
}
```

### 6. API 接口

#### 读取配置 GET /control/dns_info

响应中包含新的 `dns_request_device` 字段：

```json
{
  "dns_request_device": {
    "enabled": false,
    "user_agent": ""
  }
}
```

#### 更新配置 POST /control/dns_config

请求体支持新的配置项：

```json
{
  "dns_request_device": {
    "enabled": true,
    "user_agent": "AdGuardHome/1.0"
  }
}
```

## 影响范围

### 受影响的文件

#### 后端 (Go)

- `internal/dnsforward/config.go` - 配置结构体
- `internal/dnsforward/configvalidator.go` - 配置验证
- `internal/dnsforward/upstreams.go` - upstream 配置构建
- `internal/home/control.go` - API 接口处理
- `internal/home/dns_config.go` - DNS 配置处理（如果存在）

#### 前端 (TypeScript)

- `client/src/reducers/dnsConfig.ts` - Redux 状态管理
- `client/src/actions/dnsConfig.ts` - Redux actions
- `client/src/components/Settings/Dns/Config/Form.tsx` - 配置表单
- `client/src/components/Settings/Dns/Config/index.tsx` - 配置页面
- `client/src/__locales/zh-cn.json` - 中文翻译
- `client/src/__locales/en.json` - 英文翻译（英文为默认语言）

### 依赖项

- 无需新增外部依赖
- 使用现有的 HTTP header 设置机制

## 兼容性

- 默认 `Enabled=false`，保持向后兼容
- 现有 `ClientID` 字段保留，作为简化版本
- 如果未启用新功能，使用原有的 User-Agent 逻辑

## 测试计划

### 单元测试

1. 配置验证逻辑测试
2. User-Agent 格式化测试

### 集成测试

1. 启用/禁用功能测试
2. 自定义 User-Agent 在 DoH 请求中的验证
3. 配置文件导入/导出测试

### E2E 测试

1. 前端表单配置测试
2. 配置保存和加载测试
