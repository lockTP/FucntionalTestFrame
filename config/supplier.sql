-- ----------------------------
-- 插入供应商
-- ----------------------------
INSERT INTO supplier_info (id, supplier_name, province_id, city_id, district_id, address, phone, contact_person, account_open_bank, taxpayer_identify_num, company_address, company_telephone_num, settlement_period, company_province_id, company_city_id, company_district_id) VALUES (1, "SIT供应商", 440000, 440300, 440304, "SIT供应商售后地址", "12345678912", "SIT供应商联系人", 1, "123456", "SIT供应商公司地址", "98765432112", 1, 440000, 440300, 440304);
INSERT INTO supplier_info (id, supplier_name, province_id, city_id, district_id, address, phone, contact_person, account_open_bank, taxpayer_identify_num, company_address, company_telephone_num, settlement_period, company_province_id, company_city_id, company_district_id) VALUES (2, "SIT供应商2", 440000, 440300, 440304, "SIT供应商售后地址", "12345678912", "SIT供应商联系人", 1, "123456", "SIT供应商公司地址", "98765432112", 0, 440000, 440300, 440304);
-- ----------------------------
-- 插入地区
-- ----------------------------
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (440000, "广东省", 0, "", 1, "省");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (440300, "深圳市", 440000, "0755", 2, "市");
INSERT INTO region (id, name, parent_id, code, level, level_name) VALUES (440304, "福田区", 440300, "0755", 3, "区")
