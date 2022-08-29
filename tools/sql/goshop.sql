/*
-----------------------------------------goshop_order----------------------------------------
*/

DROP DATABASE IF EXISTS mall;
CREATE DATABASE `mall`;
USE mall;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `order_id` bigint(20) NOT NULL,
  `store_id` bigint(20) NOT NULL,
  `member_id` bigint(20) NOT NULL,
  `subtotal` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '小计(商品总金额，未扣除折扣)',
  `grand_total` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '合计(应收总金额)',
  `total_paid` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '总付款金额',
  `total_qty_ordered` int(11) NOT NULL DEFAULT '0' COMMENT '总的订购数量',
  `shipping_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '运费',
  `discount_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
  `payment_type` tinyint(4) NOT NULL COMMENT '付款类型 1,在线支付',
  `payment_status` tinyint(4) NOT NULL COMMENT '付款状态 1,未付款 2,已付款 3,部分付款',
  `payment_time` timestamp NULL DEFAULT NULL COMMENT '付款时间',
  `shipping_status` tinyint(4) NOT NULL COMMENT '发货状态 1,未发货 2,部分发货 3,已发货',
  `shipping_time` timestamp NULL DEFAULT NULL COMMENT '发货时间',
  `confirm` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否确认收货 1,否 2,是',
  `confirm_time` timestamp NULL DEFAULT NULL COMMENT '确认收货时间',
  `order_status` tinyint(4) NOT NULL COMMENT '订单状态 1,待付款 2,待审核 3,待发货 4,待收货 5,完成 6,待评价 7,取消',
  `order_type` tinyint(4) NOT NULL COMMENT '订单类型 1,普通订单',
  `refund_status` tinyint(4) NOT NULL COMMENT '退款状态 1,未退款 2,部分退款 3,全部退款',
  `return_status` tinyint(4) NOT NULL COMMENT '退货状态 1,未退货 2,部分退货 3,全部退货',
  `user_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户备注',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES (201208161539820, 0, 2, 490.00, 490.00, 0, 4, 0.00, 0.00, 1, 1, NULL, 1, NULL, 1, NULL, 1, 1, 1, 1, '', '2020-12-08 16:15:40', '2020-12-08 16:15:40', NULL);
COMMIT;

-- ----------------------------
-- Table structure for order_address
-- ----------------------------
DROP TABLE IF EXISTS `order_address`;
CREATE TABLE `order_address` (
  `order_address_id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `address_id` bigint(20) NOT NULL,
  `receiver` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收货人',
  `telephone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '电话',
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省',
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '市',
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '区县',
  `street` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '街道详细地址',
  `created_by` bigint(20) NOT NULL DEFAULT '0',
  `updated_by` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`order_address_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order_address
-- ----------------------------
BEGIN;
INSERT INTO `order_address` VALUES (201208161539834, 201208161539820, 0, '测试', '18662589099', '上海市', '上海市', '长宁区', '中山国际广场', 0, 0, '2020-12-08 16:15:40', '2020-12-08 16:15:40', NULL);
COMMIT;

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item` (
  `order_item_id` bigint(20) NOT NULL,
  `store_id` bigint(20) NOT NULL DEFAULT '0',
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `order_id` bigint(20) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品名称',
  `sku` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'SKU',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '规格图片或商品主图',
  `product_id` bigint(20) NOT NULL COMMENT '商品ID',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '销售价',
  `old_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '原始价',
  `cost_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '成本价',
  `total_payable` decimal(10,2) NOT NULL COMMENT '应付总金额',
  `total_discount_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '折扣总金额',
  `qty_ordered` int(11) NOT NULL DEFAULT '0' COMMENT '订购数量',
  `qty_shipped` int(11) NOT NULL DEFAULT '0' COMMENT '已发数量',
  `weight` decimal(10,2) NOT NULL COMMENT '重量',
  `volume` decimal(10,2) NOT NULL COMMENT '体积',
  `spec` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`order_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order_item
-- ----------------------------
BEGIN;
INSERT INTO `order_item` VALUES (201208161539827, 0, 0, 201208161539820, '梅三观短裙', '20201365', 'be65ad5693029d04ea632866515a427c.jpg', 4, 466.00, 360.00, 200.00, 466.00, 0.00, 1, 0, 0.40, 1.00, '[{\"name\":\"\",\"spec_value_id\":0,\"value\":\"\"}]', '2020-12-08 16:15:40', '2020-12-08 16:15:40');
INSERT INTO `order_item` VALUES (201208161539829, 0, 0, 201208161539820, '2020新款气质女装', '12', '2b4b0a53774a4b08ce247b4bba414806.jpg', 1, 12.00, 1.00, 1.00, 24.00, 0.00, 2, 0, 1.00, 1.00, '[{\"name\":\"尺码\",\"spec_value_id\":4,\"value\":\"xL\"},{\"name\":\"颜色\",\"spec_value_id\":7,\"value\":\"黑色\"}]', '2020-12-08 16:15:40', '2020-12-08 16:15:40');
INSERT INTO `order_item` VALUES (201208161539830, 0, 0, 201208161539820, '2020新款气质女装', '1112', '49f45fb882418f690fee13e5f5d7c47b.jpg', 1, 12.00, 1.00, 1.00, 12.00, 0.00, 1, 0, 1.00, 1.00, '[{\"name\":\"尺码\",\"spec_value_id\":6,\"value\":\"LL\"},{\"name\":\"颜色\",\"spec_value_id\":7,\"value\":\"黑色\"}]', '2020-12-08 16:15:40', '2020-12-08 16:15:40');
COMMIT;

-- ----------------------------
-- Table structure for order_payment
-- ----------------------------
DROP TABLE IF EXISTS `order_payment`;
CREATE TABLE `order_payment` (
  `order_payment_id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `shipping_amount` decimal(10,2) NOT NULL DEFAULT '0.00',
  `amount_paid` decimal(10,2) NOT NULL DEFAULT '0.00',
  `amount_ordered` decimal(10,2) NOT NULL DEFAULT '0.00',
  `payment_code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '付款方式code',
  `payment_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '付款方式名称',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`order_payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order_payment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for order_shipment
-- ----------------------------
DROP TABLE IF EXISTS `order_shipment`;
CREATE TABLE `order_shipment` (
  `order_shipment_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `store_id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `member_id` bigint(20) NOT NULL,
  `order_address_id` bigint(20) NOT NULL,
  `total_weight` decimal(10,2) NOT NULL COMMENT '总重量',
  `total_qty` int(11) NOT NULL COMMENT '总数量',
  `user_note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '客户备注',
  `carrier_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快递ID',
  `carrier_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递名称',
  `tracking_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递单号',
  `created_by` bigint(20) NOT NULL DEFAULT '0',
  `updated_by` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`order_shipment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order_shipment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for order_shipment_item
-- ----------------------------
DROP TABLE IF EXISTS `order_shipment_item`;
CREATE TABLE `order_shipment_item` (
  `order_shipment_item_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) NOT NULL,
  `order_item_id` bigint(20) NOT NULL,
  `order_shipment_id` bigint(20) NOT NULL,
  `product_id` bigint(20) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `sku` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `qty` int(11) NOT NULL DEFAULT '0',
  `weight` decimal(10,2) NOT NULL DEFAULT '0.00',
  `volume` decimal(10,2) NOT NULL DEFAULT '0.00',
  `spec` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`order_shipment_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of order_shipment_item
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;




/*
-----------------------------------------goshop_member----------------------------------------
*/

DROP DATABASE IF EXISTS goshop_member;
CREATE DATABASE `goshop_member`;
USE goshop_member;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `address_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` bigint(20) unsigned NOT NULL,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `code_prov` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '省',
  `code_city` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '市',
  `code_coun` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '区',
  `code_town` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '街道',
  `address_detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `room_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '门牌号',
  `is_default` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是默认地址（0否，1是）',
  `longitude` varchar(20) NOT NULL DEFAULT '' COMMENT '经度',
  `latitude` varchar(20) NOT NULL DEFAULT '' COMMENT '纬度',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`address_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='会员收货地址';

-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES (1, 2, '测试', '18662589099', 310000, 310100, 310105, 0, '中山国际广场', '', 1, '', '', '2020-12-08 16:12:44', '2020-12-08 16:12:44');
COMMIT;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `cart_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` bigint(20) unsigned NOT NULL,
  `product_id` bigint(20) unsigned NOT NULL,
  `product_spec_id` bigint(20) unsigned NOT NULL,
  `is_select` tinyint(1) DEFAULT '0' COMMENT '是否选中',
  `nums` int(5) unsigned DEFAULT '0' COMMENT '货品数量',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`cart_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='购物车';

-- ----------------------------
-- Records of cart
-- ----------------------------
BEGIN;
INSERT INTO `cart` VALUES (5, 3, 4, 9, 1, 5, '2020-12-08 16:30:00', '2020-12-08 16:30:02');
INSERT INTO `cart` VALUES (6, 3, 2, 7, 1, 1, '2020-12-08 16:31:42', '2020-12-08 16:31:42');
INSERT INTO `cart` VALUES (7, 2, 1, 1, 1, 1, '2020-12-08 16:37:34', '2020-12-08 16:37:34');
COMMIT;

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `member_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `register_ip` varchar(16) NOT NULL DEFAULT '' COMMENT '注册IP',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `gender` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '性别 0未知, 1女 2男',
  `id_card` varchar(18) NOT NULL DEFAULT '' COMMENT '身份证',
  `birthday` char(10) NOT NULL DEFAULT '' COMMENT '生日',
  `avatar` varchar(50) NOT NULL DEFAULT '' COMMENT '头像',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'EMAIL',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '会员状态 1正常 2 冻结',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `remark` varchar(50) NOT NULL DEFAULT '' COMMENT '备注',
  `member_level_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '会员等级ID',
  `point` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `balance` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '余额',
  `created_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `updated_by` bigint(20) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`member_id`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='会员表';

-- ----------------------------
-- Records of member
-- ----------------------------
BEGIN;
INSERT INTO `member` VALUES (1, '测试1', '18761911879', '', 'ceshi1', 0, '', '2020-12-08', '', '', 1, '2020-12-06 10:51:24', '', 0, 0, 0.00, 0, 0, '2020-12-08 13:51:24', '2020-12-08 13:51:24');
INSERT INTO `member` VALUES (2, '测试2', '18662589099', '', 'ceshi2', 1, '', '2020-12-08', '', '', 1, '2020-12-08 20:51:24', '', 0, 0, 0.00, 0, 5, '2020-12-08 16:11:36', '2020-12-08 16:14:15');
INSERT INTO `member` VALUES (3, '测试3', '18365288013', '', 'ceshi3', 0, '', '2020-12-08', '', '', 1, '2020-12-07 13:51:24', '', 0, 0, 0.00, 0, 0, '2020-12-08 16:23:11', '2020-12-08 16:23:11');
COMMIT;

-- ----------------------------
-- Table structure for member_third
-- ----------------------------
DROP TABLE IF EXISTS `member_third`;
CREATE TABLE `member_third` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` bigint(20) DEFAULT NULL COMMENT '会员id',
  `type` tinyint(1) DEFAULT '0' COMMENT '绑定第三方类型, 1.小程序',
  `open_id` varchar(255) DEFAULT NULL,
  `session_key` varchar(255) DEFAULT NULL,
  `unionid` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_member_id` (`member_id`),
  KEY `idx_open_id` (`open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='会员表第三方登录';

-- ----------------------------
-- Records of member_third
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for payment
-- ----------------------------
DROP TABLE IF EXISTS `payment`;
CREATE TABLE `payment` (
  `payment_id` varchar(20) NOT NULL COMMENT '支付单号',
  `money` decimal(10,2) DEFAULT '0.00' COMMENT '支付金额',
  `member_id` int(10) unsigned DEFAULT NULL COMMENT '用户ID',
  `type` smallint(1) unsigned NOT NULL DEFAULT '1' COMMENT '资源类型1=订单,2充值单',
  `status` tinyint(1) unsigned DEFAULT '1' COMMENT '支付状态 1=未支付 2=支付成功 3=其他',
  `payment_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '支付类型编码',
  `ip` varchar(50) DEFAULT NULL COMMENT '支付单生成IP',
  `params` varchar(200) NOT NULL COMMENT '支付的时候需要的参数，存的是json格式的一维数组',
  `payed_msg` varchar(255) DEFAULT NULL COMMENT '支付回调后的状态描述',
  `trade_no` varchar(50) DEFAULT NULL COMMENT '第三方平台交易流水号',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`payment_id`) USING BTREE,
  KEY `payment_id` (`payment_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `type` (`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付单表';

-- ----------------------------
-- Records of payment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for payment_rel
-- ----------------------------
DROP TABLE IF EXISTS `payment_rel`;
CREATE TABLE `payment_rel` (
  `payment_id` varchar(20) NOT NULL COMMENT '支付单编号',
  `source_id` varchar(20) NOT NULL COMMENT '资源编号',
  `money` decimal(8,2) unsigned NOT NULL COMMENT '金额',
  KEY `payment_id` (`payment_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='支付单明细表';

-- ----------------------------
-- Records of payment_rel
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
