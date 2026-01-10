create database if not exist A_Rede_Social;
use A_Rede_Social;

drop table if exists publicacoes;
drop table if exists seguidores;
drop table if exists usuarios;

create table usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

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
) ENGINE=INNODB;

create table publicacoes (
    id int auto_increment primary key,
    titulo varchar (50) not null,
    conteudo varchar (300) not null,
    
    autor_id int not null,
    foreign key (autor_id)
    references usuarios(id)
    on delete cascade,

    curtidas int default 0,

    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;