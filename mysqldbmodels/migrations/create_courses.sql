CREATE TABLE IF NOT EXISTS `courses` (
    `id` int NOT NULL COMMENT 'Id',
    `name` varchar(255) NOT NULL COMMENT 'Name',
    `teacher_id` int NOT NULL COMMENT 'TeacherID',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;