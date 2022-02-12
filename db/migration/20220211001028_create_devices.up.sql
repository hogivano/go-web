CREATE TABLE devices (
  id BIGINT NOT NULL AUTO_INCREMENT,
  code VARCHAR(20) NOT NULL,
  name VARCHAR(255),
  user_Id BIGINT,
  created_by BIGINT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  UNIQUE (code),
  CONSTRAINT FK_device_user_id
  FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT FK_device_created_by
  FOREIGN KEY (created_by) REFERENCES users(id)
);