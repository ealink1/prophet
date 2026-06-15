-- ============================================
-- 常用查询 SQL
-- ============================================

-- ============================================
-- 1. 数据看板统计
-- ============================================

-- 今日收入
SELECT COALESCE(SUM(amount), 0) as today_income
FROM orders
WHERE status = 'paid'
  AND date(paid_at) = date('now');

-- 累计收入
SELECT COALESCE(SUM(amount), 0) as total_income
FROM orders
WHERE status = 'paid';

-- 用户统计
SELECT
    (SELECT COUNT(*) FROM users) as total_users,
    (SELECT COUNT(*) FROM users WHERE date(last_active_at) = date('now')) as today_active,
    (SELECT COUNT(*) FROM user_history) as total_history,
    (SELECT COUNT(*) FROM user_history WHERE date(created_at) = date('now')) as today_unlocks;

-- 订单统计
SELECT
    (SELECT COUNT(*) FROM orders) as total_orders,
    (SELECT COUNT(*) FROM orders WHERE status = 'paid') as paid_orders,
    (SELECT COUNT(*) FROM orders WHERE status = 'pending') as pending_orders,
    (SELECT COUNT(*) FROM orders WHERE status = 'reviewing') as reviewing_orders;

-- ============================================
-- 2. 订单管理查询
-- ============================================

-- 订单列表（带用户和商品信息）
SELECT
    o.id as order_no,
    o.product_name,
    o.product_id,
    o.amount,
    o.original_price,
    o.status,
    o.payment_channel,
    o.note,
    u.id as user_id,
    o.created_at,
    o.expired_at,
    o.proof_uploaded_at
FROM orders o
LEFT JOIN users u ON o.user_id = u.id
ORDER BY o.created_at DESC
LIMIT 50 OFFSET 0;

-- 按状态筛选订单
SELECT * FROM orders
WHERE status = 'reviewing'  -- 待审核
ORDER BY created_at DESC;

-- 确认到账
UPDATE orders
SET status = 'paid',
    reviewed_at = datetime('now'),
    reviewed_by = 'admin_001',
    paid_at = datetime('now')
WHERE id = 'ord_xxx';

-- 拒绝订单
UPDATE orders
SET status = 'rejected',
    reviewed_at = datetime('now'),
    reviewed_by = 'admin_001',
    review_note = '凭证不清晰'
WHERE id = 'ord_xxx';

-- 过期订单（超过30分钟未支付）
UPDATE orders
SET status = 'expired'
WHERE status = 'pending'
  AND datetime(created_at, '+30 minutes') < datetime('now');

-- ============================================
-- 3. 祈福灯管理
-- ============================================

-- 灯墙数据（脱敏显示）
SELECT
    CASE
        WHEN LENGTH(display_name) > 1
        THEN SUBSTR(display_name, 1, 1) || '**'
        ELSE '**'
    END as masked_name,
    relation,
    lamp_type,
    wish,
    lit_at,
    expires_at
FROM blessing_lamps
WHERE status = 'active'
ORDER BY lit_at DESC;

-- 今日新增灯数
SELECT COUNT(*) as today_new
FROM blessing_lamps
WHERE date(lit_at) = date('now');

-- 供奉中的灯数
SELECT COUNT(*) as active_count
FROM blessing_lamps
WHERE status = 'active';

-- ============================================
-- 4. 用户分析
-- ============================================

-- 用户活跃度排行
SELECT
    u.id,
    u.nickname,
    u.lucky_code,
    COUNT(DISTINCT al.id) as activity_count,
    MAX(al.created_at) as last_active
FROM users u
LEFT JOIN activity_logs al ON u.id = al.user_id
GROUP BY u.id
ORDER BY activity_count DESC
LIMIT 20;

-- 用户消费排行
SELECT
    u.id,
    u.nickname,
    u.lucky_code,
    COUNT(o.id) as order_count,
    SUM(o.amount) as total_spent
FROM users u
LEFT JOIN orders o ON u.id = o.user_id AND o.status = 'paid'
GROUP BY u.id
ORDER BY total_spent DESC
LIMIT 20;

-- ============================================
-- 5. 商品销售分析
-- ============================================

-- 各商品销售统计
SELECT
    p.product_id,
    p.name,
    COUNT(o.id) as order_count,
    SUM(CASE WHEN o.status = 'paid' THEN o.amount ELSE 0 END) as revenue,
    SUM(CASE WHEN o.status = 'pending' THEN 1 ELSE 0 END) as pending_count
FROM products p
LEFT JOIN orders o ON p.product_id = o.product_id
GROUP BY p.product_id
ORDER BY revenue DESC;

-- 今日各商品销售
SELECT
    p.name,
    COUNT(o.id) as today_orders,
    SUM(CASE WHEN o.status = 'paid' THEN o.amount ELSE 0 END) as today_revenue
FROM products p
LEFT JOIN orders o ON p.product_id = o.product_id AND date(o.created_at) = date('now')
GROUP BY p.product_id
ORDER BY today_revenue DESC;

-- ============================================
-- 6. 禅修音乐统计
-- ============================================

-- 热门曲目
SELECT
    t.id,
    t.title,
    t.subtitle,
    t.duration,
    COUNT(p.id) as play_count,
    SUM(p.duration_listened) as total_listen_time
FROM meditation_tracks t
LEFT JOIN meditation_plays p ON t.id = p.track_id
GROUP BY t.id
ORDER BY play_count DESC
LIMIT 10;

-- ============================================
-- 7. 分享返佣统计
-- ============================================

-- 邀请排行榜
SELECT
    u.id,
    u.nickname,
    u.lucky_code,
    COUNT(r.id) as invite_count,
    SUM(r.merit_reward) as total_reward
FROM users u
LEFT JOIN referrals r ON u.id = r.referrer_id
GROUP BY u.id
HAVING invite_count > 0
ORDER BY invite_count DESC
LIMIT 20;

-- ============================================
-- 8. 灵签使用统计
-- ============================================

-- 各师父使用频率
SELECT
    master,
    COUNT(*) as usage_count,
    SUM(CASE WHEN sign_level IN ('上上', '上吉') THEN 1 ELSE 0 END) as good_signs,
    SUM(CASE WHEN sign_level = '下下' THEN 1 ELSE 0 END) as bad_signs
FROM lottery_records
GROUP BY master;

-- 各签级出现频率
SELECT
    sign_level,
    COUNT(*) as count,
    ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM lottery_records), 1) as percentage
FROM lottery_records
GROUP BY sign_level
ORDER BY count DESC;

-- ============================================
-- 9. 八字分析统计
-- ============================================

-- 各时辰出生分布
SELECT
    birth_shichen,
    COUNT(*) as count
FROM bazi_records
GROUP BY birth_shichen
ORDER BY count DESC;

-- 性别分布
SELECT
    gender,
    COUNT(*) as count
FROM bazi_records
GROUP BY gender;

-- ============================================
-- 10. 清理与维护
-- ============================================

-- 清理过期订单
DELETE FROM orders
WHERE status = 'expired'
  AND datetime(created_at, '+7 days') < datetime('now');

-- 清理过期黄历缓存
DELETE FROM almanac_cache
WHERE expires_at < datetime('now');

-- 清理30天前的活动日志
DELETE FROM activity_logs
WHERE datetime(created_at, '-30 days') < datetime('now');
