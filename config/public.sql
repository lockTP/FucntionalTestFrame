-- ----------------------------
-- 插入region
-- ----------------------------
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (110000, "北京市", 0, 010, 1, "省");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (110100, "北京市", 110000, 010, 2, "市");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (110101, "东城区", 110100, 010, 3, "区");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (110102, "西城区", 110100, 010, 3, "区");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (130000, "河北省", 0, 010, 1, "省");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (130100, "石家庄市", 130000, 0311, 3, "市");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (130102, "长安区", 130100, 0311, 3, "区");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (130103, "桥东区", 130100, 0311, 3, "区")