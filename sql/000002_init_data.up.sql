INSERT INTO `users` VALUES (1, 'donny', 'donny@arieffian.com', 'surabaya');

INSERT INTO `brands` VALUES (1, 'apple');
INSERT INTO `brands` VALUES (2, 'google');
INSERT INTO `brands` VALUES (3, 'amazon');
INSERT INTO `brands` VALUES (4, 'rasberry pi');

INSERT INTO `products` VALUES (1, 1, '123456789', 'macbook pro', 10, 1200);
INSERT INTO `products` VALUES (2, 2, '234567891','google home', 10, 1000);
INSERT INTO `products` VALUES (3, 3, '345678912','alexa speaker', 10, 1100);
INSERT INTO `products` VALUES (4, 4, '456789123','raspi', 10, 100);

INSERT INTO `transactions` VALUES (1, 1, '2021-09-01 12:00:00', 3400, 0, "");

INSERT INTO `transaction_detail` VALUES (1, 1, 1200, 1, 1200);
INSERT INTO `transaction_detail` VALUES (1, 2, 1000, 1, 1000);
INSERT INTO `transaction_detail` VALUES (1, 3, 1100, 1, 1100);