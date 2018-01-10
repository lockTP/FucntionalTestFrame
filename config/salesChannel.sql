-- ----------------------------
-- 插入销售渠道
-- ----------------------------
INSERT INTO sales_channel_info  (id, channel_names, channel_desc, channel_status) VALUES (100, "销售渠道", "销售渠道描述", 1);
INSERT INTO sales_channel_info  (id, channel_names, channel_desc, channel_status) VALUES (101, "修改销售渠道", "修改销售渠道描述", 1);
INSERT INTO sales_channel_info  (id, channel_names, channel_desc, channel_status) VALUES (102, "删除销售渠道", "删除销售渠道描述", 1);
INSERT INTO sales_channel_info  (id, channel_names, channel_desc, channel_status) VALUES (103, "复制销售渠道", "复制销售渠道描述", 1);
INSERT INTO sales_channel_info  (id, channel_names, channel_desc, channel_status) VALUES (104, "销售渠道2", "销售渠道描述2", 1);

-- ----------------------------
-- 插入SKU信息
-- ----------------------------
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (1, "master_name", "slave_name", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 11, 100, 100, 70, 200, 100, 100, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (2, "查询渠道测试数据", "查询渠道测试数据", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 12, 100, 100, 70, 200, 200, 100, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (3, "sku name", "sku slave name", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 13, 100, 100, 70, 200, 200, 100, 0, 0, "123456");
-- ----------------------------
-- 插入渠道信息
-- ----------------------------
INSERT INTO sku_sales_channel_info  (id, product_id, sku_id, sales_channel_id, price, inventory, available_inventory) VALUES (1001, 12, 2, 100, 120, 50, 30);
INSERT INTO sku_sales_channel_info  (id, product_id, sku_id, sales_channel_id, price, inventory, available_inventory) VALUES (1002, 12, 2, 104, 150, 0, -1);

-- ----------------------------
-- 插入渠道属性
-- ----------------------------
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (1, 12, 27, 100);
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (2, 12, 27, 104);
-- ----------------------------
-- 插入商品信息
-- ----------------------------
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source) VALUES (11, "sit测试商品1", "sit测试商品附属1", 1, 100, 100, 10, 200, 0, 1, 0, 0, 1, 1);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source) VALUES (12, "sit测试商品2", "sit测试商品附属2", 1, 100, 100, 10, 200, 0, 1, 0, 0, 1, 1);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source) VALUES (13, "sit测试商品3(兑换码发货类型)", "sit测试商品附属3", 1, 100, 100, 10, 200, 0, 1, 1, 0, 1, 1)