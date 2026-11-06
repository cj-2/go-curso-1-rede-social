insert into usuarios (nome, nick, email, senha)
values
    ("Carlos", "caca", "caca@exemplo.com", "$2a$10$JCmosjEVIWt5vieKBTNM7em6oBSFz26AUQdMVrgJfFCAk0HhvWhry"),
    ("Juca", "juju", "juju@exemplo.com", "$2a$10$JCmosjEVIWt5vieKBTNM7em6oBSFz26AUQdMVrgJfFCAk0HhvWhry"),
    ("Toim", "totoim", "totoim@exemplo.com", "$2a$10$JCmosjEVIWt5vieKBTNM7em6oBSFz26AUQdMVrgJfFCAk0HhvWhry"),
    ("Tico", "titico", "titico@exemplo.com", "$2a$10$JCmosjEVIWt5vieKBTNM7em6oBSFz26AUQdMVrgJfFCAk0HhvWhry"),
    ("Robison", "robin", "robin@exemplo.com", "$2a$10$JCmosjEVIWt5vieKBTNM7em6oBSFz26AUQdMVrgJfFCAk0HhvWhry");

insert into seguidores (usuario_id, seguidor_id)
values
    (1, 2),
    (3, 2),
    (2, 4),
    (2, 5);
