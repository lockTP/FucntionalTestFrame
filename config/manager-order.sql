INSERT INTO order_info (id, order_id, pay_order_id, user_id, origin_user_id, supplier_id, name, delivery_time) VALUES (1, "20171130", "123456", 1, 1, 1, "vince", "2017-12-07 16:12:32");

-- ----------------------------
-- 插入包裹信息
-- ----------------------------
INSERT INTO order_package_info (id, order_id, package_type, express_amount, express_company_id, express_company_name, express_id, remark) VALUES (1,"20171130",1,100,1,"圆通速递","123456789","这是一个备注");
INSERT INTO order_package_info (id, order_id, package_type, express_amount, express_company_id, express_company_name, express_id, remark) VALUES (2,"20171130",1,100,1,"圆通速递","123456111","这是一个备注");

-- ----------------------------
-- 插入包裹SKU信息
-- ----------------------------
INSERT INTO package_sku_info (id, package_id, sku_id, sku_count) VALUES (1,1,1001,100);
INSERT INTO package_sku_info (id, package_id, sku_id, sku_count) VALUES (2,2,1001,200)