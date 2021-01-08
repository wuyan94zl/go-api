-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES ('1', 'admin@wuyan.com', '$2a$10$hp36hYrKnH6O8DdriL9o4uiTKH/RsSG9eua2t9/LCt5x3xzhVkI6q', '无言', '2020-12-23 11:02:47', '2021-01-07 11:48:46', '15362531451');

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES ('1', '0', '权限系统', 'el-icon-lx-settings', '/', '1', '2020-12-23 11:03:09', '2021-01-07 10:21:13');
INSERT INTO `menus` VALUES ('2', '1', '菜单权限', 'el-icon-lx-qrcode', '/menu', '1', '2021-01-06 16:08:38', '2021-01-07 10:34:22');
INSERT INTO `menus` VALUES ('3', '1', '角色管理', 'el-icon-lx-friendadd', '/role', '2', '2020-12-23 11:03:24', '2021-01-07 10:32:54');
INSERT INTO `menus` VALUES ('4', '1', '用户管理', 'el-icon-lx-profile', '/admin', '3', '2020-12-23 11:03:46', '2021-01-07 10:34:36');

-- ----------------------------
-- Records of role_menus
-- ----------------------------
INSERT INTO `role_menus` VALUES ('1', '1');
INSERT INTO `role_menus` VALUES ('1', '2');
INSERT INTO `role_menus` VALUES ('1', '3');
INSERT INTO `role_menus` VALUES ('1', '4');

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES ('1', '管理员', '管理员', '2021-01-07 14:13:28', '2021-01-07 14:13:33');

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES ('1', '1');
