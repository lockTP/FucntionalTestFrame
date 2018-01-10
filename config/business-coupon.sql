-- ----------------------------
-- 插入优惠规则
-- ----------------------------
INSERT INTO coupons (id, discount_type, code_type, title, sales_channel_id, total_count, begin_time, end_time, validity_period, limit_amount, min_discount_amount, max_discount_amount, min_discount_rate, max_discount_rate, coupons_status, maximum_per_user, return_flag, scope_type) VALUES (1,0,0,"测试运费券-无码／相对有效期",0,5,20101001000000,20201001000000,10,800,500,500,0,0,2,1,0,3);
INSERT INTO coupons (id, discount_type, code_type, title, sales_channel_id, total_count, begin_time, end_time, validity_period, limit_amount, min_discount_amount, max_discount_amount, min_discount_rate, max_discount_rate, coupons_status, maximum_per_user, return_flag, scope_type) VALUES (2,0,2,"测试-统一优惠码／相对有效期",0,5,20101001000000,20201001000000,0,4000,500,500,0,0,2,1,0,0);



-- ----------------------------
-- 插入优惠码
-- ----------------------------
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (1,1,0,"",500,0,0,1,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (2,1,0,"",500,0,0,1,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (3,1,0,"",500,00,1001,2,"2017-11-09 00:00:00","2018-11-09 00:00:00");
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (4,1,0,"",500,0,1002,3,"2017-11-09 00:00:00","2018-11-09 00:00:00");
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (5,1,0,"",500,0,0,0,NOW(),NOW());

INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (6,2,0,"NFYHS123B",500,0,0,1,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (7,2,0,"NFYHS123B",500,0,0,1,NOW(),NOW());
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (8,2,0,"NFYHS123B",500,0,1001,2,"2017-11-09 00:00:00","2018-12-09 00:00:00");
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES (9,2,0,"NFYHS123B",500,0,1002,3,"2017-11-09 00:00:00","2018-12-09 00:00:00");
INSERT INTO coupons_code (id, coupons_id, sales_channel_id, coupons_code, discount_amount, discount_rate, user_id, code_status, begin_time, end_time) VALUES(10,2,0,"NFYHS123B",500,0,0,1,NOW(),NOW());


-- ----------------------------
-- 插入分类
-- ----------------------------
INSERT INTO category_info (id, parent_id, category_name) VALUES (1, 0,  "parent category");
INSERT INTO category_info (id, parent_id, category_name) VALUES (2, 1,  "category");
INSERT INTO product_category (product_id, category_id) VALUES (1001, 2);
INSERT INTO product_category (product_id, category_id) VALUES (1002, 2);


-- ----------------------------
-- 插入商品
-- ----------------------------
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1001, "sit测试商品", "sit测试商品附属", 1, 100, 100, 10, 200, 1, 1, 0, 0, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1002, "sit测试商品上架", "sit测试商品附属上架", 1, 100, 100, 10, 200, 1, 2, 0, 0, 1, 1, 0);

INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (1, "master_name", "slave_name", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1001, 100, 1000, 70, 200, 100, 100, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (2, "master_name2", "slave_name2", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1002, 100, 6000, 70, 200, 160, 120, 0, 0, "123456");