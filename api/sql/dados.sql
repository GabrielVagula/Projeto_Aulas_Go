insert into usuarios (nome, nick, email, senha)
values
("Usuário 1","usuario_1","usuario1@gmail.com","$2a$10$L/PuEAFlxSn3uurDFRt3oO8uuDBmoP24HrYUO2M/iX0naHBY567UK"),
("Usuário 2","usuario_2","usuario2@gmail.com","$2a$10$L/PuEAFlxSn3uurDFRt3oO8uuDBmoP24HrYUO2M/iX0naHBY567UK"),
("Usuário 3","usuario_3","usuario3@gmail.com","$2a$10$L/PuEAFlxSn3uurDFRt3oO8uuDBmoP24HrYUO2M/iX0naHBY567UK");

insert into seguidores (usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicação do usuario 1", "Essa é a publicação teste do usuário 1!", 1),
("Publicação do usuario 2", "Essa é a publicação teste do usuário 2!", 2),
("Publicação do usuario 3", "Essa é a publicação teste do usuário 3!", 3);
