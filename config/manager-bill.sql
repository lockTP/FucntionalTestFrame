-- ----------------------------
-- 插入订单信息
-- ----------------------------
INSERT INTO order_info (id, order_id, pay_order_id, user_id, origin_user_id, supplier_id, name, delivery_time) VALUES (1, "20171130", "123456", 1, 1, 1, "vince", "2017-12-07 16:12:32");

-- ----------------------------
-- 插入售后单信息
-- ----------------------------
INSERT INTO service_info (id, user_id, user_name, user_phone, service_id, supplier_id) VALUES (1, 1, "SIT Tester", "12345678912", "123456123", 1);

-- ----------------------------
-- 插入售后sku信息
-- ----------------------------
INSERT INTO service_sku (id, order_id, service_id) VALUES (1, "20171130", "123456123");
-- ----------------------------
-- 插入差异项信息
-- ----------------------------
INSERT INTO order_balance_info (id, supplier_id, service_id, order_id, type_name, balance_amount, balance_attachment, operator, source, remark, created_at) VALUES (1, 1, "123456123", "20171130", "差异项补录", "100", "url", "SIT Tester", 1, "客服录入差异项", 20171001000000);
INSERT INTO order_balance_info (id, supplier_id, service_id, order_id, type_name, balance_amount, balance_attachment, operator, source, remark, created_at) VALUES (2, 1, "123456122", "20171130", "差异项补录2", "100", "url", "SIT Tester", 2, "供应商录入差异项", 20171001000000)