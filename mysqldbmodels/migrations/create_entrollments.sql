CREATE TABLE IF NOT EXISTS `entrollments` (
    `id` int NOT NULL COMMENT 'Id',
    `course_id` int NOT NULL COMMENT 'Courseid',
	`student_id` int NOT NULL COMMENT 'Studentid',
     PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;