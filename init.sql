CREATE DATABASE ewallet_db;

\c ewallet_db;

CREATE TABLE users (
	id bigserial primary key,
	name varchar not null,
	email varchar not null,
	password varchar not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp null,
	unique(email)
);

CREATE TABLE password_tokens (
	id bigserial primary key,
	user_id bigint not null,
	password_token varchar not null,
	expired_at timestamp default now() + interval '10 minutes',
	foreign key(user_id) references users(id)
);

CREATE SEQUENCE prefixed_seq;

CREATE TABLE wallets (
	id bigserial primary key,
	wallet_number varchar(13) unique not null default '777' || to_char(nextval('prefixed_seq'::regclass), 'FM0000000'),
	balance decimal default 0,
	user_id bigint not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp null,
	foreign key (user_id) references users(id)
);

CREATE TYPE methode AS ENUM('wallet','gacha','bank transfer','credit card','pay later');

CREATE TABLE history_transactions (
	id bigserial primary key,
	recipient_wallet_id bigint not null,
	sender_wallet_id bigint not null,
	amount decimal not null,
	source_of_fund methode not null,
	description varchar(35) null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp null
);

CREATE TABLE bank_account (
	id bigserial primary key,
	account_number bigint not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp null
);

CREATE TABLE credit_card_account (
	id bigserial primary key,
	cc_number bigint not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp null
);

CREATE TABLE pay_later_account (
	id bigserial primary key,
	user_id bigint not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp null,
	foreign key(user_id) references users(id)
);

INSERT INTO users ("name", email, "password", created_at, updated_at, deleted_at) VALUES
	('test', 'test5@test.com', '$2a$12$/d347Oj/mA1jGBqO46yQCu27/IrUEqg2je5B5.E4dvTQCM./qxQkC', '2024-02-23 15:14:18.999882', '2024-02-23 15:14:18.999882', NULL),
	('test', 'test4@test.com', '$2a$12$IymtlNvkp17C9SdLCb.4TOOqOYFFcDOHTmNMua0yDK1QaJQNa/XqS', '2024-02-23 15:14:26.490714', '2024-02-23 15:14:26.490714', NULL),
	('test', 'test3@test.com', '$2a$12$tvqwFg7DP/pberu7LxRD4uc1hkcDTWEloQHqcvtN2WCTHnHwEN6Qq', '2024-02-23 15:14:31.642305', '2024-02-23 15:14:31.642305', NULL),
	('test', 'test2@test.com', '$2a$12$YGaF.B5N9m7.r3/jnSHqnea4x7OmZ35MrtncdvAlFOXWLzajKvIvy', '2024-02-23 15:14:37.345584', '2024-02-23 15:14:37.345584', NULL),
	('test', 'test1@test.com', '$2a$12$gm66jhywLkWwQqQOXmn5q.riSkV8cpybB/wz8c3sLXjUyoVoIGLtK', '2024-02-23 15:14:41.406622', '2024-02-23 15:14:41.406622', NULL);

INSERT INTO history_transactions (recipient_wallet_id, sender_wallet_id, amount, source_of_fund, description, created_at, updated_at, deleted_at) VALUES
	(1, 1, 100000, 'credit card', '', '2024-02-23 15:16:00.066975', '2024-02-23 15:16:00.066975', NULL),
	(1, 1, 100000, 'bank transfer', '', '2024-02-23 15:16:10.691772', '2024-02-23 15:16:10.691772', NULL),
	(1, 1, 100000, 'pay later', '', '2024-02-23 15:16:21.258226', '2024-02-23 15:16:21.258226', NULL),
	(1, 2, 100000, 'pay later', '', '2024-02-23 15:16:23.314091', '2024-02-23 15:16:23.314091', NULL),
	(3, 1, 100000, 'wallet', '', '2024-02-23 15:16:44.976723', '2024-02-23 15:16:44.976723', NULL),
	(3, 1, 100000, 'wallet', '', '2024-02-23 15:16:46.971549', '2024-02-23 15:16:46.971549', NULL),
	(2, 1, 100000, 'wallet', '', '2024-02-23 15:16:57.645442', '2024-02-23 15:16:57.645442', NULL),
	(5, 1, 100000, 'wallet', '', '2024-02-23 15:17:04.029185', '2024-02-23 15:17:04.029185', NULL),
	(3, 2, 100000, 'bank transfer', '', '2024-02-23 15:18:15.178116', '2024-02-23 15:18:15.178116', NULL),
	(3, 2, 100000, 'credit card', '', '2024-02-23 15:18:22.163864', '2024-02-23 15:18:22.163864', NULL),
	(3, 3, 100000, 'pay later', '', '2024-02-23 15:18:28.161817', '2024-02-23 15:18:28.161817', NULL),
	(1, 3, 10000, 'wallet', '', '2024-02-23 15:18:50.559241', '2024-02-23 15:18:50.559241', NULL),
	(4, 3, 10000, 'wallet', '', '2024-02-23 15:18:58.157861', '2024-02-23 15:18:58.157861', NULL),
	(2, 3, 10000, 'wallet', '', '2024-02-23 15:19:11.889645', '2024-02-23 15:19:11.889645', NULL),
	(5, 4, 100000, 'pay later', '', '2024-02-23 15:20:12.621977', '2024-02-23 15:20:12.621977', NULL),
	(5, 3, 100000, 'bank transfer', '', '2024-02-23 15:20:18.944032', '2024-02-23 15:20:18.944032', NULL),
	(5, 3, 100000, 'credit card', '', '2024-02-23 15:20:24.870245', '2024-02-23 15:20:24.870245', NULL),
	(2, 5, 10000, 'wallet', '', '2024-02-23 15:20:38.864664', '2024-02-23 15:20:38.864664', NULL),
	(4, 5, 10000, 'wallet', '', '2024-02-23 15:20:46.438748', '2024-02-23 15:20:46.438748', NULL),
	(2, 5, 10000, 'wallet', '', '2024-02-23 15:20:51.915594', '2024-02-23 15:20:51.915594', NULL);


INSERT INTO bank_account (account_number, created_at, updated_at, deleted_at) VALUES
	(2333232, '2024-02-23 15:16:10.691772', '2024-02-23 15:16:10.691772', NULL),
	(2333232, '2024-02-23 15:18:15.178116', '2024-02-23 15:18:15.178116', NULL),
	(2333232, '2024-02-23 15:20:18.944032', '2024-02-23 15:20:18.944032', NULL);


INSERT INTO credit_card_account (cc_number, created_at, updated_at, deleted_at) VALUES
	(2333232, '2024-02-23 15:16:00.066975', '2024-02-23 15:16:00.066975', NULL),
	(2333232, '2024-02-23 15:18:22.163864', '2024-02-23 15:18:22.163864', NULL),
	(2333232, '2024-02-23 15:20:24.870245', '2024-02-23 15:20:24.870245', NULL);
