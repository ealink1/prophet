-- ============================================
-- 菩提苑 - 数据库 Schema
-- 基于前台+后台截图分析生成
-- ============================================

-- 使用 SQLite 语法（Cloudflare D1 兼容）

-- ============================================
-- 1. 用户与认证
-- ============================================

CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,                           -- 用户ID，如 u_41068f92dd29
    device_id TEXT UNIQUE NOT NULL,                -- 设备ID，如 web_m1abc2de_f3g4h5i6
    phone TEXT,                                    -- 手机号（可选绑定）
    nickname TEXT DEFAULT '有缘人',                 -- 昵称
    lucky_code TEXT UNIQUE,                        -- 吉祥号/邀请码，如 ABC123
    avatar TEXT,                                   -- 头像URL
    invite_code TEXT,                              -- 被邀请时使用的邀请码
    referred_by TEXT,                              -- 邀请人用户ID
    merit_balance INTEGER DEFAULT 0,               -- 福报金余额（分）
    total_merit_earned INTEGER DEFAULT 0,          -- 累计获得福报金
    total_merit_withdrawn INTEGER DEFAULT 0,       -- 累计提现福报金
    free_lottery_daily INTEGER DEFAULT 3,          -- 每日免费灵签次数
    free_dream_daily INTEGER DEFAULT 5,            -- 每日免费解梦次数
    free_bazi_daily INTEGER DEFAULT 1,             -- 每日免费八字次数
    last_active_at TEXT,                           -- 最后活跃时间
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_users_device_id ON users(device_id);
CREATE INDEX idx_users_lucky_code ON users(lucky_code);
CREATE INDEX idx_users_phone ON users(phone);

-- ============================================
-- 2. 商品定价
-- ============================================

CREATE TABLE IF NOT EXISTS products (
    product_id TEXT PRIMARY KEY,                   -- 商品ID，如 blessing_lamp, extra_lottery
    name TEXT NOT NULL,                            -- 商品名称，如 "祈福供灯·平安灯"
    category TEXT NOT NULL,                        -- 分类：blessing/lottery/dream/bazi/palmistry/naming/divination
    description TEXT,                              -- 商品描述
    original_price REAL NOT NULL,                  -- 原价（元）
    price REAL NOT NULL,                           -- 实际售价（元）
    badge TEXT,                                    -- 标签，如 "首单特惠"/"明星产品"
    benefits TEXT,                                 -- 权益描述JSON，如 ["12月运势","师父开示"]
    is_active INTEGER DEFAULT 1,                   -- 是否上架
    sort_order INTEGER DEFAULT 0,                  -- 排序
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now'))
);

-- 初始商品数据
INSERT INTO products (product_id, name, category, original_price, price, badge, benefits) VALUES
-- 祈福灯类
('blessing_lamp', '祈福供灯', 'blessing', 6.60, 6.60, NULL, '["点亮功德灯","供奉指定时长","功德灯墙展示"]'),
('blessing平安灯', '祈福供灯·平安灯', 'blessing', 3.90, 3.72, NULL, '["平安灯","供奉指定时长"]'),
('blessing智慧灯', '祈福供灯·智慧灯', 'blessing', 9.90, 9.80, NULL, '["智慧灯","供奉指定时长"]'),
('blessing姻缘灯', '祈福供灯·姻缘灯', 'blessing', 5.90, 5.71, NULL, '["姻缘灯","供奉指定时长"]'),
('blessing财福灯', '祈福供灯·财福灯', 'blessing', 3.90, 3.82, NULL, '["财福灯","供奉指定时长"]'),
-- 灵签类
('extra_lottery', '关帝灵签·加抽一次', 'lottery', 2.90, 2.86, NULL, '["加抽一次灵签","师父开示"]'),
-- 解梦类
('extra_dream', '周公解梦·加抽一次', 'dream', 3.90, 3.78, NULL, '["加抽一次解梦"]'),
-- 八字类
('unlock_bazi', '八字精批·完整解读', 'bazi', 3.90, 3.71, NULL, '["完整十神/大运/流年","古籍引用"]'),
('single_liunian', '流年运势详批', 'bazi', 9.90, 9.90, '首单特惠', '["12月逐月运势","贵人/桃花/财禄提示","师父开示"]'),
('single_bazi_deep', '八字精批深度版', 'bazi', 19.90, 19.90, '明星产品', '["完整十神/大运/流年","古籍引用","PDF报告"]'),
-- 手相类
('unlock_palmistry', '手相图解·完整解读', 'palmistry', 6.60, 6.45, NULL, '["拍照上传","手纹细看","手纹命理详解"]'),
-- 起名类
('unlock_naming', '宝宝起名·解锁全部30名', 'naming', 66.00, 65.97, NULL, '["30个候选名","音韵/笔画/五行评分","典故出处"]'),
('single_naming_pro', '宝宝起名VIP', 'naming', 49.90, 49.90, '热销', '["30个候选名","音韵/笔画/五行评分","典故出处"]'),
('single_company', '公司起名', 'naming', 99.90, 99.90, '企业版', '["行业五行匹配","品牌寓意","5个候选方案"]'),
-- 六爻类
('extra_divination', '六爻占卜·加抽一次', 'divination', 2.90, 2.84, NULL, '["加抽一次六爻"]'),
-- 合婚
('single_hehun', '两人合婚', 'bazi', 29.90, 29.90, '情感推荐', '["双方八字配对","五行互补分析","古籍参考"]');

-- ============================================
-- 3. 订单与支付
-- ============================================

CREATE TABLE IF NOT EXISTS orders (
    id TEXT PRIMARY KEY,                           -- 订单ID，如 ord_20260611153852_1c5b70
    user_id TEXT NOT NULL,                         -- 用户ID
    product_id TEXT NOT NULL,                      -- 商品ID
    product_name TEXT NOT NULL,                    -- 商品名称（冗余）
    amount REAL NOT NULL,                          -- 实际支付金额
    original_price REAL NOT NULL,                  -- 原价
    status TEXT DEFAULT 'pending',                 -- 状态：pending(待支付)/reviewing(待审核)/paid(已支付)/expired(已过期)/refunded(已退款)/rejected(已拒绝)
    payment_channel TEXT DEFAULT 'personal_qr',    -- 支付通道：personal_qr(个人码)
    payment_method TEXT,                           -- 支付方式：wechat/alipay/bank
    proof_image TEXT,                              -- 支付凭证图片URL
    proof_uploaded_at TEXT,                        -- 凭证上传时间
    reviewed_at TEXT,                              -- 审核时间
    reviewed_by TEXT,                              -- 审核人
    review_note TEXT,                              -- 审核备注
    paid_at TEXT,                                  -- 支付确认时间
    expired_at TEXT,                               -- 过期时间
    note TEXT,                                     -- 备注
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_created_at ON orders(created_at);
CREATE INDEX idx_orders_product_id ON orders(product_id);

-- ============================================
-- 4. 祈福灯
-- ============================================

CREATE TABLE IF NOT EXISTS blessing_lamps (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,                         -- 点灯用户ID
    order_id TEXT,                                 -- 关联订单ID
    lamp_type TEXT NOT NULL,                       -- 灯类型：peace(平安)/wisdom(智慧)/love(姻缘)/wealth(财福)
    for_person TEXT NOT NULL,                      -- 为谁祈福
    relation TEXT NOT NULL,                        -- 关系：father/mother/spouse/child/grandchild/friend/self
    wish TEXT,                                     -- 心愿（最多80字）
    display_name TEXT,                             -- 显示在灯墙的称呼
    duration_hours INTEGER DEFAULT 24,             -- 供奉时长（小时）
    status TEXT DEFAULT 'active',                  -- 状态：active(供奉中)/expired(已过期)
    lit_at TEXT DEFAULT (datetime('now')),         -- 点灯时间
    expires_at TEXT,                               -- 过期时间
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_blessing_lamps_user_id ON blessing_lamps(user_id);
CREATE INDEX idx_blessing_lamps_status ON blessing_lamps(status);
CREATE INDEX idx_blessing_lamps_lit_at ON blessing_lamps(lit_at);

-- 祈福灯统计视图
CREATE VIEW IF NOT EXISTS v_blessing_stats AS
SELECT
    COUNT(*) as total_lamps,
    SUM(CASE WHEN date(lit_at) = date('now') THEN 1 ELSE 0 END) as today_new,
    SUM(CASE WHEN status = 'active' THEN 1 ELSE 0 END) as active_count
FROM blessing_lamps;

-- ============================================
-- 5. 灵签记录
-- ============================================

CREATE TABLE IF NOT EXISTS lottery_records (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    master TEXT NOT NULL,                          -- 师父：huiming/mingxin/xuanzhen
    question TEXT NOT NULL,                        -- 用户提问
    sign_number TEXT NOT NULL,                     -- 签号，如 "第X签"
    sign_title TEXT,                               -- 签题
    sign_level TEXT NOT NULL,                      -- 签级：上上/上吉/中吉/中平/下下
    sign_poem TEXT,                                -- 签诗
    sign_analysis TEXT,                            -- 签文解析
    master_reading TEXT,                           -- 师父开示
    lucky_items TEXT,                              -- 吉利事项JSON
    advice TEXT,                                   -- 建议
    is_free INTEGER DEFAULT 0,                     -- 是否免费次数
    order_id TEXT,                                 -- 关联订单ID（付费时）
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_lottery_records_user_id ON lottery_records(user_id);
CREATE INDEX idx_lottery_records_created_at ON lottery_records(created_at);

-- ============================================
-- 6. 八字记录
-- ============================================

CREATE TABLE IF NOT EXISTS bazi_records (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    master TEXT NOT NULL,                          -- 师父
    birth_year INTEGER NOT NULL,
    birth_month INTEGER NOT NULL,
    birth_day INTEGER NOT NULL,
    birth_shichen TEXT NOT NULL,                   -- 时辰：zi/chou/yin/mao/chen/si/wu/wei/shen/you/xu/hai
    gender TEXT NOT NULL,                          -- male/female
    bazi_result TEXT,                              -- 八字排盘结果JSON
    wuxing_analysis TEXT,                          -- 五行分析JSON
    shishen TEXT,                                  -- 十神JSON
    dayun TEXT,                                    -- 大运JSON
    liunian TEXT,                                  -- 流年JSON
    personality TEXT,                              -- 性格分析
    career TEXT,                                   -- 事业分析
    wealth TEXT,                                   -- 财运分析
    relationship TEXT,                             -- 感情分析
    health TEXT,                                   -- 健康分析
    master_reading TEXT,                           -- 师父开示
    is_free INTEGER DEFAULT 0,
    order_id TEXT,
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_bazi_records_user_id ON bazi_records(user_id);

-- ============================================
-- 7. 解梦记录
-- ============================================

CREATE TABLE IF NOT EXISTS dream_records (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    dream_description TEXT NOT NULL,               -- 梦境描述
    category TEXT,                                 -- 梦境分类
    interpretation TEXT,                           -- 解梦结果
    lucky_level TEXT,                              -- 吉/凶/平
    advice TEXT,                                   -- 建议
    is_free INTEGER DEFAULT 0,
    order_id TEXT,
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_dream_records_user_id ON dream_records(user_id);

-- ============================================
-- 8. 手相记录
-- ============================================

CREATE TABLE IF NOT EXISTS palmistry_records (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    image_url TEXT NOT NULL,                       -- 手相图片URL
    palm_analysis TEXT,                            -- 手纹分析JSON
    reading TEXT,                                  -- 详解内容
    is_free INTEGER DEFAULT 0,
    order_id TEXT,
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_palmistry_records_user_id ON palmistry_records(user_id);

-- ============================================
-- 9. 起名记录
-- ============================================

CREATE TABLE IF NOT EXISTS naming_records (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    baby_name TEXT,                                -- 宝宝姓氏
    gender TEXT NOT NULL,                          -- male/female
    birth_info TEXT,                               -- 出生信息JSON
    style TEXT,                                    -- 风格：诗意/刚毅/儒雅/清逸/典雅/温润
    candidates TEXT,                               -- 候选名列表JSON
    selected_name TEXT,                            -- 用户选中的名字
    is_company INTEGER DEFAULT 0,                  -- 是否公司起名
    is_free INTEGER DEFAULT 0,
    order_id TEXT,
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_naming_records_user_id ON naming_records(user_id);

-- ============================================
-- 10. 六爻记录
-- ============================================

CREATE TABLE IF NOT EXISTS divination_records (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    question TEXT NOT NULL,                        -- 所问之事
    hexagram TEXT,                                 -- 卦象
    hexagram_name TEXT,                            -- 卦名
    analysis TEXT,                                 -- 解卦内容
    is_free INTEGER DEFAULT 0,
    order_id TEXT,
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_divination_records_user_id ON divination_records(user_id);

-- ============================================
-- 11. 禅修音乐
-- ============================================

CREATE TABLE IF NOT EXISTS meditation_tracks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,                           -- 曲目名称
    subtitle TEXT,                                 -- 副标题
    icon TEXT DEFAULT '🧘',                        -- emoji图标
    url TEXT NOT NULL,                             -- 音频文件URL
    duration INTEGER NOT NULL,                     -- 时长（秒）
    category TEXT,                                 -- 分类
    sort_order INTEGER DEFAULT 0,
    is_active INTEGER DEFAULT 1,
    play_count INTEGER DEFAULT 0,                  -- 播放次数
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_meditation_tracks_active ON meditation_tracks(is_active);

-- 禅修播放记录
CREATE TABLE IF NOT EXISTS meditation_plays (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    track_id TEXT NOT NULL,
    duration_listened INTEGER DEFAULT 0,           -- 实际聆听时长（秒）
    completed INTEGER DEFAULT 0,                   -- 是否完整播放
    merit_earned INTEGER DEFAULT 0,                -- 获得功德（秒）
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (track_id) REFERENCES meditation_tracks(id)
);

CREATE INDEX idx_meditation_plays_user_id ON meditation_plays(user_id);

-- ============================================
-- 12. 黄历缓存
-- ============================================

CREATE TABLE IF NOT EXISTS almanac_cache (
    date TEXT PRIMARY KEY,                         -- 日期 YYYY-MM-DD
    data TEXT NOT NULL,                            -- 黄历数据JSON
    created_at TEXT DEFAULT (datetime('now')),
    expires_at TEXT                                -- 缓存过期时间
);

-- ============================================
-- 13. 活动追踪
-- ============================================

CREATE TABLE IF NOT EXISTS activity_logs (
    id TEXT PRIMARY KEY,
    user_id TEXT,
    device_id TEXT NOT NULL,
    event_type TEXT NOT NULL,                      -- 事件类型：page_view/custom
    path TEXT,                                     -- 页面路径
    title TEXT,                                    -- 页面标题
    referrer TEXT,                                 -- 来源
    extra TEXT,                                    -- 额外数据JSON
    created_at TEXT DEFAULT (datetime('now'))
);

CREATE INDEX idx_activity_logs_device_id ON activity_logs(device_id);
CREATE INDEX idx_activity_logs_created_at ON activity_logs(created_at);
CREATE INDEX idx_activity_logs_event_type ON activity_logs(event_type);

-- ============================================
-- 14. 历史记录（用户端）
-- ============================================

CREATE TABLE IF NOT EXISTS user_history (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    kind TEXT NOT NULL,                            -- 类型：lottery/bazi/dream/palmistry/naming/divination
    title TEXT,                                    -- 标题
    subtitle TEXT,                                 -- 副标题
    payload TEXT,                                  -- 详细内容JSON
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_user_history_user_id ON user_history(user_id);
CREATE INDEX idx_user_history_kind ON user_history(kind);

-- ============================================
-- 15. 邀请/分享返佣
-- ============================================

CREATE TABLE IF NOT EXISTS referrals (
    id TEXT PRIMARY KEY,
    referrer_id TEXT NOT NULL,                     -- 邀请人用户ID
    referred_id TEXT NOT NULL,                     -- 被邀请人用户ID
    invite_code TEXT NOT NULL,                     -- 邀请码
    merit_reward INTEGER DEFAULT 0,                -- 奖励福报金（分）
    status TEXT DEFAULT 'pending',                 -- pending/confirmed/paid
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (referrer_id) REFERENCES users(id),
    FOREIGN KEY (referred_id) REFERENCES users(id)
);

CREATE INDEX idx_referrals_referrer_id ON referrals(referrer_id);
CREATE INDEX idx_referrals_invite_code ON referrals(invite_code);

-- 福报金流水
CREATE TABLE IF NOT EXISTS merit_transactions (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    type TEXT NOT NULL,                            -- type：earn(获得)/spend(消费)/withdraw(提现)/reward(奖励)
    amount INTEGER NOT NULL,                       -- 金额（分）
    source TEXT,                                   -- 来源：referral/order/meditation/system
    reference_id TEXT,                             -- 关联ID（订单ID/邀请ID等）
    note TEXT,                                     -- 备注
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_merit_transactions_user_id ON merit_transactions(user_id);

-- ============================================
-- 16. 后台管理
-- ============================================

-- 管理员账号
CREATE TABLE IF NOT EXISTS admin_users (
    id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT DEFAULT 'admin',                     -- admin/super_admin
    is_active INTEGER DEFAULT 1,
    last_login_at TEXT,
    created_at TEXT DEFAULT (datetime('now'))
);

-- 审计日志
CREATE TABLE IF NOT EXISTS audit_logs (
    id TEXT PRIMARY KEY,
    admin_id TEXT NOT NULL,
    action TEXT NOT NULL,                          -- 操作：order_confirm/order_reject/user_update/product_update
    target_type TEXT,                              -- 目标类型：order/user/product
    target_id TEXT,                                -- 目标ID
    detail TEXT,                                   -- 操作详情JSON
    ip_address TEXT,
    created_at TEXT DEFAULT (datetime('now')),
    FOREIGN KEY (admin_id) REFERENCES admin_users(id)
);

CREATE INDEX idx_audit_logs_admin_id ON audit_logs(admin_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- ============================================
-- 17. 系统配置
-- ============================================

CREATE TABLE IF NOT EXISTS system_config (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    description TEXT,
    updated_at TEXT DEFAULT (datetime('now'))
);

-- 初始配置
INSERT INTO system_config (key, value, description) VALUES
('site_name', '菩提苑', '站点名称'),
('site_description', '心诚则灵。为家人点一盏祈福灯，求一支关帝灵签，看一卦命理八字。', '站点描述'),
('free_lottery_daily', '3', '每日免费灵签次数'),
('free_dream_daily', '5', '每日免费解梦次数'),
('free_bazi_daily', '1', '每日免费八字次数'),
('blessing_default_price', '6.60', '祈福灯默认价格（元）'),
('blessing_default_duration', '24', '祈福灯默认供奉时长（小时）'),
('order_expire_minutes', '30', '订单超时未支付自动过期（分钟）'),
('referral_reward_rate', '0.10', '分享返佣比例（10%）'),
('merit_per_yuan', '100', '每元消费获得功德（分）');

-- ============================================
-- 18. 收款码管理
-- ============================================

CREATE TABLE IF NOT EXISTS payment_qrcodes (
    id TEXT PRIMARY KEY,
    channel TEXT NOT NULL,                         -- wechat/alipay/bank
    image_url TEXT NOT NULL,                       -- 收款码图片URL
    receiver_name TEXT DEFAULT '菩提苑',            -- 收款方名称
    is_active INTEGER DEFAULT 1,
    created_at TEXT DEFAULT (datetime('now'))
);
