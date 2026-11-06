-- create database devbook;
-- use devbook;
-- create table usuarios (id int auto_increment primary key, nome varchar(50) not null, email varchar(50) not null) engine=INNODB;
-- show tables;
-- create user 'golang'@'%' identified by 'golang'; # golang é a senha

-- # liberanto todos os privilégios para um usuário
-- grant all privileges on devbook.* to 'golang'@'%';

create database if not exists devbook;
use devbook;

drop table if exists usuarios;
drop table if exists seguidores;

create table usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE = INNODB;

alter table usuarios modify senha varchar(100) not null;

create table seguidores(
    usuario_id int not null,
    foreign key (usuario_id) 
    references usuarios(id) 
    on delete cascade,
    
    seguidor_id int not null,
    foreign key (seguidor_id) 
    references usuarios(id) 
    on delete cascade,

    primary key (usuario_id, seguidor_id)
) ENGINE = INNODB;