-- ----------------------------
-- 插入分类
-- ----------------------------
INSERT INTO category_info (id, parent_id, category_name) VALUES (1, 0,  "parent category");
INSERT INTO category_info (id, parent_id, category_name) VALUES (2, 1,  "category");
INSERT INTO category_info (id, parent_id, category_name) VALUES (3, 1,  "category3");
INSERT INTO category_info (id, parent_id, category_name) VALUES (4, 1,  "category4");
INSERT INTO product_category (product_id, category_id) VALUES (1001, 2);
INSERT INTO product_category (product_id, category_id) VALUES (1002, 2);
INSERT INTO product_category (product_id, category_id) VALUES (1003, 3);
-- ----------------------------
-- 插入前端品类
-- ----------------------------
INSERT INTO front_category_info (id, parent_id, category_name, image, sort_order) VALUES (1, 0,  "parent category", "", 1);
INSERT INTO front_category_info (id, parent_id, category_name, image, sort_order) VALUES (2, 1,  "category", "xxxxxy.jpg", 1);
INSERT INTO front_category_info (id, parent_id, category_name, image, sort_order) VALUES (3, 1,  "category2", "xxxxx.jpg", 1);
-- ----------------------------
-- 插入属性
-- ----------------------------
INSERT INTO attr_info (id, attr_name, remark, attr_style, attr_range) VALUES (1, "描述", "描述信息", 0, 0);
INSERT INTO attr_info (id, attr_name, remark, attr_style, attr_range) VALUES (2, "购买须知", "购买须知信息", 0, 0);
INSERT INTO attr_info (id, attr_name, remark, attr_style, attr_range) VALUES (22, "优惠信息", "优惠信息属性", 2,2);
INSERT INTO attr_info (id, attr_name, remark, attr_style, attr_range) VALUES (1001, "属性1", "属性1描述", 0, 1);
INSERT INTO attr_info (id, attr_name, remark, attr_style, attr_range) VALUES (1002, "删除属性测试", "属性描述", 0, 1);
-- ----------------------------
-- 插入属性值
-- ----------------------------
INSERT INTO attr_value_info (id, attr_id, attr_value, attr_value_display) VALUES (1, 1001, "这是一个属性值", "这是一个属性值(显示)");
INSERT INTO attr_value_info (id, attr_id, attr_value, attr_value_display) VALUES (2, 1002, "这是一个要删除的属性值","这是一个要删除的属性值(显示)");

-- ----------------------------
-- 插入属性类型
-- ----------------------------
INSERT INTO attr_type (id, name, remark) VALUES (1, "属性类型1", "这是一个属性类型1");
INSERT INTO attr_type (id, name, remark) VALUES (2, "属性类型2", "这是一个属性类型2");
INSERT INTO attr_type (id, name, remark) VALUES (1001, "属性类型3", "这是一个属性类型3");
-- ----------------------------
-- 插入前后端品类映射
-- ----------------------------
INSERT INTO front_category_mapping (id, front_category_id, related_id) VALUES (1, 2, 2);
INSERT INTO front_category_mapping (id, front_category_id, related_id) VALUES (2, 3, 4);
-- ----------------------------
-- 插入类目与属性关联
-- ----------------------------
-- INSERT INTO category_attr_info (id, attr_id, attr_name, category_id, attr_type_id, attr_mandatory) VALUES (1, 1001, "属性1", 2, 101, 0);
-- INSERT INTO category_attr_info (id, attr_id, attr_name, category_id, attr_type_id, attr_mandatory) VALUES (2, 1002, "这是一个要删除的属性值", 2, 1, 0);

-- ----------------------------
-- 插入类目与属性关联（new）
-- ----------------------------
INSERT INTO category_attr_info_new (id, attr_id, attr_name, category_id, attr_mandatory, attr_public) VALUES (1, 1001, "属性1", 2, 0, 0);
INSERT INTO category_attr_info_new (id, attr_id, attr_name, category_id, attr_mandatory) VALUES (2, 1002, "这是一个要删除的属性值", 2, 0);
INSERT INTO category_attr_info_new (id, attr_id, attr_name, category_id, attr_mandatory, attr_public) VALUES (3, 22, "优惠信息", 2, 1, 1);
-- ----------------------------
-- 插入类目与属性值关联（new）
-- ----------------------------
INSERT INTO category_attr_value_info (id, attr_id, category_id, attr_value, attr_value_display) VALUES (1, 1001, 2,  "这是一个属性值", "这是一个属性显示值");
INSERT INTO category_attr_value_info (id, attr_id, category_id, attr_value, attr_value_display) VALUES (2, 22, 2,  "1", "优惠券1");

-- ----------------------------
-- 插入商品
-- ----------------------------
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, product_type, status) VALUES (1001, "sit测试商品", "sit测试商品附属", 1, 100, 100, 10, 200, 1, 1, 0, 0, 1, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1002, "sit测试商品上架", "sit测试商品附属上架", 1, 100, 100, 10, 200, 1, 2, 0, 0, 1, 1, 0);
INSERT INTO product_info (id, master_name, slave_name, supplier_id, min_price, max_price, display_sales_count, score, product_status, default_sku_id, delivery_type, virtual_product, auto_delivery, product_source, status) VALUES (1003, "sit测试商品非VIP", "sit测试商品非VIP", 2, 100, 100, 10, 200, 0, 1, 0, 0, 1, 0, 0);

-- ----------------------------
-- 插入商品与属性关系
-- ----------------------------
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (1, 1002, 1, "1002商品描述");
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (2, 1002, 2, "1002商品购买须知");
INSERT INTO product_attr (id, product_id, attr_id, attr_value) VALUES (3, 1002, 1001, "这是一个属性值");
INSERT INTO product_attr (id, product_id, attr_id, attr_value, sales_channel_id) VALUES (4, 1002, 22, 1, 1);
INSERT INTO product_attr (id, product_id, attr_id, attr_value, sales_channel_id) VALUES (5, 1002, 22, 2, 0);
INSERT INTO product_attr (id, product_id, attr_id, attr_value, sales_channel_id) VALUES (6, 1002, 27, 3, 0);
-- ----------------------------
-- 插入选项
-- ----------------------------
INSERT INTO option_info (id, option_name) VALUES (1001, "选项1");
INSERT INTO option_info (id, option_name) VALUES (1002, "选项2");
INSERT INTO option_info (id, option_name) VALUES (1003, "选项3");

-- ----------------------------
-- 插入选项值
-- ----------------------------
INSERT INTO option_value_info (id, option_id, option_value) VALUES (1, 1001, "选项值11");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (2, 1001, "选项值12");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (3, 1001, "选项值13");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (4, 1001, "选项值14");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (5, 1002, "选项值21");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (6, 1002, "选项值22");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (7, 1003, "选项值31");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (8, 1003, "选项值32");
INSERT INTO option_value_info (id, option_id, option_value) VALUES (9, 1003, "选项值33");

-- ----------------------------
-- 插入SKU
-- ----------------------------
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (1, "master_name", "slave_name", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1001, 100, 100, 70, 200, 100, 100, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (2, "master_name2", "slave_name2", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1002, 100, 100, 70, 200, 160, 120, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (3, "master_name3", "slave_name3", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1002, 100, 100, 70, 200, 160, 120, 0, 0, "123456");
INSERT INTO sku_info  (id, master_name, slave_name, image, product_id, market_price, price, cost_price, score, inventory, available_inventory, sales_count, display_sales_count, sku_code) VALUES (4, "master_name4", "slave_name4", "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg", 1003, 100, 100, 70, 200, 160, 120, 0, 0, "123456");

-- ----------------------------
-- 插入选项与sku的关联（测试选项删除）
-- ----------------------------
INSERT INTO sku_option (id, sku_id, option_id, option_value) VALUES (1, 1, 1002, "选项值21");
INSERT INTO sku_option (id, sku_id, option_id, option_value) VALUES (2, 1, 1002, "选项值22");
INSERT INTO sku_option (id, sku_id, option_id, option_value) VALUES (3, 2, 1002, "选项值21");
INSERT INTO sku_option (id, sku_id, option_id, option_value) VALUES (4, 2, 1002, "选项值22");

-- ----------------------------
-- 插入供应商
-- ----------------------------
INSERT INTO supplier_info (id, supplier_name, supplier_desc) VALUES (1, "Aura供应商", "Aura供应商描述");

-- ----------------------------
-- 插入商品与属性类型关系
-- ----------------------------
INSERT INTO product_attr_type (id, product_id, attr_type_id, status) VALUES (1,1002, 1, 1);

-- ----------------------------
-- 插入属性与属性类型关系
-- ----------------------------
INSERT INTO attr_type_relation (id, attr_id, attr_type_id, status) VALUES (1,1, 1, 0);
INSERT INTO attr_type_relation (id, attr_id, attr_type_id, status) VALUES (2,2, 1, 0);
INSERT INTO attr_type_relation (id, attr_id, attr_type_id, status) VALUES (3,1002, 1, 0);
INSERT INTO attr_type_relation (id, attr_id, attr_type_id, status) VALUES (4,1002, 101, 0);
INSERT INTO attr_type_relation (id, attr_id, attr_type_id, status) VALUES (5,1001, 101, 0);
INSERT INTO attr_type_relation (id, attr_id, attr_type_id, status) VALUES (6,22, 3, 0);

-- ----------------------------
-- 插入品类与选项关联关系
-- ----------------------------
INSERT INTO category_option_info (id, option_id, option_name, category_id) VALUES (1, 1001, "选项1", 2);
INSERT INTO category_option_value_info (id, option_id, option_value, category_id) VALUES (1, 1001, "选项值1", 2);
INSERT INTO category_option_value_info (id, option_id, option_value, category_id) VALUES (2, 1001, "选项值2", 2);

-- ----------------------------
-- 插入渠道与sku关联关系
-- ----------------------------
INSERT INTO sku_sales_channel_info  (id, product_id, sku_id, sales_channel_id, price, inventory, available_inventory) VALUES (1, 1002, 2, 3, 150, 0, -1);

-- ----------------------------
-- 插入商品套餐信息表
-- ----------------------------
INSERT INTO product_set_info (id, master_product_id, master_sku_id, product_id, sku_id, sku_count) VALUES (1, 1001, 1, 1002, 2, 10)

