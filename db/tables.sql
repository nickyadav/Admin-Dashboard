CREATE TABLE tbl_app (
	app_id INT NOT NULL DEFAULT unique_rowid(),
	app_name STRING(200) NULL,
	gateway_app_id STRING(100) NULL,
	gateway_name STRING(100) NULL,
	gateway_secret STRING(100) NULL,
	country STRING(50) NULL,
	currency STRING(50) NULL,
	language STRING(10) NULL,
	privacy_policy STRING NULL,
	term_condition STRING NULL,
	otp_attempt INTEGER NULL,
	app_icon STRING(500) NULL,
	status STRING(50) NULL,
	total_question INTEGER NULL,
	created_by STRING(50) NULL,
	created_at TIMESTAMP WITH TIME ZONE NULL DEFAULT now(),
	updated_by STRING(50) NULL,
	updated_at TIMESTAMP WITH TIME ZONE NULL DEFAULT now(),
	CONSTRAINT "primary" PRIMARY KEY (app_id ASC),
	UNIQUE INDEX tbl_app_gateway_app_id_key (gateway_app_id ASC),
	UNIQUE INDEX tbl_app_gateway_name_idx (gateway_name ASC),
	FAMILY "primary" (app_id, app_name, gateway_app_id, gateway_name, gateway_secret, country, currency, language, privacy_policy, term_condition, otp_attempt, app_icon, status, total_question, created_by, created_at, updated_by, updated_at)
);

CREATE TABLE tbl_user (
	user_id INT NOT NULL DEFAULT unique_rowid(),
	username STRING(255) NOT NULL,
	password_hash STRING(255) NOT NULL,
	password_reset_token STRING(255) NULL,
	email STRING(255) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL,
	created_by STRING(50) NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	updated_by STRING(50) NULL,
	user_fullname STRING(200) NULL,
	status STRING(20) NULL,
	role_id INTEGER NULL,
	reset_token_expire TIMESTAMP WITH TIME ZONE NULL,
	auth_key STRING(32) NULL,
	CONSTRAINT "primary" PRIMARY KEY (user_id ASC),
	UNIQUE INDEX tbl_user_email_key (email ASC),
	UNIQUE INDEX tbl_user_password_reset_token_key (password_reset_token ASC),
	UNIQUE INDEX tbl_user_username_key (username ASC),
	FAMILY "primary" (user_id, username, password_hash, password_reset_token, email, created_at, created_by, updated_at, updated_by, user_fullname, status, role_id, reset_token_expire, auth_key)
);

-- Initial user entry with username: admin and password: admin
INSERT INTO "tbl_user"
("username", "auth_key", "password_hash", "password_reset_token",
 "email", "created_at", "updated_at", "created_by", "updated_by", "user_fullname", "status", "role_id")
VALUES('admin', NULL, '$2a$10$AixZzoIwshMXMdTyRr1hU.80AFrY5v.XHmC4IurwzOj3FGIp7GdFS', NULL,
'yadav.hanik@gmail.com', '2018-12-22 09:48:58.60293+00:00', '2018-12-22 10:00:01.584793+00:00',
'SUPER ADMIN', 'SUPER ADMIN', 'Hanik', 'ACTIVE',NULL);