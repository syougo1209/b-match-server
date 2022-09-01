CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`     varchar(255) NOT NULL COMMENT 'ユーザー名',
    `created_at`  DATETIME(6) NOT NULL COMMENT 'ユーザー登録日',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `conversation`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '会話の識別子',
    `last_message_id` BIGINT UNSIGNED NOT NULL COMMENT '最後に送信されたメッセージのid',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー同士の会話';

CREATE TABLE `message`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '会話の識別子',
    `send_user_id` BIGINT UNSIGNED NOT NULL COMMENT 'メッセージを送ったユーザーのid',
    `conversation_id` BIGINT UNSIGNED NOT NULL COMMENT '紐ずく会話のid',
    `type`     TINYINT UNSIGNED NOT NULL COMMENT 'メッセージのテキスト',
    `text`     VARCHAR(1024)  COMMENT 'メッセージのテキスト',
    `created_at`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_send_user_id_in_message`
        FOREIGN KEY (`send_user_id`) REFERENCES `user` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_conversation_id_in_message`
        FOREIGN KEY (`conversation_id`) REFERENCES `conversation` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー同士の会話';

CREATE TABLE `conversation_state`
(
    `conversation_id` BIGINT UNSIGNED NOT NULL COMMENT '紐ずく会話のid',
    `from_user_id` BIGINT UNSIGNED NOT NULL COMMENT 'このテーブルの会話の状態をもつユーザーのid',
    `to_user_id` BIGINT UNSIGNED NOT NULL COMMENT 'メッセージをやりとりしているユーザーのid',
    `unread_messages_count` INT UNSIGNED NOT NULL COMMENT '未読のメッセージ数',
    `last_read_message_id`  BIGINT UNSIGNED NOT NULL COMMENT '最後に既読したメッセージのid',
    PRIMARY KEY (`conversation_id`,`from_user_id`, `to_user_id`),
    CONSTRAINT `fk_conversation_id_in_conversation_state`
        FOREIGN KEY (`conversation_id`) REFERENCES `conversation` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_from_user_id_in_conversation_state`
        FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_to_user_id_in_conversation_state`
        FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`)
            ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザーの会話の状態を保持するテーブル';
