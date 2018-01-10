/*
Navicat MySQL Data Transfer

Source Server         : grpc-db
Source Server Version : 50509
Source Host           : 192.168.1.204:3306
Source Database       : d_aura_jike

Target Server Type    : MYSQL
Target Server Version : 50509
File Encoding         : 65001

Date: 2017-08-11 16:10:15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(20) NOT NULL DEFAULT '' COMMENT '密码',
  `create_time` varchar(20) NOT NULL DEFAULT '' COMMENT '创建日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) NOT NULL DEFAULT '' COMMENT '标题',
  `author` varchar(20) NOT NULL DEFAULT '' COMMENT '作者',
  `content` text NOT NULL COMMENT '内容',
  `point` varchar(20) NOT NULL DEFAULT '' COMMENT '点赞',
  `create_time` varchar(20) NOT NULL DEFAULT '' COMMENT '创建日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for attr_info
-- ----------------------------
DROP TABLE IF EXISTS `attr_info`;
CREATE TABLE `attr_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `attr_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `attr_style` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '属性的值输入风格 0:文本输入, 1：单选, 2：多选',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1010 DEFAULT CHARSET=utf8 COMMENT='属性定义表';

-- ----------------------------
-- Table structure for attr_type
-- ----------------------------
DROP TABLE IF EXISTS `attr_type`;
CREATE TABLE `attr_type` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr type id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1009 DEFAULT CHARSET=utf8 COMMENT='属性类型定义表';

-- ----------------------------
-- Table structure for attr_value_info
-- ----------------------------
DROP TABLE IF EXISTS `attr_value_info`;
CREATE TABLE `attr_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr value id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `attr_value` text NOT NULL COMMENT '值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_attr_id` (`attr_id`)
) ENGINE=InnoDB AUTO_INCREMENT=289 DEFAULT CHARSET=utf8 COMMENT='公共属性值表';

-- ----------------------------
-- Table structure for category_attr_info
-- ----------------------------
DROP TABLE IF EXISTS `category_attr_info`;
CREATE TABLE `category_attr_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `attr_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `category_id` int(10) unsigned NOT NULL COMMENT '分类id',
  `attr_type_id` int(10) unsigned NOT NULL COMMENT '属性类型id',
  `attr_mandatory` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '必须输入: 0:非必须, 1：必须',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=697 DEFAULT CHARSET=utf8 COMMENT='分类属性定义表';

-- ----------------------------
-- Table structure for category_attr_value_info
-- ----------------------------
DROP TABLE IF EXISTS `category_attr_value_info`;
CREATE TABLE `category_attr_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'attr value id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'attr id',
  `category_id` int(10) unsigned NOT NULL COMMENT '分类id',
  `attr_value` text NOT NULL COMMENT '值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=437 DEFAULT CHARSET=utf8 COMMENT='分类属性值表';

-- ----------------------------
-- Table structure for category_info
-- ----------------------------
DROP TABLE IF EXISTS `category_info`;
CREATE TABLE `category_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `category_name` varchar(60) NOT NULL DEFAULT '' COMMENT '名称',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=311 DEFAULT CHARSET=utf8 COMMENT='分类表';

-- ----------------------------
-- Table structure for category_option_info
-- ----------------------------
DROP TABLE IF EXISTS `category_option_info`;
CREATE TABLE `category_option_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '选项 id',
  `option_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `category_id` int(10) unsigned NOT NULL COMMENT '分类id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=222 DEFAULT CHARSET=utf8 COMMENT='分类选项定义表';

-- ----------------------------
-- Table structure for category_option_value_info
-- ----------------------------
DROP TABLE IF EXISTS `category_option_value_info`;
CREATE TABLE `category_option_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '选项value id',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '选项 id',
  `category_id` int(10) unsigned NOT NULL COMMENT '分类id',
  `option_value` text NOT NULL COMMENT '值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=582 DEFAULT CHARSET=utf8 COMMENT='分类选项值表';

-- ----------------------------
-- Table structure for channel_id_info
-- ----------------------------
DROP TABLE IF EXISTS `channel_id_info`;
CREATE TABLE `channel_id_info` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'channel id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1004 DEFAULT CHARSET=utf8 COMMENT='销售渠道定义表';

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author` varchar(20) NOT NULL DEFAULT '' COMMENT '作者',
  `content` varchar(20) NOT NULL DEFAULT '' COMMENT '内容',
  `articleId` int(11) NOT NULL COMMENT '文章id',
  `commentId` int(11) NOT NULL COMMENT '评论id',
  `create_time` varchar(20) NOT NULL DEFAULT '' COMMENT '创建日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for commment_info
-- ----------------------------
DROP TABLE IF EXISTS `commment_info`;
CREATE TABLE `commment_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `father_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'Sku id',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `context` varchar(0) NOT NULL DEFAULT '' COMMENT '名称',
  `grade_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '评论等级',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_sku_id` (`sku_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品评论表';

-- ----------------------------
-- Table structure for coupon_assign
-- ----------------------------
DROP TABLE IF EXISTS `coupon_assign`;
CREATE TABLE `coupon_assign` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'item id',
  `promotion_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '促销规则_id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '客户_id',
  `client_ip` int(10) unsigned DEFAULT '0' COMMENT '客户的Ip',
  `created_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发生的时间',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='客户抢到的秒杀/优惠券信息';

-- ----------------------------
-- Table structure for coupons
-- ----------------------------
DROP TABLE IF EXISTS `coupons`;
CREATE TABLE `coupons` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '优惠规则_id',
  `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠券类型',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '促销描述',
  `total_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠、优惠券的总数',
  `distribution_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠、优惠券已分配的数量',
  `begin_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '生效起始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '生效结束时间',
  `limit_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '限制金额，单位分',
  `min_discount_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最小优惠的金额，单位分',
  `max_discount_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最大优惠的金额，单位分',
  `min_discount_rate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最大优惠的折扣，0-100',
  `max_discount_rate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最大优惠的折扣，0-100',
  `user_scope` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:不限制用户',
  `coupons_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 有效 1:失效',
  `maximum_per_user` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '每人限制数',
  `return_flag` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:不可退 1:可退',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='优惠券规则表';

-- ----------------------------
-- Table structure for coupons_product_scope
-- ----------------------------
DROP TABLE IF EXISTS `coupons_product_scope`;
CREATE TABLE `coupons_product_scope` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `coupons_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '优惠规则_id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'product_id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='优惠商品范围表';

-- ----------------------------
-- Table structure for coupons_sn
-- ----------------------------
DROP TABLE IF EXISTS `coupons_sn`;
CREATE TABLE `coupons_sn` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '优惠码_id',
  `coupons_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '优惠规则_id',
  `coupons_code` varchar(63) NOT NULL DEFAULT '' COMMENT '优惠码',
  `discount_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '面值金额，单位分',
  `discount_rate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠的折扣，0-100',
  `distribution_count` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未领取 >=1:已经领取',
  `use_count` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未使用 >=1:已经使用',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='优惠券表';

-- ----------------------------
-- Table structure for dataoption
-- ----------------------------
DROP TABLE IF EXISTS `dataoption`;
CREATE TABLE `dataoption` (
  `s` varchar(100) NOT NULL DEFAULT '',
  `a` varchar(100) NOT NULL DEFAULT '',
  `v` varchar(100) NOT NULL DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for datasku
-- ----------------------------
DROP TABLE IF EXISTS `datasku`;
CREATE TABLE `datasku` (
  `id` varchar(100) NOT NULL DEFAULT '',
  `title` varchar(100) NOT NULL,
  `c2` varchar(100) NOT NULL,
  `c1` varchar(100) NOT NULL,
  `s` varchar(100) NOT NULL,
  `v1` varchar(100) NOT NULL,
  `n1` varchar(100) NOT NULL,
  `v2` varchar(100) NOT NULL,
  `n2` varchar(100) NOT NULL,
  `v3` varchar(100) NOT NULL,
  `n3` varchar(100) NOT NULL,
  `v4` varchar(100) NOT NULL,
  `n4` varchar(100) NOT NULL,
  `v5` varchar(100) NOT NULL,
  `n5` varchar(100) NOT NULL,
  `v6` varchar(100) NOT NULL,
  `n6` varchar(100) NOT NULL,
  `v7` varchar(100) NOT NULL,
  `n7` varchar(100) NOT NULL,
  `v8` varchar(100) NOT NULL,
  `n8` varchar(100) NOT NULL,
  `v9` varchar(100) NOT NULL,
  `n9` varchar(100) NOT NULL,
  `v10` varchar(100) NOT NULL,
  `n10` varchar(100) NOT NULL,
  `v11` varchar(100) NOT NULL,
  `n11` varchar(100) NOT NULL,
  `v12` varchar(100) NOT NULL,
  `n12` varchar(100) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for delivery_type
-- ----------------------------
DROP TABLE IF EXISTS `delivery_type`;
CREATE TABLE `delivery_type` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'delivery type id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='发货类型定义表';

-- ----------------------------
-- Table structure for ecoupon_info
-- ----------------------------
DROP TABLE IF EXISTS `ecoupon_info`;
CREATE TABLE `ecoupon_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `promotion_id` int(10) unsigned NOT NULL COMMENT '促销规则_id',
  `eCoupon_sn` varchar(63) NOT NULL COMMENT '优惠券兑换码',
  `channel_id` int(10) unsigned NOT NULL COMMENT '渠道_id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '客户_id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='兑换码优惠券息表';

-- ----------------------------
-- Table structure for enum_data_info
-- ----------------------------
DROP TABLE IF EXISTS `enum_data_info`;
CREATE TABLE `enum_data_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `field_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '字段id',
  `field_desc` varchar(63) NOT NULL DEFAULT '' COMMENT '字段描述',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_field_id` (`field_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='系统枚举字段定义表';

-- ----------------------------
-- Table structure for enum_data_value_info
-- ----------------------------
DROP TABLE IF EXISTS `enum_data_value_info`;
CREATE TABLE `enum_data_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `field_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '字段id',
  `field_value_desc` varchar(63) NOT NULL DEFAULT '' COMMENT '字段枚举值描述',
  `field_value` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '字段枚举值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_field_id` (`field_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8 COMMENT='系统枚举值表';

-- ----------------------------
-- Table structure for exchange_code_info
-- ----------------------------
DROP TABLE IF EXISTS `exchange_code_info`;
CREATE TABLE `exchange_code_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `supplier_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '供应商id',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'product_id',
  `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '0:自定义兑换码，1：第三方采购码',
  `batch_num` int(11) NOT NULL DEFAULT '0' COMMENT '批次号',
  `begin_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '兑换开始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '兑换结束时间',
  `total_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总数量',
  `allocate_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '已分配数量',
  `exchanged_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '已兑换数量',
  `freeze_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '已冻结数量',
  `invalid_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '已失效数量',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='兑换码汇总信息表';

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'skuid',
  `create_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时的价格',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='收藏表';

-- ----------------------------
-- Table structure for front_category_info
-- ----------------------------
DROP TABLE IF EXISTS `front_category_info`;
CREATE TABLE `front_category_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `category_name` varchar(60) NOT NULL DEFAULT '' COMMENT '名称',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '缩略图',
  `sort_order` int(3) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=246 DEFAULT CHARSET=utf8 COMMENT='前端分类表';

-- ----------------------------
-- Table structure for front_category_mapping
-- ----------------------------
DROP TABLE IF EXISTS `front_category_mapping`;
CREATE TABLE `front_category_mapping` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `front_category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前端类目id',
  `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前端分类类型: 0:关联后端分类, 1:关联活动id',
  `related_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '关联对象id, 例如：后端分类id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8 COMMENT='前端分类表映射表';

-- ----------------------------
-- Table structure for interesting
-- ----------------------------
DROP TABLE IF EXISTS `interesting`;
CREATE TABLE `interesting` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户账号',
  `interest_id` int(10) unsigned DEFAULT '0' COMMENT '兴趣id=目录id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='兴趣表';

-- ----------------------------
-- Table structure for operation_history
-- ----------------------------
DROP TABLE IF EXISTS `operation_history`;
CREATE TABLE `operation_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `account_name` varchar(30) NOT NULL DEFAULT '' COMMENT '账号名，只支持英文数字',
  `api_path` varchar(255) NOT NULL DEFAULT '' COMMENT 'api路径',
  `operation` text NOT NULL COMMENT '操作日志',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_account_name` (`account_name`),
  KEY `idx_api_path` (`api_path`)
) ENGINE=InnoDB AUTO_INCREMENT=321 DEFAULT CHARSET=utf8 COMMENT='操作日志表';

-- ----------------------------
-- Table structure for option_info
-- ----------------------------
DROP TABLE IF EXISTS `option_info`;
CREATE TABLE `option_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `option_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `ind_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8 COMMENT='option定义表';

-- ----------------------------
-- Table structure for option_value_info
-- ----------------------------
DROP TABLE IF EXISTS `option_value_info`;
CREATE TABLE `option_value_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'option value id',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '名称',
  `option_value` text NOT NULL COMMENT '值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_option_id` (`option_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1008 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for order_delivery_result_info
-- ----------------------------
DROP TABLE IF EXISTS `order_delivery_result_info`;
CREATE TABLE `order_delivery_result_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `delivery_result_info` text NOT NULL COMMENT '发货结果信息json串，由业务自行设置和解析',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8 COMMENT='订单发货结果信息';

-- ----------------------------
-- Table structure for order_exchange_code
-- ----------------------------
DROP TABLE IF EXISTS `order_exchange_code`;
CREATE TABLE `order_exchange_code` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'product_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `exchange_code` varchar(64) NOT NULL DEFAULT '' COMMENT '兑换码',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_exchange_code` (`exchange_code`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单中的虚拟商品兑换码';

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '客户_id',
  `origin_user_id` bigint(20) NOT NULL DEFAULT '0',
  `supplier_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '供应商id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '收货人姓名',
  `country_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收货人地址----国家',
  `province_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收货人地址----省',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收货人地址----市',
  `district_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收货人地址----区',
  `address` varchar(100) NOT NULL DEFAULT '' COMMENT '收货人地址----详细地址',
  `postcode` varchar(10) NOT NULL DEFAULT '' COMMENT '收货人地址----邮编',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '收货人地址----电话',
  `total_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总金额单位分',
  `total_score` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总积分',
  `total_youhui_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠金额单位分',
  `score_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分抵扣金额单位分',
  `epay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '余额、现金券、卖座卡抵扣金额单位分',
  `external_pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '外部支付/第三方支付金额单位分',
  `end_pay_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付截止时间',
  `pay_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付时间',
  `order_status` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态0:等待支付，1：支付成功（兼容保留不再使用），2：用户取消，3：订单失效，4：已发货，5：已退费， 6：发货中 ,7: 支付中，8:等待发货,9:发货失败,10:退费中,11:确认收货 12: 售后中, 13:交易完成',
  `pay_status` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态',
  `delivery_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发货时间',
  `is_from_cart` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否来源购物车',
  `channel_id` int(11) NOT NULL DEFAULT '0' COMMENT '终端类型',
  `sales_channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品渠道信息',
  `virtual_product` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否虚拟商品0：普通商品 1：虚拟商品',
  `auto_delivery` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否自动发货0：不自动发货 1：自动发货',
  `delivery_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发货类型',
  `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '֧�����ͣ�0:���ֽ�1��֧������2��΢��..., �ο�pay_type ��',
  `express_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '快递费用单位分',
  `express_company_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '快递公司_id',
  `express_company_name` varchar(100) NOT NULL DEFAULT '' COMMENT '快递公司名称',
  `express_id` varchar(30) NOT NULL DEFAULT '' COMMENT '快递单号',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_order_id` (`order_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_end_pay_time` (`end_pay_time`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_pay_time` (`pay_time`),
  KEY `idx_name` (`name`),
  KEY `idx_phone` (`phone`),
  KEY `idx_order_status` (`order_status`),
  KEY `idx_origin_user_id` (`origin_user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16989 DEFAULT CHARSET=utf8 COMMENT='订单表';

-- ----------------------------
-- Table structure for order_product_attr
-- ----------------------------
DROP TABLE IF EXISTS `order_product_attr`;
CREATE TABLE `order_product_attr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'id',
  `attr_type_id` int(10) unsigned NOT NULL COMMENT '属性类型id',
  `attr_value` text NOT NULL COMMENT '属性值',
  `attr_name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_product_type_id` (`order_id`,`product_id`,`attr_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28565 DEFAULT CHARSET=utf8 COMMENT='订单商品属性值表';

-- ----------------------------
-- Table structure for order_product_ext_info
-- ----------------------------
DROP TABLE IF EXISTS `order_product_ext_info`;
CREATE TABLE `order_product_ext_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'product_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `product_ext_info` text NOT NULL COMMENT '商扩展信息json串，发货附加信息等由业务自行设置和解析',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_product_sku` (`order_id`,`product_id`,`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=889 DEFAULT CHARSET=utf8 COMMENT='订单商品扩展信息';

-- ----------------------------
-- Table structure for order_sku
-- ----------------------------
DROP TABLE IF EXISTS `order_sku`;
CREATE TABLE `order_sku` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'product_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'product Name',
  `sku_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku Name',
  `sku_code` varchar(255) NOT NULL DEFAULT '' COMMENT '货号',
  `sku_image` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku 缩略图',
  `sku_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku数量',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku价格',
  `final_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku交易价格',
  `cost_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '成本价格',
  `score` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku积分/麦点',
  `product_source` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品渠道信息 0：普通商品  1：会员商品',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17251 DEFAULT CHARSET=utf8 COMMENT='订单中的sku信息';

-- ----------------------------
-- Table structure for order_sku_epay
-- ----------------------------
DROP TABLE IF EXISTS `order_sku_epay`;
CREATE TABLE `order_sku_epay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `epay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '0:优惠,  1:积分, 2:余额, 3:现金券, 4:卖座卡',
  `epay_id` varchar(255) NOT NULL DEFAULT '' COMMENT '优惠_id, 商品的（非现金）券_id',
  `epay_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分数量,  点卡数量等',
  `epay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku使用的优惠券等的金额单位分',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12201 DEFAULT CHARSET=utf8 COMMENT='订单中的sku的优惠以及使用的非现金支付信息';

-- ----------------------------
-- Table structure for pay
-- ----------------------------
DROP TABLE IF EXISTS `pay`;
CREATE TABLE `pay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
  `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单_id',
  `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信..., 参考pay_type 表',
  `pay_account` varchar(100) NOT NULL DEFAULT '' COMMENT '支付账号, 第三方的openid',
  `pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额单元分',
  `pay_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付时间',
  `pay_seq` varchar(100) NOT NULL DEFAULT '' COMMENT '支付流水号',
  `pay_status` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态, 0:未支付, 1:已成功支付, 2:支付失败',
  `notify_status` int(11) NOT NULL DEFAULT '0' COMMENT '支付状态, 0:未通知, 1:已成通知 ',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`pay_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1841 DEFAULT CHARSET=utf8 COMMENT='支付信息表';

-- ----------------------------
-- Table structure for pay_action
-- ----------------------------
DROP TABLE IF EXISTS `pay_action`;
CREATE TABLE `pay_action` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
  `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单_id',
  `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信, ... 参考pay_type表',
  `pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额单元分',
  `pay_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '支付时间',
  `action` int(11) NOT NULL DEFAULT '0' COMMENT '0:None, 1:支付, 2:关闭支付 99:其他',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`pay_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3569 DEFAULT CHARSET=utf8 COMMENT='支付记录表';

-- ----------------------------
-- Table structure for pay_order_info
-- ----------------------------
DROP TABLE IF EXISTS `pay_order_info`;
CREATE TABLE `pay_order_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '支付订单_id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额单元分',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_pay_order_id` (`pay_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='支付订单关系表';

-- ----------------------------
-- Table structure for pay_type
-- ----------------------------
DROP TABLE IF EXISTS `pay_type`;
CREATE TABLE `pay_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信',
  `name` varchar(50) DEFAULT '' COMMENT '支付名称',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='支付类型表';

-- ----------------------------
-- Table structure for product_attr
-- ----------------------------
DROP TABLE IF EXISTS `product_attr`;
CREATE TABLE `product_attr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `attr_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'id',
  `attr_value` text NOT NULL COMMENT '属性值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_product_type_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9653 DEFAULT CHARSET=utf8 COMMENT='商品属性表';

-- ----------------------------
-- Table structure for product_attr_type
-- ----------------------------
DROP TABLE IF EXISTS `product_attr_type`;
CREATE TABLE `product_attr_type` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `attr_type_id` int(10) unsigned NOT NULL COMMENT '属性类型id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=699 DEFAULT CHARSET=utf8 COMMENT='商品属性类型表';

-- ----------------------------
-- Table structure for product_category
-- ----------------------------
DROP TABLE IF EXISTS `product_category`;
CREATE TABLE `product_category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `category_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_product_category_category_id` (`category_id`),
  KEY `idx_product_category_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1283 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for product_exchange_code
-- ----------------------------
DROP TABLE IF EXISTS `product_exchange_code`;
CREATE TABLE `product_exchange_code` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `supplier_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '供应商id',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'product_id',
  `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '0:自定义兑换码，1：第三方采购码',
  `exchange_code` varchar(64) NOT NULL DEFAULT '' COMMENT '兑换码',
  `flag` int(11) NOT NULL DEFAULT '0' COMMENT '0:未使用，1：已分配，2：已兑换 3：已冻结 4：已失效',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `begin_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '兑换开始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '兑换结束时间',
  `exchange_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '兑换时间',
  `batch_num` int(11) NOT NULL DEFAULT '0' COMMENT '批次号',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_sku_id` (`sku_id`),
  KEY `aa` (`exchange_code`)
) ENGINE=InnoDB AUTO_INCREMENT=2255 DEFAULT CHARSET=utf8 COMMENT='订单商品兑换码表';

-- ----------------------------
-- Table structure for product_id_change
-- ----------------------------
DROP TABLE IF EXISTS `product_id_change`;
CREATE TABLE `product_id_change` (
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'product_id',
  `goodsId` varchar(12) NOT NULL DEFAULT '' COMMENT '商品id',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品id转换表';

-- ----------------------------
-- Table structure for product_image
-- ----------------------------
DROP TABLE IF EXISTS `product_image`;
CREATE TABLE `product_image` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `image` varchar(255) DEFAULT NULL,
  `sort_order` int(3) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_image` (`product_id`,`image`),
  KEY `product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for product_info
-- ----------------------------
DROP TABLE IF EXISTS `product_info`;
CREATE TABLE `product_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `master_name` varchar(255) NOT NULL DEFAULT '' COMMENT '主名称',
  `slave_name` varchar(255) NOT NULL DEFAULT '' COMMENT '副名称',
  `supplier_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '供应商id',
  `default_sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '默认sku id',
  `release_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  `product_status` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '状态：0-下架、1-上架',
  `product_source` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品渠道信息 0：普通商品 1：会员商品',
  `min_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku最低销售价格',
  `max_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku最高销售数量',
  `display_sales_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '各个sku销量的和，实际销售数量*9， 管理系统可以更改',
  `sort_sales_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最近30天各个sku销量的和，实际销售数量*9，管理系统不可更改',
  `score` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `virtual_product` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否虚拟商品0：普通商品 1：虚拟商品',
  `auto_delivery` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否自动发货0：不自动发货 1：自动发货',
  `delivery_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '发货类型',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_supplier_id` (`supplier_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1781 DEFAULT CHARSET=utf8 COMMENT='商品表';

-- ----------------------------
-- Table structure for product_source_info
-- ----------------------------
DROP TABLE IF EXISTS `product_source_info`;
CREATE TABLE `product_source_info` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'product_source id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='商品渠道定义表';

-- ----------------------------
-- Table structure for promotion
-- ----------------------------
DROP TABLE IF EXISTS `promotion`;
CREATE TABLE `promotion` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '促销规则_id',
  `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '促销类型',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '促销描述',
  `total_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠、优惠券的总数',
  `distribution_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠、优惠券已分配的数量',
  `begin_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '生效起始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '生效结束时间',
  `is_active` tinyint(1) NOT NULL DEFAULT '0' COMMENT '激活标志',
  `minimum_order_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单满额减最低金额，单位分',
  `minimum_product_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品满件减最低数量',
  `discount_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠的金额，单位分',
  `discount_rate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠的折扣，0-100',
  `reward_ecoupon_promotion_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '赠送优惠券的规则_id',
  `isIterate` tinyint(1) NOT NULL DEFAULT '0' COMMENT '优惠迭代标志',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='促销规则表';

-- ----------------------------
-- Table structure for promotion_assoicated_product
-- ----------------------------
DROP TABLE IF EXISTS `promotion_assoicated_product`;
CREATE TABLE `promotion_assoicated_product` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `promotion_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '组合销售促销规则_id',
  `union_products` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品_id',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品价格',
  `rate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '折扣0-100',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='满优惠中促销商品信息';

-- ----------------------------
-- Table structure for promotion_price_detail
-- ----------------------------
DROP TABLE IF EXISTS `promotion_price_detail`;
CREATE TABLE `promotion_price_detail` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `promotion_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '促销规则_id',
  `begin_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售数量范围起始值',
  `end_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售数量范围结束值',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售数量范围内的价格',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='单品优惠价格随销售数量变化';

-- ----------------------------
-- Table structure for promotion_restrict
-- ----------------------------
DROP TABLE IF EXISTS `promotion_restrict`;
CREATE TABLE `promotion_restrict` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `promotion_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'skuid',
  `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分类id',
  `user_grade` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户级别',
  `zone_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '地区id',
  `is_include` tinyint(1) NOT NULL DEFAULT '0' COMMENT '允许/禁止',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='促销规则中的商品信息';

-- ----------------------------
-- Table structure for refund
-- ----------------------------
DROP TABLE IF EXISTS `refund`;
CREATE TABLE `refund` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
  `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单 id',
  `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信, ... 参考pay_type表',
  `refund_id` varchar(32) NOT NULL DEFAULT '' COMMENT '退费 id',
  `refund_seq` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方退费流水 id',
  `refund_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '退费金额',
  `refund_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '退费时间',
  `refund_status` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '0:未退费, 1:已成功退费, 2:退费失败',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`pay_order_id`),
  KEY `idx_refund_id` (`refund_id`),
  KEY `idx_refund_seq` (`refund_seq`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8 COMMENT='退费信息表';

-- ----------------------------
-- Table structure for refund_action
-- ----------------------------
DROP TABLE IF EXISTS `refund_action`;
CREATE TABLE `refund_action` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'app id',
  `pay_order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型，0:非现金，1：支付宝，2：微信, ... 参考pay_type表',
  `refund_id` varchar(32) NOT NULL DEFAULT '' COMMENT '退费 id',
  `refund_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '退费金额',
  `refund_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '退费时间',
  `action` int(11) NOT NULL DEFAULT '0' COMMENT '0:退费 99:其他',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`pay_order_id`),
  KEY `idx_refund_id` (`refund_id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8 COMMENT='退费记录表';

-- ----------------------------
-- Table structure for region
-- ----------------------------
DROP TABLE IF EXISTS `region`;
CREATE TABLE `region` (
  `id` int(6) unsigned NOT NULL,
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '省城区名',
  `parent_id` int(6) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `code` varchar(5) NOT NULL DEFAULT '' COMMENT '省城区码',
  `level` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '级别标识 1:省 2: 市 3:区',
  `level_name` char(3) NOT NULL DEFAULT '' COMMENT '级别名',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='区域信息表';

-- ----------------------------
-- Table structure for sales_channel_info
-- ----------------------------
DROP TABLE IF EXISTS `sales_channel_info`;
CREATE TABLE `sales_channel_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '商品购买渠道',
  `channel_names` varchar(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `channel_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '渠道描述',
  `channel_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '渠道状态0是正常，1是禁用',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8 COMMENT='渠道信息表';

-- ----------------------------
-- Table structure for score_mall_goods_id
-- ----------------------------
DROP TABLE IF EXISTS `score_mall_goods_id`;
CREATE TABLE `score_mall_goods_id` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goods_id` varchar(12) NOT NULL DEFAULT '' COMMENT '商品ID',
  PRIMARY KEY (`id`),
  KEY `idx_goodsId` (`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=328 DEFAULT CHARSET=utf8 COMMENT='商品id与积分商城商品goods id关系表';

-- ----------------------------
-- Table structure for service_history
-- ----------------------------
DROP TABLE IF EXISTS `service_history`;
CREATE TABLE `service_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `service_id` varchar(32) NOT NULL DEFAULT '' COMMENT '服务单号 id',
  `service_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '服务状态',
  `operator` varchar(255) NOT NULL DEFAULT '' COMMENT '操作员信息',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '处理备注, 包括买家发货快递信息',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='售后处理历史表';

-- ----------------------------
-- Table structure for service_image
-- ----------------------------
DROP TABLE IF EXISTS `service_image`;
CREATE TABLE `service_image` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `service_id` varchar(32) NOT NULL DEFAULT '' COMMENT '服务单号 id',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '用户申请售后的举证图片',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='申请售后的image表';

-- ----------------------------
-- Table structure for service_info
-- ----------------------------
DROP TABLE IF EXISTS `service_info`;
CREATE TABLE `service_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `service_id` varchar(32) NOT NULL DEFAULT '' COMMENT '服务单号 id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `pay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型, 参考pay_type表',
  `service_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '售后类型, 0:退费, 1:退货, 2:换货',
  `total_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `external_pay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '外部支付/第三方支付总金额单位分',
  `refund_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '退费金额',
  `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '售后原因',
  `source` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '申请来源, 0:客户,1:客服人员,2:系统自动',
  `service_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '服务状态, 0:已申请,1:待买家发货,2:待商家收货,3:待商家发货,4:售后处理中,5:售后关闭,6:已退款,7:已换货',
  `refund_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '申请时间',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_service_id` (`service_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8 COMMENT='申请售后信息表';

-- ----------------------------
-- Table structure for service_sku
-- ----------------------------
DROP TABLE IF EXISTS `service_sku`;
CREATE TABLE `service_sku` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `service_id` varchar(32) NOT NULL DEFAULT '' COMMENT '服务单号 id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `refund_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '退货后的退费金额',
  `exchange_sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '换货后的sku_id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8 COMMENT='申请售后的sku退换信息表';

-- ----------------------------
-- Table structure for service_sku_epay
-- ----------------------------
DROP TABLE IF EXISTS `service_sku_epay`;
CREATE TABLE `service_sku_epay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `order_id` varchar(32) NOT NULL DEFAULT '' COMMENT '订单_id',
  `service_id` varchar(32) NOT NULL DEFAULT '' COMMENT '服务单号 id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `epay_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '0:优惠,  1:积分, 2:余额, 3:现金券, 4:卖座卡',
  `epay_id` varchar(255) NOT NULL DEFAULT '' COMMENT '优惠_id, 商品的（非现金）券_id',
  `epay_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分数量,  点卡数量等',
  `epay_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku使用的优惠券等的金额单位分',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_service_id` (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COMMENT='售后单非现金资源退还信息';

-- ----------------------------
-- Table structure for shopping_cart
-- ----------------------------
DROP TABLE IF EXISTS `shopping_cart`;
CREATE TABLE `shopping_cart` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '客户_id',
  `sku_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku _id',
  `product_name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名字',
  `sku_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku缩略图',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku缩略图',
  `sku_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku 数量',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品加入购物车的价格',
  `score` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品加入购物车的积分价格',
  `relation_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '关联商品_id/换购或赠送等',
  `promotion_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '被采用的关联商品的优惠规则_id',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=30674 DEFAULT CHARSET=utf8 COMMENT='购物车';

-- ----------------------------
-- Table structure for sku_image
-- ----------------------------
DROP TABLE IF EXISTS `sku_image`;
CREATE TABLE `sku_image` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sku_id` bigint(20) unsigned NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `sort_order` int(3) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9520 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for sku_info
-- ----------------------------
DROP TABLE IF EXISTS `sku_info`;
CREATE TABLE `sku_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'skuid',
  `master_name` varchar(255) NOT NULL DEFAULT '' COMMENT '主名称',
  `slave_name` varchar(255) NOT NULL DEFAULT '' COMMENT '副名称',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '缩略图',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `market_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '市场价格',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售价格',
  `cost_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '成本价',
  `score` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `inventory` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '库存量',
  `available_inventory` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '可用库存',
  `sales_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售数量',
  `display_sales_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '显示销售数量',
  `sku_code` varchar(255) NOT NULL DEFAULT '' COMMENT '货号',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5248 DEFAULT CHARSET=utf8 COMMENT='SKU表';

-- ----------------------------
-- Table structure for sku_option
-- ----------------------------
DROP TABLE IF EXISTS `sku_option`;
CREATE TABLE `sku_option` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `option_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'id',
  `option_value` text NOT NULL COMMENT '属性值',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_sku_id` (`sku_id`),
  KEY `idx_option_id` (`option_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1377 DEFAULT CHARSET=utf8 COMMENT='sku option表';

-- ----------------------------
-- Table structure for sku_sales_channel_info
-- ----------------------------
DROP TABLE IF EXISTS `sku_sales_channel_info`;
CREATE TABLE `sku_sales_channel_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'product_id',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `sales_channel_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品购买渠道',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售价格',
  `inventory` int(10) NOT NULL DEFAULT '0' COMMENT '库存量',
  `available_inventory` int(10) NOT NULL DEFAULT '0' COMMENT '可用库存量',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_product_sku` (`product_id`,`sku_id`),
  KEY `idx_sales_channel_id` (`sales_channel_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1011 DEFAULT CHARSET=utf8 COMMENT='商品渠道定价表';

-- ----------------------------
-- Table structure for status_change
-- ----------------------------
DROP TABLE IF EXISTS `status_change`;
CREATE TABLE `status_change` (
  `order_status` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态,0:未支付，1：支付成功，2：用户取消，3：订单失效，4：已发货，5：已退费',
  `old_status` int(4) NOT NULL DEFAULT '1' COMMENT '-1 失败 1 待支付 2成功 3 发货中 4 订单失效  5 支付成功 发货失败  6退费  ',
  `old_payStatus` int(4) DEFAULT '0' COMMENT '支付状态 1 支付成功  2 回滚成功  3 回滚失败',
  PRIMARY KEY (`order_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单状态转换表';

-- ----------------------------
-- Table structure for supplier_info
-- ----------------------------
DROP TABLE IF EXISTS `supplier_info`;
CREATE TABLE `supplier_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '供应商id',
  `supplier_name` varchar(60) NOT NULL DEFAULT '' COMMENT '名称',
  `supplier_desc` varchar(1023) NOT NULL DEFAULT '' COMMENT '供应商描述',
  `grade_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '等级',
  `country_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '国家',
  `province_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '省份',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市',
  `district_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '地区/县',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '电话',
  `contact_person` varchar(30) NOT NULL DEFAULT '' COMMENT '联系人姓名',
  `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '供应商类型',
  `license` varchar(30) NOT NULL DEFAULT '' COMMENT '营业执照',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `bank_name` varchar(30) NOT NULL DEFAULT '' COMMENT '开户银行',
  `bank_account` varchar(50) NOT NULL DEFAULT '' COMMENT '银行账号',
  `account_name` varchar(30) NOT NULL DEFAULT '' COMMENT '账户名',
  `pay_method` varchar(30) NOT NULL DEFAULT '' COMMENT '结算方式',
  `pay_method_date` varchar(30) NOT NULL DEFAULT '' COMMENT '结算日期',
  `cooperation_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '合作时间',
  `supplier_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态（正常/禁用） 0:正常 1:禁用',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8 COMMENT='供应商信息表';

-- ----------------------------
-- Table structure for system_param
-- ----------------------------
DROP TABLE IF EXISTS `system_param`;
CREATE TABLE `system_param` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(32) NOT NULL COMMENT '参数名字',
  `value` varchar(255) NOT NULL COMMENT '参数值',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='系统静态、动态参数表';

-- ----------------------------
-- Table structure for user_address
-- ----------------------------
DROP TABLE IF EXISTS `user_address`;
CREATE TABLE `user_address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户账号',
  `country_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '国家',
  `province_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '省份',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市',
  `district_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '地区/县',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '详细地址',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '电话',
  `postcode` varchar(10) NOT NULL DEFAULT '' COMMENT '邮编',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '收件人名字',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否默认地址(1:默认地址)',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=550 DEFAULT CHARSET=utf8 COMMENT='地址表';

-- ----------------------------
-- Table structure for user_bind_history
-- ----------------------------
DROP TABLE IF EXISTS `user_bind_history`;
CREATE TABLE `user_bind_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT '第三方接入id',
  `open_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '第三方接入类型',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `sales_channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售渠道',
  `term_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置经度',
  `latitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置纬度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8 COMMENT='用户绑定第三方账号历史表';

-- ----------------------------
-- Table structure for user_change_password_history
-- ----------------------------
DROP TABLE IF EXISTS `user_change_password_history`;
CREATE TABLE `user_change_password_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `old_password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `sales_channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售渠道',
  `term_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置经度',
  `latitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置纬度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8 COMMENT='用户修改密码历史表';

-- ----------------------------
-- Table structure for user_change_pay_password_history
-- ----------------------------
DROP TABLE IF EXISTS `user_change_pay_password_history`;
CREATE TABLE `user_change_pay_password_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `old_pay_password` varchar(50) NOT NULL DEFAULT '' COMMENT '支付密码',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `sales_channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售渠道',
  `term_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置经度',
  `latitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置纬度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8 COMMENT='用户修改支付密码历史表';

-- ----------------------------
-- Table structure for user_coupons
-- ----------------------------
DROP TABLE IF EXISTS `user_coupons`;
CREATE TABLE `user_coupons` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `coupons_sn_id` bigint(20) unsigned NOT NULL COMMENT '优惠码_id',
  `coupons_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '优惠规则_id',
  `discount_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '面值金额，单位分',
  `discount_rate` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠的折扣，0-100',
  `use_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 未使用 1:已经使用',
  `begin_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '生效起始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '生效结束时间',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户优惠券表';

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户账号',
  `account_name` varchar(30) NOT NULL DEFAULT '' COMMENT '账号名，只支持英文数字',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `nick_name` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `pay_password` varchar(50) NOT NULL DEFAULT '' COMMENT '支付密码',
  `head_icon` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别',
  `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  `mail` varchar(50) NOT NULL DEFAULT '‘’' COMMENT '注册邮箱',
  `grade_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '等级',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `real_name` varchar(30) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `marriage` tinyint(1) NOT NULL DEFAULT '0' COMMENT '婚姻状况',
  `income` varchar(30) NOT NULL DEFAULT '' COMMENT '月收入',
  `identity` varchar(50) NOT NULL DEFAULT '' COMMENT '身份证',
  `education` tinyint(4) NOT NULL DEFAULT '0' COMMENT '教育',
  `career` smallint(6) NOT NULL DEFAULT '0' COMMENT '行业',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=139 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Table structure for user_login_history
-- ----------------------------
DROP TABLE IF EXISTS `user_login_history`;
CREATE TABLE `user_login_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `login_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：password，2:手机验证码  3:第三方登陆',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `sales_channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售渠道',
  `term_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置经度',
  `latitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置纬度',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3597 DEFAULT CHARSET=utf8 COMMENT='用户登陆历史表';

-- ----------------------------
-- Table structure for user_open_id_info
-- ----------------------------
DROP TABLE IF EXISTS `user_open_id_info`;
CREATE TABLE `user_open_id_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT '第三方接入id',
  `uu_id` varchar(50) NOT NULL DEFAULT '' COMMENT '第三方接入UU_id',
  `open_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '第三方登陆类型 0：自有账号(未使用) openid无效；1：微信，2：支付宝 3：微博 4:绑定手机号(openId就是mobile)',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8 COMMENT='用户第三方账号绑定信息表';

-- ----------------------------
-- Table structure for user_register_info
-- ----------------------------
DROP TABLE IF EXISTS `user_register_info`;
CREATE TABLE `user_register_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户账号',
  `client_ip` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '注册ip',
  `equipment_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端设备id',
  `channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '渠道 0:未知 1:运营  2：集客',
  `sales_channel_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销售渠道',
  `term_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0:未知，1：ios，2:android',
  `version` varchar(32) NOT NULL DEFAULT '' COMMENT '客户端版本号',
  `app_store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用商店名称, 0-maizuo, m360, AppStore,…',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市_id',
  `longitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置经度',
  `latitude` float NOT NULL DEFAULT '0' COMMENT 'gps 位置纬度',
  `carrier` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '运营商',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8 COMMENT='用户注册信息表';

-- ----------------------------
-- Procedure structure for loadSku
-- ----------------------------
DROP PROCEDURE IF EXISTS `loadSku`;
DELIMITER ;;
CREATE DEFINER=`sec`@`%` PROCEDURE `loadSku`()
BEGIN
  
		DECLARE rt VARCHAR(100) DEFAULT " ";

		DECLARE loopFlag TINYINT(1) DEFAULT 0;

		DECLARE vid VARCHAR(100) DEFAULT "";
		DECLARE vtitle VARCHAR(100) DEFAULT "";
		DECLARE vs VARCHAR(100) DEFAULT "";
		DECLARE vc1 VARCHAR(100) DEFAULT "";
		DECLARE vc2 VARCHAR(100) DEFAULT "";
		DECLARE vv1 VARCHAR(100) DEFAULT "";
		DECLARE vv2 VARCHAR(100) DEFAULT "";
		DECLARE vv3 VARCHAR(100) DEFAULT "";
		DECLARE vv4 VARCHAR(100) DEFAULT "";
		DECLARE vv5 VARCHAR(100) DEFAULT "";
		DECLARE vv6 VARCHAR(100) DEFAULT "";
		DECLARE vv7 VARCHAR(100) DEFAULT "";
		DECLARE vv8 VARCHAR(100) DEFAULT "";
		DECLARE vv9 VARCHAR(100) DEFAULT "";
		DECLARE vv10 VARCHAR(100) DEFAULT "";
		DECLARE vv11 VARCHAR(100) DEFAULT "";
		DECLARE vv12 VARCHAR(100) DEFAULT "";

		DECLARE skuCount INT DEFAULT 0;


		
		BEGIN
		DECLARE cur CURSOR FOR  SELECT id, title, s, c1, c2, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12 FROM dataSku ;

		DECLARE continue handler for not found set loopFlag = 1;

		open cur;

		TableLoop:
		LOOP
				FETCH cur INTO  vid, vtitle, vs, vc1, vc2, vv1, vv2, vv3, vv4, vv5, vv6, vv7, vv8, vv9, vv10, vv11, vv12;
				IF loopFlag = 1 THEN
						LEAVE TableLoop;
				END IF;

				set skuCount = 0;

				set @productId  = 0; 
				SELECT g.id  into @productId  FROM d_score_mall.goods g WHERE g.goodsId = vid;

				set @supplierId  = 0; 
				SELECT id  into @supplierId  FROM supplier_info  WHERE supplier_name = vs;		

				set @optionSetId  = 0; 
				SELECT id  into @optionSetId  FROM option_set_info WHERE option_set_name = vc1;

-- for V1
				set skuCount = skuCount + 1;
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv1;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId, @optionId, vv1, NOW());
				end if;


-- for V2
				if LENGTH(vv2) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv2;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*1 , @optionId, vv2, NOW());
				end if;
				end If;
				
-- for V3
				if LENGTH(vv3) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv3;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*2 , @optionId, vv3, NOW());
				end if;
				end If;

-- for V4
				if LENGTH(vv4) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv4;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*3 , @optionId, vv4, NOW());
				end if;
				end if;

-- for V5
				if LENGTH(vv5) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv5;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*4 , @optionId, vv5, NOW());
				end if;
				end If;

-- for V6
				if LENGTH(vv6) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv6;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*5 , @optionId, vv6, NOW());
				end if;
				end If;



-- for V7
				if LENGTH(vv7) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv7;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*6 , @optionId, vv7, NOW());
				end if;
				end If;

-- for V8
				if LENGTH(vv8) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv8;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*7 , @optionId, vv8, NOW());
				end if;
				end If;

-- for V9
				if LENGTH(vv9) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv9;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*8 , @optionId, vv9, NOW());
				end if;
				end If;

-- for V10
				if LENGTH(vv10) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv10;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*9 , @optionId, vv10, NOW());
				end if;
				end If;

-- for V11
				if LENGTH(vv11) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv11;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*10 , @optionId, vv11, NOW());
				end if;
				end If;

-- for V12
				if LENGTH(vv12) > 0 THEN
				set @optionName  = ""; 
				SELECT  a into @optionName  FROM dataOption WHERE s = vc1 and v =  vv12;
				set @optionId  = 0; 
				SELECT id  into @optionId  FROM option_info WHERE option_name = @optionName and option_set_id = @optionSetId;
				if @optionId > 0 then
				set skuCount = skuCount + 1;
				insert into sku_option (sku_id, option_id, option_value, created_at) value (@productId + 400*11 , @optionId, vv12, NOW());
				end if;
				end If;


				update product_info set display_sales_count = display_sales_count * skuCount where id = @productId;

				update product_info set supplier_id = @supplierId where id = @productId;

				set loopFlag = 0;
		END LOOP TableLoop;

		close cur;

		END; 




END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for loadSql
-- ----------------------------
DROP PROCEDURE IF EXISTS `loadSql`;
DELIMITER ;;
CREATE DEFINER=`sec`@`%` PROCEDURE `loadSql`()
BEGIN
  




-- 2. option part

insert into option_set_info (option_set_name, created_at)( select distinct s, now() from dataOption); 
insert into option_info (option_set_id, option_name, created_at)(  select o.id, t.a, now() from   option_set_info o,  ( select distinct s, a from dataOption d ) t  where o.option_set_name = t.s);
insert into option_value_info(option_id, option_value, created_at)
     (select t.id, d.v, now()  from dataOption d, 
            (select s.option_set_name, o.option_name, o.id from option_set_info s, option_info o where s.id = o.option_set_id) t 
            where d.s = t.option_set_name and d.a = t.option_name);

-- 3. category and supplier

insert into category_info (parent_id, category_name, sort_order, created_at) (select distinct 0, c2, 1, now() from dataSku);
insert into category_info (parent_id, category_name, sort_order, created_at) (select distinct c.id, c1, 100, now() from dataSku d, category_info c where d.c2 = c.category_name);

insert into supplier_info ( supplier_name, supplier_desc, created_at) 
    ( select distinct s, s, now()  from dataSku );

-- 4. product and sku

insert into product_info(id, default_sku_id, release_time, product_status,  master_name, slave_name, display_sales_count,min_price,max_price,score,created_at)
    (select g.id, g.id, saleStartTime, if(isShelves = 0 || flag = 0, 0,1),  goodsName, goodsTitle, soldCount,goodsPrice,goodsPrice,point,saleStartTime from d_score_mall.goods g, dataSku d where g.goodsId = d.id);

insert into product_attr(product_id, attr_id, attr_value, created_at)( select g.id, 1, goodsDetail, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id);

insert into product_category(product_id, category_id) 
(select g.id, c.id from d_score_mall.goods g, dataSku d, category_info c where g.goodsId = d.id and d.c1 = c.category_name and c.parent_id !=0);


insert into sku_info(id, master_name, image, product_id, market_price, price, score, inventory, sales_count, sku_code,created_at )
    ( select g.id, d.v1, goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n1, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id);
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*1 , d.v2,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n2, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v2) > 1);
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*2 , d.v3,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n3, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v3) > 1);
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*3 , d.v4,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n4, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v4) > 1);
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*4 , d.v5,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n5, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v5) > 1);
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*5 , d.v6,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n6, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v6) > 1);
                
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*6 , d.v7,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n7, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v7) >1);
                
                
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*7 , d.v8,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n8, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v8) > 0);                                                                
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*8 , d.v9,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n9, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v9) > 0);    
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*9 , d.v10,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n10, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v10) > 0);    
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*10 , d.v11,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n11, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v11) > 0);    
insert into sku_info(id, master_name,  image, product_id, market_price, price, score, inventory, sales_count, sku_code, created_at )
    ( select g.id + 400*11 , d.v12,  goodsIcon, g.id, goodsPrice, goodsPrice, point, totalInventory, soldCount, d.n12, now() 
                from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v12) > 0);    

           

-- v1;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id);

-- v2;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*1, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v2) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*1, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v2) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*1, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v2) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*1, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v2) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*1, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v2) > 1);

-- v3;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*2, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v3) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*2, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v3) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*2, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v3) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*2, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v3) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*2, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v3) > 1);

-- v4;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*3, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v4) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*3, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v4) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*3, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v4) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*3, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v4) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*3, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v4) > 1);

-- v5;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*4, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v5) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*4, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v5) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*4, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v5) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*4, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v5) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*4, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v5) > 1);

-- v6;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*5, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v6) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*5, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v6) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*5, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v6) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*5, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v6) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*5, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v6) > 1);

-- v7;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*6, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v7) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*6, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v7) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*6, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v7) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*6, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v7) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*6, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v7) > 1);

-- v8;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*7, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v8) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*7, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v8) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*7, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v8) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*7, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v8) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*7, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v8) > 1);

-- v9;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*8, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v9) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*8, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v9) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*8, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v9) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*8, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v9) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*8, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v9) > 1);

-- v10;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*9, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v10) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*9, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v10) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*9, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v10) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*9, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v10) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*9, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v10) > 1);

-- v11;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*10, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v11) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*10, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v11) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*10, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v11) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*10, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v11) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*10, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v11) > 1);

-- v12;
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*11, substring_index( substring_index(goodsPics, '|', 1), '|', -1), 1, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v12) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*11, substring_index( substring_index(goodsPics, '|', 2), '|', -1), 2, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v12) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*11, substring_index( substring_index(goodsPics, '|', 3), '|', -1), 3, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v12) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*11, substring_index( substring_index(goodsPics, '|', 4), '|', -1), 4, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v12) > 1);
insert ignore into sku_image (sku_id, image, sort_order, created_at ) (select g.id + 400*11, substring_index( substring_index(goodsPics, '|', 5), '|', -1), 5, now() from d_score_mall.goods g, dataSku d where g.goodsId = d.id and length(d.v12) > 1);


-- 5. some sepcial operation

update sku_info set display_sales_count = sales_count;
-- update category_info set status = 1 where id < 5;

update category_info set image = 'http://pic.maizuo.com/manager/5498d70c7a2097d9682bc2222e4f837f.png' where category_name = '3D眼镜' and parent_id !=0;
update category_info set image = 'http://pic.maizuo.com/manager/10d8c7eb0c6cab2512d83c5092f8e4ce.png' where category_name = '充电装备' and parent_id !=0;
update category_info set image = 'http://pic.maizuo.com/manager/09b6b58c85a04a659f33564cdee40251.png' where category_name = '手机配件' and parent_id !=0;
update category_info set image = 'http://pic.maizuo.com/manager/b3bb83a38650776e9c7663d9d1df7222.png' where category_name = '家居百货' and parent_id !=0;

update category_info set image = 'http://pic.maizuo.com/manager/330fd7294fb220a1180c8289c36f4087.png' where category_name = 'U盘数据线' and parent_id !=0;
update category_info set image = 'http://pic.maizuo.com/manager/3a1887c5473a29ce880754eaa5e4a6ca.png' where category_name = '服饰箱包' and parent_id !=0;
update category_info set image = 'http://pic.maizuo.com/manager/2344ba3341df2113bdc63fcd48a6d70f.png' where category_name = '模型玩具' and parent_id !=0;
update category_info set image = 'http://pic.maizuo.com/manager/b6fcb59334926886764c92ff4f9f8b0e.png' where category_name = '车载用品' and parent_id !=0;



-- 6. member product,  138 product and two new products
insert into product_info(id, default_sku_id, supplier_id, release_time, product_status, master_name, slave_name, display_sales_count,min_price,max_price,score,created_at)
    (select id + 1000, default_sku_id + 1000, supplier_id, release_time, product_status, master_name, slave_name, display_sales_count,min_price,max_price,score,created_at from product_info  where id = 138);

insert into product_attr(product_id, attr_id, attr_value, created_at)
( select product_id + 1000, attr_id, attr_value, created_at from product_attr where product_id = 138);

insert into product_category(product_id, category_id) 
(select product_id + 1000, category_id from product_category where product_id = 138);

insert into sku_info(id, master_name, image, product_id, market_price, price, score, inventory, display_sales_count, sales_count, sku_code,created_at )
    ( select id + 1000, master_name, image, product_id + 1000, market_price, price, score, inventory, sales_count, sales_count, sku_code,created_at
                from sku_info where id = 138);

insert into sku_image (sku_id, image, sort_order, created_at )
    (select sku_id + 1000, image, sort_order, created_at from sku_image  where sku_id = 138);

insert into product_attr(product_id, attr_id, attr_value, created_at) value ( 1138, 3, 1, now()  );
insert into product_attr(product_id, attr_id, attr_value, created_at) value ( 38, 3, 1, now()  );
insert into product_attr(product_id, attr_id, attr_value, created_at) value ( 39, 3, 1, now()  );




END
;;
DELIMITER ;
