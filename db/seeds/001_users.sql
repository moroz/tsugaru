begin;

  truncate users;
  insert into users (email, password_hash) values ('user@example.com', '$argon2id$v=19$m=65536,t=3,p=4$6yxcJO9hMVMbcl9P3EvxwA$mAmLrQvb57aQmyDMa7usx71t786Cze7VkMAlJWBFDt4');

  insert into users (email) values ('nopassword@example.com');

commit;
