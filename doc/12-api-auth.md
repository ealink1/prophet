# API 与认证系统

## API 基础配置

```javascript
BASE_URL = "https://denied-acplished-micro-drawn.trycloudflare.com/api/v1"
```

Cloudflare Tunnel 暴露的后端服务。

## 请求规范

- **方法**：所有请求使用 `POST`
- **Content-Type**：`application/json`
- **认证**：`Authorization: Bearer <token>`（需要认证的接口）
- **统一响应格式**：

```javascript
{
  code: 0,        // 0=成功, 40101=认证过期
  data: {...},    // 响应数据
  message: "..."  // 错误信息
}
```

## 认证流程

### 1. 匿名初始化

```javascript
POST /auth/anonymous/init
Body: { device_id: "web_xxxx_xxxx" }
Response: {
  token: "jwt_token",
  user: { id, nickname, lucky_code, ... }
}
```

- 首次访问自动调用
- 生成设备ID：`web_<时间戳36进制>_<随机8位>`
- 返回 JWT Token 和用户信息
- Token 存入 `localStorage.lingji_token_v2`
- 用户信息存入 `localStorage.lingji_user_v2`

### 2. 吉祥号恢复

```javascript
POST /auth/restore/by-lucky-code
Body: { lucky_code: "XXXXXX", device_id: "..." }
```

### 3. 手机号恢复

```javascript
POST /auth/restore/by-phone
Body: { phone: "13800138000", device_id: "..." }
```

### 4. 绑定手机号

```javascript
POST /auth/bind-phone
Body: { phone: "13800138000", ... }
```

### 5. 获取当前用户

```javascript
POST /auth/me
Headers: { Authorization: "Bearer <token>" }
Response: { user: { ... } }
```

## 设备标识

```javascript
device_id = localStorage.lingji_device_id
// 格式：web_<base36(timestamp)>_<random8chars>
// 示例：web_m1abc2de_f3g4h5i6
```

## Token 过期处理

- API 返回 `code: 40101` 时：
  1. 清除 `localStorage` 中的 token 和 user
  2. 触发 `lingji:auth-expired` 自定义事件
  3. 用户需重新登录/恢复

## 已认证 API 列表

| 接口 | 用途 |
|------|------|
| `/auth/me` | 获取当前用户信息 |
| `/auth/history/push` | 推送历史记录 |
| `/auth/history/list` | 获取历史记录列表 |
| `/auth/history/clear` | 清除历史记录 |
| `/message/send` | 发送消息反馈 |
| `/referral/apply` | 应用邀请码 |
| `/referral/me` | 获取邀请信息 |
| `/referral/withdraw` | 提现 |

## 未认证 API 列表

| 接口 | 用途 |
|------|------|
| `/almanac/today` | 获取今日黄历 |
| `/almanac/week` | 获取未来七日 |
| `/meditation/catalog` | 获取禅修曲目 |
| `/auth/anonymous/init` | 匿名初始化 |
| `/auth/restore/*` | 账号恢复 |

## localStorage 键值

| Key | 用途 | 格式 |
|-----|------|------|
| `lingji_token_v2` | JWT Token | string |
| `lingji_user_v2` | 用户信息 | JSON string |
| `lingji_device_id` | 设备ID | string |
| `lingji_invite_code` | 邀请码 | string |
