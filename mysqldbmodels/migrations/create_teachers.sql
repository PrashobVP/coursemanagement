CREATE TABLE IF NOT EXISTS `teachers` (
    `id` int NOT NULL COMMENT 'Id',
    `name` varchar(255) NOT NULL COMMENT 'Name',
     PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;