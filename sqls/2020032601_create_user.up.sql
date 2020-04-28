create table `user` (
 `id` int not null auto_increment
,`name` varchar(20)
,`age` int not null
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp on update current_timestamp
,primary key(`id`)
,unique(`name`)
) Engine=InnoDB;
