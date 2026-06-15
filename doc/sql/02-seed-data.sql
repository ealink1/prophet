-- ============================================
-- 种子数据 - 禅修音乐曲目
-- ============================================

INSERT INTO meditation_tracks (id, title, subtitle, icon, url, duration, category, sort_order) VALUES
('track_001', '般若波罗蜜多心经', '梵音净化 · 静心安神', '🧘', '/audio/xinjing.mp3', 360, '经典', 1),
('track_002', '大悲咒', '慈悲加持 · 消灾祈福', '🙏', '/audio/dabeizhou.mp3', 480, '经典', 2),
('track_003', '六字大明咒', '嗡嘛呢叭咪吽 · 净化身心', '📿', '/audio/liuzi.mp3', 300, '咒语', 3),
('track_004', '禅院钟声', '古刹清音 · 远离尘嚣', '🔔', '/audio/zhongsheng.mp3', 420, '自然', 4),
('track_005', '流水禅音', '山涧流水 · 禅意悠然', '🌊', '/audio/liushui.mp3', 540, '自然', 5),
('track_006', '木鱼声声', '节奏平稳 · 入定良伴', '🪵', '/audio/muyu.mp3', 600, '节奏', 6),
('track_007', '莲花生大师心咒', '密法加持 · 消除业障', '🪷', '/audio/lianhua.mp3', 450, '咒语', 7),
('track_008', '文殊菩萨心咒', '智慧增长 · 断除烦恼', '🗡️', '/audio/wenshu.mp3', 360, '咒语', 8),
('track_009', '药师佛心咒', '消灾延寿 · 除病苦', '💊', '/audio/yaoshi.mp3', 400, '咒语', 9),
('track_010', '阿弥陀佛圣号', '往生净土 · 莲花化生', ' Amituofo', '/audio/amituofo.mp3', 500, '佛号', 10);

-- ============================================
-- 种子数据 - 管理员账号
-- ============================================

INSERT INTO admin_users (id, username, password_hash, role) VALUES
('admin_001', 'admin', '$2b$10$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx', 'super_admin');

-- ============================================
-- 种子数据 - 收款码
-- ============================================

INSERT INTO payment_qrcodes (id, channel, image_url, receiver_name) VALUES
('qr_001', 'wechat', '/images/qrcode_wechat.png', '菩提苑'),
('qr_002', 'alipay', '/images/qrcode_alipay.png', '菩提苑');

-- ============================================
-- 种子数据 - 系统配置
-- ============================================

INSERT INTO system_config (key, value, description) VALUES
('api_base_url', 'https://denied-accomplished-micro-drawn.trycloudflare.com/api/v1', '后端API地址'),
('pwa_name', '菩提苑', 'PWA应用名称'),
('pwa_theme_color', '#1a1410', 'PWA主题色'),
('share_title', '菩提苑 · 为家人祈福求灵签', '分享标题'),
('share_description', '心诚则灵。为家人点一盏祈福灯，求一支关帝灵签，看一卦命理八字。', '分享描述'),
('share_image', 'https://putiyuan.pages.dev/share-cover.svg', '分享封面图');
