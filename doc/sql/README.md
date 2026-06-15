# 菩提苑 SQL 数据库文件

## 文件说明

| 文件 | 内容 | 行数 |
|------|------|------|
| `00-schema.sql` | 完整数据库 Schema（18 张表 + 1 视图） | ~350 行 |
| `01-queries.sql` | 常用查询语句（数据看板/订单/用户/商品/统计/维护） | ~250 行 |
| `02-seed-data.sql` | 种子数据（禅修曲目/管理员/收款码/系统配置） | ~50 行 |

## 数据库表总览

| # | 表名 | 说明 |
|---|------|------|
| 1 | `users` | 用户表（设备ID/手机号/吉祥号/福报金） |
| 2 | `products` | 商品表（定价/分类/权益） |
| 3 | `orders` | 订单表（状态流转/支付凭证/审核） |
| 4 | `blessing_lamps` | 祈福灯表（灯类型/为谁祈福/心愿） |
| 5 | `lottery_records` | 灵签记录（签号/签诗/师父开示） |
| 6 | `bazi_records` | 八字记录（生辰/排盘/五行/大运/分析） |
| 7 | `dream_records` | 解梦记录（梦境描述/解梦结果） |
| 8 | `palmistry_records` | 手相记录（图片/手纹分析） |
| 9 | `naming_records` | 起名记录（候选名/评分/选中） |
| 10 | `divination_records` | 六爻记录（卦象/解卦） |
| 11 | `meditation_tracks` | 禅修曲目（音频URL/时长） |
| 12 | `meditation_plays` | 禅修播放记录（聆听时长/功德） |
| 13 | `almanac_cache` | 黄历缓存（日期/数据JSON） |
| 14 | `activity_logs` | 活动日志（页面访问追踪） |
| 15 | `user_history` | 用户历史记录（各功能使用记录） |
| 16 | `referrals` | 邀请记录（邀请人/被邀请人/奖励） |
| 17 | `merit_transactions` | 福报金流水（获得/消费/提现） |
| 18 | `admin_users` | 管理员账号 |
| 19 | `audit_logs` | 审计日志 |
| 20 | `system_config` | 系统配置 |
| 21 | `payment_qrcodes` | 收款码 |
| — | `v_blessing_stats` | 祈福灯统计视图 |

## 技术规范

- **数据库**：SQLite（Cloudflare D1 兼容）
- **字符集**：UTF-8
- **时间格式**：ISO 8601（`datetime('now')`）
- **ID 格式**：
  - 用户：`u_<8位字符>`
  - 订单：`ord_<时间戳>_<8位hex>`
  - 其他：`<类型>_<序号>`
- **金额单位**：元（real 类型）
- **状态字段**：使用字符串枚举

## 订单状态流转

```
pending(待支付) → reviewing(待审核) → paid(已支付)
                 ↓
              rejected(已拒绝)
                 ↓
              expired(已过期)
```

## 关键索引

- `users.device_id` — 设备登录查询
- `users.lucky_code` — 邀请码查询
- `orders.user_id` — 用户订单查询
- `orders.status` — 状态筛选
- `blessing_lamps.status` — 灯墙展示
- `activity_logs.created_at` — 时间范围查询
