-- ----------------------------
-- 插入订单
-- ----------------------------
INSERT INTO order_info (id, order_id, pay_order_id, user_id, origin_user_id, supplier_id) VALUES (1, "20171211", "12345611", 1002, 0, 1);
INSERT INTO order_info (id, order_id, pay_order_id, user_id, origin_user_id, supplier_id, pay_status) VALUES (2, "20171213", "12345613", 1002, 0, 1, 1);
-- ----------------------------
-- 插入订单SKU
-- ----------------------------
INSERT INTO order_sku (id, order_id, product_id, sku_id, product_name, sku_name, sku_code, sku_image, sku_count, price, market_price) VALUES (1, "20171211", 1001, 1, "订单商品", "订单商品sku", "111", "zzz.jpg", 10, 1000, 1200);
-- ----------------------------
-- 插入购物车车
-- ----------------------------
INSERT INTO shopping_cart (id, user_id, sku_id, product_name, image, sku_name, sku_count, price, score, relation_id, promotion_id) VALUES (1, 1002, 1, "sit测试商品1", "xxx.jpg", "master_name", 2, 1000, 1000, 0,0);
INSERT INTO shopping_cart (id, user_id, sku_id, product_name, image, sku_name, sku_count, price, score, relation_id, promotion_id) VALUES (2, 1002, 2, "sit测试商品2", "yyy.jpg", "master_name2", 2, 1000, 1000, 0,0);

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
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1001, "sit测试商品1", "sit测试商品附属", 1, 100, 100, 10, 200, 1, 1, 0, 0, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1002, "sit测试商品2", "sit测试商品附属上架", 1, 100, 100, 10, 200, 1, 2, 0, 0, 1, 1, 0);

INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (1, "master_name", "slave_name", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1001, 100, 1000, 70, 200, 100, 100, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (2, "master_name2", "slave_name2", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1002, 100, 6000, 70, 200, 160, 120, 0, 0, "123456")