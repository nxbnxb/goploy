package model

import (
	"database/sql"
	"errors"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Pagination struct
type Pagination struct {
	Page  uint64 `json:"page"`
	Rows  uint64 `json:"rows"`
	Total uint64 `json:"total"`
}

// state type
const (
	Fail = iota
	Success
)

// state type
const (
	Disable = iota
	Enable
)

// DB init when the program start
var DB *sql.DB

// Init DB open
func Init() {
	dbType := os.Getenv("DB_TYPE")
	dbConn := os.Getenv("DB_CONN")
	var err error
	DB, err = sql.Open(dbType, dbConn)
	if err != nil {
		log.Fatal(err)
	}
}

// PaginationFrom param return pagination struct
func PaginationFrom(param url.Values) (Pagination, error) {
	page, err := strconv.ParseUint(param.Get("page"), 10, 64)
	if err != nil {
		return Pagination{}, errors.New("invalid page")
	}
	rows, err := strconv.ParseUint(param.Get("rows"), 10, 64)
	if err != nil {
		return Pagination{}, errors.New("invalid rows")
	}
	pagination := Pagination{Page: page, Rows: rows}
	return pagination, nil
}

const ddl string = "CREATE DATABASE IF NOT EXISTS `goploy`;CREATE TABLE IF NOT EXISTS `goploy`.`log`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '日志类型',  `ip` int(10) UNSIGNED NOT NULL DEFAULT 0,  `desc` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',  PRIMARY KEY (`id`) USING BTREE,  INDEX `idx_create_time`(`create_time`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`project`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `group_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '项目名称',  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '项目仓库地址',  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '项目部署路径',  `symlink_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '软链源路径',  `environment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '生产环境' COMMENT '部署环境',  `branch` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'master' COMMENT '分支',  `after_pull_script_mode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '脚本类型',  `after_pull_script` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '脚本路径',  `after_deploy_script_mode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '脚本类型',  `after_deploy_script` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '脚本路径',  `rsync_option` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'rsync 参数',  `auto_deploy` tinyint(4) UNSIGNED NOT NULL DEFAULT 1 COMMENT '0=>关闭 1=>Webhook',  `state` tinyint(4) UNSIGNED NOT NULL DEFAULT 1 COMMENT '0=>失效 1=>生效',  `deploy_state` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0=>未构建 1=>构建中 2=>成功 3=>失败',  `publisher_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `publisher_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `last_publish_token` char(36) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',  `notify_type` tinyint(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1=>企业微信 2=>钉钉',  `notify_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推送目标，目前只支持webhook',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,   PRIMARY KEY (`id`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`project_server`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `project_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `server_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE,  UNIQUE INDEX `uk_project_server`(`project_id`, `server_id`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`project_user`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `project_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE,  UNIQUE INDEX `uk_project_user`(`project_id`, `user_id`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`project_task` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `project_id` int(10) unsigned NOT NULL DEFAULT '0',  `commit_id` char(40) NOT NULL DEFAULT '',  `date` datetime DEFAULT NULL,  `state` tinyint(4) unsigned NOT NULL DEFAULT '1',  `is_run` tinyint(4) unsigned NOT NULL DEFAULT '0',  `creator_id` int(10) unsigned NOT NULL DEFAULT '0',  `creator` varchar(255) NOT NULL DEFAULT '',  `editor_id` int(10) unsigned NOT NULL DEFAULT '0',  `editor` varchar(255) NOT NULL DEFAULT '',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE,  KEY `index_project_update` (`project_id`,`update_time`) USING BTREE COMMENT 'project_id,update_time') ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`group`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`publish_trace` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `token` char(36) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',  `project_id` int(10) unsigned NOT NULL DEFAULT '0',  `project_group_id` int(10) unsigned NOT NULL DEFAULT '0',  `project_name` varchar(255) NOT NULL DEFAULT '',  `detail` longtext NOT NULL,  `state` tinyint(4) unsigned NOT NULL DEFAULT '1',  `publisher_id` int(10) unsigned NOT NULL DEFAULT '0',  `publisher_name` varchar(255) NOT NULL DEFAULT '',  `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1拉代码前脚本，2.git获取代码，3拉代码后脚本，4部署前脚本，5部署日志，6部署后脚本',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  `ext` longtext NOT NULL,  PRIMARY KEY (`id`) USING BTREE,  KEY `idx_project_id` (`project_id`) USING BTREE COMMENT 'project_id') ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;CREATE TABLE IF NOT EXISTS `goploy`.`server`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `group_id` int(10) UNSIGNED NOT NULL DEFAULT 0,  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `port` smallint(10) UNSIGNED NOT NULL DEFAULT 22,  `owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `last_publish_token` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  `state` tinyint(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '0=>失效 1=>生效',  PRIMARY KEY (`id`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`crontab` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `command` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `command_md5` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'command md5 for replace',  `creator_id` int(10) unsigned NOT NULL DEFAULT '0',  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `editor_id` int(10) unsigned NOT NULL DEFAULT '0',  `editor` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE,  UNIQUE KEY `uk_command_md5` (`command_md5`) USING BTREE) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`crontab_server` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `crontab_id` int(10) unsigned NOT NULL,  `server_id` int(10) unsigned NOT NULL,  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE,  UNIQUE KEY `idx_crontab_server` (`crontab_id`,`server_id`) USING BTREE) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`template` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `package_id_str` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,  `script` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`package` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `size` int(10) unsigned NOT NULL DEFAULT '0',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  PRIMARY KEY (`id`) USING BTREE) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`install_trace` (  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,  `token` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `server_id` int(10) unsigned NOT NULL DEFAULT '0',  `server_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `detail` longtext NOT NULL,  `state` tinyint(4) unsigned NOT NULL DEFAULT '1',  `operator_id` int(10) unsigned NOT NULL DEFAULT '0',  `operator_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1rsync 2ssh 3script',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  `ext` text NOT NULL,  PRIMARY KEY (`id`) USING BTREE,  KEY `idx_project_id` (`server_id`) USING BTREE COMMENT 'project_id') ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;CREATE TABLE IF NOT EXISTS `goploy`.`user`  (  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `password` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `mobile` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',  `role` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'member',  `manage_group_str` text CHARACTER SET utf8mb4 NOT NULL COMMENT '管理分组逗号分割(空串没有管理权限，all管理全部)',  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '0:=被禁用  1=正常',  `insert_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  `last_login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',  PRIMARY KEY (`id`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;"

const dml string = "INSERT INTO `goploy`.`user` (`account`, `password`, `name`, `mobile`, `role`, `manage_group_str`, `state`) VALUES ('admin', '$2a$10$89ZJ2xeJj35GOw11Qiucr.phaEZP4.kBX6aKTs7oWFp1xcGBBgijm', '超管', '18023496666', 'admin' ,'all', 1);"

func ImportSQL(db *sql.DB) error {
	for _, query := range strings.Split(ddl, ";") {
		if len(query) == 0 {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	for _, query := range strings.Split(dml, ";") {
		if len(query) == 0 {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}