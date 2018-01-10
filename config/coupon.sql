-- ----------------------------
-- 插入分类
-- ----------------------------
INSERT INTO category_info (id, parent_id, category_name) VALUES (1, 0,  "parent category");
INSERT INTO category_info (id, parent_id, category_name) VALUES (2, 1,  "category");
INSERT INTO category_info (id, parent_id, category_name) VALUES (3, 1,  "second category");
INSERT INTO product_category (product_id, category_id) VALUES (1001, 2);
INSERT INTO product_category (product_id, category_id) VALUES (1002, 2);
INSERT INTO product_category (product_id, category_id) VALUES (1003, 2);
INSERT INTO product_category (product_id, category_id) VALUES (1004, 3);

-- ----------------------------
-- 插入商品
-- ----------------------------
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1001, "sit测试商品", "sit测试商品附属", 1, 100, 100, 10, 200, 0, 1, 0, 0, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1002, "sit测试商品上架", "sit测试商品附属上架", 1, 100, 100, 10, 200, 1, 2, 0, 0, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1003, "测试商品（不在优惠券中的）", "测试商品附属（不在优惠券中的）", 1, 100, 100, 10, 200, 1, 3, 0, 0, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1004, "测试商品1004", "测试商品附属", 1, 100, 100, 10, 200, 1, 4, 0, 0, 1, 1, 0);
-- ----------------------------
-- 插入SKU
-- ----------------------------
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (1, "master_name1", "slave_name1", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1001, 100, 100, 70, 200, 100, 100, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (2, "master_name2", "slave_name2", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1002, 100, 100, 70, 200, 160, 120, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (3, "master_name3", "slave_name3", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1003, 100, 100, 70, 200, 160, 120, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (4, "master_name4", "slave_name4", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1004, 100, 100, 70, 200, 160, 120, 0, 0, "123456");
-- ----------------------------
-- 插入优惠规则
-- ----------------------------
INSERT INTO coupons (id, discount_type, code_type, title, sales_channel_id, total_count, begin_time, end_time, validity_period, limit_amount, min_discount_amount, max_discount_amount, min_discount_rate, max_discount_rate, coupons_status, maximum_per_user, return_flag, scope_type) VALUES (1,0,2,"sit优惠券描述",0,10,20201001000000,20201001000000,0,40,50,50,0,0,0,5,0,2);
INSERT INTO coupons (id, discount_type, code_type, coupons_code, title, sales_channel_id, total_count, begin_time, end_time, validity_period, limit_amount, min_discount_amount, max_discount_amount, min_discount_rate, max_discount_rate, coupons_status, maximum_per_user, return_flag, scope_type) VALUES (2,3,2,"这是一个重复优惠券码！", "title优惠券描述",0,4,NOW(),NOW(),0,40,0,0,60,90,2,5,0,2);
INSERT INTO coupons (id, discount_type, code_type, title, sales_channel_id, total_count, begin_time, end_time, validity_period, limit_amount, min_discount_amount, max_discount_amount, min_discount_rate, max_discount_rate, coupons_status, maximum_per_user, return_flag, scope_type) VALUES (3,0,2,"sit优惠券更新描述",0,2,NOW(),NOW(),10,40,50,50,0,0,0,5,0,2);
INSERT INTO coupons (id, discount_type, code_type, title, sales_channel_id, total_count, begin_time, end_time, validity_period, limit_amount, min_discount_amount, max_discount_amount, min_discount_rate, max_discount_rate, coupons_status, maximum_per_user, return_flag, scope_type) VALUES (4,0,2,"sit优惠券更新描述(未开始)",0,2,20201001000000,20201001000000,0,40,50,50,0,0,0,5,0,1);
-- ----------------------------
-- 插入优惠规则与商品的关联关系
-- ----------------------------
INSERT INTO coupons_scope (id, coupons_id, type, object_id) VALUES (1, 1, 1, 1001);
INSERT INTO coupons_scope (id, coupons_id, type, object_id) VALUES (2, 1, 1, 1002);
INSERT INTO coupons_scope (id, coupons_id, type, object_id) VALUES (3, 1, 0, 2);
INSERT INTO coupons_scope (id, coupons_id, type, object_id) VALUES (4, 4, 0, 2);
-- ----------------------------
-- 插入商品优惠属性
-- ----------------------------
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (1, 1001, 22, 4);
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (2, 1002, 22, 4);
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (3, 1001, 22, 1);
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (4, 1001, 23, 1);
-- ----------------------------
-- 插入优惠码
-- ----------------------------
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (1,2,0,"这是一个重复优惠券码！",50,50,0,0,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (2,2,0,"这是一个重复优惠券码！",50,50,0,1,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (3,2,0,"这是一个重复优惠券码！",50,50,1001,2,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (4,2,0,"这是一个重复优惠券码！",50,50,1002,3,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (5,3,0,"ddddas",50,50,1001,2,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (6,3,0,"ddddas",50,50,1002,3,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (7,4,0,"ttttt",50,50,0,2,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (8,4,0,"ttttt",50,50,0,3,NOW(),NOW())